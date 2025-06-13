package Baro

import "time"

type BaroData struct {
	ID           string    `json:"id"`
	Activation   time.Time `json:"activation"`
	StartString  string    `json:"startString"`
	Expiry       time.Time `json:"expiry"`
	Active       bool      `json:"active"`
	Character    string    `json:"character"`
	Location     string    `json:"location"`
	Inventory    []any     `json:"inventory"`
	PsID         string    `json:"psId"`
	EndString    string    `json:"endString"`
	InitialStart time.Time `json:"initialStart"`
	Schedule     []any     `json:"schedule"`
}
