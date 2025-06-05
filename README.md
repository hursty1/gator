helpful commands:

psql -U postgres -h localhost -d gator

\c gator

\dt shows tables

sqlc generate


goose postgres postgres://postgres:postgres@localhost:5432/gator down

goose postgres postgres://postgres:postgres@localhost:5432/gator up




# Instructions

Requirements 
* Postgres installed
* Go 21 + 



## Commands
register - create user usage: gator register [user]
login - login to a user: gator login [username]
config - debug to print config file. Located at ~./.gatorconfig (contains username and postgres connection info see examlpe)
reset - resets Datbase tables
users - lists users 
agg - collects posts from followed feeds for the active user args: interval for checking 5s, 10m, 2h etc saves post to local db
addfeed - add a new feed to the system 
feeds - lists all feeds currently registered
follow - follow a existing feed
following - print feeds current user is following
unfollow - unfollow a followed feed
browse - view saved posts
