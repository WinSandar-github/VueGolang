package models

import (
	"database/sql"

	userdata "backend-golang/src/user"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Db *sql.DB
}

func (userModel UserModel) GetProduct() (product []userdata.Product, err error) {
	rows, err := userModel.Db.Query("SELECT id, name,price from products")
	if err != nil {
		return nil, err
	} else {
		var products []userdata.Product
		for rows.Next() {
			var id int64
			var name string
			var price string
			err2 := rows.Scan(&id, &name, &price)
			if err2 != nil {
				return nil, err2
			} else {
				product := userdata.Product{
					Id:    id,
					Name:  name,
					Price: price,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

//Create user
func (userModel UserModel) CreateUser(user *userdata.User) (err error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	result, err := userModel.Db.Exec("insert into users(name,email,password,phone,address) VALUES(?,?,?,?,?)", user.Name, user.Email, hash, user.Phone, user.Address)
	if err != nil {
		return err
	} else {
		user.Id, _ = result.LastInsertId()
		return nil
	}
}

//Create product
func (userModel UserModel) CreateProduct(product *userdata.Product) (err error) {
	result, err := userModel.Db.Exec("insert into products(name,price) VALUES(?,?)", product.Name, product.Price)
	if err != nil {
		return err
	} else {
		product.Id, _ = result.LastInsertId()
		return nil
	}
}

//update product
func (userModel UserModel) UpdateProduct(product *userdata.Product) (int64, error) {
	result, err := userModel.Db.Exec("update products set name=?,price=? where id=?", product.Name, product.Price, product.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

//delete product
func (userModel UserModel) DeleteProduct(id int64) (int64, error) {
	result, err := userModel.Db.Exec("delete from products where id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
