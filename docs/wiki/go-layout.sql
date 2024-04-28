/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80035 (8.0.35)
 Source Host           : 127.0.0.1:3306
 Source Schema         : go-layout

 Target Server Type    : MySQL
 Target Server Version : 80035 (8.0.35)
 File Encoding         : 65001

 Date: 05/01/2024 22:49:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `desc` varchar(180) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  `deleted_at` datetime NULL DEFAULT NULL,
  `create_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (1, 'p', '0', '/v1/get/dictList', 'GET', '', '', '', '获取菜单列表', '2023-06-25 22:15:26', '2023-06-25 22:19:44', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (2, 'p', '0', '/v1/get/allMenuList', 'GET', '', '', '', '获取全部菜单列表', '2023-06-25 22:15:37', '2023-06-25 22:19:40', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (3, 'p', '0', '/v1/get/roleMenuList', 'GET', '', '', '', '获取角色菜单列表', '2023-06-25 22:15:59', '2023-06-25 22:19:37', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (4, 'p', '0', '/v1/get/roleCasbinList', 'POST', '', '', '', '获取角色权限列表', '2023-06-25 22:16:18', '2023-06-25 22:19:34', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (5, 'p', '0', '/v1/get/setting', 'GET', '', '', '', '获用户设置', '2023-06-25 22:16:34', '2023-06-25 22:20:10', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (6, 'p', '0', '/v1/update/setting', 'POST', '', '', '', '更新用户设置', '2023-06-25 22:16:47', '2023-06-25 22:20:04', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (7, 'p', '0', '/v1/add/roleMenu', 'POST', '', '', '', '添加角色菜单', '2023-06-25 22:17:35', '2023-06-25 22:20:20', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (8, 'p', '0', '/v1/delete/roleMenu', 'POST', '', '', '', '删除角色菜单', '2023-06-25 22:17:50', '2023-06-25 22:20:27', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (9, 'p', '0', '/v1/update/roleMenu', 'POST', '', '', '', '更新角色菜单', '2023-06-25 22:18:04', '2023-06-25 22:20:34', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (10, 'p', '0', '/v1/add/roleCasbin', 'POST', '', '', '', '添加角色权限', '2023-06-25 22:18:20', '2023-06-25 22:20:42', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (11, 'p', '0', '/v1/get/userInfo', 'GET', '', '', '', '获取用户信息', '2023-07-04 14:07:40', '2023-07-04 14:07:40', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (12, 'p', '0', '/v1/update/userInfo', 'POST', '', '', '', '更新用户信息', '2023-07-04 14:09:10', '2023-07-04 14:09:10', NULL, 'admin', 'admin');
INSERT INTO `casbin_rule` VALUES (13, 'p', '0', '/v1/update/password', 'POST', '', '', '', '更新用户密码', '2023-07-04 14:09:33', '2023-07-04 14:09:33', NULL, 'admin', 'admin');

-- ----------------------------
-- Table structure for debug_perf
-- ----------------------------
DROP TABLE IF EXISTS `debug_perf`;
CREATE TABLE `debug_perf`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'uid',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `motto` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '座右铭',
  `text` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  `deleted_at` datetime NULL DEFAULT NULL,
  `create_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_key_debug_perf`(`uid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '性能测试' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of debug_perf
-- ----------------------------

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `redirect` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '重定向',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件地址',
  `parent_id` int NOT NULL DEFAULT 0 COMMENT '父id',
  `role_id_group` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色组',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `hidden` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否隐藏',
  `keep_alive` tinyint(1) NOT NULL DEFAULT 1 COMMENT 'keepAlive',
  `sort` int NOT NULL DEFAULT 0 COMMENT '排序',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  `deleted_at` datetime NULL DEFAULT NULL,
  `create_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menu`(`name` ASC, `path` ASC, `parent_id` ASC, `role_id_group` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '路由表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, '/layout', 'layout', '/dashboard', 'layout/index.vue', 0, '0', '', '', 0, 1, 0, '2023-06-25 21:58:52', '2023-06-25 22:01:23', NULL, 'admin', 'admin');
INSERT INTO `menu` VALUES (2, '/dashboard', 'dashboard', '', 'views/dashboard/index.vue', 1, '0', '仪表盘', 'odometer', 0, 1, 1, '2023-06-25 22:02:44', '2023-06-25 22:02:44', NULL, 'admin', 'admin');
INSERT INTO `menu` VALUES (3, '/admin', 'admin', '/casbin', '', 1, '0', '管理员', 'user', 0, 1, 2, '2023-06-25 22:03:42', '2023-06-25 22:03:42', NULL, 'admin', 'admin');
INSERT INTO `menu` VALUES (4, '/casbin', 'casbin', '', 'views/admin/casbin/index.vue', 3, '0', '权限管理', 'platform', 0, 1, 1, '2023-06-25 22:04:27', '2023-06-25 22:07:25', NULL, 'admin', 'admin');
INSERT INTO `menu` VALUES (5, '/menus', 'menus', '', 'views/admin/menus/index.vue', 3, '0', '菜单管理', 'tickets', 0, 1, 2, '2023-06-25 22:07:20', '2023-06-25 22:07:34', NULL, 'admin', 'admin');
INSERT INTO `menu` VALUES (6, '/about', 'about', '', 'views/about/index.vue', 1, '0', '关于', 'star', 0, 1, 3, '2023-06-25 22:08:43', '2023-06-25 22:12:04', NULL, 'admin', 'admin');
INSERT INTO `menu` VALUES (7, '/person', 'person', '', 'views/person/index.vue', 1, '0', '个人信息', 'message', 1, 1, 4, '2023-06-25 22:09:42', '2023-06-25 22:51:58', NULL, 'admin', 'admin');
INSERT INTO `menu` VALUES (8, '/404', '404', '', 'views/error/index.vue', 1, '0', '404', 'InfoFilled', 1, 1, 5, '2023-06-25 22:10:47', '2023-06-25 22:51:50', NULL, 'admin', 'admin');
INSERT INTO `menu` VALUES (9, '/:catchAll(.*)', '', '/404', '', 0, '0', '', '', 0, 1, 0, '2023-06-25 22:11:52', '2023-06-25 22:12:10', NULL, 'admin', 'admin');

-- ----------------------------
-- Table structure for setting
-- ----------------------------
DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'uid',
  `lang` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '语言',
  `side_mode_color` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '侧边栏颜色',
  `collapse` tinyint(1) NOT NULL DEFAULT 0 COMMENT '侧边栏折叠',
  `breadcrumb` tinyint(1) NOT NULL DEFAULT 1 COMMENT '面包屑',
  `default_router` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '默认路由',
  `active_text_color` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活跃文本颜色',
  `active_background_color` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活跃文本背景色',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  `deleted_at` datetime NULL DEFAULT NULL,
  `create_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_settings`(`uid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '前端layout设置' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of setting
-- ----------------------------
INSERT INTO `setting` VALUES (1, '000000001', 'zh-cn', '#000000', 0, 1, 'dashboard', '#096DE6', '#484fda', '2023-06-25 21:55:26', '2023-06-25 22:52:56', NULL, 'admin', 'admin');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'uid',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `nickname` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `motto` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '座右铭',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加盐',
  `birth` date NOT NULL DEFAULT '2006-01-02' COMMENT '出生日期',
  `avatar` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'avatar.jpg' COMMENT '头像',
  `role_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色Id',
  `role_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
  `wechat` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '微信号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `state` int NOT NULL COMMENT '用户状态:(0:初始,1:使用,2:停用,3:删除)',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  `deleted_at` datetime NULL DEFAULT NULL,
  `create_user` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `update_user` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '修改人',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_key_value`(`username` ASC, `uid` ASC, `role_id` ASC, `state` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '000000001', 'admin', '小不点', '前进的苦，是在于人要逼着自己开辟新的痛苦领域', 'ce8c73b00381ac97ceba4feaf6f67b4d', '+t!dV$', '2006-01-02', 'https://i.postimg.cc/15038Rxn/4.png', '0', '管理员', '13269110806', 'xiaohubai', 'xiaohubai@outlook.com', 0, '2023-06-25 21:39:42', '2023-06-25 21:52:22', NULL, 'admin', 'admin');

SET FOREIGN_KEY_CHECKS = 1;
