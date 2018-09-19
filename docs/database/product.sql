-- ----------------------------
-- Table structure for brand
-- ----------------------------
CREATE TABLE `product_brand` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '品牌id',
  `name` VARCHAR(50) NOT NULL COMMENT '品牌名',
  `item_id` INT UNSIGNED NOT NULL COMMENT '分类id',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of brand
-- ----------------------------
INSERT INTO `product_brand` VALUES ('1', '李宁', '2');
INSERT INTO `product_brand` VALUES ('2', '耐克', '2');
INSERT INTO `product_brand` VALUES ('3', '李宁', '4');
INSERT INTO `product_brand` VALUES ('4', '耐克', '4');
INSERT INTO `product_brand` VALUES ('5', '红蜻蜓', '1');
INSERT INTO `product_brand` VALUES ('6', '东方骆驼', '4');
INSERT INTO `product_brand` VALUES ('7', '婷美', '4');
INSERT INTO `product_brand` VALUES ('8', 'Chanel', '5');
INSERT INTO `product_brand` VALUES ('9', 'CoCo', '7');
INSERT INTO `product_brand` VALUES ('10', 'Amani', '8');

-- ----------------------------
-- Table structure for item
-- ----------------------------
CREATE TABLE `product_item` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `name` VARCHAR(50) NOT NULL COMMENT '分类名',
  `parent_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父类id',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of item
-- ----------------------------
INSERT INTO `product_item` VALUES ('1', '服装/鞋包', 0);
INSERT INTO `product_item` VALUES ('2', '男装', '1');
INSERT INTO `product_item` VALUES ('3', '流行男鞋', '1');
INSERT INTO `product_item` VALUES ('4', '女装', '1');
INSERT INTO `product_item` VALUES ('5', '箱包', 0);
INSERT INTO `product_item` VALUES ('6', '双肩包', '5');
INSERT INTO `product_item` VALUES ('7', '单肩包', '5');
INSERT INTO `product_item` VALUES ('8', '行李箱', '5');
INSERT INTO `product_item` VALUES ('9', '其他', 0);

-- ----------------------------
-- 产品表
-- ----------------------------
CREATE TABLE `product_info` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '品牌id',
  `name` VARCHAR(50) NOT NULL COMMENT '产品名',
  `brand_id` INT NOT NULL COMMENT '品牌id',
  `item_id` INT NOT NULL COMMENT '分类id',
  `postage` INT NOT NULL DEFAULT 0 COMMENT '邮费',
  `buy_count` INT NOT NULL DEFAULT 0 COMMENT '已售数量',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '商品状态 在架 下架',
  `detail` longtext NOT NULL DEFAULT NULL COMMENT '商品详情',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES ('1', '球鞋', '1', 1, '1');
INSERT INTO `product` VALUES ('2', '网球', '1', 2, '9');
INSERT INTO `product` VALUES ('3', '衬衫', '1', 1, '1');
INSERT INTO `product` VALUES ('4', '袜子', '1', 1, '1');
INSERT INTO `product` VALUES ('5', '球鞋', '2', 2, '1');
INSERT INTO `product` VALUES ('6', 'Air', '2', 2, '1');
INSERT INTO `product` VALUES ('7', '袜子', '2', 2, '1');
INSERT INTO `product` VALUES ('8', '乒乓球', '2', 1, '9');
INSERT INTO `product` VALUES ('9', '高跟鞋', '5', 1, '1');
INSERT INTO `product` VALUES ('10', '水晶鞋', '5', 2, '1');
INSERT INTO `product` VALUES ('11', '凉鞋', '7', 2, '1');
INSERT INTO `product` VALUES ('12', '皮鞋', '7', 2, '1');
INSERT INTO `product` VALUES ('13', '高跟鞋', '7', 2, '1');
INSERT INTO `product` VALUES ('14', 'Air1', '4', '6', '1');
INSERT INTO `product` VALUES ('15', 'Air2', '4', '6', '1');
INSERT INTO `product` VALUES ('16', 'Air3', '4', '6', '1');

-- ----------------------------
-- 产品属性map表(产品选择所包含的属性)
-- ----------------------------
CREATE TABLE `product_property_value_map` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `product_id` INT UNSIGNED NOT NULL COMMENT '产品id',
  `property_id` INT UNSIGNED NOT NULL COMMENT '属性id',
  `property_value_id` INT UNSIGNED NOT NULL COMMENT '属性值id',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of product_pro
-- ----------------------------
INSERT INTO `product_property_value_map` VALUES ('9', '1', '1', '1');
INSERT INTO `product_property_value_map` VALUES ('10', '1', '2', '1');
INSERT INTO `product_property_value_map` VALUES ('11', '1', '3', '2');
INSERT INTO `product_property_value_map` VALUES ('12', '1', '4', '2');
INSERT INTO `product_property_value_map` VALUES ('13', '1', '5', '2');
INSERT INTO `product_property_value_map` VALUES ('14', '1', '6', '3');
INSERT INTO `product_property_value_map` VALUES ('15', '1', '7', '3');
INSERT INTO `product_property_value_map` VALUES ('16', '1', '8', '3');

-- ----------------------------
-- Table structure for product_sku
-- ----------------------------
CREATE TABLE `product_sku` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `product_id` INT NOT NULL COMMENT '产品id',
  `number` INT NOT NULL DEFAULT 0 COMMENT '库存量',
  `price` decimal(10,4) DEFAULT NULL,
  `name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'sku名称 默认可以是产品名属性拼接',
  `properties` varchar(300) NOT NULL DEFAULT '' COMMENT 'product_properties 1;2;3',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of product_sku
-- ----------------------------
INSERT INTO `product_sku` VALUES ('1', '1', '53', '528.5000', '李宁-A79-12球鞋（红，亚麻，北美）', '1;3;6;');
INSERT INTO `product_sku` VALUES ('2', '1', '23', '238.9000', '李宁-A79-12球鞋（绿，塑料，中国）', '2;4;8;');
INSERT INTO `product_sku` VALUES ('3', '1', '10', '370.0000', '李宁-A21-11球鞋（绿，针织，南非）', '2;5;7;');
INSERT INTO `product_sku` VALUES ('4', '2', '12', '123.3000', '李宁-B1-12网球（红，针织，南非）', '1;5;7;');
INSERT INTO `product_sku` VALUES ('5', '2', '19', '250.0000', '李宁-B3-18网球（绿，针织，南非）', '2;5;7;');
INSERT INTO `product_sku` VALUES ('6', '3', '12', '200.0000', '李宁-A-12衬衫（绿，针织，南非）', '2;5;7;');
INSERT INTO `product_sku` VALUES ('7', '4', '10', '10.0000', '李宁-C-12袜子（绿，针织，南非）', '2;5;7;');
INSERT INTO `product_sku` VALUES ('9', '5', '5', '498.5000', '耐克-A79-12球鞋（绿，亚麻，北美）', '2;3;6;');
INSERT INTO `product_sku` VALUES ('10', '5', '25', '498.5000', '耐克-A79-12球鞋（绿，塑料，北美）', '2;4;6;');
INSERT INTO `product_sku` VALUES ('11', '5', '20', '498.5000', '耐克-A79-12球鞋（绿，塑料，南非）', '2;4;7;');
INSERT INTO `product_sku` VALUES ('13', '14', '1', '238.9000', '耐克Air1-A102-23球鞋（绿，亚麻，中国）', '2;3;8;');
INSERT INTO `product_sku` VALUES ('14', '14', '2', '250.9000', '耐克Air1-A102-23球鞋（绿，针织，北美）', '2;5;6;');
INSERT INTO `product_sku` VALUES ('15', '15', '22', '200.9000', '耐克Air2-A102-23球鞋（绿，针织，南非）', '2;5;7;');
INSERT INTO `product_sku` VALUES ('16', '16', '12', '200.9000', '耐克Air3-A21-11球鞋（红，亚麻，北美）', '1;3;6;');
INSERT INTO `product_sku` VALUES ('17', '12', '12', '200.9000', '婷美-A21-11凉鞋（红，亚麻，南非）', '1;3;7;');
INSERT INTO `product_sku` VALUES ('18', '12', '12', '238.9000', '婷美-A21-11凉鞋（红，亚麻，中国）', '1;3;8;');
INSERT INTO `product_sku` VALUES ('19', '12', '12', '320.9000', '婷美-A21-11凉鞋（红，塑料，南非）', '2;4;8;');

-- ----------------------------
-- 属性表
-- ----------------------------
CREATE TABLE `product_property` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` VARCHAR(32) NOT NULL COMMENT '属性名',
  `item_id` INT NOT NULL COMMENT '分类id',
  `is_color` TINYINT NOT NULL DEFAULT 0 COMMENT '是否是颜色属性',
  `is_enum` TINYINT NOT NULL DEFAULT 0 COMMENT '是否是枚举: 建立属性的时候指明可取值',
  `is_input` TINYINT NOT NULL DEFAULT 0 COMMENT '是否是输入属性',
  `is_sale` TINYINT NOT NULL DEFAULT 0 COMMENT '是否是销售属性：可sku的属性，非销售属性只做展示等，比如是否保修方式等',
  `is_must` TINYINT NOT NULL DEFAULT 0 COMMENT '是否是必须属性',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of pro_name
-- ----------------------------
INSERT INTO `product_property` VALUES ('1', '颜色', '1', '0', '1', '0', '0', '0');
INSERT INTO `product_property` VALUES ('2', '材质', '1', '0', '0', '0', '0', '0');
INSERT INTO `product_property` VALUES ('3', '厂商', '1', '0', '0', '0', '0', '0');
INSERT INTO `product_property` VALUES ('4', '尺码', '1', '0', '0', '0', '0', '0');
INSERT INTO `product_property` VALUES ('5', '面值', '5', '0', '0', '0', '0', '0');
INSERT INTO `product_property` VALUES ('6', '渠道', '5', '0', '0', '0', '0', '0');
INSERT INTO `product_property` VALUES ('15', '鞋跟', '3', '0', '0', '0', '0', '0');

-- ----------------------------
-- Table structure for pro_value
-- ----------------------------
CREATE TABLE `product_property_value` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `value` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '属性值，可保存多种类型的值',
  `property_id` INT NOT NULL COMMENT '属性id',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of pro_value
-- ----------------------------
INSERT INTO `product_property_value` VALUES ('1', '红', '1');
INSERT INTO `product_property_value` VALUES ('2', '绿', '1');
INSERT INTO `product_property_value` VALUES ('3', '亚麻', '2');
INSERT INTO `product_property_value` VALUES ('4', '塑料', '2');
INSERT INTO `product_property_value` VALUES ('5', '针织', '2');
INSERT INTO `product_property_value` VALUES ('6', '北美工厂店', '3');
INSERT INTO `product_property_value` VALUES ('7', '南非工厂店', '3');
INSERT INTO `product_property_value` VALUES ('8', '中国制造', '3');
INSERT INTO `product_property_value` VALUES ('9', '高脚鞋跟', '15');
INSERT INTO `product_property_value` VALUES ('10', '平底鞋跟', '15');
INSERT INTO `product_property_value` VALUES ('11', '尖顶鞋跟', '15');



CREATE TABLE `product_img` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `path` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '图片地址',
  `product_id` INT NOT NULL COMMENT '产品id',
  `is_cover` TINYINT NOT NULL DEFAULT 0 COMMENT '是否是主图',
  `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;