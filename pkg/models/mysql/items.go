package mysql

import (
	"database/sql"
	"errors"

	"github.com/hellorichardpham/onlyfarms/pkg/models"
)

type ItemModel struct {
	DB *sql.DB
}

func (model *ItemModel) Get(id int) (*models.Item, error) {
	query := `SELECT id, name, price, packageId FROM items
		WHERE id = ?`
	item := &models.Item{}

	row := model.DB.QueryRow(query, id)
	err := row.Scan(&item.ID, &item.Name, &item.Price, &item.PackageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return item, nil
}

func (model *ItemModel) GetItemsByPackageId(packageId int) ([]*models.Item, error) {
	query := `SELECT id, name, price, packageId FROM items
		WHERE packageId = ?`

	rows, err := model.DB.Query(query, packageId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*models.Item{}

	for rows.Next() {
		item := &models.Item{}
		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.PackageID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
