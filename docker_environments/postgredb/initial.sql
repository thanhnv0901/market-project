
-- https://mysql.tutorials24x7.com/blog/guide-to-design-database-for-shopping-cart-in-mysql
DROP TABLE IF EXISTS "companies";
CREATE TABLE IF NOT EXISTS "companies"(
	id bigserial PRIMARY KEY,
	name VARCHAR(100) DEFAULT '',
	tax_number VARCHAR(50) DEFAULT '',
	address VARCHAR(200) DEFAULT '',
	country VARCHAR(10) DEFAULT '',
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null

);

DROP TABLE IF EXISTS "products";
CREATE TABLE IF NOT EXISTS "products"(
	id bigserial PRIMARY KEY,
	name VARCHAR(100) DEFAULT '',
	quantity BIGINT DEFAULT 0,
	unit VARCHAR(20) DEFAULT 'peace',
	price DOUBLE PRECISION DEFAULT 0.0,
	price_unit VARCHAR(20) DEFAULT 'dollar',
	user_id bigserial NOT NULL,
	company_id BIGINT DEFAULT -1,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null
);

DROP TABLE IF EXISTS "users";
CREATE TABLE IF NOT EXISTS "users"(
	id bigserial primary key,
	name VARCHAR(100) DEFAULT '',
	phone VARCHAR(50) DEFAULT '',
	country VARCHAR(10) DEFAULT '',
	username VARCHAR(50) NOT NULL,
	password_hash VARCHAR(50) NOT NULL,
	active BOOLEAN DEFAULT true,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null,
	lasttime_login TIMESTAMP DEFAULT now()
);

DROP TABLE IF EXISTS "orders";
CREATE TABLE IF NOT EXISTS "orders"(
	id bigserial PRIMARY KEY,
	discount DOUBLE PRECISION DEFAULT 0.0,
	total_price	DOUBLE PRECISION DEFAULT 0.0,
	name VARCHAR(100) DEFAULT '',
	phone VARCHAR(20) DEFAULT '',
	address VARCHAR(100) DEFAULT '',
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT  null,
	notice TEXT
);

DROP TABLE IF EXISTS "order_products";
CREATE TABLE IF NOT EXISTS "order_products"(
	product_id bigserial NOT NULL,
	order_id bigserial NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null,
	quantity SMALLINT DEFAULT 0,
	price DOUBLE PRECISION DEFAULT 0.0,
	discount DOUBLE PRECISION DEFAULT 0.0
);

DROP TABLE IF EXISTS "transactions";
CREATE TABLE IF NOT EXISTS "transactions"(
	id bigserial PRIMARY KEY,
	user_id bigserial NOT NULL,
	order_id bigserial NOT NULL,
	code VARCHAR(50),
	status SMALLINT DEFAULT 0,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null
);

DROP TABLE IF EXISTS "carts";
CREATE TABLE IF NOT EXISTS "carts"(
	id bigserial PRIMARY KEY,
	user_id bigserial NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null

);

DROP TABLE IF EXISTS "cart_products";
CREATE TABLE IF NOT EXISTS "cart_products"(
	cart_id bigserial NOT NULL,
	product_id bigserial NOT NULL,
	price DOUBLE PRECISION DEFAULT 0.0,
	discount DOUBLE PRECISION DEFAULT 0.0,
	quantity SMALLINT DEFAULT 0,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null
);

DROP TABLE IF EXISTS "product_metas";
CREATE TABLE IF NOT EXISTS "product_metas"(
	id smallserial PRIMARY KEY,
	product_id bigserial NOT NULL,
	key VARCHAR(50) DEFAULT '',
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null
);

DROP TABLE IF EXISTS "categories";
CREATE TABLE IF NOT EXISTS "categorys"(
	id smallserial PRIMARY KEY,
	title TEXT DEFAULT '',
	description TEXT DEFAULT '',
	CONTENT TEXT DEFAULT '',
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null
);

DROP TABLE IF EXISTS "product_categories";
CREATE TABLE IF NOT EXISTS "product_categorys"(
	product_id bigserial NOT NULL,
	category_id smallserial NOT NULL,
	title TEXT DEFAULT '',
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP DEFAULT null
);






-- ############


INSERT INTO companies (  name, tax_number, address, country) VALUES
('Routine', '0132968XA', '145 Quang Trung, Phuong 10, Go Vap District', 'vn'),
('Couple TX', '5619942YT', '150 Quang Trung, Phường 10, Go Vap District', 'vn'),
('Nike', '1295495QW', 'Tang 3, Diamond Plaza, 34, Duong Le Duan, Ben Nge, District 1', 'vn'),
('Adidas', '4956135KH', '77 Thích Bửu Đăng, Phường 1, Gò Vấp', 'vn'),
('Apple', '3459213QW', '488 Lý Thái Tổ, Phường 10, Quận 10, Thành phố Hồ Chí Minh', 'vn');

INSERT INTO "users"( "name", phone, country, username, password_hash) VALUES
('Hana', '15852135910', 'us', 'userhana', '123456'),
('John', '14842918707', 'us', 'userjohn', '123456'),
('James', '14842918962', 'us', 'userjames', '123456'),
('Robert', '18143008665', 'us', 'userrobert', '123456'),
('David', '15852135910', 'us', 'userdavid', '123456');

INSERT INTO products( "name",quantity, unit, price, price_unit, user_id, company_id ) VALUES
('Trousers', 100, 'peace',50,'dollar',1,1),
('T-Shirt', 100, 'peace',30,'dollar',2,2),
('Snakers', 100, 'pair',100,'dollar',3,3),
('Hat', 100, 'peace',20,'dollar',4,4),
('Apple Watch', 100, 'peace',300,'dollar',5,5);




