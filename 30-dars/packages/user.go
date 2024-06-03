package packages

import "database/sql"


type User struct{
	Id int
	Username string
	Email string
	Password string
}

type UserRepo struct{
	Db *sql.DB
}

func NewUserRepo(db *sql.DB)*UserRepo{
	return &UserRepo{Db: db}
}

func(u *UserRepo) Create(user User)error{
	tr, err := u.Db.Begin()
	if err != nil{
		return err
	}
	defer tr.Commit()

	_, err = u.Db.Exec("Insert into Users(username, email, password) values($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil{
		return err
	}

	return nil
}


func(u *UserRepo) ReadAll()([]User, error){
	tr, err := u.Db.Begin()
	if err != nil{
		return nil, err
	}
	defer tr.Commit()

	rows, err := u.Db.Query(`select id, username, email from Users`)
	if err != nil{
		return nil, err 
	}
	users := []User{}
	for rows.Next(){
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil{
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func(U *UserRepo) ReadUser(id int)(User, error){
	tr, err := U.Db.Begin()
	if err != nil{
		return User{}, err
	}
	defer tr.Commit()
	var user User
	err = U.Db.QueryRow("Select id, username, email from Users where id = $1", id).Scan(&user.Id, &user.Username, &user.Email)
	if err != nil{
		return User{}, nil
	}
	return user, nil
}

func(U *UserRepo) Update(user User)error{
	tr, err := U.Db.Begin()
	if err != nil{
		return err
	}
	defer tr.Commit()

	_, err = U.Db.Exec(`Update Users Set username = $1, email = $2, password = $3 where id = $4`, 
	user.Username, user.Email, user.Password, user.Id)
	if err != nil{
		return err
	}
	return nil
}


func(U *UserRepo) Delete(id int)error{
	tr, err := U.Db.Begin()
	if err != nil{
		return err
	}
	defer tr.Commit()
	
	_, err = U.Db.Exec("Delete from User_products where user_id = id")
	if err != nil{
		return err
	}
	_, err = U.Db.Exec(`Delete from Users where id = $1`, id)
	if err != nil{
		return err
	}
	return nil
}
