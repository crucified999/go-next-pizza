-- Откат данных вариантов продуктов

DELETE FROM product_variant_options WHERE variant_type_id IN (
    SELECT id FROM product_variant_types WHERE name IN ('Размер', 'Тесто', 'Количество')
);

DELETE FROM product_variant_types WHERE name IN ('Размер', 'Тесто', 'Количество');