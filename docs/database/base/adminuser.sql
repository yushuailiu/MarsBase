CREATE TABLE `admin_user_info`(
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_id` INT NOT NULL COMMENT '角色id',
  `username` VARCHAR(256) NOT NULL COMMENT '用户名',
  `nickname` varchar(256) NOT NULL default '' COMMENT '昵称',
  `password` VARCHAR(128) NOT NULL COMMENT '密码',
  `email` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '电话',
  `avatar` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '头像',
  `intro` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '个人介绍',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态: 正常或禁止等',
  `last_ip` INT NOT NULL DEFAULT 0 COMMENT '上次登陆ip',
  `sex` TINYINT NOT NULL DEFAULT 0 COMMENT '0:保密 1:男 2:女',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`),
  unique key username (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `admin_user_role`(
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` VARCHAR(64) NOT NULL COMMENT '角色名',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态 0:可使用 1:禁止使用',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `admin_user_role` (`id`, `name`)
VALUES
  (1, 'master'),
  (2, 'manager'),
  (3, 'user');

INSERT INTO `admin_user_info` (`role_id`, `username`, `nickname`, `password`)
VALUES (1, 'demo', 'demo', 'ICFpDErnPEgTHXFqZXNCBHn51o9W/ZDbeRE4UU5OSaztNbJDMVKBg0iN0OsfKo7yLijFKbsf42lOS+58Ojiluw==');