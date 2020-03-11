package model
import "log"
type User struct {
	Username string `xorm:"username"`
	Password string `xorm:"password"`
	Email string `xorm:"password"`
	Md5 string `xorm:"password"`
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
