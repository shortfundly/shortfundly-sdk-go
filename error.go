package shortfundly

import "fmt"

// Error defines the error handling
type Error struct {
	Status     bool   `json:"status"`
	ErrMessage string `json:"error"`
}

// Error returns a string representing the error, satisfying the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("Error: %v", e.ErrMessage)
}
