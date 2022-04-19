package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"backend-golang/src/config"
	"backend-golang/src/models"
	userdata "backend-golang/src/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = "mysecret"

func GetProduct(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		products, err2 := userModel.GetProduct()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}
func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user userdata.User
	err := json.NewDecoder(request.Body).Decode(&user)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		err2 := userModel.CreateUser(&user)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, user)
		}
	}
}
func CreateProduct(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var product userdata.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		err2 := userModel.CreateProduct(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}
func UpdateProduct(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var product userdata.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		_, err2 := userModel.UpdateProduct(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}
func DeleteProduct(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		_, err2 := userModel.DeleteProduct(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, nil)
		}
	}
}

func Login(response http.ResponseWriter, request *http.Request) {

	var user userdata.User
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		}
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)

		if userModel.Db.QueryRow("select id from users where email=?", keyVal["email"]).Scan(&user.Id) != nil {
			respondWithError(response, http.StatusBadRequest, "Please Check Email Again")
		}
		exist := &userdata.User{}
		if userModel.Db.QueryRow("select password from users where email=?", keyVal["email"]).Scan(&exist.Password) == nil {
			err3 := bcrypt.CompareHashAndPassword([]byte(exist.Password), []byte(keyVal["password"]))
			if err3 != nil {
				respondWithError(response, http.StatusUnauthorized, "incorrect password")
			} else {

				claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
					Issuer:    strconv.Itoa(int(user.Id)),
					ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				})
				tokenString, err2 := claims.SignedString([]byte(secretKey))

				if err2 != nil {
					respondWithError(response, http.StatusBadRequest, "could not login")
				} else {
					respondWithJson(response, http.StatusOK, userdata.JWTToken{Token: tokenString})

				}
			}

		}

	}

}

func GetUser(response http.ResponseWriter, request *http.Request) {

	cookies := request.Header.Get("key")

	token, err1 := jwt.ParseWithClaims(cookies, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err1 != nil {
		respondWithError(response, http.StatusUnauthorized, "unauthorized")
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user userdata.User
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "Please try again")
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		userModel.Db.QueryRow("select name from users where id=?", claims.Issuer).Scan(&user.Name)
		respondWithJson(response, http.StatusOK, user.Name)
	}

}

func respondWithError(w http.ResponseWriter, code int, msg string) {

	respondWithJson(w, code, map[string]string{"error": msg})

}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
