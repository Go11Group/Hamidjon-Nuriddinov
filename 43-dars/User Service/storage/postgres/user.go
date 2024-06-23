package postgres

import (
	"database/sql"
	"mymode/model"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (U *UserRepo) CreateUser(user model.CreateUser) error {
	_, err := U.Db.Exec(`INSERT INTO Users(name, phone, age) VALUES($1, $2, $3)`, user.Name, user.Phone, user.Age)
	return err
}

func (U *UserRepo) GetByIdUser(id string) (model.CreateUser, error) {
	user := model.CreateUser{}
	err := U.Db.QueryRow(`SELECT * FROM Users WHERE user_id = $1`, id).Scan(&user.UserId, &user.Name, &user.Phone, &user.Age)
	return user, err
}

func (U *UserRepo) GetAllUsers() ([]model.CreateUser, error) {
	users := []model.CreateUser{}
	rows, err := U.Db.Query(`SELECT * FROM Users`)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		user := model.CreateUser{}
		err = rows.Scan(&user.UserId, &user.Name, &user.Phone, &user.Age)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (U *UserRepo) UpdateUser(id string, user model.CreateUser) error {
	_, err := U.Db.Exec(`UPDATE Users SET name = $1, phone = $2, age = $3 WHERE user_id = $4`, user.Name, user.Phone, user.Age, id)
	return err
}

func (U *UserRepo) DeleteUser(id string) error {
	_, err := U.Db.Exec(`DELETE FROM Users WHERE user_id = $1`, id)
	return err
}
