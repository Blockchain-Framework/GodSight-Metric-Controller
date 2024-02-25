package v1

import (
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/config"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/modules/api/v1/controllers"
	"github.com/Blockchain-Framework/controller/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {

	router := chi.NewRouter()

	//init data
	router.Get("/init", controllers.Init)

	// Organization
	router.Get("/organizations", controllers.ListOrganizations)
	router.Post("/organizations", controllers.CreateOrganization)
	router.With(middleware.ParamValidate(config.Conf.Validator, controllers.OrganizationIdRule)).Put("/organizations/{organization-id}", controllers.UpdateOrganization)
	router.With(middleware.ParamValidate(config.Conf.Validator, controllers.OrganizationIdRule)).Get("/organizations/{organization-id}", controllers.OrganizationInfo)
	router.With(middleware.ParamValidate(config.Conf.Validator, controllers.OrganizationIdRule)).Post("/organizations/{organization-id}", controllers.ActivateOrganization)
	router.With(middleware.ParamValidate(config.Conf.Validator, controllers.OrganizationIdRule)).Delete("/organizations/{organization-id}", controllers.DeleteOrganization)

	//Users
	router.Get("/users", controllers.ListUsers)
	router.Post("/users", controllers.CreateUser)
	router.With(middleware.ParamValidate(config.Conf.Validator, controllers.UserIdRule)).Put("/users/{user-id}", controllers.UpdateUser)
	router.With(middleware.ParamValidate(config.Conf.Validator, controllers.UserIdRule)).Get("/users/{user-id}", controllers.UserInfo)
	router.With(middleware.ParamValidate(config.Conf.Validator, controllers.UserIdRule)).Post("/users/{user-id}", controllers.ActivateUser)
	router.With(middleware.ParamValidate(config.Conf.Validator, controllers.UserIdRule)).Delete("/users/{user-id}", controllers.DeleteUser)

	return router
}
