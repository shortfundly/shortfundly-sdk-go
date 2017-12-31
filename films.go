package shortfundly

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Films defines the films data
type Films struct {
	Count  int64         `json:"count"`
	Result []FilmResults `json:"results"`
}

// FilmResults defines the film results
type FilmResults struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	Thumb       string `json:"thumb"`
	Views       string `json:"views"`
	Liked       string `json:"liked"`
	Duration    string `json:"duration"`
	Nsfw        string `json:"nsfw"`
	Owner       string `json:"owner"`
	RegID       string `json:"reg_id"`
	RegionName  string `json:"region_name"`
	CatName     string `json:"cat_name"`
	Source      string `json:"source"`
	Category    string `json:"category"`
	Userthumb   string `json:"userthumb"`
	Role        string `json:"role"`
	Followers   string `json:"followers"`
	Type        string `json:"type"`
	Date        string `json:"date"`
}

// GetTrendingFilms returns the film data which has been updated recently and has high views
func (s *Shortfundly) GetTrendingFilms(pageNo int) (*Films, error) {
	var (
		r      CRequest
		params = url.Values{}
	)
	if pageNo == 0 {
		r = CRequest{
			Method: "GET",
			Path:   "film/trending_films",
		}
	} else {
		params.Set("p", strconv.Itoa(pageNo))
		r = CRequest{
			Method: "GET",
			Path:   fmt.Sprintf("film/trending_films?%v", params.Encode()),
		}
	}
	films := &Films{}
	err := s.SendRequest(r, &films)
	return films, err
}

// GetTopRatedFilms returns the high rating films
func (s *Shortfundly) GetTopRatedFilms(pageNo int) (*Films, error) {
	var (
		r      CRequest
		params = url.Values{}
	)
	if pageNo == 0 {
		r = CRequest{
			Method: "GET",
			Path:   "film/toprated",
		}
	} else {
		params.Set("p", strconv.Itoa(pageNo))
		r = CRequest{
			Method: "GET",
			Path:   fmt.Sprintf("film/toprated?%v", params.Encode()),
		}
	}
	films := &Films{}
	err := s.SendRequest(r, &films)
	return films, err
}

// GetMostViewedFilms returns film data which has more number of views
func (s *Shortfundly) GetMostViewedFilms(pageNo int) (*Films, error) {
	var (
		r      CRequest
		params = url.Values{}
	)
	if pageNo == 0 {
		r = CRequest{
			Method: "GET",
			Path:   "film/most_viewed",
		}
	} else {
		params.Set("p", strconv.Itoa(pageNo))
		r = CRequest{
			Method: "GET",
			Path:   fmt.Sprintf("film/most_viewed?%v", params.Encode()),
		}
	}
	films := &Films{}
	err := s.SendRequest(r, &films)
	return films, err
}

// GetMostLikedFilms returns film data which has more number of likes
func (s *Shortfundly) GetMostLikedFilms(pageNo int) (*Films, error) {
	var (
		r      CRequest
		params = url.Values{}
	)
	if pageNo == 0 {
		r = CRequest{
			Method: "GET",
			Path:   "film/most_liked",
		}
	} else {
		params.Set("p", strconv.Itoa(pageNo))
		r = CRequest{
			Method: "GET",
			Path:   fmt.Sprintf("film/most_liked?%v", params.Encode()),
		}
	}
	films := &Films{}
	err := s.SendRequest(r, &films)
	return films, err
}

/*
GetRecentFilms returns the recent films with respect to available languages

Available languages :
	1. tamil
	2. malayalam
	3. telugu
	4. kannada

	Note : language field is valid only for recent source
*/
func (s *Shortfundly) GetRecentFilms(language string, pageNo int) (*Films, error) {
	var (
		r      CRequest
		params = url.Values{}
	)
	language = strings.ToLower(language)
	switch language {
	case "all":
		if pageNo == 0 {
			r = CRequest{
				Method: "GET",
				Path:   "film/recent_films",
			}
		} else {
			params.Set("p", strconv.Itoa(pageNo))
			r = CRequest{
				Method: "GET",
				Path:   fmt.Sprintf("film/recent_films?%v", params.Encode()),
			}
		}
	default:
		if pageNo == 0 {
			r = CRequest{
				Method: "GET",
				Path:   fmt.Sprintf("film/recent_%s", language),
			}
		} else {
			params.Set("p", strconv.Itoa(pageNo))
			r = CRequest{
				Method: "GET",
				Path:   fmt.Sprintf("film/recent_%s?%v", language, params.Encode()),
			}
		}
	}
	films := &Films{}
	err := s.SendRequest(r, &films)
	return films, err
}