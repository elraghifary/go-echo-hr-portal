package time

import (
	"time"
)

var (
	loc *time.Location
	Now = now
)

// Initiate time package
func Init() (err error) {
	loc, err = time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}

	return nil
}

// Returns time location
func GetLocation() *time.Location {
	return loc
}

// Returns current time by current location
func now() time.Time {
	return time.Now().In(loc)
}
