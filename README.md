# shortfundly-sdk-go
A Golang SDK for Shortfundly API

```go

package main

import (
	"fmt"

	shortfundly "github.com/shortfundly/shortfundly-sdk-go"
)

func main() {

	// Initialise your application key to start the application
	s := shortfundly.New("YOUR_API_KEY")

	// trending gets the trending film data
	trending, err := s.GetFilms("trending_films")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(trending)

	// recent gets the recent uploaded film data in the pageNo of 265
	recent, err := s.GetFilms("recent_films", 265)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(recent)

	// recentTamil gets the recent tamil uploaded film data
	recentTamil, err := s.GetRecentFilms("tamil")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(recentTamil)

	// recentMalayalam gets the recent malayalam uploaded film data
	recentMalayalam, err := s.GetRecentFilms("malayalam", 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(recentMalayalam)
}


```