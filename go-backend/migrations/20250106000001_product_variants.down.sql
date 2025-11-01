-- Откат миграции для системы вариантов продуктов

DROP INDEX IF EXISTS idx_order_product_variants_order_id;
DROP INDEX IF EXISTS idx_cart_product_variants_cart_id;
DROP INDEX IF EXISTS idx_product_variant_options_variant_type_id;
DROP INDEX IF EXISTS idx_product_variant_types_product_id;

DROP TABLE IF EXISTS order_product_variants;
DROP TABLE IF EXISTS cart_product_variants;
DROP TABLE IF EXISTS product_variant_options;
DROP TABLE IF EXISTS product_variant_types;