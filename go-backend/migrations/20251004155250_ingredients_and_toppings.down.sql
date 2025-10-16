DROP TABLE products_with_size;

ALTER TABLE products_with_size
DROP CONSTRAINT IF EXISTS product_with_size_product_id_fk;