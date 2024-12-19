package views

import (
	"github.com/elaurentium/exilium-blog-backend/src/models"
	"github.com/elaurentium/exilium-blog-backend/src/models/response"
)

func convertDomainToResponse(userDomain models.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
		//Password: userDomain.GetPassword(),
	}
}
