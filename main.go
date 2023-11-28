package main

import (
	"fmt"
	"hackernews/client"
)

const BaseURL = "https://hacker-news.firebaseio.com/v0/"

func main() {
	c := client.New(BaseURL)

	stories, err := c.GetTopStoriesIDs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Stories found: %d\n", len(stories))

	for i, story := range stories {
		fmt.Printf("processing story: %d\n", i)
		item, err := c.GetItem(story)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", item.Title)
	}
}
