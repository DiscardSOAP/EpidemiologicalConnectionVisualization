package model
import "log"
type User struct {
	Username string `xorm:"username"`
	Password string `xorm:"password"`
	Email string `xorm:"email"`
	Md5 string `xorm:"md5"`
	Nickname string `xorm:"nickname"`
	Organization string `xorm:"organization"`
	Birth string `xorm:"birth"`
	Description string `xorm:"description"`
}

func (u User) Get() *User {
	exist, _ := dbEngine.Get(&u)
	log.Println("Find User ",u.Username,u.Password);
	if exist {
		log.Println("User exists")
		return &u
	}
	log.Println("User unexists")
	return nil
}

func (u User) Insert() *User {
	_, err := dbEngine.Insert(&u)
	log.Println("Insert ",u.Username,u.Password);
	if err == nil {
		log.Println("insert success")
		return &u
	} else {
		log.Println("insert error")
		return nil
	}
}

func (u User) Update() {
	_, err := dbEngine.Where("username = ?",u.Username).AllCols().Update(&u)
	if err == nil {
		log.Println("update success")
	} else {
		log.Println("update error")
	}
}