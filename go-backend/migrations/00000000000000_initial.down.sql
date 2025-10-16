-- Drop foreign key constraints first
ALTER TABLE categories DROP CONSTRAINT IF EXISTS product_product_id_fk;
ALTER TABLE toppings DROP CONSTRAINT IF EXISTS topping_product_id_fk;
ALTER TABLE pizza_ingredients DROP CONSTRAINT IF EXISTS ingredient_product_id_fk;
ALTER TABLE orders DROP CONSTRAINT IF EXISTS order_user_id_fk;
ALTER TABLE carts DROP CONSTRAINT IF EXISTS cart_user_id_fk;
ALTER TABLE products_in_cart DROP CONSTRAINT IF EXISTS products_product_id_fk;
ALTER TABLE products_in_cart DROP CONSTRAINT IF EXISTS products_cart_id_fk;
ALTER TABLE combos_in_cart DROP CONSTRAINT IF EXISTS combos_combo_id_fk;
ALTER TABLE combos_in_cart DROP CONSTRAINT IF EXISTS combos_cart_id_fk;
ALTER TABLE combo_replaces DROP CONSTRAINT IF EXISTS combo_replaces_product_to_replace_id_fk;
ALTER TABLE combo_replaces DROP CONSTRAINT IF EXISTS combo_replaces_default_product_id_fk;
ALTER TABLE combo_replaces DROP CONSTRAINT IF EXISTS combo_replaces_combo_id_fk;
ALTER TABLE combos_in_order DROP CONSTRAINT IF EXISTS combos_combo_id_fk;
ALTER TABLE combos_in_order DROP CONSTRAINT IF EXISTS combos_order_id_fk;
ALTER TABLE products_in_order DROP CONSTRAINT IF EXISTS products_product_id_fk;
ALTER TABLE products_in_order DROP CONSTRAINT IF EXISTS products_order_id_fk;
ALTER TABLE combo_products DROP CONSTRAINT IF EXISTS combo_products_combo_id_fk;
ALTER TABLE combo_products DROP CONSTRAINT IF EXISTS combo_products_product_id_fk;

-- Drop tables
DROP TABLE IF EXISTS combos_in_order;
DROP TABLE IF EXISTS products_in_order;
DROP TABLE IF EXISTS combos_in_cart;
DROP TABLE IF EXISTS products_in_cart;
DROP TABLE IF EXISTS combo_replaces;
DROP TABLE IF EXISTS combo_products;
DROP TABLE IF EXISTS combos;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS pizza_ingredients;
DROP TABLE IF EXISTS toppings;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS users;