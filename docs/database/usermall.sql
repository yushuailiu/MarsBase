CREATE TABLE `user_info`(
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `address_id` INT NOT NULL DEFAULT  0 COMMENT '默认地址id',
  `nickname` VARCHAR(256) NOT NULL COMMENT '昵称',
  `email` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '电话',
  `avatar` VARCHAR(1024) NOT NULL COMMENT '头像',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态: 正常或禁止等',
  `last_ip` INT NOT NULL DEFAULT 0 COMMENT '上次登陆ip',
  `intro` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '个人介绍',
  `sex` TINYINT NOT NULL DEFAULT 0 COMMENT '0:保密 1:男 2:女',
  `city` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '城市',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `user_auth`(
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` INT NOT NULL COMMENT '用户id',
  `login_type` VARCHAR(256) NOT NULL COMMENT '登录类型: username, weixin, weibo 等',
  `identifier` VARCHAR(256) NOT NULL COMMENT '登录标识:用户名或则微信账号等',
  `credential` VARCHAR(256) NULL DEFAULT '' COMMENT  '登录密码或token等',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态 0:可使用 1:禁止使用',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;