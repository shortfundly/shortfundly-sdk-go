package shortfundly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	// host is the default host of Shortfundly's API.
	host = "http://api.shortfundly.com"
	// defaultTimeout is the default timeout duration used on HTTP requests.
	defaultTimeout = 5 * time.Second
)

// Error defines the error handling
type Error struct {
	Status     bool   `json:"status"`
	ErrMessage string `json:"error"`
}

// Shortfundly defines the Shortfundly client
type Shortfundly struct {
	APIKey  string
	Host    string
	Timeout time.Duration
}

// CRequest defines the request given by the client to Shortfundly
type CRequest struct {
	Method string
	Path   string
}

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

// New returns a new Shortfundly's API client credentials which can be used to make the requests.
func New(key string) *Shortfundly {
	return &Shortfundly{
		APIKey:  key,
		Host:    host,
		Timeout: defaultTimeout,
	}
}

// Error returns a string representing the error, satisfying the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("Error: %v", e.ErrMessage)
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

// SendRequest makes a request to Shortfundly's API
func (s *Shortfundly) SendRequest(r CRequest, data interface{}) error {
	req, err := http.NewRequest(r.Method, fmt.Sprintf("%s/%s", s.Host, r.Path), strings.NewReader(""))
	if err != nil {
		return Error{ErrMessage: fmt.Sprintf("Unable to create the request: %s", err)}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", s.APIKey)

	httpClient := &http.Client{Timeout: s.Timeout}

	resp, err := httpClient.Do(req)
	if err != nil {
		return Error{ErrMessage: fmt.Sprintf("Failed to make request: %s", err)}
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Error{ErrMessage: fmt.Sprintf("Could not read response: %s", err)}
	}

	if resp.StatusCode == http.StatusOK {
		return json.Unmarshal(contents, &data)
	}

	errg := Error{}
	json.Unmarshal(contents, &errg)
	return Error{ErrMessage: errg.ErrMessage}
}
