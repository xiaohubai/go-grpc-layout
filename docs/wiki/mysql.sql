create table
    `user`(
        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
        `uid` varchar(20) NOT NULL DEFAULT '' COMMENT 'uid',
        `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
        `nickname` varchar(20) NOT NULL DEFAULT '' COMMENT '昵称',
        `motto` varchar(191) NOT NULL DEFAULT '' COMMENT '座右铭',
        `password` varchar(191) NOT NULL DEFAULT '' COMMENT '密码',
        `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '加盐',
        `birth` date NOT NULL DEFAULT '2006-01-02' COMMENT '出生日期',
        `avatar` varchar(191) NOT NULL DEFAULT 'avatar.jpg' COMMENT '头像',
        `role_id` varchar(20) NOT NULL DEFAULT '' COMMENT '角色Id',
        `role_name` varchar(20) NOT NULL DEFAULT '' COMMENT '角色名称',
        `phone` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
        `wechat` varchar(191) NOT NULL DEFAULT '' COMMENT '微信号',
        `email` varchar(40) NOT NULL DEFAULT '' COMMENT '邮箱',
        `state` int(10) NOT NULL DEFAULT '0' COMMENT '用户状态:(0:初始,1:使用,2:停用,3:删除)',
        `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
        `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
        `deleted_at` datetime DEFAULT NULL,
        `create_user` varchar(20) NOT NULL DEFAULT '' COMMENT '创建人',
        `update_user` varchar(20) NOT NULL DEFAULT '' COMMENT '修改人',
        PRIMARY KEY (`id`),
        KEY `idx_key_value` (
            `username`,
            `uid`,
            `role_id`,
            `state`
        )
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表';

CREATE TABLE
    `casbin_rule` (
        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
        `ptype` varchar(100) NOT NULL DEFAULT '' COMMENT '',
        `v0` varchar(100) NOT NULL DEFAULT '' COMMENT '角色',
        `v1` varchar(100) NOT NULL DEFAULT '' COMMENT '接口地址',
        `v2` varchar(100) NOT NULL DEFAULT '' COMMENT 'method',
        `v3` varchar(100) NOT NULL DEFAULT '' COMMENT '',
        `v4` varchar(100) NOT NULL DEFAULT '' COMMENT '',
        `v5` varchar(100) NOT NULL DEFAULT '' COMMENT '',
        `desc` varchar(180) NOT NULL DEFAULT '' COMMENT '描述',
        `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
        `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
        `deleted_at` datetime DEFAULT NULL,
        `create_user` varchar(20) NOT NULL DEFAULT '' COMMENT '创建人',
        `update_user` varchar(20) NOT NULL DEFAULT '' COMMENT '修改人',
        PRIMARY KEY (`id`),
        UNIQUE KEY `idx_casbin_rule` (
            `ptype`,
            `v0`,
            `v1`,
            `v2`,
            `v3`,
            `v4`,
            `v5`
        )
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限表';


CREATE TABLE
    `menu` (
        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
        `path` varchar(100) NOT NULL DEFAULT '' COMMENT '路由',
        `name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
        `redirect` varchar(100) NOT NULL DEFAULT '' COMMENT '重定向',
        `component` varchar(100) NOT NULL DEFAULT '' COMMENT '文件地址',
        `parent_id` int(10) NOT NULL DEFAULT 0 COMMENT '父id',
        `role_id_group` varchar(18) NOT NULL DEFAULT '' COMMENT '角色组',
        `title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
        `icon` varchar(100) NOT NULL DEFAULT '' COMMENT '图标',
        `hidden` BOOLEAN NOT NULL DEFAULT 0 COMMENT '是否隐藏',
        `keep_alive` BOOLEAN NOT NULL DEFAULT 1 COMMENT 'keepAlive',
        `sort` int(10) NOT NULL DEFAULT 0 COMMENT '排序',
        `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
        `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
        `deleted_at` datetime DEFAULT NULL,
        `create_user` varchar(20) NOT NULL DEFAULT '' COMMENT '创建人',
        `update_user` varchar(20) NOT NULL DEFAULT '' COMMENT '修改人',
        PRIMARY KEY (`id`),
        KEY `idx_menu` (`name`, `path`, `parent_id`, `role_id_group`)
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '路由表';

CREATE TABLE
    `setting` (
        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
        `uid` varchar(100) NOT NULL DEFAULT '' COMMENT 'uid',
        `lang` varchar(100) NOT NULL DEFAULT '' COMMENT '语言',
        `side_mode_color` varchar(100) NOT NULL DEFAULT '' COMMENT '侧边栏颜色',
        `collapse` BOOLEAN NOT NULL DEFAULT 0 COMMENT '侧边栏折叠',
        `breadcrumb` BOOLEAN NOT NULL DEFAULT 1 COMMENT '面包屑',
        `default_router` varchar(20) NOT NULL DEFAULT '' COMMENT '默认路由',
        `active_text_color` varchar(20) NOT NULL DEFAULT '' COMMENT '活跃文本颜色',
        `active_background_color` varchar(20) NOT NULL DEFAULT '' COMMENT '活跃文本背景色',
        `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
        `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
        `deleted_at` datetime DEFAULT NULL,
        `create_user` varchar(20) NOT NULL DEFAULT '' COMMENT '创建人',
        `update_user` varchar(20) NOT NULL DEFAULT '' COMMENT '修改人',
        PRIMARY KEY (`id`),
        UNIQUE KEY `idx_settings` (`uid`)
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '前端layout设置';


CREATE TABLE
    `debug_perf` (
        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
        `uid` varchar(100) NOT NULL DEFAULT '' COMMENT 'uid',
        `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
        `motto` varchar(100) NOT NULL DEFAULT '' COMMENT '座右铭',
        `text` varchar(100) NOT NULL DEFAULT '' COMMENT '内容',
        `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
        `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
        `deleted_at` datetime DEFAULT NULL,
        `create_user` varchar(20) NOT NULL DEFAULT '' COMMENT '创建人',
        `update_user` varchar(20) NOT NULL DEFAULT '' COMMENT '修改人',
        PRIMARY KEY (`id`),
        KEY `idx_key_debug_perf` (`uid`)
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '性能测试';
