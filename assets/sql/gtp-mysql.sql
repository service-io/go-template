-- create a table for gtp
drop table if exists `gtp`;
create table `gtp`
(
  `id`          bigint       not null auto_increment,
  `title`       varchar(255) not null comment '主题',
  `description` varchar(1000)         default null comment '描述',
  `start_at`    timestamp    not null default current_timestamp comment '开始时间',
  `end_at`      timestamp    not null default current_timestamp comment '结束时间',
  `status`      tinyint      not null default 0 comment '0: 发布 1: 暂存 2: 已结束 3: 已失效',
  `top`         tinyint      not null default 0 comment '0: 不置顶 1: 置顶',
  `creator_by`  bigint                default null comment '创建人员 ID',
  `updater_by`  bigint                default null comment '更新人员 ID',
  `created_at`  timestamp             default current_timestamp comment '创建时间',
  `updated_at`  timestamp             default current_timestamp on update current_timestamp comment '更新时间',
  primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='gtp主表';
