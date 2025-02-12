package category_repo

import (
	"database/sql"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepo {
	return CategoryRepo{db: db}
}

func (r CategoryRepo) GetCategoriesByIDs(categoryIDs []int64) (categories []string, err error) {
	categories = make([]string, len(categoryIDs))

	for i, id := range categoryIDs {
		row := r.db.QueryRow("SELECT category_name from Category WHERE category_id = ?", id)
		err = row.Scan(&categories[i])

		if err != nil {
			return categories, err
		}
	}
	return categories, nil
}
