package event

type UserInsertEvent struct {
	FullName    string
	Identity    string
	PhoneNumber string
	Email       string
	Gender      string
}
