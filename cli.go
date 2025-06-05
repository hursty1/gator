package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/hursty1/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) !=1 {
			fmt.Println("username is required")
			os.Exit(1)
		}
	
	usr, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Username %s does not exist.\n", usr.Name)
		os.Exit(1)
	}
	err = s.config.SetUser(usr.Name)
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", cmd.args[0])
	return nil
}

func showConfig(s *state, cmd command) error {
	fmt.Println("User name", s.config.Current_user_name)
	fmt.Println("DB URL", s.config.Db_url)
	return nil
}


func register(s *state, cmd command) error {
	if len(cmd.args) !=1 {
			fmt.Println("username is required")
			os.Exit(1)
		}
	// createUser := database.CreateUserParams{uuid.New(), }
	newUser := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	}
	
	usr, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil{
		return err
	}
	fmt.Printf("User %s has been created.\n", usr.Name)
	err = handlerLogin(s,cmd)
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

func reset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Printf("Error resetting: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Users were reset.")
	
	return nil
}

func list(s*state, cmd command, usr database.User) error {
	usr_list, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error getting users: %s\n", err)
		os.Exit(1)
	}
	for _, usr := range usr_list {
		if usr.Name == s.config.Current_user_name {
			fmt.Printf("* %s (current)\n", usr.Name)
		} else {
			fmt.Printf("* %s\n", usr.Name)
		}
	}
	return nil
}

func agg(s *state, cmd command) error {
	if len(cmd.args) !=1 {
		fmt.Println("Time between feeds fetch is required")
		os.Exit(1)
	}
	dur, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		fmt.Printf("Format of duration was incorrect.\n")
		os.Exit(1)
	}
	fmt.Printf("Collecting feeds every %s\n", dur)
	ticker := time.NewTicker(dur)
	for ; ; <-ticker.C {
		fmt.Printf("Fetching Feeds\n")
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Errorf("Error %s", err)
			return err
		}
	}
	return nil
}

func addfeed(s *state, cmd command, usr database.User) error {
	if len(cmd.args) != 2 {
			fmt.Println("Args: name, url are required.")
			os.Exit(1)
	}
	feedObj, err := createFeed(cmd.args[0], cmd.args[1], s) 
	if err != nil {
		os.Exit(1)
	}
	// feed, err := fetchFeed(context.Background(), feedObj.Url)
	fmt.Printf("%+v\n", feedObj)
	// add feed follows
	ff, err := feedFollow(feedObj.Url, s, usr)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%+v\n", ff)
	return nil
}

func feeds(s *state, cmd command) error {
	//prints all the feeds in the database to the console
	feeds, err := s.db.FetchAllFeeds(context.Background())
	if err != nil {
		fmt.Errorf("Error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", feeds)
	return nil
}

func follow(s *state, cmd command, usr database.User) error {
	if len(cmd.args) !=1 {
			fmt.Println("url is required")
			os.Exit(1)
	}
	ff, err := feedFollow(cmd.args[0], s, usr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", ff)
	return nil

}

func following(s *state, cmd command, usr database.User) error {
	
	feed_follows, err := s.db.GetFeedFollowForUser(context.Background(), usr.ID)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		os.Exit(1)
		return err
	}
	for _, feed := range feed_follows {
		fmt.Printf("Feed Name: %s\n", feed.Name_2)
		fmt.Printf("Feed Url: %s\n", feed.Url)
	}
	
	fmt.Printf("User: %s\n", usr.Name)
	return nil
}

func unfollow(s *state, cmd command, usr database.User) error {
	err := unfollowFeed(cmd.args[0],usr, s)
	if err != nil {
		fmt.Println("Failed to delete feed")
		os.Exit(1)
	}
	return nil
}

func browse(s *state, cmd command, usr database.User) error {
	limit := 2
	var err error
	if len(cmd.args) ==1 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Println("Limit is not a valid int")
			os.Exit(1)
		}
	} 

	err = fetchPosts(s, usr, limit)
	if err != nil {
		fmt.Printf("Error is: %s", err)
		os.Exit(1)
	}
	return nil
}