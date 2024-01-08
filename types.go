package main

import "time"

// result
type Record struct {
	Id        int64     `json:"id" db:"id"`
	Name      string    `json:"name,omitempty" db:"name"`
	Marks     int       `json:"totalMarks" db:"totalMarks"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
}

// request json
type RecordRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

// response json
type RecordResponse struct {
	Code    int       `json:"code"`
	Msg     string    `json:"msg"`
	Records []*Record `json:"records"`
}
