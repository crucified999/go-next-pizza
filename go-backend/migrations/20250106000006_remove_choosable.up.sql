-- Удаление поля choosable из таблицы products
-- Теперь информация о том, что продукт выбираемый, определяется наличием вариантов

ALTER TABLE products DROP COLUMN IF EXISTS choosable;