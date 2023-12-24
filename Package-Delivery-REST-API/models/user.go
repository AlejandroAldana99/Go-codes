package models

type UserData struct {
	UserID   MyObjectID `json:"userid" bson:"_id,omitempty"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Role     string     `json:"role"`
}
