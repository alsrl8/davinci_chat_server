package types

import "errors"

type NewUserRequest struct {
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
	Password  string `json:"password"`
}

type ValidateUserRequest struct {
	UserEmail string `json:"userEmail"`
}

func (req *ValidateUserRequest) Validate() error {
	if req.UserEmail == "" {
		return errors.New("userEmail is required")
	}
	return nil
}

type LoginRequest struct {
	UserEmail string `json:"userEmail"`
	Password  string `json:"password"`
}

type SendInvitationRequest struct {
	UserName string `json:"userName"`
}
