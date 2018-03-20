package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mattheath/kala/bigflake"
	"github.com/monzo/gocassa"
)

const (
	host     = "127.0.0.1"
	username = ""
	password = ""
)

var postsFlakeSeries gocassa.FlakeSeriesTable

// Post represents a social media post
type Post struct {
	ID   string
	Body string
}

func init() {
	log.Printf("Connecting to Cassandra: %s", host)

	keyspace, err := gocassa.ConnectToKeySpace("posts", []string{host}, username, password)
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}

	log.Printf("Connected to Cassandra: %s", host)

	postsFlakeSeries = keyspace.FlakeSeriesTable(
		"posts_flakeseries",
		"ID",
		time.Hour,
		Post{},
	).WithOptions(gocassa.Options{
		TableName: "posts_flakeseries",
	})
}

func createFlakeID(prefix string) (string, error) {
	minter, err := bigflake.New(100)
	if err != nil {
		return "", err
	}

	flake, err := minter.Mint()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s_%s", prefix, flake.Base62WithPadding(22)), nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("usage: go run main.go <body>")
	}

	body := args[1]
	id, err := createFlakeID("post")
	if err != nil {
		log.Fatalf("Failed to generate flake ID: %v", err)
	}

	post := &Post{
		ID:   id,
		Body: body,
	}

	if err := postsFlakeSeries.Set(post).Run(); err != nil {
		log.Fatalf("Failed to save post: %v", err)
	}

	log.Printf("Created post: %+v", post)
}
