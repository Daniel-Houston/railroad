package railroad

import "testing"

func TestRegister(t *testing.T) {
	tests := []struct {
		name string
		req  UserRegistrationRequest
		res  bool
	}{
		{
			name: "valid request",
			req: UserRegistrationRequest{
				Username: "dhouston",
				Email:    "dan@daniel.me",
			},
			res: true,
		},
		{
			name: "username too long",
			req: UserRegistrationRequest{
				Username: "dhoustondhouston",
				Email:    "dan@daniel.me",
			},
			res: false,
		},
		{
			name: "email wrong format",
			req: UserRegistrationRequest{
				Username: "dhouston",
				Email:    "dan",
			},
			res: false,
		},
		{
			name: "user already exists",
			req: UserRegistrationRequest{
				Username: "chouston",
				Email:    "cam@daniel.me",
			},
			res: false,
		},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			svc := UserService{
				db: map[string]UserRegistrationRequest{},
			}
			res := svc.Register(test.req)

			if want, got := test.res, res; want != got {
				t.Errorf("Wrong result. Want %v got %v", want, got)
			}
		})
	}
}
