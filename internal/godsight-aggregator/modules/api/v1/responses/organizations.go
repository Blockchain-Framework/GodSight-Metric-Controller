package responses

import (
	"context"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/models"
	log "github.com/Blockchain-Framework/controller/pkg/logger"
	"github.com/go-playground/validator/v10"
)

func ListOrganizationsResponseBody(ctx context.Context, organizationsData []models.Organization) ([]OrganizationResponse, error) {

	log.Debug(ctx).Msg("Create Response for List Organization")

	organizations := make([]OrganizationResponse, 0)
	for _, item := range organizationsData {

		organization := OrganizationResponse{
			OrganizationId:   item.OrganizationId,
			OrganizationName: item.OrganizationName,
			Type:             item.Type,
			State:            item.State,
			AddressLine1:     item.AddressLine1,
			AddressLine2:     item.AddressLine2,
			Country:          item.Country,
			City:             item.City,
			Zip:              item.Zip,
			ContactPerson:    item.ContactPerson,
			ContactDetails:   item.ContactDetails,
			Website:          item.Website,
			OnboardedAt:      item.OnboardTimestamp,
			ActivatedAt:      item.ActivatedTimestamp,
			ActiveUserCount:  item.ActiveUserCount,
			UserCount:        item.UserCount,
			Comments:         item.Comments,
		}

		validate := validator.New()

		err := validate.Struct(organization)
		if err != nil {
			return nil, err
		}

		organizations = append(organizations, organization)

	}

	return organizations, nil
}

func CreateOrganizationResponse(ctx context.Context) (ResponseData, error) {

	log.Debug(ctx).Msg("Create Response for Create Organization")

	responseData := ResponseData{
		Message: "successfully created new organization",
		State:   0,
	}

	return responseData, nil
}

func UpdateOrganizationResponse(ctx context.Context) (ResponseData, error) {

	log.Debug(ctx).Msg("Create Response for Update Organization")

	responseData := ResponseData{
		Message: "successfully updated organization",
		State:   0,
	}

	return responseData, nil
}

func OrganizationInfoResponseBody(ctx context.Context, organizationData models.Organization) (OrganizationResponse, error) {

	log.Debug(ctx).Msg("Create Response for List Organization")

	organization := OrganizationResponse{
		OrganizationId:   organizationData.OrganizationId,
		OrganizationName: organizationData.OrganizationName,
		Type:             organizationData.Type,
		State:            organizationData.State,
		AddressLine1:     organizationData.AddressLine1,
		AddressLine2:     organizationData.AddressLine2,
		Country:          organizationData.Country,
		City:             organizationData.City,
		Zip:              organizationData.Zip,
		ContactPerson:    organizationData.ContactPerson,
		ContactDetails:   organizationData.ContactDetails,
		Website:          organizationData.Website,
		OnboardedAt:      organizationData.OnboardTimestamp,
		ActivatedAt:      organizationData.ActivatedTimestamp,
		ActiveUserCount:  organizationData.ActiveUserCount,
		UserCount:        organizationData.UserCount,
		Comments:         organizationData.Comments,
	}

	validate := validator.New()

	err := validate.Struct(organization)
	if err != nil {
		return OrganizationResponse{}, err
	}

	return organization, nil
}

func ActivateOrganizationResponse(ctx context.Context) (ResponseData, error) {

	log.Debug(ctx).Msg("Create Response for Activate Organization")

	responseData := ResponseData{
		Message: "successfully activate organization",
		State:   0,
	}

	return responseData, nil
}

func DeleteOrganizationResponse(ctx context.Context) (ResponseData, error) {

	log.Debug(ctx).Msg("Create Response for Delete Organization")

	responseData := ResponseData{
		Message: "successfully delete organization",
		State:   0,
	}

	return responseData, nil
}
