CREATE TABLE `order_info` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `pay_date` DATETIME NULL DEFAULT NULL COMMENT '付款时间',
  `user_id` INT NOT NULL COMMENT '买家id',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '订单状态',
  `total_price` decimal(10,4) NOT NULL DEFAULT 0 COMMENT '订单总价',
  `pay_price` decimal(10,4) NOT NULL DEFAULT 0 COMMENT '实际付款',
  `freight` DECIMAL(10,4) NOT NULL DEFAULT 0 COMMENT '邮费',
  `remark` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '备注',
  `address` VARCHAR(128) NOT NULL COMMENT '地址',
  `address_id` INT NOT NULL NULL COMMENT '地址id',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `order_detail` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `order_id` INT NOT NULL COMMENT '父订单id',
  `product_name` VARCHAR(256) NOT NULL COMMENT '商品标题',
  `total_price` DECIMAL(10,4) NOT NULL DEFAULT 0 COMMENT '价格',
  `num` INT NOT NULL DEFAULT 0 COMMENT '商品数量',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;