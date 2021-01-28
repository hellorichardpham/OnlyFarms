package daos

import (
	"database/sql"

	"github.com/hellorichardpham/onlyfarms/user-service/models"
)

//User Get info from the Users MySQL table
type User struct {
	DB *sql.DB
}

//If a function is lowercase, it won't be exported to other packages.
//If it's uppercase, it can be accessed in the Main package.

//Get get a user based on their ID
func (orm *User) Get(ID int) (*models.User, error) {
	query := "SELECT id, name, email, password, created, active FROM users WHERE id = ?"
	user := &models.User{}
	row := orm.DB.QueryRow(query, ID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Created, &user.Active)

	if err != nil {
		return nil, err
	}
	return user, nil
}

//Insert add a new user into the database
func (orm *User) Insert(name, email, password string) error {
	query := "INSERT INTO users (name, email, password, created) VALUES (?, ?, ?, UTC_TIMESTAMP())"
	_, err := orm.DB.Exec(query, name, email, password)
	if err != nil {
		return err
	}
	return nil
}

//Authenticate authenticate a user based on email and password
func (orm *User) Authenticate(email, password string) (int, error) {
	var id int
	query := "SELECT id FROM users WHERE email = ? AND password = ? AND active = true"
	row := orm.DB.QueryRow(query, email, password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
