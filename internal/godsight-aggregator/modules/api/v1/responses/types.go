package responses

import (
	"encoding/json"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/models"
	"time"
)

type OrganizationResponse struct {
	OrganizationId   string       `json:"organization_id" validate:"required"`
	OrganizationName string       `json:"organization_name" validate:"required"`
	Type             string       `json:"organization_type" validate:"required"`
	State            models.State `json:"organization_state" validate:"required"`
	AddressLine1     string       `json:"address_line_1"`
	AddressLine2     string       `json:"address_line_2"`
	City             string       `json:"city"`
	Country          string       `json:"country"`
	Zip              string       `json:"zip"`
	ContactPerson    string       `json:"contact_person"`
	ContactDetails   string       `json:"contact_details"`
	Website          string       `json:"website"`
	OnboardedAt      time.Time    `json:"onboarded_at" validate:"required"`
	ActivatedAt      time.Time    `json:"activated_at"`
	ActiveUserCount  int          `json:"active_user_count"`
	UserCount        int          `json:"user_count"`
	Comments         string       `json:"comments"`
}

type OrganizationListResponse struct {
	Organizations []OrganizationResponse `json:"organizations"`
}

type OrganizationInfoResponse struct {
	Organization OrganizationResponse `json:"organization"`
}

type ResponseMeta struct {
	Page     int `json:"page"`
	NextPage int `json:"next_page"`
	Count    int `json:"count"`
}

type ResponseOrganizationCreate struct {
	OrganizationId string `json:"organization_id" validate:"required,uuid"`
}

type ResponseOrganizationActivate struct {
	State models.State `json:"state" validate:"required"`
}

type UserResponse struct {
	OrganizationID string          `json:"organization_id" validate:"required"`
	UserID         string          `json:"user_id" validate:"required"`
	OtherNames     string          `json:"other_names"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	NameInitials   string          `json:"name_initials"`
	Email          string          `json:"email" validate:"required"`
	State          models.State    `json:"user_state"`
	Roles          json.RawMessage `json:"roles"`
	OnboardedAt    time.Time       `json:"onboarded_at" validate:"required"`
	ActivatedAt    time.Time       `json:"activated_at"`
	ContactNumber  string          `json:"contact_number"`
}

type ResponseData struct {
	Message string
	State   int
}
