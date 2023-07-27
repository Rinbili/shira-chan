/*
 Navicat Premium Data Transfer

 Source Server         : *
 Source Server Type    : MySQL
 Source Server Version : 80026
 Source Host           : *
 Source Schema         : shirachan_dev

 Target Server Type    : MySQL
 Target Server Version : 80026
 File Encoding         : 65001

 Date: 21/07/2023 00:46:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Orders
-- ----------------------------
DROP TABLE IF EXISTS `Orders`;
CREATE TABLE `Orders`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `content` varchar(15000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `contact` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `type` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `is_closed` tinyint(1) NOT NULL DEFAULT 0,
  `is_finished` tinyint(1) NOT NULL DEFAULT 0,
  `evaluation` double NULL DEFAULT NULL,
  `hope_at` bigint NOT NULL,
  `created_at` bigint NOT NULL,
  `updated_at` bigint NOT NULL,
  `user_requested` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `Orders_Users_requested`(`user_requested` ASC) USING BTREE,
  CONSTRAINT `Orders_Users_requested` FOREIGN KEY (`user_requested`) REFERENCES `Users` (`id`) ON DELETE SET NULL ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of Orders
-- ----------------------------
INSERT INTO `Orders` VALUES (1, '11', '如题', 'Rinke_Hu', '0', 0, 1, 0, 1686196529, 1688563093, 1689785201, 4294967301);
INSERT INTO `Orders` VALUES (2, '电脑爆炸', '如题', '18676288160', '0', 1, 0, 5, 1686282929, 1686023729, 1689785302, 4294967301);
INSERT INTO `Orders` VALUES (3, '11', '111', '13242598786', '0', 0, 0, 0, 1688010929, 1689254293, 1689784883, 4294967301);
INSERT INTO `Orders` VALUES (4, '11', '111', '13242598786', '1', 1, 0, 0, 1687146929, 1689254293, 1689785303, 4294967301);
INSERT INTO `Orders` VALUES (5, '11', '111', '13242598786', '1', 1, 0, 5, 1686023042, 1686400525, 1689785304, 4294967301);
INSERT INTO `Orders` VALUES (6, '11', '111', '13242598786', '1', 1, 0, 5, 1686023042, 1689254293, 1689785305, 4294967301);
INSERT INTO `Orders` VALUES (7, '121', '1212', '13242598786', '3', 1, 0, 5, 1686401576, 1686401580, 1689785306, 4294967301);
INSERT INTO `Orders` VALUES (8, '12', '12', '13242598786', '4', 1, 0, NULL, 1686401836, 1686401841, 1689785307, 4294967301);
INSERT INTO `Orders` VALUES (9, '212', '121', '13242598786', '5', 1, 0, NULL, 1686401925, 1689081493, 1689785309, 4294967301);
INSERT INTO `Orders` VALUES (10, '1212', '121', '13242598786', '4', 1, 0, NULL, 1687525190, 1686402001, 1689785317, 4294967301);
INSERT INTO `Orders` VALUES (12, '1212', '212', '13242598786', '6', 1, 0, NULL, 1686023042, 1686402163, 1689785318, 4294967301);
INSERT INTO `Orders` VALUES (15, '213', '123', '13242598786', '6', 1, 0, NULL, 1686023042, 1689427093, 1689785320, 4294967301);
INSERT INTO `Orders` VALUES (18, '12', '121', '13242598786', '4', 0, 0, NULL, 1689427093, 1689427093, 1686413345, 4294967301);
INSERT INTO `Orders` VALUES (19, 'hhhhh', 'huohoioi', '13242598786', '1', 0, 0, NULL, 1686555066, 1689427093, 1689781619, 4294967301);
INSERT INTO `Orders` VALUES (21, '123rewfdgbvd', 'eWGAREHTSRJYFLIG', '13242598786', '2', 0, 0, NULL, 1686555381, 1686555384, 1686555384, 4294967301);
INSERT INTO `Orders` VALUES (22, 'test_06', '1', 'Rinke', '3', 0, 0, NULL, 1686626939, 1686627235, 1689781621, 4294967296);
INSERT INTO `Orders` VALUES (23, '23123', '123123', '13242598786', '3', 0, 0, NULL, 1689493583, 1689493587, 1689784899, 4294967301);
INSERT INTO `Orders` VALUES (24, '12312', '12312312', '13242598786', '1', 0, 1, NULL, 1689524677, 1689524680, 1689786717, 4294967301);

-- ----------------------------
-- Table structure for Users
-- ----------------------------
DROP TABLE IF EXISTS `Users`;
CREATE TABLE `Users`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `passwd` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `is_admin` tinyint(1) NOT NULL DEFAULT 0,
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` bigint NOT NULL,
  `updated_at` bigint NOT NULL,
  `is_secretary` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `phone`(`phone` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4294967305 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of Users
-- ----------------------------
INSERT INTO `Users` VALUES (4294967296, 'Rinke', '$2a$10$T4JlmSdYg3/BunIUInYplOKUZuXTEx/rPml3AY2g0ntjUIYts6Iza', '18676288163', 1, 1, 1686023729, 1686023729, 1);
INSERT INTO `Users` VALUES (4294967298, 'ssrinke', '$2a$10$T4JlmSdYg3/BunIUInYplOKUZuXTEx/rPml3AY2g0ntjUIYts6Iza', '18676288162', 0, 1, 1686023729, 1686023729, 1);
INSERT INTO `Users` VALUES (4294967299, 'Rinbili', '$2a$10$T4JlmSdYg3/BunIUInYplOKUZuXTEx/rPml3AY2g0ntjUIYts6Iza', '18676288161', 0, 0, 1686023729, 1686023729, 0);
INSERT INTO `Users` VALUES (4294967301, '张', '$2a$10$znlFYyPBMIKw5HLMOoS0/uXF8.otaIN6BkT12UGnIUkhHg2T6/TCy', '13242598786', 1, 1, 1686023729, 1686023729, 1);
INSERT INTO `Users` VALUES (4294967302, '0716', '$2a$10$m1Ye0nbDf6LBhkyzi/FZvep8Ifpr0ErGuaxX64icSC7K3o8Il8ZEq', '18676000000', 0, 1, 1689482641, 1689482641, 1);
INSERT INTO `Users` VALUES (4294967304, '010t', '$2a$10$sKHElcJAtIbMU.lQiRthK.e1Cnk.sS7WHP93zp6DYVnlDj3IRG4Jq', '01000000000', 0, 1, 1689737549, 1689752535, 0);

-- ----------------------------
-- Table structure for ent_types
-- ----------------------------
DROP TABLE IF EXISTS `ent_types`;
CREATE TABLE `ent_types`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `type`(`type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ent_types
-- ----------------------------
INSERT INTO `ent_types` VALUES (1, 'Orders');
INSERT INTO `ent_types` VALUES (2, 'Users');

-- ----------------------------
-- Table structure for order_receiver
-- ----------------------------
DROP TABLE IF EXISTS `order_receiver`;
CREATE TABLE `order_receiver`  (
  `order_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`order_id`, `user_id`) USING BTREE,
  INDEX `order_receiver_user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `order_receiver_order_id` FOREIGN KEY (`order_id`) REFERENCES `Orders` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `order_receiver_user_id` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_receiver
-- ----------------------------
INSERT INTO `order_receiver` VALUES (1, 4294967301);
INSERT INTO `order_receiver` VALUES (3, 4294967301);
INSERT INTO `order_receiver` VALUES (23, 4294967301);
INSERT INTO `order_receiver` VALUES (24, 4294967301);

SET FOREIGN_KEY_CHECKS = 1;
