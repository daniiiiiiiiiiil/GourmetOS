CREATE INDEX idx_employees_role ON employees(role);

CREATE INDEX idx_tables_is_occupied ON tables(is_occupied);

CREATE INDEX idx_customers_phone ON customers(phone);

CREATE INDEX idx_dishes_category ON dishes(category);
CREATE INDEX idx_dishes_cuisine ON dishes(cuisine);

CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_created_at ON orders(created_at);
CREATE INDEX idx_orders_table_id ON orders(table_id);
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
CREATE INDEX idx_orders_waiter_id ON orders(waiter_id);       

CREATE INDEX idx_payments_order_id ON payments(order_id);

CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_dish_id ON order_items(dish_id);
