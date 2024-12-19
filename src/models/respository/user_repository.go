package respository

import (
	"github.com/HunCoding/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/elaurentium/exilium-blog-backend/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepositoryInterface interface {
	CreateUser(userDomain models.UserDomainInterface) (*models.UserDomainInterface, *rest_err.RestErr)
}
