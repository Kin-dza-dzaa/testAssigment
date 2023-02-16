package dto

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Salt struct {
	Salt string `json:"salt"`
}

type UserDb struct {
	Email    string `json:"email" bson:"email"`
	Salt     string `json:"salt" bson:"salt"`
	Password string `json:"password" bson:"password"`
}
