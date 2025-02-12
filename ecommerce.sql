USE simidb;

CREATE TABLE IF NOT EXISTS Store (
    Store_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Name varchar(60) NOT NULL,
    Logo varchar(100) NOT NULL,
    Banner varchar(100) NOT NULL,
    Address varchar(100),
    Cellphone varchar(20)
);

CREATE TABLE IF NOT EXISTS Social_media (
    Social_media_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Name varchar(60),
    Url varchar(200),
    Store_id INT UNSIGNED,
    FOREIGN KEY (Store_id) REFERENCES Store(Store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Payment_method (
    Payment_method_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Type varchar(60),
    CVU varchar(40),
    Alias varchar(40),
    Store_id INT UNSIGNED,
    FOREIGN KEY (Store_id) REFERENCES Store(Store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Category (
     Category_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
     Category_name varchar(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS Discount (
    Discount_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Type varchar(20) NOT NULL, #percent|fixed|installments
    Amount int NOT NULL
);

CREATE TABLE IF NOT EXISTS Product (
    Product_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Name varchar(45) NOT NULL,
    Main_description varchar(45) NOT NULL,
    Long_description varchar(500),
    Price int NOT NULL,
    Category_id INT UNSIGNED NOT NULL,
    Store_id INT UNSIGNED NOT NULL,
    Discount_id INT UNSIGNED,
    FOREIGN KEY (Category_id) REFERENCES Category(Category_id) ON DELETE CASCADE,
    FOREIGN KEY (Store_id) REFERENCES Store(Store_id) ON DELETE CASCADE,
    FOREIGN KEY (Discount_id) REFERENCES Discount(Discount_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Size (
    Size_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Size VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS Color (
     Color_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
     Color VARCHAR(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS Product_detail (
    Product_detail_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Product_id INT UNSIGNED,
    Size_id INT UNSIGNED,
    Color_id INT UNSIGNED,
    Stock_count INT NOT NULL,
    FOREIGN KEY (Product_id) REFERENCES Product(Product_id),
    FOREIGN KEY (Size_id) REFERENCES Size(Size_id),
    FOREIGN KEY (Color_id) REFERENCES Color(Color_id)
);

CREATE TABLE IF NOT EXISTS Product_image (
    Product_image_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Url varchar(100) NOT NULL,
    Product_id INT UNSIGNED,
    FOREIGN KEY (Product_id) REFERENCES Product(Product_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS User (
    User_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Name varchar(45) NOT NULL,
    Email varchar(45) NOT NULL UNIQUE,
    Cellphone varchar(45) NOT NULL,
    Address varchar(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS Product_order (
    Order_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    User_id INT UNSIGNED,
    FOREIGN KEY (User_id) REFERENCES User(User_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Product_order_detail (
    Product_order_detail_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Count int NOT NULL,
    Product_id INT UNSIGNED,
    Order_id INT UNSIGNED,
    FOREIGN KEY (Product_id) REFERENCES Product(Product_id) ON DELETE CASCADE,
    FOREIGN KEY (Order_id) REFERENCES Product_order(Order_id) ON DELETE CASCADE
);

INSERT INTO Store
(name, address, logo, banner, cellphone)
VALUES ('Nike', 'Alto palermo', '/logo-nike.png', '/banner-nike.png', '1165696294');

INSERT INTO Social_media
(Name, Url, Store_id)
VALUES ('Instagram','https://www.instagram.com/nike?igsh=MXJkaDh0NHdvNmE4Nw==', 1);

INSERT INTO Social_media
(Name, Url, Store_id)
VALUES ('Twitter','https://x.com/nike', 1);

INSERT INTO Category
(category_name)
VALUES ('Zapatillas');

INSERT INTO Category
(category_name)
VALUES ('Ropa');

INSERT INTO Category
(category_name)
VALUES ('Accesorios');

INSERT INTO Size
(Size_id, Size)
VALUES (1, 'S');

INSERT INTO Size
(Size_id, Size)
VALUES (2, 'M');

INSERT INTO Size
(Size_id, Size)
VALUES (3, 'L');

INSERT INTO Size
(Size_id, Size)
VALUES (4, 'XL');

INSERT INTO Size
(Size_id, Size)
VALUES (5, '40');

INSERT INTO Size
(Size_id, Size)
VALUES (6, '41');

INSERT INTO Size
(Size_id, Size)
VALUES (7, '42');

INSERT INTO Size
(Size_id, Size)
VALUES (8, 'UNIQUE');

INSERT INTO Color
(Color_id, Color)
VALUES (1, 'Red');

INSERT INTO Color
(Color_id, Color)
VALUES (2, 'Green');

INSERT INTO Color
(Color_id, Color)
VALUES (3, 'Blue');

INSERT INTO Color
(Color_id, Color)
VALUES (4, 'Black');

INSERT INTO Color
(Color_id, Color)
VALUES (5, 'White');

INSERT INTO Product
(name, main_description, price, category_id, store_id, discount_id)
VALUES ('Nike jordan low', 'Zapatilla nike de cuero', 250, 1, 1, null);

INSERT INTO Product
(name, main_description, price, category_id, store_id, discount_id)
VALUES ('Nike AIR', 'Zapatilla nike de cuero', 185, 1, 1, null);

INSERT INTO Product
(name, description, price, category_id, store_id, discount_id)
VALUES ('Remera Nike jordan', 'Remera nike de algodon', 65, 2, 1, null);

INSERT INTO Product
(name, description, price, category_id, store_id, discount_id)
VALUES ('Mochila Nike', 'Mochila nike de tela', 110, 3, 1, null);

INSERT INTO Product
(name, description, price, count, category_id, store_id, discount_id)
VALUES ('Medias Nike', 'Pack de medias Nike x 3', 35, 10, 3, 1, null);

INSERT INTO Product_image
(url, product_id)
VALUES ('/nike-jordan-low.jpg', 1);

INSERT INTO Product_image
(url, product_id)
VALUES ('/nike-air.jpg', 2);

INSERT INTO Product_image
(url, product_id)
VALUES ('/reme-nike-jordan.jpg', 3);

INSERT INTO Product_image
(url, product_id)
VALUES ('/mochila-nike.jpg', 4);

INSERT INTO Product_image
(url, product_id)
VALUES ('/medias-nike.jpg', 5);

INSERT INTO Product_detail
(Product_id, Size_id, Color_id, Stock_count)
VALUES (1, 5, 4, 10);

INSERT INTO Product_detail
(Product_id, Size_id, Color_id, Stock_count)
VALUES (2, 6, 5, 10);

INSERT INTO Product_detail
(Product_id, Size_id, Color_id, Stock_count)
VALUES (3, 3, 4, 10);

INSERT INTO Product_detail
(Product_id, Size_id, Color_id, Stock_count)
VALUES (4, 5, 4, 10);

INSERT INTO Product_detail
(Product_id, Size_id, Color_id, Stock_count)
VALUES (5, 8, 4, 10);
