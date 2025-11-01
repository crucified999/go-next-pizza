package sql_storage

// import (
// 	"database/sql"
// 	"encoding/json"

// 	"github.com/go-next-pizza/internal/app/model"
// )

// type ProductVariantRepository struct {
// 	db *sql.DB
// }

// func NewProductVariantRepository(db *sql.DB) *ProductVariantRepository {
// 	return &ProductVariantRepository{db: db}
// }

// // GetProductVariants получает все варианты для продукта
// func (r *ProductVariantRepository) GetProductVariants(productId int) ([]model.ProductVariant, error) {
// 	query := `
// 		SELECT vt.id, vt.name, vt.product_id, vt.is_required, vt.display_order
// 		FROM product_variant_types vt
// 		WHERE vt.product_id = $1
// 		ORDER BY vt.display_order
// 	`
	
// 	rows, err := r.db.Query(query, productId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var variants []model.ProductVariant
// 	for rows.Next() {
// 		var variant model.ProductVariant
// 		err := rows.Scan(
// 			&variant.Id,
// 			&variant.Name,
// 			&variant.ProductId,
// 			&variant.IsRequired,
// 			&variant.DisplayOrder,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Получаем опции для этого варианта
// 		options, err := r.getVariantOptions(variant.Id)
// 		if err != nil {
// 			return nil, err
// 		}
// 		variant.Options = options

// 		variants = append(variants, variant)
// 	}

// 	return variants, nil
// }

// // getVariantOptions получает опции для варианта
// func (r *ProductVariantRepository) getVariantOptions(variantTypeId int) ([]model.ProductVariantOption, error) {
// 	query := `
// 		SELECT id, variant_type_id, name, value, price_modifier, is_default, display_order
// 		FROM product_variant_options
// 		WHERE variant_type_id = $1
// 		ORDER BY display_order
// 	`
	
// 	rows, err := r.db.Query(query, variantTypeId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var options []model.ProductVariantOption
// 	for rows.Next() {
// 		var option model.ProductVariantOption
// 		err := rows.Scan(
// 			&option.Id,
// 			&option.VariantTypeId,
// 			&option.Name,
// 			&option.Value,
// 			&option.PriceModifier,
// 			&option.IsDefault,
// 			&option.DisplayOrder,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		options = append(options, option)
// 	}

// 	return options, nil
// }

// // AddProductVariantToCart добавляет продукт с вариантами в корзину
// func (r *ProductVariantRepository) AddProductVariantToCart(cartId, productId int, variantData map[string]string, quantity int, basePrice int) error {
// 	// Вычисляем итоговую цену
// 	calculatedPrice := r.calculatePrice(basePrice, variantData)
	
// 	// Конвертируем variantData в JSON
// 	variantDataJSON, err := json.Marshal(variantData)
// 	if err != nil {
// 		return err
// 	}

// 	query := `
// 		INSERT INTO cart_product_variants (cart_id, product_id, variant_data, quantity, calculated_price)
// 		VALUES ($1, $2, $3, $4, $5)
// 	`
	
// 	_, err = r.db.Exec(query, cartId, productId, variantDataJSON, quantity, calculatedPrice)
// 	return err
// }

// // GetCartProductVariants получает все продукты с вариантами из корзины
// func (r *ProductVariantRepository) GetCartProductVariants(cartId int) ([]model.CartProductVariant, error) {
// 	query := `
// 		SELECT cpv.id, cpv.cart_id, cpv.product_id, cpv.variant_data, cpv.quantity, cpv.calculated_price,
// 		       p.id, p.title, p.description, p.price, p.image, p.amount, p.weight
// 		FROM cart_product_variants cpv
// 		JOIN products p ON cpv.product_id = p.id
// 		WHERE cpv.cart_id = $1
// 	`
	
// 	rows, err := r.db.Query(query, cartId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var variants []model.CartProductVariant
// 	for rows.Next() {
// 		var variant model.CartProductVariant
// 		var variantDataJSON []byte
// 		var product model.Product
		
// 		err := rows.Scan(
// 			&variant.Id,
// 			&variant.CartId,
// 			&variant.ProductId,
// 			&variantDataJSON,
// 			&variant.Quantity,
// 			&variant.CalculatedPrice,
// 			&product.Id,
// 			&product.Title,
// 			&product.Description,
// 			&product.Price,
// 			&product.Image,
// 			&product.Amount,
// 			&product.Weight,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Парсим JSON с данными вариантов
// 		err = json.Unmarshal(variantDataJSON, &variant.VariantData)
// 		if err != nil {
// 			return nil, err
// 		}

// 		variant.Product = &product
// 		variants = append(variants, variant)
// 	}

// 	return variants, nil
// }

// // RemoveProductVariantFromCart удаляет продукт с вариантами из корзины
// func (r *ProductVariantRepository) RemoveProductVariantFromCart(cartId, productId int, variantData map[string]string) error {
// 	variantDataJSON, err := json.Marshal(variantData)
// 	if err != nil {
// 		return err
// 	}

// 	query := `
// 		DELETE FROM cart_product_variants 
// 		WHERE cart_id = $1 AND product_id = $2 AND variant_data = $3
// 	`
	
// 	_, err = r.db.Exec(query, cartId, productId, variantDataJSON)
// 	return err
// }

// // calculatePrice вычисляет итоговую цену с учетом модификаторов вариантов
// func (r *ProductVariantRepository) calculatePrice(basePrice int, variantData map[string]string) int {
// 	// Здесь можно добавить логику для получения модификаторов цены из базы данных
// 	// Пока что просто возвращаем базовую цену
// 	_ = variantData // пока не используется
// 	return basePrice
// }

// // AddProductVariantToOrder добавляет продукт с вариантами в заказ
// func (r *ProductVariantRepository) AddProductVariantToOrder(orderId, productId int, variantData map[string]string, quantity int, calculatedPrice int) error {
// 	variantDataJSON, err := json.Marshal(variantData)
// 	if err != nil {
// 		return err
// 	}

// 	query := `
// 		INSERT INTO order_product_variants (order_id, product_id, variant_data, quantity, calculated_price)
// 		VALUES ($1, $2, $3, $4, $5)
// 	`
	
// 	_, err = r.db.Exec(query, orderId, productId, variantDataJSON, quantity, calculatedPrice)
// 	return err
// }

// // GetOrderProductVariants получает все продукты с вариантами из заказа
// func (r *ProductVariantRepository) GetOrderProductVariants(orderId int) ([]model.OrderProductVariant, error) {
// 	query := `
// 		SELECT opv.id, opv.order_id, opv.product_id, opv.variant_data, opv.quantity, opv.calculated_price,
// 		       p.id, p.title, p.description, p.price, p.image, p.amount, p.weight
// 		FROM order_product_variants opv
// 		JOIN products p ON opv.product_id = p.id
// 		WHERE opv.order_id = $1
// 	`
	
// 	rows, err := r.db.Query(query, orderId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var variants []model.OrderProductVariant
// 	for rows.Next() {
// 		var variant model.OrderProductVariant
// 		var variantDataJSON []byte
// 		var product model.Product
		
// 		err := rows.Scan(
// 			&variant.Id,
// 			&variant.OrderId,
// 			&variant.ProductId,
// 			&variantDataJSON,
// 			&variant.Quantity,
// 			&variant.CalculatedPrice,
// 			&product.Id,
// 			&product.Title,
// 			&product.Description,
// 			&product.Price,
// 			&product.Image,
// 			&product.Amount,
// 			&product.Weight,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Парсим JSON с данными вариантов
// 		err = json.Unmarshal(variantDataJSON, &variant.VariantData)
// 		if err != nil {
// 			return nil, err
// 		}

// 		variant.Product = &product
// 		variants = append(variants, variant)
// 	}

// 	return variants, nil
// }