package packages

import "database/sql"


type Product struct{
	Id int
	Name string
	Description string
	Price float64
	Stock_quantity int
}

type ProductRepo struct{
	Db *sql.DB
}

func NewProductRepo(db *sql.DB)*ProductRepo{
	return &ProductRepo{Db: db}
}

func(P *ProductRepo) Create(product Product)error{
	tr, err := P.Db.Begin()
	if err != nil{
		return err
	}
	defer tr.Commit()

	_, err = P.Db.Exec("Insert into Products(name, description, price, stock_quantity) values ($1, $2, $3, $4)", product.Name, product.Description, product.Price, product.Stock_quantity)
	if err != nil{
		return err
	}

	return nil
}

func(P *ProductRepo) ReadAll()([]Product, error){
	tr, err := P.Db.Begin()
	if err != nil{
		return nil, err
	}
	defer tr.Commit()

	rows, err := P.Db.Query(`select * from Products`)
	if err != nil{
		return nil, err 
	}
	products := []Product{}
	for rows.Next(){
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock_quantity)
		if err != nil{
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func(P *ProductRepo) ReadUser(id int)(Product, error){
	tr, err := P.Db.Begin()
	if err != nil{
		return Product{}, err
	}
	defer tr.Commit()
	var product Product
	err = P.Db.QueryRow("Select * from Products where id = $1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock_quantity)
	if err != nil{
		return Product{}, nil
	}
	return product, nil
}

func(U *ProductRepo) Update(product Product)error{
	tr, err := U.Db.Begin()
	if err != nil{
		return err
	}
	defer tr.Commit()

	_, err = U.Db.Exec(`Update Products Set name = $1, description = $2, price = $3, stock_quantity = $4 where id = $5`, 
	product.Name, product.Description, product.Price, product.Stock_quantity, product.Id)
	if err != nil{
		return err
	}
	return nil
}


func(P *ProductRepo) Delete(id int)error{
	tr, err := P.Db.Begin()
	if err != nil{
		return err
	}
	defer tr.Commit()

	_, err = P.Db.Exec("Delete from User_products where product_id = $1", id)
	if err != nil{
		return err
	}
	_, err = P.Db.Exec(`Delete from Products where id = $1`, id)
	if err != nil{
		return err
	}
	return nil
}
