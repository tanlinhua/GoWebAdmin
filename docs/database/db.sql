-- 管理员表
DROP TABLE IF EXISTS `go_admin`;
CREATE TABLE `go_admin` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
  `user_name` varchar(32) NOT NULL COMMENT '登录名',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `role` int(11) DEFAULT 0 COMMENT '角色ID',
  `pid` int(11) DEFAULT 0 COMMENT '上级ID',
  `status` int(4) NOT NULL DEFAULT 0 COMMENT '状态:0禁用/1启用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `last_login_ip` varchar(20) DEFAULT NULL COMMENT '最后登录IP',
  UNIQUE KEY `user_name` (`user_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

INSERT INTO `go_admin` VALUES ('1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '0', '0', '1', '2020-12-12 00:00:00', '2020-12-12 00:00:00', null, null);

-- 角色表
DROP TABLE IF EXISTS `go_role`;
CREATE TABLE `go_role` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `role_name` varchar(40) NOT NULL COMMENT '角色名称',
  `role_desc` varchar(40) DEFAULT NULL COMMENT '角色描述',
  `per_id` varchar(255) DEFAULT NULL COMMENT '权限ids: 1,2,5'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- 角色权限表
DROP TABLE IF EXISTS `go_permission`;
CREATE TABLE `go_permission` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(30) NOT NULL COMMENT '权限名称',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父id',
  `uri` varchar(50) NOT NULL DEFAULT '' COMMENT 'API路由',
  `method` varchar(10) NOT NULL DEFAULT '' COMMENT '路由请求方法(GET/POST)',
  `icon` varchar(10) DEFAULT NULL COMMENT '主菜单图标',
  `level` int(11) NOT NULL DEFAULT '1' COMMENT '权限等级[1,2,3]'
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COMMENT='权限表';

INSERT INTO `go_permission` VALUES ('1', '主页', '0', '/admin/main', 'GET', 'xe68e;', '1');
INSERT INTO `go_permission` VALUES ('2', '控制台', '1', '/admin/console', 'GET', null, '2');
INSERT INTO `go_permission` VALUES ('3', '修改密码', '2', '/admin/cpw', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('4', '权限配置', '0', '', '', 'xe672;', '1');
INSERT INTO `go_permission` VALUES ('5', '用户管理', '4', '/admin/adm/view', 'GET', null, '2');
INSERT INTO `go_permission` VALUES ('6', '增加用户', '5', '/admin/adm/add', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('7', '删除用户', '5', '/admin/adm/del', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('8', '修改用户', '5', '/admin/adm/update', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('9', '查询用户', '5', '/admin/adm/get', 'GET', null, '3');
INSERT INTO `go_permission` VALUES ('10', '角色管理', '4', '/admin/role/view', 'GET', null, '2');
INSERT INTO `go_permission` VALUES ('11', '增加角色', '10', '/admin/role/add', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('12', '删除角色', '10', '/admin/role/del', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('13', '修改角色', '10', '/admin/role/update', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('14', '查询角色', '10', '/admin/role/get', 'GET', null, '3');
INSERT INTO `go_permission` VALUES ('15', '权限管理', '4', '/admin/per/view', 'GET', null, '2');
INSERT INTO `go_permission` VALUES ('16', '查询权限', '15', '/admin/per/get', 'GET', null, '3');
INSERT INTO `go_permission` VALUES ('17', '系统配置', '0', '', '', 'xe716;', '1');
INSERT INTO `go_permission` VALUES ('18', '系统参数', '17', '/admin/params/view', 'GET', null, '2');
INSERT INTO `go_permission` VALUES ('19', '增加参数', '18', '/admin/params/add', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('20', '删除参数', '18', '/admin/params/del', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('21', '修改参数', '18', '/admin/params/update', 'POST', null, '3');
INSERT INTO `go_permission` VALUES ('22', '查询参数', '18', '/admin/params/get', 'GET', null, '3');
INSERT INTO `go_permission` VALUES ('23', '操作日志', '17', '/admin/adminlog/view', 'GET', null, '2');
INSERT INTO `go_permission` VALUES ('24', '查看日志', '23', '/admin/adminlog/get', 'GET', null, '3');

-- 管理员日志表
DROP TABLE IF EXISTS `go_admin_log`;
CREATE TABLE `go_admin_log` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '操作者ID',
  `uri` varchar(255) DEFAULT NULL COMMENT '对应资源',
  `title` varchar(255) DEFAULT '' COMMENT '日志标题',
  `body` varchar(1024) NOT NULL COMMENT '日志内容',
  `ip` varchar(50) DEFAULT '' COMMENT 'IP',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  KEY `uid` (`uid`) USING BTREE
  -- FOREIGN KEY (`uid`) REFERENCES go_admin(`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='管理员日志表';

-- 系统配置表
DROP TABLE IF EXISTS `go_sys_params`;
CREATE TABLE `go_sys_params` (
  `id` int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `type` int(4) NOT NULL DEFAULT 0 COMMENT '类型:0后台不可编辑/1可编辑',
  `key` varchar(100) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `remarks` varchar(255) DEFAULT NULL,
  UNIQUE KEY `key` (`key`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='系统配置表';

-- 用户表
DROP TABLE IF EXISTS `go_user`;
CREATE TABLE `go_user` (
  `id` int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `kf_id` int unsigned NOT NULL DEFAULT 0 COMMENT '客服ID',
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
