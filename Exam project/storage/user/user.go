package users

import (
	"database/sql"
	"time"
)

type NewUserRepo struct {
	Db *sql.DB
}

func NewUser(db *sql.DB) *NewUserRepo {
	return &NewUserRepo{Db: db}
}


func (U *NewUserRepo) Create(user User) error {
	_, err := U.Db.Exec(`INSERT INTO 
						Users(name, email, birthday, password) 
						VALUES($1, $2, $3, $4)`, user.Name, user.Email, user.Birthday, user.Password)
	return err
}


func (U *NewUserRepo) Read(id string) (User, error){
	user := User{}
	err := U.Db.QueryRow(`SELECT 
							  name, email, birthday, created_at FROM Users 
						  WHERE 
							  user_id = $1`, id).Scan(&user.Name, &user.Email, &user.Birthday, &user.CreateAt)
	return user, err
}


func (U *NewUserRepo) Update(user User, id string) error {
	_, err := U.Db.Exec(`UPDATE Users SET 
							name = $1, email = $2, birthday = $3,  password = $4
						WHERE 
							deleted_at is null user_id = $5`,
						user.Name, user.Email, user.Birthday, user.Password, id)
	return err
}


func (U *NewUserRepo) Delete(id string) error {
	_, err := U.Db.Exec(`UPDATE Users SET 
							deleted_at = $1
						WHERE 
							deleted_at is null and user_id = $2`, time.Now(), id)
	return err
}


func (U *NewUserRepo) GetAll(limit int, offset int) ([]User, error){
	users := []User{}
	rows, err := U.Db.Query(`SELECT 
								name, email, birthday, created_at 
							FROM Users
							LIMIT $1 
							OFFSET $2`, limit, offset)
	if err != nil{
		return users, err
	}

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Name, &user.Email, &user.Birthday, &user.CreateAt)
		if err != nil{
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}
