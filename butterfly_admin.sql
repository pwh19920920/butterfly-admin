/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50735
 Source Host           : localhost:3306
 Source Schema         : butterfly_admin

 Target Server Type    : MySQL
 Target Server Version : 50735
 File Encoding         : 65001

 Date: 20/01/2022 17:39:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_menu`;
CREATE TABLE `t_sys_menu`  (
  `id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单路径',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单图标',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单组件',
  `sort` int(10) NOT NULL DEFAULT 0 COMMENT '菜单排序',
  `option` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单操作',
  `parent` bigint(20) NOT NULL DEFAULT 0 COMMENT '上级目录',
  `deleted` tinyint(1) NOT NULL DEFAULT 0,
  `route` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由路径',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单代码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_menu
-- ----------------------------
INSERT INTO `t_sys_menu` VALUES (1332302770434215918, '2020-11-30 20:38:28', '2021-11-30 20:03:05', '系统管理', '/sys', 'crown', '', 200, '[]', 0, 0, '/1332302770434215918', 'sys');
INSERT INTO `t_sys_menu` VALUES (1332302770434215920, '2021-10-15 06:58:22', '2022-01-20 17:35:11', '菜单管理', '/sys/sysMenu', 'smile', './SysMenu', 1, '[]', 1332302770434215918, 0, '/1332302770434215918/1332302770434215920', 'sysMenu');
INSERT INTO `t_sys_menu` VALUES (1332302770434215922, '2021-10-13 02:11:28', '2022-01-20 17:23:40', '用户管理', '/sys/sysUser', 'smile', './SysUser', 2, '[]', 1332302770434215918, 0, '/1332302770434215918/1332302770434215922', 'sysUser');
INSERT INTO `t_sys_menu` VALUES (1332302770434215924, '2021-10-12 18:11:56', '2022-01-20 17:34:58', '角色管理', '/sys/sysRole', 'smile', './SysRole', 1, '[{\"id\":405,\"name\":\"xxx\",\"value\":\"xxxx\",\"method\":\"POST\",\"path\":\"/test\"}]', 1332302770434215918, 0, '/1332302770434215918/1332302770434215924', 'sysRole');

-- ----------------------------
-- Table structure for t_sys_menu_option
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_menu_option`;
CREATE TABLE `t_sys_menu_option`  (
  `id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint(1) NOT NULL DEFAULT 0,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限串',
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'URL方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'URL路径',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '唯一编码',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unq_code`(`code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单操作表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_menu_option
-- ----------------------------
INSERT INTO `t_sys_menu_option` VALUES (1447759564626726912, '2021-10-12 18:47:33', '2022-01-20 17:35:10', 0, '菜单查看', 'sys:menu:query', 'GET', '/api/sys/menu', 'caa126a343b0e1cef0774b637c246af3', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1448238719118086145, '2021-10-12 19:04:17', '2022-01-20 17:35:10', 0, '菜单新增', 'sys:menu:create', 'POST', '/api/sys/menu', '79102b6efd1174afdf1732d9e7e80629', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1463883940220637184, '2021-10-15 07:17:21', '2022-01-20 17:23:40', 0, '用户查询', 'sys:user:query', 'GET', '/api/sys/user', '0924d00bac6e4d1b9e10040e095a980f', 1332302770434215922);
INSERT INTO `t_sys_menu_option` VALUES (1463883940220637185, '2021-10-15 07:17:21', '2022-01-20 17:23:40', 0, '用户修改', 'sys:user:modify', 'PUT', '/api/sys/user', 'e75e13959a8f6577fee78e1bc61d3e10', 1332302770434215922);
INSERT INTO `t_sys_menu_option` VALUES (1463883940220637186, '2021-10-15 07:17:21', '2022-01-20 17:23:40', 0, '用户创建', 'sys:menu:create', 'POST', '/api/sys/user', '80b0eae883b924868c44df5295b8ee33', 1332302770434215922);
INSERT INTO `t_sys_menu_option` VALUES (1484094198515765282, '2021-10-15 07:20:37', '2022-01-20 17:34:58', 0, '角色查询', 'sys:role:query', 'GET', '/api/sys/role', '92cd13408fbd6512a4e5328c800d5439', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1484094198515765283, '2021-10-15 07:20:37', '2022-01-20 17:34:58', 0, '角色创建', 'sys:role:create', 'POST', '/api/sys/role', '0f1f61330f3ef4b03bf5632bbcc5737f', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1484094198515765284, '2021-10-15 07:20:37', '2022-01-20 17:34:58', 0, '角色修改', 'sys:role:modify', 'PUT', '/api/sys/role', 'f83fd0b3bd1902b67676b1a64a78b309', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1484094198515765285, '2021-10-15 07:20:37', '2022-01-20 17:34:58', 0, '角色删除', 'sys:role:delete', 'DELETE', '/api/sys/role/:id', '0f14a64e7dff47937164f64dfa01dbbf', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1484094198515765286, '2021-10-15 07:20:37', '2022-01-20 17:34:58', 0, '角色权限查询', 'sys:role:queryPermission', 'GET', '/api/sys/role/permission/:roleId', '381b1a498606e82226a0604d8c853e65', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1484097160667467832, '2021-10-15 07:16:02', '2021-10-15 07:16:02', 0, '菜单修改', 'sys:menu:modify', 'PUT', '/api/sys/menu', '6a7b949b19c27a1f9ee3753e06c3ecf5', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1484097160667467833, '2021-10-15 07:16:02', '2021-10-15 07:16:02', 0, '菜单删除', 'sys:menu:delete', 'DELETE', '/api/sys/menu/:id', 'c249795688bf6e62fe7b16ba1d539540', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1484097160667467834, '2021-10-15 07:16:02', '2021-10-15 07:16:02', 0, '菜单操作', 'sys:menu:option', 'GET', '/api/sys/menu/option/:id', 'dfb98e82d0666d936314879ab3cbe37d', 1332302770434215920);

-- ----------------------------
-- Table structure for t_sys_permission
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_permission`;
CREATE TABLE `t_sys_permission`  (
  `id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单',
  `role_id` bigint(20) NOT NULL COMMENT '角色',
  `option` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作',
  `deleted` tinyint(1) NOT NULL DEFAULT 0,
  `independent` tinyint(2) NOT NULL DEFAULT 0 COMMENT '是否独立',
  `half` tinyint(2) NOT NULL DEFAULT 0 COMMENT '是否虚拟选中',
  `root` tinyint(2) NOT NULL COMMENT '是否为跟',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_permission
-- ----------------------------
INSERT INTO `t_sys_permission` VALUES (1484097040798453802, '2022-01-21 01:22:35', '2022-01-21 01:22:35', 1332302770434215922, 1, '', 0, 1, 0, 0);
INSERT INTO `t_sys_permission` VALUES (1484097040798453803, '2022-01-21 01:22:35', '2022-01-20 17:38:16', 1332302770434215920, 1, '', 0, 1, 0, 0);
INSERT INTO `t_sys_permission` VALUES (1484097040798453804, '2022-01-21 01:22:35', '2022-01-21 01:22:35', 1332302770434215924, 1, '', 0, 1, 0, 0);
INSERT INTO `t_sys_permission` VALUES (1484097040798453805, '2022-01-21 01:22:35', '2022-01-21 01:22:35', 1332302770434215918, 1, '', 0, 1, 0, 1);
INSERT INTO `t_sys_permission` VALUES (1484097040798453806, '2022-01-21 01:22:35', '2022-01-20 17:38:19', 1332302770434215920, 1, '1447759564626726912,1448238719118086145,1484097160667467832,1484097160667467833,1484097160667467834', 0, 0, 0, 0);
INSERT INTO `t_sys_permission` VALUES (1484097040798453807, '2022-01-21 01:22:35', '2022-01-21 01:34:33', 1332302770434215924, 1, '1484094198515765282,1484094198515765283,1484094198515765284,1484094198515765285,1484094198515765286', 0, 0, 0, 0);
INSERT INTO `t_sys_permission` VALUES (1484097040798453808, '2022-01-21 01:22:35', '2022-01-21 01:22:35', 1332302770434215922, 1, '1463883940220637184,1463883940220637185,1463883940220637186,1463883940220637187', 0, 0, 0, 0);

-- ----------------------------
-- Table structure for t_sys_role
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_role`;
CREATE TABLE `t_sys_role`  (
  `id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `deleted` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_role
-- ----------------------------
INSERT INTO `t_sys_role` VALUES (1, '2021-11-06 20:51:40', '2022-01-20 17:34:42', 'super_admin', 0);

-- ----------------------------
-- Table structure for t_sys_token
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_token`;
CREATE TABLE `t_sys_token`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `subject` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 187 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统令牌表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_token
-- ----------------------------
INSERT INTO `t_sys_token` VALUES (186, '2022-01-20 17:05:46', '2022-01-20 17:05:46', '270182e4-c946-41b4-b72e-a7cf41c968a0', 1, '1dad828a-e9f9-4a5c-9dfe-25279be12b9c', 0);

-- ----------------------------
-- Table structure for t_sys_user
-- ----------------------------
DROP TABLE IF EXISTS `t_sys_user`;
CREATE TABLE `t_sys_user`  (
  `id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码盐',
  `deleted` tinyint(1) NOT NULL DEFAULT 0,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `roles` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniq_username`(`username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_user
-- ----------------------------
INSERT INTO `t_sys_user` VALUES (1, '2020-11-24 15:49:07', '2021-11-29 14:58:35', 'admin', '593d4632a8c70251d0e9be4b1799bcc1', '54099a65-a235-158c-d610-74d2ff4c789b', 0, '王小二', 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png', '1', 'pengweihuang@we.cn', '18650036719');

SET FOREIGN_KEY_CHECKS = 1;
