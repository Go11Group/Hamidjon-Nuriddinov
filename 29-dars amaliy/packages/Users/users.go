package users

import (
	"database/sql"
)

type User struct {
	Id        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
	Gender    string
	Age       int
}

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (U *UserRepo) Insert(user User) error {
	_, err := U.Db.Exec(`Insert Into Userss values($1, $2, $3, $4, $5, $6)`, user.Id, user.FirstName, user.LastName, user.Email, user.Gender, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func (U *UserRepo) ReadAll() ([]User, error) {
	users := []User{}
	rows, err := U.Db.Query(`Select * from Userss`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Gender, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (U *UserRepo) GetByID(id string)(User, error){
	user := User{}
	err := U.Db.QueryRow(`Select * from Userss where id = $1`, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Gender, &user.Age)
	if err != nil{
		return User{}, err
	}
	return user, nil
}

func(U *UserRepo) Update(user User)error{
	_, err := U.Db.Exec(`Update Userss 
	Set first_name = $1, last_name = $2, email = $3, gender = $4, age = $5 
	where id = $6`, user.FirstName, user.LastName, user.Email, user.Gender, user.Age, user.Id)
	if err != nil{
		return err
	}
	return nil
}

func (U *UserRepo) Delete(id string)error{
	_, err := U.Db.Exec(`Delete from Userss where id = $1`, id)
	if err != nil{
		return err
	}
	return nil
}