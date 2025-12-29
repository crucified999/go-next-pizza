CREATE TABLE users (
  id bigserial PRIMARY KEY,
  email varchar UNIQUE,
  -- encrypted_password varchar NOT NULL,
  phone varchar UNIQUE,
  name varchar
);

CREATE TABLE sms_codes (
  id bigserial PRIMARY KEY,
  phone varchar NOT NULL,
  code varchar NOT NULL,
  expires_at timestamp NOT NULL
);

CREATE TABLE carts (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL,
  total_count int
);

CREATE TABLE products (
  id bigserial not null primary key,
  title varchar,
  description varchar,
  price int,
  image varchar,
  amount varchar,
  weight int
);

CREATE TABLE categories (
  id bigserial not null primary key,
  category_id bigint not null,
  product_id bigint
);

CREATE TABLE toppings (
  id bigserial not null primary key,
  title varchar not null,
  product_id bigint,
  image varchar,
  price int
);

CREATE TABLE pizza_ingredients (
  id bigserial not null primary key,
  title varchar not null,
  product_id bigint,
  replacable int
);

CREATE TABLE orders (
  id bigserial not null primary key,
  user_id bigint not null,
  payment_method varchar,
  delivery_address varchar,
  delivery_time timestamp,
  status varchar,
  total_price int not null,
  created_at timestamp not null
);

CREATE TABLE combos (
  id bigserial not null primary key,
  title varchar,
  description varchar,
  price int,
  image varchar
);

CREATE TABLE combo_products (
  combo_id bigint not null,
  product_id bigint not null
);

CREATE TABLE combo_replaces (
  combo_id bigint not null,
  default_product_id bigint not null,
  product_to_replace_id bigint
);

CREATE TABLE products_in_cart (
  cart_id bigint,
  product_id bigint,
  amount varchar,
  count int
);

CREATE TABLE pizzas_in_cart (
  cart_id bigint,
  pizza_id bigint,
  size varchar,
  dough varchar,
  toppings_mask bigint,
  count int
);

CREATE TABLE toppings_in_cart (
  cart_id bigint,
  product_id bigint,
  topping_id bigint
);

CREATE TABLE combos_in_cart (
  cart_id bigint,
  combo_id bigint,
  amount int
);

CREATE TABLE products_in_order (
  order_id bigint,
  product_id bigint,
  amount int
);

CREATE TABLE combos_in_order (
  order_id bigint,
  combo_id bigint,
  amount int
);

ALTER TABLE combo_products
ADD CONSTRAINT combo_products_combo_id_fk FOREIGN KEY (combo_id) REFERENCES combos(id);

ALTER TABLE combo_products
ADD CONSTRAINT combo_products_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id);

ALTER TABLE products_in_order
ADD CONSTRAINT products_order_id_fk FOREIGN KEY (order_id) REFERENCES orders(id);

ALTER TABLE products_in_order
ADD CONSTRAINT products_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id);

ALTER TABLE combos_in_order
ADD CONSTRAINT combos_order_id_fk FOREIGN KEY (order_id) REFERENCES orders(id);

ALTER TABLE combos_in_order
ADD CONSTRAINT combos_combo_id_fk FOREIGN KEY (combo_id) REFERENCES combos(id);

ALTER TABLE combo_replaces
ADD CONSTRAINT combo_replaces_combo_id_fk FOREIGN KEY (combo_id) REFERENCES combos(id);

ALTER TABLE combo_replaces
ADD CONSTRAINT combo_replaces_default_product_id_fk FOREIGN KEY (default_product_id) REFERENCES products(id);

ALTER TABLE combo_replaces
ADD CONSTRAINT combo_replaces_product_to_replace_id_fk FOREIGN KEY (product_to_replace_id) REFERENCES products(id);

ALTER TABLE combos_in_cart
ADD CONSTRAINT combos_cart_id_fk FOREIGN KEY (cart_id) REFERENCES carts(id);

ALTER TABLE combos_in_cart
ADD CONSTRAINT combos_combo_id_fk FOREIGN KEY (combo_id) REFERENCES combos(id);

ALTER TABLE products_in_cart
ADD CONSTRAINT products_cart_id_fk FOREIGN KEY (cart_id) REFERENCES carts(id);

ALTER TABLE products_in_cart
ADD CONSTRAINT products_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id);

ALTER TABLE carts
ADD CONSTRAINT cart_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE orders
ADD CONSTRAINT order_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE pizza_ingredients
ADD CONSTRAINT ingredient_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id);

ALTER TABLE toppings
ADD CONSTRAINT topping_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id);

ALTER TABLE categories
ADD CONSTRAINT product_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id);