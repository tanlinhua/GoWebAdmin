-- 管理员表
DROP TABLE IF EXISTS `go_admin`;
CREATE TABLE `go_admin` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
  `user_name` varchar(32) NOT NULL COMMENT '登录名',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `role` int(4) DEFAULT 0 COMMENT '1管理/2客服',
  `status` int(4) NOT NULL DEFAULT 0 COMMENT '状态:0禁用/1启用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `last_login_ip` varchar(20) DEFAULT NULL COMMENT '最后登录IP',
  UNIQUE KEY `user_name` (`user_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='管理员表';

INSERT INTO `go_admin` VALUES ('1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '1', '1', '2020-12-12 00:00:00', '2020-12-12 00:00:00', null, null);
INSERT INTO `go_admin` VALUES ('2', 'kefu01', 'e10adc3949ba59abbe56e057f20f883e', '2', '1', '2020-12-12 00:00:00', '2020-12-12 00:00:00', null, null);

-- 用户表
DROP TABLE IF EXISTS `go_user`;
CREATE TABLE `go_user` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `kf_id` int(11) NOT NULL DEFAULT 0 COMMENT '客服ID',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号码',
  `user_name` varchar(60) DEFAULT NULL COMMENT '账号',
  `nick_name` varchar(50) DEFAULT NULL COMMENT '昵称',
  `password` varchar(60) DEFAULT NULL COMMENT '密码',
  `pay_code` varchar(10) DEFAULT NULL COMMENT '支付密码',
  `status` int(4) NOT NULL DEFAULT '0' COMMENT '状态:0禁用/1启用',
  `money` int(11) DEFAULT '0' COMMENT '账户金额',
  `header` varchar(255) DEFAULT NULL COMMENT '头像',
  `token` varchar(300) DEFAULT NULL,
  `device_type` int(4) DEFAULT NULL COMMENT '设备类型:1安卓,2苹果',
  `device_info` varchar(255) DEFAULT NULL COMMENT '设备详情',
  `device_key` varchar(255) DEFAULT NULL COMMENT '唯一设备码',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `reg_ip` varchar(20) DEFAULT NULL COMMENT '注册ip',
  `last_login_ip` varchar(20) DEFAULT NULL COMMENT '最后登录IP',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY `user_name` (`user_name`) USING BTREE,
  UNIQUE KEY `phone` (`phone`) USING BTREE,
  KEY `kf_id` (`kf_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- 系统配置表
DROP TABLE IF EXISTS `go_sys_params`;
CREATE TABLE `go_sys_params` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `type` int(4) NOT NULL DEFAULT 0 COMMENT '类型:0后台不可编辑/1可编辑',
  `key` varchar(100) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `remarks` varchar(255) DEFAULT NULL,
  UNIQUE KEY `key` (`key`) USING BTREE
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='系统配置表';