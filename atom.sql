-- ----------------------------
-- 用户表
-- ----------------------------
CREATE TABLE `user` (
	`user_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '用户id',
	`user_type` VARCHAR(50) NOT NULL DEFAULT 'USER' COMMENT '用户类型：USER-普通用户；SUPER_ADMIN-超级管理员' COLLATE 'utf8mb4_0900_ai_ci',
	`username` VARCHAR(50) NOT NULL COMMENT '用户名' COLLATE 'utf8mb4_0900_ai_ci',
	`password` VARCHAR(1000) NOT NULL COMMENT '密码' COLLATE 'utf8mb4_0900_ai_ci',
	`nickname` VARCHAR(50) NOT NULL COMMENT '昵称' COLLATE 'utf8mb4_0900_ai_ci',
	`avatar` VARCHAR(255) NULL DEFAULT NULL COMMENT '头像' COLLATE 'utf8mb4_0900_ai_ci',
	`email` VARCHAR(50) NULL DEFAULT NULL COMMENT '邮箱' COLLATE 'utf8mb4_0900_ai_ci',
	`phone` VARCHAR(11) NULL DEFAULT NULL COMMENT '手机号' COLLATE 'utf8mb4_0900_ai_ci',
	`gender` CHAR(1) NOT NULL DEFAULT '0' COMMENT '性别：0-未知；1-男；2-女' COLLATE 'utf8mb4_0900_ai_ci',
	`login_last_ip` VARCHAR(50) NULL DEFAULT NULL COMMENT '最后一次登录ip' COLLATE 'utf8mb4_0900_ai_ci',
	`login_last_time` DATETIME NULL DEFAULT NULL COMMENT '最后一次登录时间',
	`status` CHAR(1) NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用' COLLATE 'utf8mb4_0900_ai_ci',
	`create_by` VARCHAR(50) NULL DEFAULT NULL COMMENT '创建人' COLLATE 'utf8mb4_0900_ai_ci',
	`create_time` DATETIME NULL DEFAULT NULL COMMENT '创建时间',
	`update_by` VARCHAR(50) NULL DEFAULT NULL COMMENT '更新人' COLLATE 'utf8mb4_0900_ai_ci',
	`update_time` DATETIME NULL DEFAULT NULL COMMENT '更新时间',
	`delete_time` DATETIME NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`user_id`) USING BTREE
)
COMMENT='用户表'
COLLATE='utf8mb4_0900_ai_ci'
ENGINE=InnoDB;

-- ----------------------------
-- 初始化-用户表数据
-- ----------------------------
INSERT INTO `user` (`user_id`, `user_type`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `gender`, `login_last_ip`, `login_last_time`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `delete_time`) VALUES (1, 'SUPER_ADMIN', 'admin', '$2a$10$e271Fk3XyghRsKzwAy66du.PtNDUiBC7XF0Za1PfhDy0TEHUGQw2m', '超级管理员', '', '', '', '0', '', NULL, '0', 'system', '2024-11-26 11:31:12', '', '2024-11-26 11:31:12', NULL);

-- ----------------------------
-- 角色表
-- ----------------------------
CREATE TABLE `role` (
	`role_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '角色id',
	`role_name` VARCHAR(50) NOT NULL COMMENT '角色名称' COLLATE 'utf8mb4_0900_ai_ci',
	`status` CHAR(1) NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用' COLLATE 'utf8mb4_0900_ai_ci',
	`create_by` VARCHAR(50) NULL DEFAULT NULL COMMENT '创建人' COLLATE 'utf8mb4_0900_ai_ci',
	`create_time` DATETIME NULL DEFAULT NULL COMMENT '创建时间',
	`update_by` VARCHAR(50) NULL DEFAULT NULL COMMENT '更新人' COLLATE 'utf8mb4_0900_ai_ci',
	`update_time` DATETIME NULL DEFAULT NULL COMMENT '更新时间',
	`delete_time` DATETIME NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`role_id`) USING BTREE
)
COMMENT='角色表'
COLLATE='utf8mb4_0900_ai_ci'
ENGINE=InnoDB;

-- ----------------------------
-- 菜单权限表
-- ----------------------------
CREATE TABLE `menu` (
	`menu_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '菜单id',
	`menu_name` VARCHAR(50) NOT NULL COMMENT '菜单名称' COLLATE 'utf8mb4_0900_ai_ci',
	`parent_id` INT(10) NOT NULL DEFAULT '0' COMMENT '父菜单id',
	`menu_type` CHAR(1) NOT NULL COMMENT '菜单类型：M-目录；C-菜单；B-按钮' COLLATE 'utf8mb4_0900_ai_ci',
	`sort` INT(10) NOT NULL DEFAULT '0' COMMENT '排序',
	`component` VARCHAR(255) NULL DEFAULT NULL COMMENT '组件' COLLATE 'utf8mb4_0900_ai_ci',
	`icon` VARCHAR(255) NULL DEFAULT NULL COMMENT '菜单图标' COLLATE 'utf8mb4_0900_ai_ci',
	`path` VARCHAR(255) NULL DEFAULT NULL COMMENT '请求地址路径' COLLATE 'utf8mb4_0900_ai_ci',
	`method` VARCHAR(10) NULL DEFAULT NULL COMMENT '请求方式：GET|POST|PUT|DELETE' COLLATE 'utf8mb4_0900_ai_ci',
	`status` CHAR(1) NOT NULL DEFAULT '0' COMMENT '状态：0-正常；1-停用' COLLATE 'utf8mb4_0900_ai_ci',
	`create_by` VARCHAR(50) NULL DEFAULT NULL COMMENT '创建人' COLLATE 'utf8mb4_0900_ai_ci',
	`create_time` DATETIME NULL DEFAULT NULL COMMENT '创建时间',
	`update_by` VARCHAR(50) NULL DEFAULT NULL COMMENT '更新人' COLLATE 'utf8mb4_0900_ai_ci',
	`update_time` DATETIME NULL DEFAULT NULL COMMENT '更新时间',
	`delete_time` DATETIME NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`menu_id`) USING BTREE
)
COMMENT='菜单权限表'
COLLATE='utf8mb4_0900_ai_ci'
ENGINE=InnoDB;

-- ----------------------------
-- 用户角色表
-- ----------------------------
CREATE TABLE `user_role` (
	`user_id` INT(10) NOT NULL COMMENT '用户id',
	`role_id` INT(10) NOT NULL COMMENT '角色id',
	PRIMARY KEY (`user_id`, `role_id`) USING BTREE
)
COMMENT='用户角色表'
COLLATE='utf8mb4_0900_ai_ci'
ENGINE=InnoDB;

-- ----------------------------
-- 角色菜单权限表
-- ----------------------------
CREATE TABLE `role_menu` (
	`role_id` INT(10) NOT NULL COMMENT '角色id',
	`menu_id` INT(10) NOT NULL COMMENT '菜单权限id',
	PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
)
COMMENT='角色菜单权限表'
COLLATE='utf8mb4_0900_ai_ci'
ENGINE=InnoDB;

-- ----------------------------
-- 登录日志表
-- ----------------------------
CREATE TABLE `login_log` (
	`login_log_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '登录日志id',
	`login_username` VARCHAR(50) NOT NULL COMMENT '登录用户名' COLLATE 'utf8mb4_0900_ai_ci',
	`login_ip` VARCHAR(50) NOT NULL COMMENT '登录ip地址' COLLATE 'utf8mb4_0900_ai_ci',
	`login_location` VARCHAR(255) NOT NULL COMMENT '登录地点' COLLATE 'utf8mb4_0900_ai_ci',
	`browser` VARCHAR(50) NOT NULL COMMENT '浏览器类型' COLLATE 'utf8mb4_0900_ai_ci',
	`os` VARCHAR(50) NOT NULL COMMENT '系统' COLLATE 'utf8mb4_0900_ai_ci',
	`status` CHAR(1) NOT NULL DEFAULT '0' COMMENT '登录状态：0-成功；1-失败' COLLATE 'utf8mb4_0900_ai_ci',
	`message` VARCHAR(255) NOT NULL COMMENT '提示消息' COLLATE 'utf8mb4_0900_ai_ci',
	`login_time` DATETIME NULL DEFAULT NULL COMMENT '登陆时间',
	PRIMARY KEY (`login_log_id`) USING BTREE
)
COMMENT='登录日志表'
COLLATE='utf8mb4_0900_ai_ci'
ENGINE=InnoDB;

-- ----------------------------
-- 操作日志表
-- ----------------------------
CREATE TABLE `operation_log` (
	`operation_log_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '操作日志id',
	`path` VARCHAR(255) NOT NULL COMMENT '请求地址路径' COLLATE 'utf8mb4_0900_ai_ci',
	`method` VARCHAR(10) NOT NULL COMMENT '请求方式' COLLATE 'utf8mb4_0900_ai_ci',
	`param` VARCHAR(2000) NOT NULL DEFAULT '' COMMENT '请求参数' COLLATE 'utf8mb4_0900_ai_ci',
	`result` VARCHAR(2000) NOT NULL DEFAULT '' COMMENT '请求结果' COLLATE 'utf8mb4_0900_ai_ci',
	`operation_ip` VARCHAR(50) NOT NULL COMMENT '操作主机' COLLATE 'utf8mb4_0900_ai_ci',
	`operation_location` VARCHAR(255) NOT NULL COMMENT '操作地点' COLLATE 'utf8mb4_0900_ai_ci',
	`operation_name` VARCHAR(50) NOT NULL COMMENT '操作人员' COLLATE 'utf8mb4_0900_ai_ci',
	`operation_time` DATETIME NULL DEFAULT NULL COMMENT '操作时间',
	PRIMARY KEY (`operation_log_id`) USING BTREE
)
COMMENT='操作日志表'
COLLATE='utf8mb4_0900_ai_ci'
ENGINE=InnoDB;
