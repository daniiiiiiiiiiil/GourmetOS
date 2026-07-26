

CREATE TABLE IF NOT EXISTS employees (
                                         employee_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
                                         name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    role VARCHAR(50) NOT NULL,
    shift VARCHAR(20),
    hire_date DATE NOT NULL,
    salary DECIMAL(10, 2),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT PK_employees_employee_id PRIMARY KEY (employee_id)
    );

CREATE TABLE IF NOT EXISTS tables (
                                      table_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
                                      number INT UNIQUE NOT NULL,
                                      capacity INT NOT NULL,
                                      location VARCHAR(50),
    is_occupied BOOLEAN DEFAULT FALSE,
    is_reserved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT PK_tables_table_id PRIMARY KEY (table_id)
    );

CREATE TABLE IF NOT EXISTS customers (
                                         customer_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
                                         name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE,
    address TEXT,
    birth_date DATE,
    loyalty_level VARCHAR(20) DEFAULT 'bronze',
    total_orders INT DEFAULT 0,
    total_spent DECIMAL(10, 2) DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT PK_customers_customer_id PRIMARY KEY (customer_id)
    );

CREATE TABLE IF NOT EXISTS ingredients (
                                           ingredient_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
                                           name VARCHAR(100) UNIQUE NOT NULL,
    unit VARCHAR(20) NOT NULL,
    stock_quantity DECIMAL(10, 2) DEFAULT 0,
    min_stock DECIMAL(10, 2) DEFAULT 0,
    max_stock DECIMAL(10, 2) DEFAULT 0,
    cost_per_unit DECIMAL(10, 2) DEFAULT 0,
    supplier VARCHAR(100),
    expiry_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT PK_ingredients_ingredient_id PRIMARY KEY (ingredient_id)
    );

CREATE TABLE IF NOT EXISTS dishes (
                                      dish_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
                                      name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(50),
    cuisine VARCHAR(50),
    cooking_time INT DEFAULT 15,
    is_available BOOLEAN DEFAULT TRUE,
    is_vegetarian BOOLEAN DEFAULT FALSE,
    is_vegan BOOLEAN DEFAULT FALSE,
    is_gluten_free BOOLEAN DEFAULT FALSE,
    calories INT,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT PK_dishes_dish_id PRIMARY KEY (dish_id)
    );

CREATE TABLE IF NOT EXISTS orders (
                                      order_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
                                      table_id INT,
                                      customer_id INT,
                                      waiter_id INT,
                                      status VARCHAR(30) DEFAULT 'created',
    total_amount DECIMAL(10, 2) DEFAULT 0,
    discount_amount DECIMAL(10, 2) DEFAULT 0,
    final_amount DECIMAL(10, 2) DEFAULT 0,
    payment_method VARCHAR(30),
    payment_status VARCHAR(20) DEFAULT 'pending',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT PK_orders_order_id PRIMARY KEY (order_id),
    CONSTRAINT FK_orders_table FOREIGN KEY (table_id) REFERENCES tables(table_id) ON DELETE SET NULL,
    CONSTRAINT FK_orders_customer FOREIGN KEY (customer_id) REFERENCES customers(customer_id) ON DELETE SET NULL,
    CONSTRAINT FK_orders_waiter FOREIGN KEY (waiter_id) REFERENCES employees(employee_id) ON DELETE SET NULL
    );

CREATE TABLE IF NOT EXISTS payments (
                                        payment_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
                                        order_id INT NOT NULL,
                                        amount DECIMAL(10, 2) NOT NULL,
    method VARCHAR(30) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    transaction_id VARCHAR(100),
    card_last4 VARCHAR(4),
    crypto_address VARCHAR(100),
    receipt_url VARCHAR(255),
    paid_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT PK_payments_payment_id PRIMARY KEY (payment_id),
    CONSTRAINT FK_payments_order FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS order_items (
                                           order_item_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
                                           order_id INT NOT NULL,
                                           dish_id INT NOT NULL,
                                           quantity INT NOT NULL DEFAULT 1,
                                           price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT PK_order_items_order_item_id PRIMARY KEY (order_item_id),
    CONSTRAINT FK_order_items_order FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE,
    CONSTRAINT FK_order_items_dish FOREIGN KEY (dish_id) REFERENCES dishes(dish_id) ON DELETE CASCADE
    );

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_employees_updated_at BEFORE UPDATE ON employees FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_tables_updated_at BEFORE UPDATE ON tables FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_customers_updated_at BEFORE UPDATE ON customers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_ingredients_updated_at BEFORE UPDATE ON ingredients FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_dishes_updated_at BEFORE UPDATE ON dishes FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_orders_updated_at BEFORE UPDATE ON orders FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_payments_updated_at BEFORE UPDATE ON payments FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_order_items_updated_at BEFORE UPDATE ON order_items FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
