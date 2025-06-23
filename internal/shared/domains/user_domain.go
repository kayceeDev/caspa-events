package domains

import "github.com/kayceeDev/caspa-events/ent-go/ent"

type User struct {
	ID       uint
	Username string
	Email    string
	Password string
}

type UserResponse struct {
	*ent.User
	Edges interface{} `json:"edges,omitempty"`
}

func (u *UserResponse) GetUUID() string {
	if u == nil || u.User == nil {
		return ""
	}
	return u.User.UUID.String()
}

func ParseUser(user *ent.User) *UserResponse {
	if user == nil {
		return nil
	}
	return &UserResponse{
		User:  user,
		Edges: nil,
	}
}

func ParseUsers(users []*ent.User) []*UserResponse {
	if users == nil {
		return nil
	}
	userResponses := make([]*UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = ParseUser(user)
	}
	return userResponses
}
