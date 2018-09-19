CREATE TABLE `admin_menu` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` VARCHAR(128) NOT NULL COMMENT '目录名',
  `parent_name` VARCHAR(128) NOT NULL COMMENT '上级目录id',
  `app` VARCHAR(128) NOT NULL COMMENT '对应APP名称',
  `page` VARCHAR(128) NOT NULL COMMENT '页面名',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态: 正常或禁止等',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
);

CREATE TABLE `admin_menu_map` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `menu_id` INT NOT NULL COMMENT '目录id',
  `role_id` INT NOT NULL COMMENT '角色id',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态: 正常或禁止等',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
);

INSERT INTO `admin_menu` (`name`, `app`, `page`, `parent_name`)
VALUES
('Dashboard', 'Dashboard', 'Statistics', 'Dashboard');