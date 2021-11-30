CREATE TABLE `web-service-gin`.`items`
( `id` INT NOT NULL AUTO_INCREMENT , `title` VARCHAR(2500) NOT NULL , `price` FLOAT NOT NULL ,
`date_created` DATE NOT NULL , `date_updated` DATE NOT NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB;

INSERT INTO `items` (`title`, `price`, `date_created`) VALUES
('Harry Potter and the Philosopher’s Stone', '102.34', '2021-11-29');

SELECT * FROM `items`