package enum

import "errors"

// Status represents the status type
type Status string

const (
	Approved Status = "approved"
	Declined Status = "declined"
	Waiting  Status = "waiting"
)

// IsValid checks if the status is valid
func (s Status) IsValid() error {
	switch s {
	case Approved, Declined, Waiting:
		return nil
	}
	return errors.New("invalid status")
}

// String returns the string representation of the status
func (s Status) String() string {
	return string(s)
}
