package user


type User struct {
	Username string
	Sex string
	Age int
	AvatarUrl string
}

//构造函数
func NewUser(username string, sex string, age int, avatar string) *User {
	user := &User{
		Username:"zhanghao",
		Age:18,
		Sex:"mail",
		AvatarUrl:"http://www.baodu.com",
	}
	return user

	/* 方法2
	user := new(User)
	user.Username = "ljf"
	user.Age = 18
	user.Sex = "male"
	user.AvatarUrl = "www.baidu.com"
	 */
}
