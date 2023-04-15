package models

import (
	"strings"
	"time"
)

type Tax struct {
	TotalFee int            `json:"total_fee"`
	History  map[string]int `json:"history"`
}

type NewCongestionTax struct {
	City        string      `json:"city" validate:"required"`
	VehicleType string      `json:"vehicle_type" validate:"required,oneof=car emergency bus diplomat motorcycle military foreign tractor"`
	Times       []CivilTime `json:"times" validate:"required"`
}
type CivilTime struct {
	time.Time
}

const CustomTimeLayout = "2006-01-02 15:04:05"

func (c *CivilTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	t, err := time.Parse(CustomTimeLayout, str) //parse time
	if err != nil {
		return err
	}
	c.Time = t //set result using the pointer
	return nil
}
