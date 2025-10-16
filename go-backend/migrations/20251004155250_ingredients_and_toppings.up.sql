CREATE TABLE products_with_size (
  product_id bigint not null,
  size varchar,
  dough_type int,
  image varchar,
  weight int,
  price int
);

ALTER TABLE products_with_size
ADD CONSTRAINT product_with_size_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id);