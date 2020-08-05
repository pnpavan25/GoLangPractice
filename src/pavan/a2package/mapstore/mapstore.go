package mapstore

import (
	"fmt"
	"pavan/a2package/domain"
)

type MapStore struct {
	store map[string]domain.Customer // An in-memory ms.store
}

// Factory method gives a new instance of MapStore
// This is for caller packages, not for MapStore
func NewMapStore() *MapStore {
	return &MapStore { store: make(map[string]domain.Customer)}
}

// MapStore create domain.Customer
func (ms MapStore) Create(c domain.Customer) error {
	if _, ok := ms.store[c.ID]; ok {
		return fmt.Errorf("given domain.Customer %s already exists", c.ID)
	}
	ms.store[c.ID] = c
	return nil
}

// MapStore update domain.Customer
func (ms MapStore) Update(cid string, c domain.Customer) error {
	if _, ok := ms.store[cid]; !ok {
		return fmt.Errorf("given domain.Customer %s not exists", cid)
	}
	ms.store[cid] = c
	return nil
}

// MapStore delete domain.Customer
func (ms MapStore) Delete(cid string) error {
	if _, ok := ms.store[cid]; !ok {
		return fmt.Errorf("given domain.Customer %s not exists", cid)
	}
	delete(ms.store, cid)
	return nil
}

// MapStore creates domain.Customer
func (ms MapStore) GetById(cid string) (domain.Customer, error) {
	if _, ok := ms.store[cid]; !ok {
		return domain.Customer{}, fmt.Errorf("given domain.Customer %s not exists", cid)
	}
	return ms.store[cid], nil
}

// MapStore creates domain.Customer
func (ms MapStore) GetAll() ([]domain.Customer, error) {
	allCustomers := make([]domain.Customer, len(ms.store))
	for _, customer := range ms.store {
		allCustomers = append(allCustomers, customer)
	}
	return allCustomers, nil
}