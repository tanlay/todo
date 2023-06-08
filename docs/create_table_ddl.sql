CREATE DATABASE `todo`
CREATE TABLE `todo` (
                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                        `task` varchar(500) NOT NULL COMMENT '任务名',
                        `category` varchar(50) NOT NULL COMMENT '分类',
                        `status` tinyint(1) NOT NULL COMMENT '是否已经完成，0：未完成，1：完成',
                        `create_at` int(20) NOT NULL COMMENT '创建时间',
                        `completed_at` int(20) NOT NULL COMMENT '完成时间',
                        PRIMARY KEY (`id`),
                        KEY `idx_todo_status` (`status`),
                        KEY `idx_todo_create_at` (`create_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;