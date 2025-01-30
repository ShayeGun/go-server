package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetID(i string) error {
	u.ID = i
	return nil
}

func (u *User) SetEmail(e string) error {
	u.Email = e
	return nil
}

func (u *User) SetPassword(p string) error {
	u.Password = p
	return nil
}
