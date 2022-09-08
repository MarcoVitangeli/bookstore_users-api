package users

// core domain
/**
Domain is the layer that represents an entity
from the domain that this API belongs
*/
type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
