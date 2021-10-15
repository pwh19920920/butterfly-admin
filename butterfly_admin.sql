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

 Date: 15/10/2021 15:33:05
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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_menu
-- ----------------------------
INSERT INTO `t_sys_menu` VALUES (1332302770434215918, '2020-11-30 12:38:28', '2021-10-14 04:31:24', '系统管理', '/sys', 'crown', '', 2, '[]', 0, 0, '/1332302770434215918', 'sys');
INSERT INTO `t_sys_menu` VALUES (1332302770434215920, '2021-10-14 22:58:22', '2021-10-14 23:16:02', '菜单管理', '/sys/sysMenu', 'smile', './SysMenu', 1, '[]', 1332302770434215918, 0, '/1332302770434215918/1332302770434215920', 'sysMenu');
INSERT INTO `t_sys_menu` VALUES (1332302770434215922, '2021-10-12 10:11:28', '2021-10-14 23:17:21', '用户管理', '/sys/sysUser', 'smile', './SysUser', 2, '[]', 1332302770434215918, 0, '/1332302770434215918/1332302770434215922', 'sysUser');
INSERT INTO `t_sys_menu` VALUES (1332302770434215924, '2021-10-12 02:11:56', '2021-10-14 23:20:37', '角色管理', '/sys/sysRole', 'smile', './SysRole', 1, '[{\"id\":405,\"name\":\"xxx\",\"value\":\"xxxx\",\"method\":\"POST\",\"path\":\"/test\"}]', 1332302770434215918, 0, '/1332302770434215918/1332302770434215924', 'sysRole');
INSERT INTO `t_sys_menu` VALUES (1332302770434215926, '2021-10-09 10:12:47', '2021-10-14 04:31:49', '监控管理', '/job', 'smile', '', 1, '[]', 0, 0, '/1332302770434215926', 'job');
INSERT INTO `t_sys_menu` VALUES (1332302770434215928, '2021-10-11 10:13:06', '2021-10-14 23:21:59', '数据源管理', '/job/database', 'smile', './JobDatabase', 2, '[]', 1332302770434215926, 0, '/1332302770434215926/1332302770434215928', 'jobDatabase');
INSERT INTO `t_sys_menu` VALUES (1332302770434215930, '2021-10-11 02:13:51', '2021-10-13 15:44:57', '任务管理', '/job/task', 'smile', './SysRole', 1, '[]', 1332302770434215926, 0, '/1332302770434215926/1332302770434215930', '');

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_menu_option
-- ----------------------------
INSERT INTO `t_sys_menu_option` VALUES (1447759564626726912, '2021-10-12 18:47:33', '2021-10-14 15:16:01', 0, '菜单查看', 'sys:menu:query', 'GET', '/api/sys/menu', 'caa126a343b0e1cef0774b637c246af3', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1447760000133894145, '2021-10-12 11:04:17', '2021-10-14 23:16:02', 1, '菜单新增', 'sys:menu:create', 'POST', '/api/sys/menu', '8094790fdce7d048ebcf43979339aae0', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1448238719118086145, '2021-10-12 19:04:17', '2021-10-14 15:16:01', 0, '菜单新增', 'sys:menu:create', 'POST', '/api/sys/menu', '79102b6efd1174afdf1732d9e7e80629', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1448668927012900866, '2021-10-14 23:16:02', '2021-10-14 23:16:02', 0, '菜单修改', 'sys:menu:modify', 'PUT', '/api/sys/menu', '6b3b68431579eaf6b3d4a69a0ec18b08', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1448668927017095168, '2021-10-14 23:16:02', '2021-10-14 23:16:02', 0, '菜单删除', 'sys:menu:delete', 'DELETE', '/api/sys/menu/:id', 'dca2d4b95c306fbd21a73e38e82ff007', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1448668927017095169, '2021-10-14 23:16:02', '2021-10-14 23:16:02', 0, '菜单操作', 'sys:menu:option', 'GET', '/api/sys/menu/option/:id', '315a218f5d7b8f961d41604facedab91', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1448668927017095170, '2021-10-14 23:16:02', '2021-10-14 23:16:02', 0, '菜单获取', 'sys:menu:queryWithOption', 'GET', '/api/sys/menu/withOption', '9ed6bf710b5b9802534783571be07050', 1332302770434215920);
INSERT INTO `t_sys_menu_option` VALUES (1448669259440852992, '2021-10-14 23:17:21', '2021-10-14 23:17:21', 0, '用户查询', 'sys:user:query', 'GET', '/api/sys/user', 'c4b0789ebee8c17c0ff07829eec8670a', 1332302770434215922);
INSERT INTO `t_sys_menu_option` VALUES (1448669259440852993, '2021-10-14 23:17:21', '2021-10-14 23:17:21', 0, '用户修改', 'sys:user:modify', 'PUT', '/api/sys/user', '2b6a72c2eeb70feb99b4cd0b65e954eb', 1332302770434215922);
INSERT INTO `t_sys_menu_option` VALUES (1448669259440852994, '2021-10-14 23:17:21', '2021-10-14 23:17:21', 0, '用户创建', 'sys:menu:create', 'POST', '/api/sys/user', 'afbf478129b34727685739d3ca5d606a', 1332302770434215922);
INSERT INTO `t_sys_menu_option` VALUES (1448670082216497152, '2021-10-14 23:20:37', '2021-10-14 23:20:37', 0, '角色查询', 'sys:role:query', 'GET', '/api/sys/role', '442aa2720fc50a72a35b53cb5a5695eb', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1448670082216497153, '2021-10-14 23:20:37', '2021-10-14 23:20:37', 0, '角色创建', 'sys:role:create', 'POST', '/api/sys/role', 'c422e4a0129da9c465ae422280bf8838', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1448670082216497154, '2021-10-14 23:20:37', '2021-10-14 23:20:37', 0, '角色修改', 'sys:role:modify', 'PUT', '/api/sys/role', '6c200c9d08b929494efea5790f3e5fa0', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1448670082216497155, '2021-10-14 23:20:37', '2021-10-14 23:20:37', 0, '角色删除', 'sys:role:delete', 'DELETE', '/api/sys/role/:id', 'f5ea9ec30eead23392954aa9d5759183', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1448670082216497156, '2021-10-14 23:20:37', '2021-10-14 23:20:37', 0, '查询全部角色', 'sys:role:queryAll', 'GET', '/api/sys/role/all', 'd75094205cf8cbb9927e08b0cc24be84', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1448670082216497157, '2021-10-14 23:20:37', '2021-10-14 23:20:37', 0, '角色权限查询', 'sys:role:queryPermission', 'GET', '/api/sys/role/permission/:roleId', '65ce3dd61cfdb0f3f3b4e7bfafb59a37', 1332302770434215924);
INSERT INTO `t_sys_menu_option` VALUES (1448670425583194112, '2021-10-14 23:21:59', '2021-10-14 23:21:59', 0, '数据源查询', 'job:database:query', 'GET', '/api/job/database', 'e865145bec3c73a01decb991c0be6418', 1332302770434215928);
INSERT INTO `t_sys_menu_option` VALUES (1448670425583194113, '2021-10-14 23:21:59', '2021-10-14 23:21:59', 0, '数据源查看', 'job:database:create', 'POST', '/api/job/database', '4dcb9ec7ce7f6ea5f69ee7e2549d44f3', 1332302770434215928);
INSERT INTO `t_sys_menu_option` VALUES (1448670425583194114, '2021-10-14 23:21:59', '2021-10-14 23:21:59', 0, '数据源更新', 'job:database:modify', 'PUT', '/api/job/database', '17c32e7b6eadfa6a170890ccf66aa7d5', 1332302770434215928);

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_permission
-- ----------------------------
INSERT INTO `t_sys_permission` VALUES (1448914621250408448, '2021-10-15 07:25:00', '2021-10-15 07:25:00', 1332302770434215922, 1, '', 0, 1, 0);
INSERT INTO `t_sys_permission` VALUES (1448914621250408449, '2021-10-15 07:25:00', '2021-10-15 07:25:00', 1332302770434215918, 1, '', 0, 1, 0);
INSERT INTO `t_sys_permission` VALUES (1448914621250408450, '2021-10-15 07:25:00', '2021-10-15 07:25:00', 1332302770434215920, 1, '', 0, 1, 0);
INSERT INTO `t_sys_permission` VALUES (1448914621250408451, '2021-10-15 07:25:00', '2021-10-15 07:25:00', 1332302770434215924, 1, '', 0, 1, 0);
INSERT INTO `t_sys_permission` VALUES (1448914621250408452, '2021-10-15 07:25:00', '2021-10-15 07:25:00', 1332302770434215920, 1, '1447759564626726912,1448238719118086145,1448668927012900866,1448668927017095168,1448668927017095169,1448668927017095170', 0, 0, 0);
INSERT INTO `t_sys_permission` VALUES (1448914621250408453, '2021-10-15 07:25:00', '2021-10-15 07:25:00', 1332302770434215922, 1, '1448669259440852992,1448669259440852993,1448669259440852994', 0, 0, 0);
INSERT INTO `t_sys_permission` VALUES (1448914621250408454, '2021-10-15 07:25:00', '2021-10-15 07:25:00', 1332302770434215924, 1, '1448670082216497152,1448670082216497153,1448670082216497154,1448670082216497155,1448670082216497156,1448670082216497157', 0, 0, 0);

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_role
-- ----------------------------
INSERT INTO `t_sys_role` VALUES (1, '2021-10-18 20:51:40', '2021-10-15 15:32:20', 'super_admin', 0);
INSERT INTO `t_sys_role` VALUES (1447459031953182720, '2021-10-11 15:08:20', '2021-10-11 15:26:45', 'test', 1);
INSERT INTO `t_sys_role` VALUES (1447459092053364736, '2021-10-15 07:08:35', '2021-10-11 15:26:42', 'xx', 1);

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
) ENGINE = InnoDB AUTO_INCREMENT = 102 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


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
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unq_username`(`username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_sys_user
-- ----------------------------
INSERT INTO `t_sys_user` VALUES (1, '2020-11-24 07:49:07', '2021-10-12 22:55:22', 'admin', '593d4632a8c70251d0e9be4b1799bcc1', '54099a65-a235-158c-d610-74d2ff4c789b', 0, '王小二', 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png', '1');

SET FOREIGN_KEY_CHECKS = 1;
