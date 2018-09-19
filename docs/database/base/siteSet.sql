CREATE TABLE `site_set` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `site_name` VARCHAR(64) NOT NULL COMMENT '站点名',
  `meta_keywords` VARCHAR(1024) NOT NULL COMMENT '角色id',
  `meta_description` VARCHAR(1024) NOT NULL COMMENT '站点描述',
  `site_version`  VARCHAR(16) NOT NULL DEFAULT 0 COMMENT '网站版本',
  `avatar` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '站点logo',
  `favicon` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT 'favicon',
  `copyright` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT 'Copyright',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
);