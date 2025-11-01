-- Восстановление поля choosable в таблице products
-- Восстанавливаем поле choosable (все продукты будут помечены как не выбираемые)

ALTER TABLE products ADD COLUMN choosable boolean DEFAULT false;