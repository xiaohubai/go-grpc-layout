create table `user`(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` varchar(180) NOT NULL DEFAULT '' COMMENT 'uid',
    `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
    `nickname` varchar(20) NOT NULL DEFAULT '' COMMENT '昵称',
    `motto` varchar(180) NOT NULL DEFAULT '' COMMENT '座右铭',
    `password` varchar(180) NOT NULL DEFAULT '' COMMENT '密码',
    `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '加盐',
    `birth` date NOT NULL DEFAULT '2006-01-02' COMMENT '出生日期',
    `avatar` varchar(20) NOT NULL DEFAULT 'avatar.jpg' COMMENT '头像',
    `role_id` varchar(20) NOT NULL DEFAULT '' COMMENT '角色Id',
    `role_name` varchar(20) NOT NULL DEFAULT '' COMMENT '角色名称',
    `phone` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `wechat` varchar(180) NOT NULL DEFAULT '' COMMENT '微信号',
    `email` varchar(40) NOT NULL DEFAULT '' COMMENT '邮箱',
    `state` int(10) NOT NULL DEFAULT '0' COMMENT '用户状态:(0:初始,1:使用,2:停用,3:删除)',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
    `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
    `create_user` varchar(20) NOT NULL DEFAULT '' COMMENT '创建人',
    `update_user` varchar(20) NOT NULL DEFAULT '' COMMENT '修改人',
    PRIMARY KEY (`id`),
    KEY `idx_key_value` (
        `username`,
        `uid`,
        `role_id`,
        `state`
    )
) ENGINE = InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表';

CREATE TABLE `casbin_rule` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `ptype` varchar(100) NOT NULL DEFAULT '' COMMENT '',
    `v0` varchar(100) NOT NULL DEFAULT '' COMMENT '角色',
    `v1` varchar(100) NOT NULL DEFAULT '' COMMENT '接口地址',
    `v2` varchar(100) NOT NULL DEFAULT '' COMMENT 'method',
    `v3` varchar(100) NOT NULL DEFAULT '' COMMENT '',
    `v4` varchar(100) NOT NULL DEFAULT '' COMMENT '',
    `v5` varchar(100) NOT NULL DEFAULT '' COMMENT '',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
    `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
    `create_user` varchar(20) NOT NULL DEFAULT '' COMMENT '创建人',
    `update_user` varchar(20) NOT NULL DEFAULT '' COMMENT '修改人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
) ENGINE = InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限表';

CREATE TABLE `menu` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `path` varchar(100) NOT NULL DEFAULT '' COMMENT '路由',
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
    `redirect` varchar(100) NOT NULL DEFAULT '' COMMENT '重定向',
    `component` varchar(100) NOT NULL DEFAULT '' COMMENT '文件地址',
    `parentId` int(10) NOT NULL DEFAULT 0 COMMENT '父id',
    `roleIDs` varchar(18) NOT NULL DEFAULT '' COMMENT '角色组',
    `title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
    `icon` varchar(100) NOT NULL DEFAULT '' COMMENT '图标',
    `hidden` BOOLEAN NOT NULL DEFAULT 0 COMMENT '是否隐藏',
    `keepAlive` BOOLEAN NOT NULL DEFAULT 1 COMMENT 'keepAlive',
    `sort` int(10) NOT NULL DEFAULT 0 COMMENT '排序',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
    `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
    `create_user` varchar(20) NOT NULL DEFAULT '' COMMENT '创建人',
    `update_user` varchar(20) NOT NULL DEFAULT '' COMMENT '修改人',
    PRIMARY KEY (`id`),
    KEY `idx_menu` (
        `name`,
        `parentId`,
        `roleIDs`
    )
) ENGINE = InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '路由表';