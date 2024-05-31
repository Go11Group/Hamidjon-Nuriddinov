package packages

import "gorm.io/gorm"

type User1 struct {
	gorm.Model
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Age        int
	Field      string
	Gender     string
	IsEmployee bool
}

type UserRepo struct{
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB)*UserRepo{
	return &UserRepo{Db: db}
}


func(U *UserRepo) CreateTable()error{
	err := U.Db.AutoMigrate(&User1{})
	if err != nil{
		return err
	}
	return nil
}

func(U *UserRepo) Create(user User1){
	U.Db.Create(&user)
}

func(U *UserRepo) Read()([]User1){
	var users = []User1{}
	U.Db.Find(&users)
	return users
}

func(U *UserRepo) Update(user User1, id int){
	U.Db.Model(&user).Updates(User1{FirstName: user.FirstName})
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
}

func(U *UserRepo) Delete(user User1){
	U.Db.Delete(&user, 1)
}

