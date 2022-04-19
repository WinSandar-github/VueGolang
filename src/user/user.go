package userdata

import "fmt"

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email" mysql:"unique",db:"email"`
	Address  string `json:"address"`
	Password string `json:"password",db:"password"`
}

func (user User) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\nemail: %s\npassword: %s\nphone: %s\naddress: %s",
		user.Id, user.Name, user.Email, user.Password, user.Phone, user.Address)
}

type Product struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func (product Product) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\nprice: %s",
		product.Id, product.Name, product.Price)
}

type JWTToken struct {
	Token string `json:"token"`
}
