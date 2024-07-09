package data

import (
	"api/data/models"
	"database/sql"
	"log"
)

// SetupDB Method for setting up the tables
func SetupDB(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS manufacturers (" +
		"ID SERIAL PRIMARY KEY," +
		"name VARCHAR(255) NOT NULL," +
		"phone VARCHAR(255) NOT NULL UNIQUE);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (" +
		"ID SERIAL PRIMARY KEY," +
		"name VARCHAR(255) NOT NULL," +
		"price FLOAT NOT NULL," +
		"size VARCHAR(255)," +
		"color VARCHAR(255));")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS customers (" +
		"ID SERIAL PRIMARY KEY," +
		"firstName VARCHAR(255) NOT NULL," +
		"lastName VARCHAR(255) NOT NULL," +
		"phoneNumber VARCHAR(255) NOT NULL," +
		"email VARCHAR(255) NOT NULL," +
		"street VARCHAR(255)," +
		"city VARCHAR(255)," +
		"country VARCHAR(255));")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS bikes (" +
		"productID INT references products(id) NOT NULL," +
		"frameNumber VARCHAR(255) PRIMARY KEY," +
		"owner INT references customers(id));")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS productsmanufacturers (" +
		"productID INT references products(id) NOT NULL," +
		"manufacturerID INT references manufacturers(id) NOT NULL," +
		"PRIMARY KEY (productID, manufacturerID));")
	if err != nil {
		log.Fatal(err)
	}
}

// Methods for CRUD operations on the products table
func CreateProduct(db *sql.DB, product *models.Product) (int, error) {
	row := db.QueryRow("INSERT INTO products(name, price, size, color) VALUES ($1, $2, $3, $4) RETURNING id", product.Name, product.Price, product.Size, product.Color)
	if row.Err() != nil {
		return -1, row.Err()
	}

	var id int
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetProduct(db *sql.DB, id int) (models.Product, error) {
	var product models.Product
	row := db.QueryRow("SELECT * FROM products WHERE id = $1", id)
	if row.Err() != nil {
		return product, row.Err()
	}
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Size, &product.Color)
	if err != nil {
		return product, err
	}
	return product, nil
}

func GetProducts(db *sql.DB) ([]models.Product, error) {
	var products []models.Product
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return products, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Size, &product.Color)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}

func UpdateProduct(db *sql.DB, product models.Product) error {
	err := db.QueryRow("UPDATE products "+
		"SET name = $1, price = $2, size = $3, color = $4 "+
		"WHERE id = $5", product.Name, product.Price, product.Size, product.Color, product.Id)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}

func DeleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// Methods for performing CRUD on customer table
func CreateCustomer(db *sql.DB, customer models.Customer) (int, error) {
	row := db.QueryRow("INSERT INTO customers(firstname, lastname, phonenumber, email, street, city, country) VALUES ("+
		"$1, $2, $3, $4, $5, $6, $7) RETURNING id;", customer.FirstName, customer.LastName, customer.Phone, customer.Email, customer.Address.Street, customer.Address.City, customer.Address.Country)
	if row.Err() != nil {
		return -1, row.Err()
	}
	var id int
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetCustomer(db *sql.DB, id int) (models.Customer, error) {
	var customer models.Customer
	var addressStreet sql.NullString
	var addressCity sql.NullString
	var addressCountry sql.NullString

	row := db.QueryRow("SELECT * FROM customers WHERE id = $1", id)
	if row.Err() != nil {
		return customer, row.Err()
	}

	err := row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Phone, &customer.Email, &addressStreet, &addressCity, &addressCountry)
	if err != nil {
		return customer, err
	}

	customer.Address.Street = addressStreet.String
	customer.Address.City = addressCity.String
	customer.Address.Country = addressCountry.String

	return customer, nil
}

func GetCustomers(db *sql.DB) ([]models.Customer, error) {
	var customers []models.Customer
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		return customers, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var customer models.Customer
		var addressStreet sql.NullString
		var addressCity sql.NullString
		var addressCountry sql.NullString
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Phone, &customer.Email, &addressStreet, &addressCity, &addressCountry)
		if err != nil {
			return customers, err
		}
		customer.Address.Street = addressStreet.String
		customer.Address.City = addressCity.String
		customer.Address.Country = addressCountry.String
		customers = append(customers, customer)
	}
	return customers, nil
}

func UpdateCustomer(db *sql.DB, customer models.Customer) error {
	err := db.QueryRow("UPDATE customers "+
		"SET firstname = $1, lastname = $2, phonenumber = $3, email = $4, street = $5, city = $6, country = $7"+
		"WHERE id = $8", customer.FirstName, customer.LastName, customer.Phone, customer.Email, customer.Address.Street, customer.Address.City, customer.Address.Country, customer.Id)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}

func DeleteCustomer(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// Methods for performing CRUD on customer table
func CreateManufacturer(db *sql.DB, manufacturer models.Manufacturer) (int, error) {
	row := db.QueryRow("INSERT INTO manufacturers (name, phone) VALUES ($1, $2) RETURNING id;", manufacturer.Name, manufacturer.Phone)
	if row.Err() != nil {
		return -1, row.Err()
	}

	var id int
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetManufacturer(db *sql.DB, id int) (models.Manufacturer, error) {
	var manufacturer models.Manufacturer
	row := db.QueryRow("SELECT * FROM manufacturers WHERE id = $1", id)
	if row.Err() != nil {
		return manufacturer, row.Err()
	}
	err := row.Scan(&manufacturer.Id, &manufacturer.Name, &manufacturer.Phone)
	if err != nil {
		return manufacturer, err
	}
	return manufacturer, nil
}

func GetManufacturers(db *sql.DB) ([]models.Manufacturer, error) {
	var manufacturers []models.Manufacturer
	rows, err := db.Query("SELECT * FROM manufacturers")
	if err != nil {
		return manufacturers, err
	}
	for rows.Next() {
		var manufacturer models.Manufacturer
		err := rows.Scan(&manufacturer.Id, &manufacturer.Name, &manufacturer.Phone)
		if err != nil {
			return manufacturers, err
		}
		manufacturers = append(manufacturers, manufacturer)
	}
	return manufacturers, nil
}

func UpdateManufacturer(db *sql.DB, manufacturer models.Manufacturer) error {
	err := db.QueryRow("UPDATE manufacturers "+
		"SET name = $1, phone = $2 WHERE id = $3;", manufacturer.Name, manufacturer.Phone, manufacturer.Id)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}

func DeleteManufacturer(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM manufacturers WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func AssociateManufacturers(db *sql.DB, id int, manufacturers []int) error {
	for _, manufacturer := range manufacturers {
		_, err := db.Exec("INSERT INTO productsmanufacturers (productid, manufacturerid) VALUES ($1, $2)", id, manufacturer)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteAssociationManufacturers(db *sql.DB, id int, manufacturers []int) error {
	for _, manufacturer := range manufacturers {
		_, err := db.Exec("DELETE FROM productsmanufacturers WHERE productid = $1 AND manufacturerid = $2", id, manufacturer)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateBike(db *sql.DB, bike models.Bike) (string, error) {
	row := db.QueryRow("INSERT INTO bikes (productid, framenumber) VALUES ($1, $2) RETURNING framenumber;", bike.Id, bike.FrameNumber)
	if row.Err() != nil {
		return "", row.Err()
	}
	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
