package mysql

import (
	"database/sql"
	"errors"

	"github.com/hellorichardpham/onlyfarms/pkg/models"
)

// PackageModel ...
type PackageModel struct {
	DB *sql.DB
}

func (model *PackageModel) Get(id int) (*models.Package, error) {
	query := `SELECT id, name FROM packages WHERE id = ?`

	pack := &models.Package{}

	row := model.DB.QueryRow(query, id)
	err := row.Scan(&pack.ID, &pack.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return pack, nil
}
