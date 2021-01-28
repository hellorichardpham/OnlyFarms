package mysql

import (
	"database/sql"
	"errors"

	"github.com/hellorichardpham/onlyfarms/product-service/pkg/models"
)

type FarmerModel struct {
	DB *sql.DB
}

func (model *FarmerModel) Get(id int) (*models.Farmer, error) {
	query := `SELECT id, name, street1, city, state, zip, picture, description FROM farmers WHERE id = ?`
	farmer := &models.Farmer{}

	row := model.DB.QueryRow(query, id)
	err := row.Scan(&farmer.ID, &farmer.Name, &farmer.Street1, &farmer.City, &farmer.State,
		&farmer.Zip, &farmer.Picture, &farmer.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return farmer, nil
}
