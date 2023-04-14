CREATE TABLE `user` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `number` varchar(255) NOT NULL DEFAULT '' COMMENT '学号',
                        `name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名称',
                        `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',
                        `gender` char(5) NOT NULL COMMENT '男｜女｜未公开',
                        `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `del_state` tinyint NOT NULL DEFAULT '0',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `number_unique` (`number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;