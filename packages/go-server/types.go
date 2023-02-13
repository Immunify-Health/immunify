package main

import "math/rand"

type Record struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"Address"`
	PatientID int64  `json:"patientId"`
}

func NewRecord(firstName, lastName string) *Record {
	return &Record{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		PatientID: int64(rand.Intn(1000000)),
	}
}
