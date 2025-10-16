package sql_storage

import "github.com/go-next-pizza/internal/app/model"

type ComboRepository struct {
	storage *SQLStorage
}

func (cr *ComboRepository) GetCombos() ([]*model.Combo, error) {

	rows, err := cr.storage.db.Query("SELECT * FROM combos")
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	
	var combos []*model.Combo
	
	for rows.Next() {
		combo := &model.Combo{}
		if err := rows.Scan(&combo.Id, &combo.Title, &combo.Description, &combo.Price, &combo.Image); err != nil {
			return nil, err
		}

		combos = append(combos, combo)
	}

	return combos, nil
}

func (cr *ComboRepository) GetComboById(comboId int) (*model.Combo, error) {
	combo := &model.Combo{}

	if err := cr.storage.db.QueryRow("SELECT * FROM combos WHERE id = $1", comboId).Scan(
		&combo.Id,
		&combo.Title,
		&combo.Description,
		&combo.Price,
		&combo.Image,
	); err != nil {
		return nil, err
	}

	return combo, nil
}

func (cr *ComboRepository) GetComboProducts(comboId int) ([]*model.Product, error) {
	rows, err := cr.storage.db.Query("SELECT product_id FROM combo_products WHERE combo_id = $1", comboId)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	
	var products []*model.Product
	
	for rows.Next() {
		product := &model.Product{}
		if err := rows.Scan(&product.Id); err != nil {
			return nil, err
		}

		if err := cr.storage.db.QueryRow("SELECT title, description, price, image, amount, weight FROM products WHERE id = $1", product.Id).Scan(&product.Title, &product.Description, &product.Price, &product.Image, &product.Amount, &product.Weight); err != nil {
			return nil, err
		}
		
		products = append(products, product)
	}

	return products, nil
}

func (cr *ComboRepository) GetComboDefaultProducts(comboId int) ([]*model.Product, error) {
	rows, err := cr.storage.db.Query("SELECT DISTINCT default_product_id FROM combo_replaces WHERE combo_id = $1", comboId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*model.Product

	for rows.Next() {
		product := &model.Product{}

		if err := rows.Scan(&product.Id); err != nil {
			return nil, err
		}

		if err := cr.storage.db.QueryRow("SELECT title, description, price, image, amount, weight FROM products WHERE id = $1", product.Id).Scan(&product.Title, &product.Description, &product.Price, &product.Image, &product.Amount, &product.Weight); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, err

}

func (cr *ComboRepository) GetComboReplaces(comboId int) (map[int][]int, error) {
	rows, err := cr.storage.db.Query("SELECT default_product_id, product_to_replace_id FROM combo_replaces WHERE combo_id = $1", comboId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	
	replaces := make(map[int][]int)
	
	for rows.Next() {
		var defaultProductId, replaceProductId int
		if err := rows.Scan(&defaultProductId, &replaceProductId); err != nil {
			return nil, err
		}
		
		replaces[defaultProductId] = append(replaces[defaultProductId], replaceProductId)
	}
	
	return replaces, nil
}
