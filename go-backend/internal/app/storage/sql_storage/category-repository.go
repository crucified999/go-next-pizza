package sql_storage

import "github.com/go-next-pizza/internal/app/model"

type CategoryRepository struct {
	storage *SQLStorage
}

func (cr *CategoryRepository) GetCategories() ([]*model.Category, error) {
	rows, err := cr.storage.db.Query("SELECT DISTINCT category_id FROM categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		category := &model.Category{}
		if err := rows.Scan(&category.ID); err != nil {
			return nil, err
		}

		category.Title = cr.ConvertIdToCategory(category.ID)
		categories = append(categories, category)
	}

	return categories, nil
}

func (cr *CategoryRepository) ConvertIdToCategory(id int) string {
	
	switch id {
	case 1:
		return "Пиццы"
	case 2:
		return "Комбо"
	case 3:
		return "Закуски"
	case 4:
		return "Коктейли"
	case 5:
		return "Кофе"
	case 6:
		return "Напитки"
	case 7:
		return "Десерты"
	case 8:
		return "Соусы"
	default:
		return ""
	}
}
