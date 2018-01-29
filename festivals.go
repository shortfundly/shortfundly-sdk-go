package shortfundly

import (
	"fmt"
	"net/url"
	"strconv"
)

// Festivals defines the festival data
type Festivals struct {
	Count   int64             `json:"count"`
	Status  bool              `json:"status"`
	Results []FestivalResults `json:"results"`
}

// FestivalResults defines the festival content results
type FestivalResults struct {
	PkID             string `json:"pk_id"`
	Headline         string `json:"headline"`
	Theme            string `json:"theme"`
	Thumb            string `json:"thumb"`
	TimelineImg      string `json:"timeline_img"`
	EventType        string `json:"event_type"`
	YearsOfRunning   string `json:"years_of_running"`
	ContactAddress   string `json:"contact_address"`
	City             string `json:"city"`
	State            string `json:"state"`
	Pincode          string `json:"pincode"`
	Country          string `json:"country"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	ContactNumber    string `json:"contact_number"`
	OrganizersName   string `json:"organizers_name"`
	Description      string `json:"description"`
	AwardsPrize      string `json:"awards_prize"`
	Rules            string `json:"rules"`
	OpeningDate      string `json:"opening_date"`
	RegularDeadline  string `json:"regular_deadline"`
	ExtendedDeadline string `json:"extended_deadline"`
	NotificationDate string `json:"notification_date"`
	EventDate        string `json:"event_date"`
	Category         string `json:"category"`
	FbPage           string `json:"fb_page"`
	TwitterPage      string `json:"twitter_page"`
	SharedPublic     string `json:"shared_public"`
	CreatedDate      string `json:"created_date"`
	OpenState        string `json:"open_state"`
	EntryFees        string `json:"entry_fees"`
	Website          string `json:"website"`
	UserID           string `json:"user_id"`
	IsExpired        string `json:"is_expired"`
	Sharableurl      string `json:"sharableurl"`
}

// GetTodayFestival returns the festival data which has been updated today
func (s *Shortfundly) GetTodayFestival(pageNo int) (*Festivals, error) {
	return s.festiCommon(pageNo, "today")
}

// GetYesterdayFestival returns the festival data which has been updated yesterday
func (s *Shortfundly) GetYesterdayFestival(pageNo int) (*Festivals, error) {
	return s.festiCommon(pageNo, "yesterday")
}

// GetThisWeekFestival returns the festival data which has been updated this week
func (s *Shortfundly) GetThisWeekFestival(pageNo int) (*Festivals, error) {
	return s.festiCommon(pageNo, "thisweek")
}

// GetLastWeekFestival returns the festival data which has been updated last week
func (s *Shortfundly) GetLastWeekFestival(pageNo int) (*Festivals, error) {
	return s.festiCommon(pageNo, "lastweek")
}

// GetThisMonthFestival returns the festival data which has been updated this month
func (s *Shortfundly) GetThisMonthFestival(pageNo int) (*Festivals, error) {
	return s.festiCommon(pageNo, "thismonth")
}

// GetLastMonthFestival returns the festival data which has been updated last month
func (s *Shortfundly) GetLastMonthFestival(pageNo int) (*Festivals, error) {
	return s.festiCommon(pageNo, "lastmonth")
}

// festiCommon returns the festival details
func (s *Shortfundly) festiCommon(pageNo int, typeData string) (*Festivals, error) {
	var (
		r      CRequest
		params = url.Values{}
	)
	if pageNo == 0 {
		r = CRequest{
			Method: "GET",
			Path:   "festivals/searchByDate?type=" + typeData,
		}
	} else {
		params.Set("p", strconv.Itoa(pageNo))
		r = CRequest{
			Method: "GET",
			Path:   fmt.Sprintf("festivals/searchByDate?type=%v&%v", typeData, params.Encode()),
		}
	}
	fResults := &Festivals{}
	err := s.sendRequest(r, &fResults)
	return fResults, err
}
