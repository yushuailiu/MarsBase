create table if not exists `app_social_vote`(
  `id` int not null auto_increment comment '自增id',
  `user_id` int not null comment '用户id',
  `content_id` int not null comment '文章',
  `business` varchar(128) not null comment '业务类型',
  `status` tinyint not null default 0 comment '状态 0:上线 1:下线',
  `updated_at` datetime null default null on update current_timestamp comment '更新时间',
  `created_at` datetime not null default current_timestamp comment '创建时间',
  primary key (`id`),
  key business_id (`business`, `content_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;