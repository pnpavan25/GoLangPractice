package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"pavan/a3http/model"
)

type CustomerController struct {
	Store model.CustomerStore // It can be any Store
}

func (ctl *CustomerController) add (c model.Customer) error {
	err := ctl.Store.Create(c)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("New Customer has been created")
	return nil
}

func (ctl *CustomerController) update (cid string, c model.Customer)  error {
	err := ctl.Store.Update(cid, c)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("Customer has been updated")
	return nil
}

func (ctl *CustomerController) delete (cid string) error {
	err := ctl.Store.Delete(cid)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Printf("Customer with given ID %s is deleted\n", cid)
	return nil
}

func (ctl *CustomerController) getById (cid string) (model.Customer, error){
	customer, err := ctl.Store.GetById(cid)
	if err != nil {
		fmt.Println("Error: ", err)
		return model.Customer{}, err
	}
	fmt.Println("Customer details by ID: ", customer)
	return customer, nil
}

func (ctl *CustomerController) getAll () ([]model.Customer, error) {
	customers, err := ctl.Store.GetAll()
	if err != nil {
		fmt.Println("Error: ", err)
		return []model.Customer{}, err
	}
	fmt.Println("Number of customers: ", len(customers))
	fmt.Println("All customers: ", customers)
	return customers, nil
}

//HTTP Post - /api/customer
func (ctl *CustomerController) Post(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer
	// Decode the incoming customer json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create customer
	//if err := ctl.Store.Create(customer); err != nil {
	if err := ctl.add(customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

//HTTP Put - /api/customer/{id}
func (ctl *CustomerController) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var customer model.Customer
	// Decode the incoming customer json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update
	//if err := ctl.Store.Update(id, customer); err != nil {
	if err := ctl.update(id, customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//HTTP Get - /api/customers
func (ctl *CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	// Get all
	//if customers, err := ctl.Store.GetAll(); err != nil {
	if customers, err := ctl.getAll(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	} else {
		j, err := json.Marshal(customers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//HTTP Get - /api/customers/{id}
func (ctl *CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Get by id
	//if customer, err := ctl.Store.GetById(id); err != nil {
	if customer, err := ctl.getById(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//HTTP Delete - /api/customer/{id}
func (ctl *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// delete
	//if err := ctl.Store.Delete(id); err != nil {
	if err := ctl.delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (ctl *CustomerController) InitializeCustomerController()  {
	c1 := model.Customer{
		ID:    "cust101",
		Name:  "JPM",
		Email: "jpm@domain.com",
	}
	c2 := model.Customer{
		ID:    "cust102",
		Name:  "Pavan",
		Email: "pavan@ibm.com",
	}

	c3 := model.Customer{
		ID:    "cust103",
		Name:  "asdf",
		Email: "asdf@org.com",
	}
	_ = ctl.add(c1) // Create new Customer
	_ = ctl.add(c2) // Create new Customer
	_ = ctl.add(c3) // Create new Customer
}
