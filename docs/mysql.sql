create table
    `user`(
        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
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
    ) ENGINE = InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';
