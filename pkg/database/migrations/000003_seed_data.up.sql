
INSERT INTO employees (name, email, phone, role, shift, hire_date, salary)
VALUES
    ('Иван Иванов', 'ivan@restaurant.com', '+7-999-111-22-33', 'admin', 'morning', '2023-01-15', 80000),
    ('Мария Петрова', 'maria@restaurant.com', '+7-999-222-33-44', 'manager', 'evening', '2023-02-01', 70000),
    ('Алексей Смирнов', 'alex@restaurant.com', '+7-999-333-44-55', 'waiter', 'morning', '2023-03-10', 45000),
    ('Елена Козлова', 'elena@restaurant.com', '+7-999-444-55-66', 'waiter', 'evening', '2023-04-15', 45000),
    ('Сергей Попов', 'sergey@restaurant.com', '+7-999-555-66-77', 'chef', 'morning', '2023-02-20', 60000),
    ('Ольга Соколова', 'olga@restaurant.com', '+7-999-666-77-88', 'chef', 'evening', '2023-03-05', 60000),
    ('Дмитрий Волков', 'dmitry@restaurant.com', '+7-999-777-88-99', 'cashier', 'morning', '2023-05-01', 35000);

INSERT INTO tables (number, capacity, location)
VALUES
    (1, 2, 'hall'),
    (2, 2, 'hall'),
    (3, 4, 'hall'),
    (4, 4, 'hall'),
    (5, 6, 'hall'),
    (6, 2, 'terrace'),
    (7, 4, 'terrace'),
    (8, 8, 'vip'),
    (9, 2, 'vip'),
    (10, 4, 'vip');

INSERT INTO customers (name, phone, email, address, birth_date, loyalty_level)
VALUES
    ('Андрей Морозов', '+7-999-101-20-30', 'andrey@mail.com', 'ул. Ленина 10', '1990-05-15', 'gold'),
    ('Екатерина Новикова', '+7-999-202-30-40', 'ekaterina@mail.com', 'ул. Пушкина 20', '1985-08-20', 'silver'),
    ('Михаил Федоров', '+7-999-303-40-50', 'mikhail@mail.com', 'ул. Гоголя 30', '1992-12-10', 'bronze'),
    ('Наталья Орлова', '+7-999-404-50-60', 'natalia@mail.com', 'ул. Чехова 40', '1988-03-25', 'platinum'),
    ('Владимир Соловьев', '+7-999-505-60-70', 'vladimir@mail.com', 'ул. Толстого 50', '1975-07-05', 'silver');

INSERT INTO ingredients (name, unit, stock_quantity, min_stock, max_stock, cost_per_unit, supplier)
VALUES
    ('Мука', 'kg', 50, 10, 100, 80, 'Хлебосол'),
    ('Томатный соус', 'l', 30, 5, 60, 150, 'Помидорка'),
    ('Моцарелла', 'kg', 20, 3, 40, 600, 'Сырный рай'),
    ('Спагетти', 'kg', 25, 5, 50, 120, 'Паста Italiana'),
    ('Бекон', 'kg', 15, 2, 30, 800, 'Мясной двор'),
    ('Яйца', 'pcs', 100, 20, 200, 10, 'Птицефабрика'),
    ('Оливковое масло', 'l', 20, 5, 40, 500, 'Италия импорт'),
    ('Картофель', 'kg', 40, 10, 80, 50, 'Овощи-Фрукты'),
    ('Фасоль', 'kg', 15, 3, 30, 90, 'Бобовый рай'),
    ('Сыр Пармезан', 'kg', 10, 2, 20, 1200, 'Сырный рай'),
    ('Капуста', 'kg', 20, 5, 40, 40, 'Овощи-Фрукты'),
    ('Свинина', 'kg', 25, 5, 50, 450, 'Мясной двор');

INSERT INTO dishes (name, description, price, category, cuisine, cooking_time)
VALUES
    ('Пицца Маргарита', 'Томатный соус, моцарелла, базилик', 450, 'pizza', 'italian', 15),
    ('Пицца Пепперони', 'Томатный соус, моцарелла, пепперони', 520, 'pizza', 'italian', 15),
    ('Паста Карбонара', 'Спагетти, бекон, яйцо, пармезан', 380, 'pasta', 'italian', 20),
    ('Паста Болоньезе', 'Спагетти, мясной соус, пармезан', 420, 'pasta', 'italian', 25),
    ('Салат Капрезе', 'Томаты, моцарелла, базилик, оливковое масло', 280, 'salad', 'italian', 10),
    ('Лимончелло', 'Освежающий лимонный ликер', 150, 'drink', 'italian', 5),

    ('Окономияки', 'Японская пицца с капустой и свининой', 500, 'pizza', 'japanese', 25),
    ('Рамен', 'Лапша в свином бульоне с яйцом и нори', 420, 'pasta', 'japanese', 30),
    ('Салат Вакамэ', 'Водоросли с огурцом и кунжутом', 250, 'salad', 'japanese', 10),
    ('Маття', 'Зеленый чай с молоком', 180, 'drink', 'japanese', 5),

    ('Тортилья', 'Мексиканская пицца с фасолью и сальсой', 470, 'pizza', 'mexican', 20),
    ('Чили кон карне', 'Говядина с фасолью и перцем чили', 450, 'pasta', 'mexican', 35),
    ('Салат с кактусом', 'Кактус, помидоры, лук, авокадо', 300, 'salad', 'mexican', 15),
    ('Агуа Фреска', 'Освежающий фруктовый напиток', 130, 'drink', 'mexican', 5);

INSERT INTO orders (table_id, customer_id, waiter_id, status, total_amount, discount_amount, final_amount, payment_method, payment_status)
VALUES
    (1, 1, 3, 'completed', 830, 50, 780, 'card', 'paid'),
    (2, 2, 4, 'served', 900, 0, 900, 'cash', 'paid'),
    (3, 3, 3, 'paid', 650, 65, 585, 'card', 'paid'),
    (1, 4, 4, 'ready', 920, 0, 920, 'pending', 'pending');

INSERT INTO payments (order_id, amount, method, status, transaction_id, paid_at)
VALUES
    (1, 780, 'card', 'completed', 'txn_001', CURRENT_TIMESTAMP - INTERVAL '2 hours'),
    (2, 900, 'cash', 'completed', 'txn_002', CURRENT_TIMESTAMP - INTERVAL '1 hour'),
    (3, 585, 'card', 'completed', 'txn_003', CURRENT_TIMESTAMP - INTERVAL '30 minutes');