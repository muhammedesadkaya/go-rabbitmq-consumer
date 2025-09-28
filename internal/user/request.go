package user

type UserRequest struct {
	FullName    string `json:"fullName"`
	Identity    string `json:"identity"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
}
