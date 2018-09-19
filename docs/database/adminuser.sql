CREATE TABLE `admin_user_info`(
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_id` INT NOT NULL COMMENT '角色id',
  `username` VARCHAR(256) NOT NULL COMMENT '昵称',
  `password` VARCHAR(128) NOT NULL COMMENT '密码',
  `email` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '电话',
  `avatar` VARCHAR(1024) NOT NULL COMMENT '头像',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态: 正常或禁止等',
  `last_ip` INT NOT NULL DEFAULT 0 COMMENT '上次登陆ip',
  `sex` TINYINT NOT NULL DEFAULT 0 COMMENT '0:保密 1:男 2:女',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `admin_user_role`(
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` INT NOT NULL COMMENT '角色名',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态 0:可使用 1:禁止使用',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;