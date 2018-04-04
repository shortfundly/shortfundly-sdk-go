package shortfundly

import (
	"fmt"
	"net/url"
	"strconv"
)

type FilmMakers struct {
	Count   int64              `json:"count"`
	Results []FilmMakerDetails `json:"results"`
}

type FilmMakerDetails struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Avatar         string `json:"avatar"`
	DateRegistered string `json:"date_registered"`
	Type           string `json:"type"`
	Local          string `json:"local"`
	Country        string `json:"country"`
	Bio            string `json:"bio"`
	Views          string `json:"views"`
	Fblink         string `json:"fblink"`
	Twlink         string `json:"twlink"`
	Glink          string `json:"glink"`
	Gender         string `json:"gender"`
	Role           string `json:"role"`
	Likes          int64  `json:"likes"`
	Followers      int64  `json:"followers"`
	Following      string `json:"following"`
	FilmCount      string `json:"film_count"`
	ContactNumber  string `json:"contact_number"`
}

// GetFilmMakers returns the film-maker details
func (s *Shortfundly) GetFilmMakers(pageNo int) (*FilmMakers, error) {
	var (
		r      CRequest
		params = url.Values{}
	)
	if pageNo == 0 {
		r = CRequest{
			Method: "GET",
			Path:   "profile/topfilmmakers",
		}
	} else {
		params.Set("p", strconv.Itoa(pageNo))
		r = CRequest{
			Method: "GET",
			Path:   fmt.Sprintf("profile/topfilmmakers?%v", params.Encode()),
		}
	}
	fResults := &FilmMakers{}
	err := s.sendRequest(r, &fResults)
	return fResults, err
}
