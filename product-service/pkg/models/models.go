package models

import "errors"

var (
	// ErrNoRecord ...
	ErrNoRecord = errors.New("models: no matching record found")
)

//Item ...
type Item struct {
	ID        int
	Name      string
	Price     float64
	PackageID int
}

//Package ...
type Package struct {
	ID   int
	Name string
}

//Farmer ...
type Farmer struct {
	ID          int
	Name        string
	Street1     string
	City        string
	State       string
	Zip         string
	Picture     string
	Description string
}
