package domain

import "time"

type Employee struct {
	Id            int64      `json:"id"`
	NIK           string     `json:"nik"`
	Name          string     `json:"name"`
	PlaceOfBirth  string     `json:"placeOfBirth"`
	DateOfBirth   time.Time  `json:"dateOfBirth"`
	Gender        string     `json:"gender"`
	BloodType     string     `json:"bloodType"`
	Address       string     `json:"address"`
	Religion      int        `json:"religion"`
	MaritalStatus int        `json:"maritalStatus"`
	CreatedAt     time.Time  `json:"createdAt"`
	ModifiedAt    *time.Time `json:"modifiedAt"`
}
