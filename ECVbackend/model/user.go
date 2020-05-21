package model

import "ecvbackend/pkg/util"

type User struct {
	Username string `xorm:"username"`
	Password string `xorm:"password"`
	Email string `xorm:"email"`
	Md5 string `xorm:"md5"`
	Nickname string `xorm:"nickname"`
	Organization string `xorm:"organization"`
	Birth string `xorm:"birth"`
	Description string `xorm:"description"`
	Root bool `xorm:"root"`
}

func (u User) Get() *User {
	exist, _ := dbEngine.Get(&u)
	if exist {
		return &u
	}
	return nil
}

func (u User) Delete()  {
	dbEngine.Delete(&u)
}

func (u User) FindAll() []User {
	everyone := make([]User, 0)
	err := dbEngine.Find(&everyone)
	if err==nil {
		return everyone 
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

func (u User) Update() {
	_, err := dbEngine.Where("username = ?",u.Username).AllCols().Update(&u)
	if err == nil {
		util.RuntimeLog.Println("update success")
	} else {
		util.RuntimeLog.Println("update error",err)
	}
}