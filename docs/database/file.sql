create table if not exists `app_site_file`(
  `id` int not null auto_increment comment '自增id',
  `file_source` varchar(64) not null comment '文件来源',
  `file_type` varchar(64) not null comment '文件类型 图片等！',
  `user_id` int not null comment '上传者',
  `size` int not null comment '文件大小',
  `file_path` varchar(256) not null comment '文件路径',
  `updated_at` datetime null default null on update current_timestamp comment '更新时间',
  `created_at` datetime not null default current_timestamp comment '创建时间',
  primary key (`id`)
)engine=InnoDb default charset=utf8;