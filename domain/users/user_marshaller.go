package users

import "encoding/json"

type PublicUser struct {
	Id int64 `json:"id"`
	// FirstName   string `json:"first_name"` // `json:"first_name" binding:"required"`
	// LastName    string `json:"last_name"`  // `json:"last_name" binding:"required"`
	// Email       string `json:"email"`      // `json:"email" binding:"required,email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"` //`json:"-"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"` // `json:"first_name" binding:"required"`
	LastName    string `json:"last_name"`  // `json:"last_name" binding:"required"`
	Email       string `json:"email"`      // `json:"email" binding:"required,email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"` //`json:"-"`
}

// if json tag fields have same names - approach like this
func (user *User) Marshall(isPublic bool) interface{} {
	userJson, _ := json.Marshal(user)

	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJson, &publicUser)
		return publicUser
	}

	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}

// if json tag fields have different names this approach
func (user *User) Marshall2(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	return PrivateUser{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateCreated: user.DateCreated,
		Status:      user.Status,
	}
}
