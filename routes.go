package main

import (
	"api/data"
	"api/data/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func addRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", indexHandler)

	mux.HandleFunc("POST /products", createProductHandler)
	mux.HandleFunc("GET /products/{id}", getProductHandler)
	mux.HandleFunc("GET /products", getProductsHandler)
	mux.HandleFunc("PUT /products", updateProductHandler)
	mux.HandleFunc("DELETE /products/{id}", deleteProductHandler)

	mux.HandleFunc("POST /customers", createCustomerHandler)
	mux.HandleFunc("GET /customers/{id}", getCustomerHandler)
	mux.HandleFunc("GET /customers", getCustomersHandler)
	mux.HandleFunc("PUT /customers", updateCustomerHandler)
	mux.HandleFunc("DELETE /customers/{id}", deleteCustomerHandler)

	mux.HandleFunc("POST /manufacturers", createManufacturerHandler)
	mux.HandleFunc("GET /manufacturers/{id}", getManufacturerHandler)
	mux.HandleFunc("GET /manufacturers", getManufacturersHandler)
	mux.HandleFunc("PUT /manufacturers", updateManufacturerHandler)
	mux.HandleFunc("DELETE /manufacturers/{id}", deleteManufacturerHandler)

	mux.HandleFunc("POST /products/{id}", associateManufacturersHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

// Functions for manipulating products
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := data.CreateProduct(db, &product)
	if err != nil || id == -1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Product created successfully - Product Id: %d", id)))
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := data.GetProduct(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := data.GetProducts(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := data.UpdateProduct(db, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Product updated successfully")))
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = data.DeleteProduct(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Product deleted successfully - Product Id: %d", id)))
}

// Functions for manipulating customers
func createCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := data.CreateCustomer(db, customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Customer created successfully - Customer Id: %d", id)))
}

func getCustomerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := data.GetCustomer(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getCustomersHandler(w http.ResponseWriter, r *http.Request) {
	customers, err := data.GetCustomers(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(customers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func updateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := data.UpdateCustomer(db, customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Customer updated successfully")))
}

func deleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = data.DeleteCustomer(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Customer deleted successfully - Customer Id: %d", id)))
}

func createManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	var manufacturer models.Manufacturer
	if err := json.NewDecoder(r.Body).Decode(&manufacturer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := data.CreateManufacturer(db, manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Manufacturer successfully created - Manufacturer Id: %d", id)))
}

func getManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	manufacturer, err := data.GetManufacturer(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	j, err := json.Marshal(manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getManufacturersHandler(w http.ResponseWriter, r *http.Request) {
	manufacturers, err := data.GetManufacturers(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(manufacturers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func updateManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	var manufacturer models.Manufacturer
	if err := json.NewDecoder(r.Body).Decode(&manufacturer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := data.UpdateManufacturer(db, manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Manufacturer updated successfully")))
}

func deleteManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = data.DeleteManufacturer(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Manufacturer deleted successfully")))
}

func associateManufacturersHandler(w http.ResponseWriter, r *http.Request) {
	var manufacturers []int
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&manufacturers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = data.AssociateManufacturers(db, id, manufacturers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Manufacturer associated successfully")))
}
