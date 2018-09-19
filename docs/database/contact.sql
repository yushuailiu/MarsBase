CREATE TABLE `user_contact` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '收件人',
  `phone` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '电话',
  `province` INT NOT NULL COMMENT '省份',
  `city` INT NOT NULL COMMENT '城市',
  `street` INT NOT NULL COMMENT '街道',
  `postcode` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '邮编',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;