package railroad

import (
	"fmt"
	"strings"
)

type UserService struct {
	db map[string]UserRegistrationRequest
}

func NewUserService() UserService {
	return UserService{
		db: map[string]UserRegistrationRequest{},
	}
}

func (us *UserService) Register(req UserRegistrationRequest) bool {
	return Result{State: req}.
		Apply(us.ValidateRequest).
		Apply(us.CheckUserExists).
		Apply(us.SaveUserInDatabase).
		Success
}

func (us *UserService) ValidateRequest(res Result) Result {
	req, ok := res.State.(UserRegistrationRequest)
	if !ok {
		return Result{
			State:   res.State,
			Success: false,
			Failure: true,
			Err:     fmt.Errorf("state passed as parameter was not of type UserRegistrationRequest"),
		}
	}

	username := req.Username
	email := req.Email

	if len(username) > 10 {
		return Result{
			State:   nil,
			Success: false,
			Failure: true,
			Err:     fmt.Errorf("Username %v has a length of %v which is less than 10", username, len(username)),
		}
	}

	if !strings.Contains(email, "@") {
		return Result{
			State:   nil,
			Success: false,
			Failure: true,
			Err:     fmt.Errorf("Email %v must contain an @ sign ", email),
		}
	}

	return Result{
		State:   req,
		Success: true,
		Failure: false,
	}
}

func (us *UserService) CheckUserExists(res Result) Result {
	req, ok := res.State.(UserRegistrationRequest)
	if !ok {
		return Result{
			State:   res.State,
			Success: false,
			Failure: true,
			Err:     fmt.Errorf("state passed as parameter was not of type UserRegistrationRequest"),
		}
	}

	username := strings.ToLower(req.Username)

	if strings.HasPrefix(username, "c") {
		return Result{
			State:   nil,
			Success: false,
			Failure: true,
			Err:     fmt.Errorf("User with username '%v' already exists", req.Username),
		}
	}

	return Result{
		State:   req,
		Success: true,
		Failure: false,
	}
}

func (us *UserService) SaveUserInDatabase(res Result) Result {
	req, ok := res.State.(UserRegistrationRequest)
	if !ok {
		return Result{
			State:   res.State,
			Success: false,
			Failure: true,
			Err:     fmt.Errorf("state passed as parameter was not of type UserRegistrationRequest"),
		}
	}

	us.db[req.Username] = req

	return Result{
		State:   req,
		Success: true,
		Failure: false,
	}
}
