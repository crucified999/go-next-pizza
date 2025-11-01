-- Создание полных вариантов для всех продуктов

-- Варианты для всех пицц (размер и тесто)
INSERT INTO product_variant_types (name, product_id, is_required, display_order)
SELECT 'Размер', id, true, 1
FROM products 
WHERE title IN (
    'Пепперони фреш', 'Терияки', 'Чесночный цыпленок', 'Пикантные колбаски', 
    'Четыре сыра', 'Сырная', 'Чоризо фреш', 'Ветчина и сыр', 'Двойной цыпленок',
    'Креветка и песто', 'Чилл Грилл', 'Ветчина и грибы', 'Аррива!', 'Креветки со сладким чили',
    'Бефстроганов', 'Карбонара', 'Жюльен', 'Песто', 'Мясная', 'Бургер-пицца',
    'Сырный цыпленок', 'Додо', 'Пепперони', 'Гавайская', 'Цыпленок барбекю'
);

INSERT INTO product_variant_types (name, product_id, is_required, display_order)
SELECT 'Тесто', id, true, 2
FROM products 
WHERE title IN (
    'Пепперони фреш', 'Терияки', 'Чесночный цыпленок', 'Пикантные колбаски', 
    'Четыре сыра', 'Сырная', 'Чоризо фреш', 'Ветчина и сыр', 'Двойной цыпленок',
    'Креветка и песто', 'Чилл Грилл', 'Ветчина и грибы', 'Аррива!', 'Креветки со сладким чили',
    'Бефстроганов', 'Карбонара', 'Жюльен', 'Песто', 'Мясная', 'Бургер-пицца',
    'Сырный цыпленок', 'Додо', 'Пепперони', 'Гавайская', 'Цыпленок барбекю'
);

-- Опции размеров для пицц
INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, 'Маленькая', 'small', -100, true, 1
FROM product_variant_types vt
WHERE vt.name = 'Размер';

INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, 'Средняя', 'medium', 0, false, 2
FROM product_variant_types vt
WHERE vt.name = 'Размер';

INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, 'Большая', 'large', 100, false, 3
FROM product_variant_types vt
WHERE vt.name = 'Размер';

-- Опции теста для пицц
INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, 'Тонкое', 'thin', 0, true, 1
FROM product_variant_types vt
WHERE vt.name = 'Тесто';

INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, 'Толстое', 'thick', 50, false, 2
FROM product_variant_types vt
WHERE vt.name = 'Тесто';

-- Варианты для креветок (количество)
INSERT INTO product_variant_types (name, product_id, is_required, display_order)
SELECT 'Количество', id, true, 1
FROM products 
WHERE title LIKE '%Креветки%';

-- Опции количества для креветок
INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, '5 штук', '5', 0, true, 1
FROM product_variant_types vt
WHERE vt.name = 'Количество' AND vt.product_id IN (
    SELECT id FROM products WHERE title LIKE '%Креветки%'
);

INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, '9 штук', '9', 260, false, 2
FROM product_variant_types vt
WHERE vt.name = 'Количество' AND vt.product_id IN (
    SELECT id FROM products WHERE title LIKE '%Креветки%'
);

-- Варианты для хашбраунов (количество)
INSERT INTO product_variant_types (name, product_id, is_required, display_order)
SELECT 'Количество', id, true, 1
FROM products 
WHERE title LIKE '%Хашбрауны%';

-- Опции количества для хашбраунов
INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, '2 штуки', '2', 0, true, 1
FROM product_variant_types vt
WHERE vt.name = 'Количество' AND vt.product_id IN (
    SELECT id FROM products WHERE title LIKE '%Хашбрауны%'
);

INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, '3 штуки', '3', 50, false, 2
FROM product_variant_types vt
WHERE vt.name = 'Количество' AND vt.product_id IN (
    SELECT id FROM products WHERE title LIKE '%Хашбрауны%'
);

INSERT INTO product_variant_options (variant_type_id, name, value, price_modifier, is_default, display_order)
SELECT vt.id, '4 штуки', '4', 100, false, 3
FROM product_variant_types vt
WHERE vt.name = 'Количество' AND vt.product_id IN (
    SELECT id FROM products WHERE title LIKE '%Хашбрауны%'
);