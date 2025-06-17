CREATE TABLE users (
    `id` char(36) NOT NULL DEFAULT '' COMMENT '教职人员ID',
    `username` VARCHAR(50) NOT NULL UNIQUE,
    `nick_name` VARCHAR(50) NOT NULL UNIQUE,
    `passwd` varchar(255) NOT NULL COMMENT '密码（已加密）',
    `phone` varchar(20) DEFAULT '' COMMENT '联系方式',
    `email` varchar(100) DEFAULT NULL COMMENT '邮箱地址',
    `is_delete` tinyint(1) DEFAULT '0' COMMENT '是否删除（逻辑删除标记）',
    `creator` varchar(100) DEFAULT NULL COMMENT '创建者',
    `updater` varchar(100) DEFAULT NULL COMMENT '更新者',
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

