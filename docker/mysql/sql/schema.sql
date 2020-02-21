CREATE TABLE IF NOT EXISTS `test_articles` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(32) BINARY UNIQUE,
    `title` VARCHAR(32),
    PRIMARY KEY (`id`)
);

INSERT INTO test_articles(name, title) VALUE('admin', 'torimo');
