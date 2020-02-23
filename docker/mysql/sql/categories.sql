CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `create_user` int NOT NULL DEFAULT 1,
  `is_deleted` boolean DEFAULT 0,
  `create_at` datetime NOT NULL,
  `delete_at` datetime,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO
categories
(`name`,
 `create_user`,
 `create_at`)
VALUES
('test_tag', '1', now());