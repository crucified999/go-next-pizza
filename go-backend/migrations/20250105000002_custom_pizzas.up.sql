CREATE TABLE custom_pizzas (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL,
  base_pizza_id bigint NOT NULL,
  name varchar,
  total_price real NOT NULL,
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (base_pizza_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE custom_pizza_ingredients (
  id bigserial PRIMARY KEY,
  custom_pizza_id bigint NOT NULL,
  ingredient_id bigint NOT NULL,
  is_added boolean NOT NULL DEFAULT true,
  created_at timestamp DEFAULT NOW(),
  FOREIGN KEY (custom_pizza_id) REFERENCES custom_pizzas(id) ON DELETE CASCADE,
  FOREIGN KEY (ingredient_id) REFERENCES pizza_ingredients(id) ON DELETE CASCADE
);

CREATE INDEX idx_custom_pizzas_user_id ON custom_pizzas(user_id);
CREATE INDEX idx_custom_pizza_ingredients_custom_pizza_id ON custom_pizza_ingredients(custom_pizza_id);