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
	return s.filmsCommon(pageNo, "trending_films")
}

// GetTopRatedFilms returns the high rating films
func (s *Shortfundly) GetTopRatedFilms(pageNo int) (*Films, error) {
	return s.filmsCommon(pageNo, "toprated")
}

// GetMostViewedFilms returns film data which has more number of views
func (s *Shortfundly) GetMostViewedFilms(pageNo int) (*Films, error) {
	return s.filmsCommon(pageNo, "most_viewed")
}

// GetMostLikedFilms returns film data which has more number of likes
func (s *Shortfundly) GetMostLikedFilms(pageNo int) (*Films, error) {
	return s.filmsCommon(pageNo, "most_liked")
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
	language = strings.ToLower(language)
	switch language {
	case "all":
		return s.filmsCommon(pageNo, "recent_films")
	default:
		return s.filmsCommon(pageNo, "recent_"+language)
	}
}

// filmsCommon returns the film details
func (s *Shortfundly) filmsCommon(pageNo int, filmData string) (*Films, error) {
	var (
		r      CRequest
		params = url.Values{}
	)
	if pageNo == 0 {
		r = CRequest{
			Method: "GET",
			Path:   "film/" + filmData,
		}
	} else {
		params.Set("p", strconv.Itoa(pageNo))
		r = CRequest{
			Method: "GET",
			Path:   fmt.Sprintf("film/%v?%v", filmData, params.Encode()),
		}
	}
	films := &Films{}
	err := s.sendRequest(r, &films)
	return films, err
}
