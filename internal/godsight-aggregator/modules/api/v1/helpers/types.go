package helpers

import "github.com/Blockchain-Framework/controller/internal/godsight-aggregator/models"

type ResponseOrganizationList struct {
	Data struct {
		Organizations []models.Organization `json:"organizations"`
	} `json:"data"`
	Meta models.Meta `json:"meta"`
}

type ResponseOrganization struct {
	Data struct {
		Organization models.Organization `json:"organization"`
	} `json:"data"`
	Meta models.Meta `json:"meta"`
}

type ResponseOrganizationCreate struct {
	Data struct {
		OrganizationId string `json:"organization_id"`
	} `json:"data"`
	Meta models.Meta `json:"meta"`
}

type ResponseActivate struct {
	Data struct {
		State models.State `json:"state"`
	} `json:"data"`
	Meta models.Meta `json:"meta"`
}

type ResponseError struct {
	Action string `json:"action"`
	Error  error  `json:"error"`
	Status int    `json:"status"`
}

type ResponseUserList struct {
	Data struct {
		Organizations []models.User `json:"users"`
	} `json:"data"`
	Meta models.Meta `json:"meta"`
}

type ResponseUser struct {
	Data struct {
		Organization models.User `json:"user"`
	} `json:"data"`
	Meta models.Meta `json:"meta"`
}

type ResponseUserCreate struct {
	Data struct {
		UserId string `json:"user_id"`
	} `json:"data"`
	Meta models.Meta `json:"meta"`
}

type ResponseInitData struct {
	Data struct {
		OrganizationTypes []models.OrganizationType `json:"organization_types"`
	} `json:"data"`
	Meta models.Meta `json:"meta"`
}
