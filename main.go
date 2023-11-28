package main

import (
	"fmt"
	"hackernews/client"
)

const BaseURL = "https://hacker-news.firebaseio.com/v0/"

func main() {
	c := client.NewHackerNewsClient()

	fmt.Println(c)
}

func getTopStories() []string {
	return []string{}
}
