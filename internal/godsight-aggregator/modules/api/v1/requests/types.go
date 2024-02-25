package requests

import (
	"encoding/json"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/models"
)

type RequestOrganizationInsertionBody struct {
	OrganizationName string `json:"name" validate:"required"`
	Type             string `json:"type" validate:"required"`
	AddressLine1     string `json:"address_line_1"`
	AddressLine2     string `json:"address_line_2"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Zip              string `json:"zip"`
	ContactPerson    string `json:"contact_person"`
	ContactDetails   string `json:"contact_details"`
	Website          string `json:"website"`
	Comments         string `json:"comments"`
}

type RequestOrganizationUpdateBody struct {
	OrganizationName string `json:"name" validate:"required"`
	AddressLine1     string `json:"address_line_1"`
	AddressLine2     string `json:"address_line_2"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Zip              string `json:"zip"`
	ContactPerson    string `json:"contact_person"`
	ContactDetails   string `json:"contact_details"`
	Website          string `json:"website"`
	Comments         string `json:"comments"`
}

type RequestUserBody struct {
	OrganizationID string          `json:"organization_id"`
	OtherName      string          `json:"other_name"`
	FirstName      string          `json:"first_name" validate:"required"`
	LastName       string          `json:"last_name" validate:"required"`
	NameInitials   string          `json:"name_initials"`
	Email          string          `json:"email" validate:"required,email"`
	State          models.State    `json:"state" default:"DEACTIVE"`
	Roles          json.RawMessage `json:"roles"`
	ContactNumber  string          `json:"contact_number"`
}
