# Go Faker Fixtures

## Getting Started

#### Download

```shell
go get -u github.com/guiyomh/go-faker-fixtures
```
## Example

### 1. Create a Database

```sql
CREATE TABLE `order` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `seller_id` INT NULL,
  `customer_id` INT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `seller` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(60) NULL,
  `last_name` VARCHAR(65) NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `customer` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(60) NULL,
  `last_name` VARCHAR(65) NULL,
  `email` VARCHAR(95) NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `phone` (
  `number` VARCHAR(30) NOT NULL,
  `user_id` INT NOT NULL);

CREATE TABLE `address` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `street_address` VARCHAR(255) NULL,
  `city` VARCHAR(45) NULL,
  `zip_code` VARCHAR(20) NULL,
  `user_id` INT NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `pseudo` VARCHAR(45) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `isAdmin` TINYINT NOT NULL DEFAULT 0,
  `password` VARCHAR(65) NOT NULL,
  PRIMARY KEY (`id`));
```

### 2. Create fixtures
Create a folder with one or many yaml files

```yaml
---
order:
    order_{1..5}:
        seller_id: '@seller_<Current()>'
        customer_id: '@customer_<Current()>'

seller:
    seller_{1..5}:
        first_name: '<FirstName()>'
        last_name: '<LastName()>'

customer:
    customer_tpl (template):
        first_name: '<FirstName()>'
        last_name: '<LastName()>'
    customer_{1..10} (extends customer_tpl):
        email : '<Email()>'
phone:
    phone_{bob,harry,george}:
        number: '<PhoneNumber()>'
        user_id: '@user_<Current()>'
    phone_2_{bob,george}:
        number: '<PhoneNumber()>'
        user_id: '@user_<Current()>'
address:
    address_tpl (template):
        street_address: '<StreetAddress()>'
        city: '<City()>'
        zip_code: '<PostCode()>'
    address_{bob,harry,george} (extends address_tpl):
        user_id: '@user_<Current()>'
    address_{1..2} (extends address_tpl):
        user_id: '@admin_<Current()>'
user:
    user_tpl (template):
        first_name: '<FirstName()>'
        last_name: '<LastName()>'
        pseudo: '<UserName()>'
        password: '<Words(2,true)>'
        email : '<Email()>'
    admin_1:
        first_name: 'William'
        last_name: 'Wallace'
        pseudo: 'WW'
        password: 'freedommmmmmm'
        email : 'freedom@gouv.co.uk'
        isAdmin: true
    admin_{2..5} (extends user_tpl):
        isAdmin: true
    user_{bob,harry,george} (extends user_tpl):
        isAdmin: false
```
### 3. Run the commande

```bash
go-faker-fixtures load --fixtures ./fixtures --user=<your_db_user> --dbname=<your_dbname> --pass=<your_db_pass>
```