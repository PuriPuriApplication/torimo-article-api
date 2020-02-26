CREATE TABLE `follow_users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `from_user` bigint NOT NULL,
  `to_user` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO follow_users VALUES ('1', '1', '2');