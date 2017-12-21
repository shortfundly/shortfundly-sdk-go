# shortfundly-sdk-go
A Golang SDK for Shortfundly API

## Installation
```go
go get github.com/shortfundly/shortfundly-sdk-go
```

## Avalailable Languages
```
1. Tamil
2. Malayalam
3. Kannada
4. Telugu
```
## Example
```go

package main

import (
	"fmt"

	shortfundly "github.com/shortfundly/shortfundly-sdk-go"
)

func main() {

	// Initialise your application key to start the application
	s := shortfundly.New("YOUR_API_KEY")
	
	// Each page contains five list of film details

	// trending gets the trending film data
	trending, err := s.GetTrendingFilms()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(trending)

	// trending gets the trending film data in the page number of 2
	trending2, err := s.GetTrendingFilms(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(trending2)

	// topRated gets the high rating film data in the pageNo of 265
	topRated, err := s.GetTopRatedFilms(265)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(topRated)

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