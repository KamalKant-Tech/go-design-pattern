package users

/* Single responsibilty principle */
/* This class follow the single responsibilty principle because this does not have the function/method which do other than user work */

type IUsers interface {
	GetUserProfile() *Users
	SetUserProfile(name string, id int, addr string)
	GetUserName() string
}

type Users struct {
	IUsers
	name    string
	id      int
	address string
}

func InitializeUsers() *Users {
	user := &Users{}
	return user
}

func (u *Users) GetUserName() string {
	return u.name
}

func (u *Users) SetUserProfile(name string, id int, addr string) {
	u.address = addr
	u.name = name
	u.id = id
}
