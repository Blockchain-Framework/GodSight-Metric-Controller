package responses

import (
	"context"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/models"
	log "github.com/Blockchain-Framework/controller/pkg/logger"
	"github.com/go-playground/validator/v10"
)

func ListUsersResponseBody(ctx context.Context, usersData []models.User) ([]UserResponse, error) {

	log.Debug(ctx).Msg("Create Response for List Users")

	users := make([]UserResponse, 0)
	//for _, item := range usersData {
	//
	//	user := UserResponse{
	//		OrganizationID: item.OrganizationID,
	//		UserID:         item.UserID,
	//		OtherNames:     item.OtherNames,
	//		FirstName:      item.FirstName,
	//		LastName:       item.LastName,
	//		NameInitials:   item.NameInitials,
	//		Email:          item.Email,
	//		State:          item.State,
	//		Roles:          item.Roles,
	//		OnboardedAt:    item.OnboardedAt,
	//		ActivatedAt:    item.ActivatedAt,
	//		ContactNumber:  item.ContactNumber,
	//	}
	//
	//	validate := validator.New()
	//
	//	err := validate.Struct(user)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	users = append(users, user)
	//
	//}

	return users, nil
}

func CreateUserResponse(ctx context.Context) (ResponseData, error) {

	log.Debug(ctx).Msg("Create Response for User Creation")

	responseData := ResponseData{
		Message: "successfully created new user",
		State:   0,
	}

	return responseData, nil
}

func UpdateUserResponse(ctx context.Context) (ResponseData, error) {

	log.Debug(ctx).Msg("Create Response for User Update")

	responseData := ResponseData{
		Message: "successfully updated user",
		State:   0,
	}

	return responseData, nil
}

func UserInfoResponseBody(ctx context.Context, userData models.User) (UserResponse, error) {

	log.Debug(ctx).Msg("Create Response for User Info")

	user := UserResponse{
		OrganizationID: userData.OrganizationID,
		UserID:         userData.UserID,
		OtherNames:     userData.OtherNames,
		FirstName:      userData.FirstName,
		LastName:       userData.LastName,
		NameInitials:   userData.NameInitials,
		Email:          userData.Email,
		State:          userData.State,
		Roles:          userData.Roles,
		ContactNumber:  userData.ContactNumber,
	}

	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		return UserResponse{}, err
	}

	return user, nil
}

func ActivateUserResponse(ctx context.Context) (ResponseData, error) {

	log.Debug(ctx).Msg("Create Response for User Activate")

	responseData := ResponseData{
		Message: "successfully activate user",
		State:   0,
	}

	return responseData, nil
}

func DeleteUserResponse(ctx context.Context) (ResponseData, error) {

	log.Debug(ctx).Msg("Create Response for Delete User")

	responseData := ResponseData{
		Message: "successfully delete user",
		State:   0,
	}

	return responseData, nil
}
