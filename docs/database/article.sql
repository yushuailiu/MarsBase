CREATE TABLE `article_category` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` VARCHAR(128) NOT NULL COMMENT '栏目名称',
  `del` TINYINT NOT NULL DEFAULT 0 COMMENT '标识栏目是否为删除状态',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


create table if not exists `article_content`(
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `title` VARCHAR(256) NOT NULL COMMENT '文章标题',
  `content` TEXT CHARSET utf8mb4 NOT NULL COMMENT '文章内容',
  `category_id` INT NOT NULL COMMENT '分类id',
  `intro` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '文章简介',
  `author` INT NOT NULL COMMENT '作者',
  `hits_num` INT NOT NULL DEFAULT 0 COMMENT '点击数',
  `votes_num` INT NOT NULL DEFAULT 0 COMMENT '点赞数',
  `comments_num` INT NOT NULL DEFAULT 0 COMMENT '评论数',
  `ontop` TINYINT NOT NULL DEFAULT 0 COMMENT '是否置顶',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态 0:草稿 1:待审核 2:发布 3:删除',
  `extra` TEXT NULL DEFAULT NULL COMMENT '额外信息，关联图片信息，自定义字段等',
  `publish_date` DATETIME NULL DEFAULT NULL COMMENT '发表时间',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


create table if not exists `article_label` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` VARCHAR(32) NOT NULL COMMENT '标签名称',
  `del` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除 0:否 1:是',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

create table if not exists `article_label_map` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `label_id` INT NOT NULL COMMENT '标签名称',
  `article_id` INT NOT NULL DEFAULT 0 COMMENT '是否删除 0:否 1:是',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;