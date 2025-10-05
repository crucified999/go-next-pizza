CREATE TABLE pizza_sizes (
  pizza_id bigint not null,
  dough_size int,
  dough_type int,
  image varchar,
  weight int,
  price int
);

ALTER TABLE pizza_sizes
ADD CONSTRAINT pizza_product_id_fk FOREIGN KEY (pizza_id) REFERENCES products(id);