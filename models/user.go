package models

type User struct {
	Id int
	Username string
	Balance	int64
}

func (u *User) Register() (error) {
	err := DB.Create(&u)
	return err.Error
}

func (u *User) Update() (error) {
	err := DB.Save(&u)
	return err.Error
}