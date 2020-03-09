package model

type User struct {
	Username string `xorm:"username"`
	Password string `xorm:"password"`
}

func (u User) Get() *User {
	exist, _ := dbEngine.Get(&u)
	if exist {
		return &u
	}
	return nil
}

func (u User) Insert() *User {
	_, err := dbEngine.Insert(&u)
	if err == nil {
		return &u
	} else {
		return nil
	}
}
