package shortfundly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
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

// SendRequest makes a request to Shortfundly's API
func (s *Shortfundly) SendRequest(r CRequest, data interface{}) error {
	fmt.Println(fmt.Sprintf("%s/%s", s.Host, r.Path))
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

// Error returns a string representing the error, satisfying the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("Error: %v", e.ErrMessage)
}

// GetTrendingFilms returns the film data which has been updated recently and has high views
func (s *Shortfundly) GetTrendingFilms(count int) ([]FilmResults, error) {
	fmt.Println(count)
	var (
		films  = &Films{}
		result = make([]FilmResults, 0)
		err    error
		r      CRequest
		params = url.Values{}
	)
	if count == 0 {
		r = CRequest{
			Method: "GET",
			Path:   "film/trending_films",
		}
		err = s.SendRequest(r, &films)
		return films.Result, err
	} else {
		pageNo := count / 5
		pageNo++
		fmt.Println(pageNo)

		// result = append(result, films.Result...)
		resultCh := make(chan []FilmResults)
		quit := make(chan bool, 1)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			for {
				select {
				case rslts := <-resultCh:
					result = append(result, rslts...)
				case <-quit:
					wg.Done()
				}

			}
		}()

		loopWg := sync.WaitGroup{}
		fmt.Println("---", pageNo)

		for i := 1; i <= pageNo; i++ {
			loopWg.Add(1)
			go func(pgNum int) {

				params.Set("p", strconv.Itoa(pgNum))
				r = CRequest{
					Method: "GET",
					Path:   fmt.Sprintf("film/trending_films?%v", params.Encode()),
				}
				err = s.SendRequest(r, &films)
				if pgNum == pageNo {
					remainingCount := count - (pgNum-1)*5
					resultCh <- films.Result[0:remainingCount]
				} else {
					resultCh <- films.Result
				}

				loopWg.Done()
			}(i)
		}
		loopWg.Wait()
		quit <- true
		wg.Wait()
		return result, err
	}
}
