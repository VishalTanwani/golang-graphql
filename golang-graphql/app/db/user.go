package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//User is a structure
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

//UserPrint is a structure
type UserPrint struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

var users []User
var printUsers []UserPrint

//StartAddDataForUsers will get data from file add into db
func StartAddDataForUsers() {
	data, err := ioutil.ReadFile("../dump_data/users.json")
	checkErr(err)
	err = json.Unmarshal(data, &users)
	checkErr(err)
	fmt.Println("data extracted from file")
	fmt.Println("# Inserting values")
	var lastInsertID int
	for _, value := range users {
		err := DB.QueryRow("INSERT INTO users(name,password,email,role) VALUES($1,$2,$3,$4) returning id;", value.Name, value.Password, value.Email, value.Role).Scan(&lastInsertID)
		if err != nil {
			panic(err)
		}
		fmt.Println("last inserted id =", lastInsertID)
	}

}

//LoadAllUsers retunr all users from db
func LoadAllUsers() ([]UserPrint, error) {
	row, err := DB.Query("SELECT id,name,password,email,role FROM users")
	checkErr(err)
	data := []UserPrint{}
	for row.Next() {
		var u UserPrint
		row.Scan(&u.ID, &u.Name, &u.Password, &u.Email, &u.Role)
		// fmt.Println(u)
		data = append(data, u)
	}
	defer row.Close()
	return data, nil
}

//GetUserByID return user by id
func GetUserByID(id int) *UserPrint {
	row, err := DB.Query("SELECT id,name,password,email,role FROM users WHERE id = $1", id)
	checkErr(err)
	var u UserPrint
	for row.Next() {
		row.Scan(&u.ID, &u.Name, &u.Password, &u.Email, &u.Role)
	}
	defer row.Close()
	return &u
}

//CreateUser create user
func CreateUser(name, password, email, role string) (*UserPrint, error) {
	var lastInsertID int
	err := DB.QueryRow("INSERT INTO users(name,password,email,role) VALUES($1,$2,$3,$4) returning id;", name, password, email, role).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}
	fmt.Println("last inserted id =", lastInsertID)

	u := UserPrint{
		ID:       lastInsertID,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
	return &u, nil

}
