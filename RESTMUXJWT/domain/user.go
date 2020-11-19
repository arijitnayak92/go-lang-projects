package domain

type User struct {
	Id       uint64 `json:"id" bson:"id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
}
