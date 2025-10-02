CREATE TABLE users (
  id bigserial not null primary key,
  email varchar not null unique,
  encrypted_password varchar not null,
  cart_id bigserial foreign key
);

CREATE TABLE cart (
  id bigserial not null primary key
  user_id bigserial not null foreign key,
  product_id bigserial not null foreign key
);

CREATE TABLE products (
  id bigserial not null primary key,
  category_id bigserial not null foreign key, 
  topping_id bigserial foreign key
  title varchar not null,
  description varchar,
  price int not null,
  image varchar not null,
  dough_size int,
  size varchar,
  amount int,
);

CREATE TABLE combos (
  id bigserial not null primary key,
  product_id bigserial not null foreign key,
  product_to_replace_id bigserial not null foreign key
);

CREATE TABLE categories (
  id bigserial not null primary key,
  product_id bigserial not null foreign key
);

CREATE TABLE ingredients (
  id bigserial not null primary key,
  product_id bigserial not null foreign key,
  title varchar not null,
  price int not null
); 