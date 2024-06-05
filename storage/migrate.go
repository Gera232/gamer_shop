package storage

var (
	sentenceMigrate = `CREATE TABLE IF NOT EXISTS card (
	card_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	number VARCHAR(100) NOT NULL,
	security_code VARCHAR(50) NOT NULL,
	due_date VARCHAR(50) NOT NULL,
	name_owner VARCHAR(100) NOT NULL);
	
	CREATE TABLE IF NOT EXISTS province (
	province_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	postal_code INT NOT NULL);
	
	CREATE TABLE IF NOT EXISTS location (
	location_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	province_id INT NOT NULL,
	FOREIGN KEY (province_id) REFERENCES province(province_id));
	
	CREATE TABLE IF NOT EXISTS address (
	address_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	place ENUM('house','work') NOT NULL,
	street VARCHAR(50) NOT NULL,
	height VARCHAR(50) NOT NULL,
	floor INT,
	department INT,
	tower INT,
	between_streets VARCHAR(80) NOT NULL,
	observations VARCHAR(40),
	shipment BOOLEAN NOT NULL,
	province_id INT NOT NULL,
	FOREIGN KEY (province_id) REFERENCES province(province_id));
	
	CREATE TABLE IF NOT EXISTS account (
	account_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	surname VARCHAR(20) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password VARCHAR(200) NOT NULL,
	role ENUM('user','admin','worker') NOT NULL,
	address_id INT,
	FOREIGN KEY (address_id) REFERENCES address(address_id),
	card_id INT,
	FOREIGN KEY (card_id) REFERENCES card(card_id));
	
	CREATE TABLE IF NOT EXISTS trolley (
	trolley_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	account_id INT NOT NULL,
	FOREIGN KEY (account_id) REFERENCES account(account_id));
	
	CREATE TABLE IF NOT EXISTS categorie (
	categorie_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(30) NOT NULL);
	
	CREATE TABLE IF NOT EXISTS product (
	product_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	sku VARCHAR(50) NOT NULL,
	price DOUBLE NOT NULL,
	warranty INT NOT NULL,
	stock INT NOT NULL,
	shipments BOOLEAN NOT NULL,
	categorie_id INT NOT NULL,
	FOREIGN KEY (categorie_id) REFERENCES categorie(categorie_id));
	
	CREATE TABLE IF NOT EXISTS ask (
	ask_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	ask VARCHAR(80),
	answer VARCHAR(80),
	product_id INT NOT NULL,
	FOREIGN KEY (product_id) REFERENCES product(product_id));
	
	CREATE TABLE IF NOT EXISTS product_trolley (
	Product_trolley_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	amount INT NOT NULL,
	trolley_id INT NOT NULL,
	product_id INT NOT NULL,
	FOREIGN KEY (trolley_id) REFERENCES trolley(trolley_id),
	FOREIGN KEY (product_id) REFERENCES product(product_id));
	
	CREATE TABLE IF NOT EXISTS phone_number (
	phone_number_id  INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	area_code INT NOT NULL,
	phone_number VARCHAR(100) NOT NULL,
	account_id INT NOT NULL,
	FOREIGN KEY (account_id) REFERENCES account(account_id));`
)

func Migrate() error {
	stmt, err := db.Prepare(sentenceMigrate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
