package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hursty1/gator/internal/database"
)


type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil )
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return &RSSFeed{}, err
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	// Decode HTML entities in each item
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil


}

func createFeed(name, url string, s *state) (database.Feed, error) {
	current_user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return database.Feed{}, err
	}
	newFeed := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: current_user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), newFeed) 
	if err != nil {
		return database.Feed{}, err
	}
	return feed, nil
}

func feedFollow(url string, s *state, usr database.User) (database.CreateFeedFollowRow, error) {
	
	feed, err := s.db.FetchFeedByUrl(context.Background(), url)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		os.Exit(1)
		return database.CreateFeedFollowRow{}, err
	}
	ff := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: usr.ID,
		FeedID: feed.ID,
	}
	feed_follows, err := s.db.CreateFeedFollow(context.Background(), ff)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		os.Exit(1)
		return database.CreateFeedFollowRow{}, err

	}
	// fmt.Printf("%+v\n", feed_follows)
	return feed_follows, nil
}

func unfollowFeed(url string, usr database.User, s *state) error {
	feed, err := s.db.FetchFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}
	deleteFF := database.DeleteFeedFollowsParams {
		UserID: usr.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFeedFollows(context.Background(), deleteFF)
	if err != nil {
		return err
	}
	fmt.Printf("Feed %s has been unfollowed", feed.Name)
	return nil

}


func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(),nextFeed.ID)
	if err != nil {
		return err
	}
	feed, err := fetchFeed(context.Background(),nextFeed.Url)
	if err != nil {return err}

	for _, item := range feed.Channel.Item {
		// fmt.Printf("%+v\n", item.Title)
		fmt.Printf("Saving Post -----: %s\n", item.Title)
		savePosts(item,nextFeed,s)
	}
	return nil
}
func parseDateFlexible(dateStr string) (time.Time, error) {
	formats := []string{
		time.RFC1123Z, // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,  // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC822Z,  // "02 Jan 06 15:04 -0700"
		time.RFC3339,  // "2006-01-02T15:04:05Z07:00"
		"2006-01-02",  // ISO date
	}
	for _, layout := range formats {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
}
func savePosts(post RSSItem, feed database.Feed, s *state) error {
	pub_date, err := parseDateFlexible(post.PubDate)
	if err != nil {
		fmt.Printf("Failed to pase date %s\n", err)
		return err
	}
	db_post := database.CreatePostsParams{
		ID: uuid.New(),
		CreatedAt: pub_date,
		UpdatedAt: time.Now(),
		Title: fmt.Sprint(post.Title),
		Description: fmt.Sprint(post.Description),
		Url: fmt.Sprint(post.Link),
		FeedID: feed.ID,
	}
	saved_post, err := s.db.CreatePosts(context.Background(), db_post)
	if err != nil {
		fmt.Printf("Failed to save post: Error: %s\n", err)
		return err
	}
	fmt.Printf("Post %s has been saved\n", saved_post.Title)
	return nil
}

func fetchPosts(s *state, usr database.User, limit int) error {
	dbargs := database.GetPostsForUserParams{
		UserID: usr.ID,
		Limit: int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), dbargs)
	if err != nil {
		return err
	}

	for _, post := range posts {
		// fmt.Printf("%+v\n", post)
		
		fmt.Printf("%s from %s\n", post.UpdatedAt.Format("Mon Jan 2"), post.Title) 
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil

}