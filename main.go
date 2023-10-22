package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const apiKey = "49456689fd2046608bd409d8facbcff0"

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type NewsResponse struct {
	Articles []Article `json:"articles"`
}

func getTopHeadlines() ([]Article, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?apiKey=%s", apiKey)

	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}

	var newsResponse NewsResponse
	if err := json.Unmarshal(resp.Body(), &newsResponse); err != nil {
		return nil, err
	}

	return newsResponse.Articles, nil
}

func main() {
	articles, err := getTopHeadlines()
	if err != nil {
		log.Fatalf("Error fetching headlines: %v", err)
	}

	fmt.Println("Top Headlines:")
	for _, article := range articles {
		fmt.Printf("%s\n%s\n%s\n\n", article.Title, article.Description, article.URL)
	}
}



