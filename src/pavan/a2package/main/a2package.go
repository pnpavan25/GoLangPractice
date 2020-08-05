package main

import (
	"fmt"
	"pavan/a2package/domain"
	"pavan/a2package/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore // It can be any Store
}

func (cc CustomerController) Add (c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (cc CustomerController) Update (cid string, c domain.Customer) {
	err := cc.store.Update(cid, c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Customer has been updated")
}

func (cc CustomerController) Delete (cid string) {
	err := cc.store.Delete(cid)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Customer with given ID %s is deleted\n", cid)
}

func (cc CustomerController) GetById (cid string) {
	customer, err := cc.store.GetById(cid)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Customer details by ID: ", customer)
}

func (cc CustomerController) GetAll () {
	customers, err := cc.store.GetAll()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Number of customers: ", len(customers))
	fmt.Println("All customers: ", customers)
}

func main() {
	controller := CustomerController {
		store : mapstore.NewMapStore(),
	}
	c1 := domain.Customer {
		ID : "cust101",
		Name: "JPM",
		Email: "jpm@domain.com",
	}
	c2 := domain.Customer {
		ID : "cust102",
		Name: "Pavan",
		Email: "pavan@ibm.com",
	}

	c3 := domain.Customer {
		ID : "cust103",
		Name: "asdf",
		Email: "asdf@domain.com",
	}
	controller.Add(c1) // Create new Customer
	controller.Add(c2) // Create new Customer
	controller.Add(c3) // Create new Customer

	c4 := domain.Customer {
		ID : "cust103",
		Name: "zxcv",
		Email: "zxcv@domain.com",
	}

	controller.GetById("cust102") // Get Customer by ID
	controller.GetAll() // Get All customers

	controller.Update("cust103", c4) // Update new Customer
	controller.GetById("cust103") // Get Customer by ID

	controller.Delete("cust103") // Delete Customer by ID

	controller.GetAll() // Get All customers
}
