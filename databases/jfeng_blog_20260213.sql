/*
 Navicat Premium Dump SQL

 Source Server         : 115.190.253.253_mysql8.0.44
 Source Server Type    : MySQL
 Source Server Version : 80044 (8.0.44)
 Source Host           : 115.190.253.253:3306
 Source Schema         : jfeng_blog

 Target Server Type    : MySQL
 Target Server Version : 80044 (8.0.44)
 File Encoding         : 65001

 Date: 13/02/2026 11:52:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`  (
  `post_id` int NOT NULL AUTO_INCREMENT,
  `author_id` int NOT NULL COMMENT '作者ID',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文章标题',
  `slug` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'URL友好标识',
  `excerpt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '文章摘要',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '文章内容',
  `cover_image_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '封面图链接',
  `status` enum('draft','published','archived') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'published' COMMENT '状态',
  `view_count` int NULL DEFAULT 0 COMMENT '阅读量',
  `is_recommended` tinyint(1) NULL DEFAULT 0 COMMENT '是否推荐',
  `published_at` timestamp NULL DEFAULT NULL COMMENT '发布时间',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `tags_str` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标签字符串，逗号分隔',
  PRIMARY KEY (`post_id`) USING BTREE,
  UNIQUE INDEX `slug`(`slug` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '博客文章表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post
-- ----------------------------
INSERT INTO `post` VALUES (1, 1, 'Aptos Move 语言入门指南', 'aptos-move-intro', '本文将介绍 Aptos 区块链及其 Move 编程语言的基础知识，包括账户模型、资源管理和智能合约开发。', '文章完整内容...', NULL, 'published', 1523, 1, '2025-09-11 03:53:56', '2026-02-11 03:53:56', '2026-02-11 03:53:56', '区块链,Web3');
INSERT INTO `post` VALUES (2, 1, 'Rust 在 Web3 开发中的优势', 'rust-web3-advantages', '探讨 Rust 语言为何成为 Web3 开发的热门选择，以及其在区块链领域的应用场景。', '文章完整内容...', NULL, 'published', 2156, 1, '2025-09-11 03:53:56', '2026-02-11 03:53:56', '2026-02-11 03:53:56', 'Rust,Web3');
INSERT INTO `post` VALUES (3, 1, 'Kubernetes 生产环境最佳实践', 'kubernetes-best-practices', '总结 Kubernetes 在生产环境中的最佳实践，包括高可用配置、监控告警和日志管理。', '文章完整内容...', NULL, 'published', 3421, 1, '2025-09-11 03:53:56', '2026-02-11 03:53:56', '2026-02-11 03:53:56', '云原生');
INSERT INTO `post` VALUES (4, 1, 'DeFi 协议安全审计清单', 'defi-security-audit', '整理 DeFi 协议安全审计的完整清单，涵盖常见漏洞类型和审计要点。', '文章完整内容...', NULL, 'published', 1876, 1, '2025-08-11 03:53:56', '2026-02-11 03:53:56', '2026-02-11 03:53:56', '安全,区块链');
INSERT INTO `post` VALUES (5, 1, '深入理解 Solidity 0.8 新特性', 'solidity-0-8-features', '详细介绍 Solidity 0.8 版本的新特性，包括自融溢出检查、ABI解码器等。', '文章完整内容...', NULL, 'published', 1234, 0, '2025-08-11 03:53:56', '2026-02-11 03:53:56', '2026-02-11 03:53:56', 'Ethereum,Web3');
INSERT INTO `post` VALUES (6, 1, 'Go 并发编程实践', 'go-concurrency', '分享 Go 语言并发编程的最佳实践，包括 goroutine、channel 和同步原语的使用。', '文章完整内容...', NULL, 'published', 987, 0, '2025-08-11 03:53:56', '2026-02-11 03:53:56', '2026-02-11 03:53:56', 'Go');
INSERT INTO `post` VALUES (7, 1, 'IPFS 分布式存储入门', 'ipfs-introduction', '介绍 IPFS 协议及其在 Web3 应用中的使用，包括节点搭建和内容寻址。', '文章完整内容...', NULL, 'published', 0, 0, '2026-02-11 03:53:56', '2026-02-11 03:53:56', '2026-02-11 03:53:56', '');

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project`  (
  `project_id` int NOT NULL AUTO_INCREMENT,
  `author_id` int NOT NULL COMMENT '作者ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '项目名称',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '项目描述',
  `cover_image_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '项目封面图',
  `detail_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '项目详情内容',
  `repo_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '仓库链接',
  `demo_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '演示链接',
  `is_recommended` tinyint(1) NULL DEFAULT 0 COMMENT '是否推荐',
  `view_count` int NULL DEFAULT 0 COMMENT '查看次数',
  `status` enum('active','archived') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'active' COMMENT '项目状态',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `tech_stack_str` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '技术栈字符串，逗号分隔',
  PRIMARY KEY (`project_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '项目表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of project
-- ----------------------------
INSERT INTO `project` VALUES (1, 1, 'JFeng Blog', '基于 Vue 3 + TypeScript 的现代化技术博客系统，包含完整的前后端实现。', NULL, NULL, 'https://github.com/jfeng/blog', 'https://blog.jfeng.dev', 1, 0, 'active', '2026-02-11 03:54:13', '2026-02-11 03:54:13', 'Vue 3,TypeScript,Pinia,Tailwind CSS');
INSERT INTO `project` VALUES (2, 1, 'Aptos SDK Rust', 'Rust 语言实现的 Aptos 区块链开发工具包，支持交易构造、签名和链上交互。', NULL, NULL, 'https://github.com/jfeng/aptos-sdk-rust', 'https://crates.io/crates/aptos-sdk', 1, 0, 'active', '2026-02-11 03:54:13', '2026-02-11 03:54:13', 'Rust,Aptos,Blockchain');
INSERT INTO `project` VALUES (3, 1, 'DeFi Dashboard', 'DeFi 协议收益监控仪表板，支持多协议、多账户的收益聚合展示。', NULL, NULL, 'https://github.com/jfeng/defi-dashboard', 'https://defi.jfeng.dev', 1, 0, 'active', '2026-02-11 03:54:13', '2026-02-11 03:54:13', 'React,TypeScript,Web3.js');
INSERT INTO `project` VALUES (4, 1, 'Hardware Wallet Tool', 'Ledger 硬件钱包的 Rust 交互工具，支持离线签名和交易构造。', NULL, NULL, 'https://github.com/jfeng/hardware-wallet-tool', NULL, 0, 0, 'active', '2026-02-11 03:54:13', '2026-02-11 03:54:13', 'Rust,USB HID,Blockchain');

-- ----------------------------
-- Table structure for sys_tag
-- ----------------------------
DROP TABLE IF EXISTS `sys_tag`;
CREATE TABLE `sys_tag`  (
  `tag_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标签名称',
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'URL友好标识',
  `type` enum('post_category','project_tech','skill') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'post_category' COMMENT '标签类型',
  `description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标签描述',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`tag_id`) USING BTREE,
  UNIQUE INDEX `slug`(`slug` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '标签表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_tag
-- ----------------------------
INSERT INTO `sys_tag` VALUES (1, '全部', 'all', 'post_category', '所有文章', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (2, '区块链', 'blockchain', 'post_category', '区块链相关文章', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (3, 'Rust', 'rust', 'post_category', 'Rust语言相关文章', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (4, '云原生', 'cloud-native', 'post_category', '云原生技术文章', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (5, '安全', 'security', 'post_category', '安全相关文章', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (6, 'Ethereum', 'ethereum', 'post_category', '以太坊相关文章', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (7, 'Go', 'go', 'post_category', 'Go语言相关文章', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (8, 'Web3', 'web3', 'post_category', 'Web3技术文章', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (9, 'TypeScript', 'typescript', 'post_category', 'TypeScript相关文章', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (10, 'Vue 3', 'vue3', 'project_tech', 'Vue 3框架', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (11, 'TypeScript', 'typescript-tech', 'project_tech', 'TypeScript语言', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (12, 'Pinia', 'pinia', 'project_tech', 'Pinia状态管理', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (13, 'Tailwind CSS', 'tailwind-css', 'project_tech', 'Tailwind CSS框架', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (14, 'Rust', 'rust-tech', 'project_tech', 'Rust编程语言', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (15, 'Aptos', 'aptos', 'project_tech', 'Aptos区块链', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (16, 'Blockchain', 'blockchain-tech', 'project_tech', '区块链技术', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (17, 'React', 'react', 'project_tech', 'React框架', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (18, 'Web3.js', 'web3js', 'project_tech', 'Web3.js库', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (19, 'USB HID', 'usb-hid', 'project_tech', 'USB HID协议', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (20, 'Aptos', 'aptos-skill', 'skill', 'Aptos区块链开发', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (21, 'Move', 'move', 'skill', 'Move编程语言', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (22, 'Web3', 'web3-skill', 'skill', 'Web3开发', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (23, 'Rust', 'rust-skill', 'skill', 'Rust编程', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (24, 'Kubernetes', 'kubernetes', 'skill', 'Kubernetes容器编排', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (25, 'DeFi', 'defi', 'skill', '去中心化金融', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (26, 'Solidity', 'solidity', 'skill', 'Solidity智能合约', '2026-02-11 03:51:05');
INSERT INTO `sys_tag` VALUES (27, 'Go', 'go-skill', 'skill', 'Go语言开发', '2026-02-11 03:51:05');

-- ----------------------------
-- Table structure for timeline
-- ----------------------------
DROP TABLE IF EXISTS `timeline`;
CREATE TABLE `timeline`  (
  `timeline_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户ID',
  `year` year NOT NULL COMMENT '年份',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '事件标题',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '事件描述',
  `type` enum('personal','project_milestone','career') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'personal' COMMENT '事件类型',
  `sort_order` int NULL DEFAULT 0 COMMENT '排序序号',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`timeline_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '时间线表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of timeline
-- ----------------------------
INSERT INTO `timeline` VALUES (1, 1, 2020, '开始接触云原生技术', '深入学习Kubernetes及相关技术栈', 'personal', 1, '2026-02-11 03:52:34', '2026-02-11 03:52:34');
INSERT INTO `timeline` VALUES (2, 1, 2022, '转向Web3领域', '学习Rust和区块链开发，开始深入研究Web3技术', 'personal', 1, '2026-02-11 03:52:34', '2026-02-11 03:52:34');
INSERT INTO `timeline` VALUES (3, 1, 2023, '深入研究DeFi协议', '开始写作技术博客，分享Web3开发经验', 'personal', 2, '2026-02-11 03:52:34', '2026-02-11 03:52:34');
INSERT INTO `timeline` VALUES (4, 1, 2023, '开始写作技术博客', '创建个人技术博客，分享开发经验', 'personal', 1, '2026-02-11 03:52:34', '2026-02-11 03:52:34');
INSERT INTO `timeline` VALUES (5, 1, 2024, '开始专注于Aptos生态开发', '深入研究Aptos区块链和Move语言', 'personal', 1, '2026-02-11 03:52:34', '2026-02-11 03:52:34');
INSERT INTO `timeline` VALUES (6, 1, 2024, '发布多个开源区块链项目', '包括Aptos SDK、DeFi Dashboard等多个开源项目', 'project_milestone', 2, '2026-02-11 03:52:34', '2026-02-11 03:52:34');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `display_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '显示名称',
  `avatar_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '头像链接',
  `bio` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '个人简介',
  `role` enum('admin','author','guest') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'author' COMMENT '角色',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `password_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码哈希',
  `social_links` json NULL COMMENT '社交链接JSON',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `email`(`email` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'jfeng', 'JFeng', NULL, '我是一名资深后端工程师，拥有多年的分布式系统开发经验。近年来，我专注于区块链技术和云原生架构的研究与实践。', 'author', 'jfeng@example.com', NULL, '{\"email\": \"jfeng@example.com\", \"github\": \"https://github.com/jfeng\"}', '2026-02-11 03:49:54', '2026-02-11 03:49:54');

SET FOREIGN_KEY_CHECKS = 1;
