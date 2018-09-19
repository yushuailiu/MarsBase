create table if not exists `app_comment_info`(
  `id` int not null auto_increment comment '自增id',
  `body` varchar(1024) charset utf8mb4 not null comment '评论内容',
  `bodyHtml` varchar(2014) charset utf8mb4 not null DEFAULT '' comment '评论内容html结构',
  `user_id` int not null comment '用户id',
  `business` varchar(128) not null comment '业务类型',
  `content_id` int not null comment '业务id',
  `parent_id` int not null default 0 comment '评论父级评论id',
  `status` tinyint not null default 0 comment '状态 0:上线 1:下线',
  `updated_at` datetime null default null on update current_timestamp comment '更新时间',
  `created_at` datetime not null default current_timestamp comment '创建时间',
  primary key (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;