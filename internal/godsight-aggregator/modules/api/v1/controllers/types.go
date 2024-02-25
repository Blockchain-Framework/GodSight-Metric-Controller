package controllers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type organizationId struct {
	Id string `json:"organization_id"`
}

func OrganizationIdRule(r *http.Request) interface{} {
	return organizationId{
		Id: chi.URLParam(r, "organization-id"),
	}
}

type userId struct {
	Id string `json:"user_id"`
}

func UserIdRule(r *http.Request) interface{} {
	return userId{
		Id: chi.URLParam(r, "user-id"),
	}
}
