-- Создание таблиц для системы вариантов продуктов

-- Таблица для типов вариантов (размер, тесто, количество и т.д.)
CREATE TABLE product_variant_types (
    id bigserial PRIMARY KEY,
    name varchar NOT NULL, -- "Размер", "Тесто", "Количество"
    product_id bigint NOT NULL,
    is_required boolean DEFAULT true, -- обязательный ли вариант
    display_order int DEFAULT 0
);

-- Таблица для опций вариантов (малое тесто, средняя пицца, 5 креветок)
CREATE TABLE product_variant_options (
    id bigserial PRIMARY KEY,
    variant_type_id bigint NOT NULL,
    name varchar NOT NULL, -- "Малое", "Среднее", "Большое"
    value varchar NOT NULL, -- "small", "medium", "large"
    price_modifier int DEFAULT 0, -- изменение цены (+50, -20, 0)
    is_default boolean DEFAULT false,
    display_order int DEFAULT 0
);

-- Таблица для связи продуктов с их вариантами в корзине
CREATE TABLE cart_product_variants (
    id bigserial PRIMARY KEY,
    cart_id bigint NOT NULL,
    product_id bigint NOT NULL,
    variant_data jsonb NOT NULL, -- {"size": "large", "dough": "thin", "quantity": 5}
    quantity int NOT NULL DEFAULT 1,
    calculated_price int NOT NULL -- итоговая цена с учетом вариантов
);

-- Таблица для связи продуктов с их вариантами в заказах
CREATE TABLE order_product_variants (
    id bigserial PRIMARY KEY,
    order_id bigint NOT NULL,
    product_id bigint NOT NULL,
    variant_data jsonb NOT NULL,
    quantity int NOT NULL DEFAULT 1,
    calculated_price int NOT NULL
);

-- Внешние ключи
ALTER TABLE product_variant_types
ADD CONSTRAINT product_variant_types_product_id_fk 
FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;

ALTER TABLE product_variant_options
ADD CONSTRAINT product_variant_options_variant_type_id_fk 
FOREIGN KEY (variant_type_id) REFERENCES product_variant_types(id) ON DELETE CASCADE;

ALTER TABLE cart_product_variants
ADD CONSTRAINT cart_product_variants_cart_id_fk 
FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE;

ALTER TABLE cart_product_variants
ADD CONSTRAINT cart_product_variants_product_id_fk 
FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;

ALTER TABLE order_product_variants
ADD CONSTRAINT order_product_variants_order_id_fk 
FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE;

ALTER TABLE order_product_variants
ADD CONSTRAINT order_product_variants_product_id_fk 
FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;

-- Индексы для производительности
CREATE INDEX idx_product_variant_types_product_id ON product_variant_types(product_id);
CREATE INDEX idx_product_variant_options_variant_type_id ON product_variant_options(variant_type_id);
CREATE INDEX idx_cart_product_variants_cart_id ON cart_product_variants(cart_id);
CREATE INDEX idx_order_product_variants_order_id ON order_product_variants(order_id);