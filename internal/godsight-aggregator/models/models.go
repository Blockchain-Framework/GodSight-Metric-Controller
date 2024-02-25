package models

import (
	"encoding/json"
	"time"
)

type State string
type OrganizationTypeValue string

type Organization struct {
	OrganizationId     string    `json:"organization_id" validate:"required,uuid"`
	OrganizationName   string    `json:"name" validate:"required"`
	Type               string    `json:"type" validate:"required,alpha"`
	State              State     `json:"state" validate:"required"`
	AddressLine1       string    `json:"address_line_1"`
	AddressLine2       string    `json:"address_line_2"`
	City               string    `json:"city"`
	Country            string    `json:"country"`
	Zip                string    `json:"zip"`
	ContactPerson      string    `json:"contact_person"`
	ContactDetails     string    `json:"contact_details"`
	Website            string    `json:"website"`
	OnboardTimestamp   time.Time `json:"onboard_timestamp"`
	ActivatedTimestamp time.Time `json:"activated_timestamp"`
	ActiveUserCount    int       `json:"active_user_count" default:"0"`
	UserCount          int       `json:"user_count" default:"0"`
	Comments           string    `json:"comments"`
}

type User struct {
	OrganizationID     string          `json:"organization_id" validate:"required,uuid"`
	UserID             string          `json:"user_id" validate:"required,uuid"`
	OtherName          string          `json:"other_name"`
	FirstName          string          `json:"first_name"`
	LastName           string          `json:"last_name"`
	NameInitials       string          `json:"name_initials"`
	Email              string          `json:"email" validate:"required,email"`
	State              State           `json:"state"`
	Roles              json.RawMessage `json:"roles"`
	OnboardTimestamp   time.Time       `json:"onboard_timestamp"`
	ActivatedTimestamp time.Time       `json:"activated_timestamp"`
	ContactNumber      string          `json:"contact_number"`
}

type OrganizationType struct {
	TypeId string                `json:"type_id" validate:"required"`
	Name   OrganizationTypeValue `json:"name" validate:"required"`
}

type Meta struct {
	Page         int `json:"page"`
	NextPage     int `json:"next_page"`
	RequestCount int `json:"request_count"`
	Count        int `json:"count"`
}
