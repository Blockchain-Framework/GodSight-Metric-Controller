package iam_service

import "github.com/Blockchain-Framework/controller/internal/godsight-aggregator/models"

type IAMResponseOrganizationList struct {
	Action    string            `json:"action"`
	Data      organizationsData `json:"Data"`
	Timestamp string            `json:"timestamp"`
	TraceId   string            `json:"traceId"`
}

type IAMResponseOrganizationInfo struct {
	Action    string           `json:"action"`
	Data      organizationData `json:"Data"`
	Timestamp string           `json:"timestamp"`
	TraceId   string           `json:"traceId"`
}

type organizationsData struct {
	Payload organizationsPayload `json:"payload"`
	Meta    models.Meta          `json:"meta"`
}

type organizationData struct {
	Payload organizationPayload `json:"payload" validate:"dive"`
	Meta    models.Meta         `json:"meta"`
}

type organizationsPayload struct {
	Organizations []models.Organization `json:"organizations"`
}

type organizationPayload struct {
	Organization models.Organization `json:"organization"`
}

type IAMResponseOrganizationCreate struct {
	Action    string                 `json:"action"`
	Data      organizationCreateData `json:"Data"`
	Timestamp string                 `json:"timestamp"`
	TraceId   string                 `json:"traceId"`
}

type organizationCreateData struct {
	Payload struct {
		OrganizationId string `json:"organization_Id" validate:"required,uuid"`
	} `json:"payload" validate:"required,dive"`
	Meta models.Meta `json:"meta"`
}

type IAMResponseActivate struct {
	State     models.State `json:"state" validate:"required"`
	Action    string       `json:"action"`
	Timestamp string       `json:"timestamp"`
	TraceId   string       `json:"traceId"`
}

type IAMResponse struct {
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
	TraceId   string `json:"traceId"`
}

type usersPayload struct {
	Users []models.User `json:"users"`
}

type userPayload struct {
	User models.User `json:"user"`
}

type userCreateData struct {
	Payload struct {
		UserId string `json:"user_id" validate:"required,uuid"`
	} `json:"payload" validate:"required,dive"`
	Meta models.Meta `json:"meta"`
}

type IAMResponseUserList struct {
	Action    string    `json:"action"`
	Data      usersData `json:"Data"`
	Timestamp string    `json:"timestamp"`
	TraceId   string    `json:"traceId"`
}

type usersData struct {
	Payload usersPayload `json:"payload"`
	Meta    models.Meta  `json:"meta"`
}

type userData struct {
	Payload userPayload `json:"payload"`
	Meta    models.Meta `json:"meta"`
}

type IAMResponseUserCreate struct {
	Action    string         `json:"action"`
	Data      userCreateData `json:"Data"`
	Timestamp string         `json:"timestamp"`
	TraceId   string         `json:"traceId"`
}

type IAMResponseUserInfo struct {
	Action    string   `json:"action"`
	Data      userData `json:"Data"`
	Timestamp string   `json:"timestamp"`
	TraceId   string   `json:"traceId"`
}

type initPayload struct {
	OrganizationTypes []models.OrganizationType `json:"organization_types"`
}

type initData struct {
	Payload initPayload `json:"payload"`
	Meta    models.Meta `json:"meta"`
}

type IAMResponseInit struct {
	Action    string   `json:"action"`
	Data      initData `json:"Data"`
	Timestamp string   `json:"timestamp"`
	TraceId   string   `json:"traceId"`
}

type IAMResponseError struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}
