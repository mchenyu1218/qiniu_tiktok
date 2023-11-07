/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : localhost:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 06/11/2023 14:52:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for collect
-- ----------------------------
DROP TABLE IF EXISTS `collect`;
CREATE TABLE `collect`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint NOT NULL COMMENT '进行收藏操作的用户',
  `video_id` int NOT NULL COMMENT '收藏的视频',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC STATS_AUTO_RECALC = 1;

-- ----------------------------
-- Records of collect
-- ----------------------------
INSERT INTO `collect` VALUES (8, 1672489513669, 1, '2023-11-01 16:03:33', '2023-11-01 16:03:33');
INSERT INTO `collect` VALUES (9, 1672489513669, 3, '2023-11-01 16:07:09', '2023-11-01 16:07:09');
INSERT INTO `collect` VALUES (10, 1672489513669, 5, '2023-11-01 16:07:12', '2023-11-01 16:07:12');
INSERT INTO `collect` VALUES (11, 1672489513672, 5, '2023-11-01 16:07:19', '2023-11-01 16:07:19');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除(0-未删, 1-已删)',
  `video_id` bigint NOT NULL COMMENT '视频id',
  `author_id` bigint NOT NULL COMMENT '发布评论的用户id',
  `publisher_id` bigint NOT NULL COMMENT '评论的视频是哪个作者的id',
  `content` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
  `favourite_count` int NOT NULL COMMENT '评论点赞数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1670896770124 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '评论表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1644002802725, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1644343662976, 1641275264003, 1648443840420, '蔡徐坤', 54);
INSERT INTO `comment` VALUES (1644238047258, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1645820465948, 1641347509148, 1648753996219, '鸡', 3);
INSERT INTO `comment` VALUES (1646821740992, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1646230672085, 1641529965110, 1648831225038, '练习', 12);
INSERT INTO `comment` VALUES (1647173220985, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1646298962754, 1641755320246, 1648933351526, '唱', 33);
INSERT INTO `comment` VALUES (1647424380753, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1647017652341, 1642070340260, 1649175220138, '太', 98);
INSERT INTO `comment` VALUES (1648852605871, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1647021013279, 1642192401779, 1649453942303, '篮球', 19);
INSERT INTO `comment` VALUES (1654289927631, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1648109069096, 1642303003597, 1649495994019, '篮球', 83);
INSERT INTO `comment` VALUES (1655012847767, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1648131463537, 1642502600681, 1649761216969, '美', 71);
INSERT INTO `comment` VALUES (1655511697647, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1648418156239, 1642662025286, 1650997226170, '时长', 90);
INSERT INTO `comment` VALUES (1656989148608, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1648466085624, 1642906146538, 1651369992412, '鸡', 9);
INSERT INTO `comment` VALUES (1657913450888, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1648595775059, 1643559986949, 1651712363002, '时长', 46);
INSERT INTO `comment` VALUES (1659849841415, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1648600160164, 1643699118265, 1652066054671, 'music', 58);
INSERT INTO `comment` VALUES (1664517938705, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1648956827547, 1643944188310, 1652161842470, '篮球', 91);
INSERT INTO `comment` VALUES (1665516852148, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1649065526476, 1644714052623, 1652206273467, 'rap', 87);
INSERT INTO `comment` VALUES (1665675094998, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1649391303994, 1645346229343, 1652519823169, '篮球', 31);
INSERT INTO `comment` VALUES (1667559169576, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1649597943320, 1646052819692, 1652644940845, '你', 6);
INSERT INTO `comment` VALUES (1667616637242, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1649618321599, 1646060622267, 1652662439185, '两年半', 75);
INSERT INTO `comment` VALUES (1668604815396, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1649669004831, 1646434939010, 1652993634480, 'music', 76);
INSERT INTO `comment` VALUES (1669966303635, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1650155369308, 1646541441569, 1653209060773, '太', 100);
INSERT INTO `comment` VALUES (1670896770124, '2023-01-22 17:21:56', '2023-01-22 17:30:00', 0, 1650231875488, 1648161951242, 1653330031207, '两年半', 32);

-- ----------------------------
-- Table structure for fans
-- ----------------------------
DROP TABLE IF EXISTS `fans`;
CREATE TABLE `fans`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除(0-未删, 1-已删)',
  `blogger_id` bigint NOT NULL COMMENT '博主id',
  `fans_id` bigint NOT NULL COMMENT '粉丝id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1672099833458 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '粉丝表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of fans
-- ----------------------------
INSERT INTO `fans` VALUES (1672099833452, '2023-01-22 17:22:12', '2023-01-22 17:30:33', 0, 1648161951242, 1653330031207);
INSERT INTO `fans` VALUES (1672099833458, '2023-11-01 19:22:43', '2023-11-01 19:22:43', 0, 1672489513672, 1672489513671);

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '点赞操作的用户',
  `video_id` int NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of favorite
-- ----------------------------
INSERT INTO `favorite` VALUES (1, 1672489513669, 1, '2023-09-06 23:29:08', '2023-10-18 23:29:11');
INSERT INTO `favorite` VALUES (2, 1672489513669, 1, '2023-11-01 15:38:56', '2023-11-01 15:38:56');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `gender` int NOT NULL DEFAULT 0 COMMENT '男1 女0',
  `follower_count` int NOT NULL DEFAULT 0 COMMENT '粉丝数',
  `follow_count` int NOT NULL DEFAULT 0 COMMENT '关注数',
  `face` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e7/Everest_North_Face_toward_Base_Camp_Tibet_Luca_Galuzzi_2006.jpg/800px-Everest_North_Face_toward_Base_Camp_Tibet_Luca_Galuzzi_2006.jpg' COMMENT '头像',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除(0-未删, 1-已删)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1672489513673 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户数据表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1672489513669, 'hyn', '123456', 0, 0, 0, 'http://s38oif2dm.hn-bkt.clouddn.com/%E4%B8%8B%E8%BD%BD.jpg', '2023-01-24 22:45:07', '2023-10-30 22:00:50', 0);
INSERT INTO `user` VALUES (1672489513671, '02nv', '123456', 0, 0, 1, 'http://s38oif2dm.hn-bkt.clouddn.com/%E4%B8%8B%E8%BD%BD%20%281%29.jpg', '2023-01-28 18:10:55', '2023-11-01 19:22:43', 0);
INSERT INTO `user` VALUES (1672489513672, 'moyang', '$2a$10$rZSwBMICDYZ5g1DtbJakX.9XU/DY5hnc66e4lFiE7YB6Flo8YvKPG', 0, 1, 0, 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e7/Everest_North_Face_toward_Base_Camp_Tibet_Luca_Galuzzi_2006.jpg/800px-Everest_North_Face_toward_Base_Camp_Tibet_Luca_Galuzzi_2006.jpg', '2023-10-30 22:37:52', '2023-11-01 19:22:43', 0);

-- ----------------------------
-- Table structure for user_video
-- ----------------------------
DROP TABLE IF EXISTS `user_video`;
CREATE TABLE `user_video`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `video_id` bigint NOT NULL COMMENT '视频id',
  `is_favorite` tinyint NOT NULL DEFAULT 0 COMMENT '1表示已点赞 0未点赞',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除(0-未删, 1-已删)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任意用户和任意视屏之间的关联' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_video
-- ----------------------------
INSERT INTO `user_video` VALUES (1, 1672489513669, 1, 0, '2023-01-28 16:49:38', '2023-01-28 16:49:38', 0);
INSERT INTO `user_video` VALUES (32, 1672489513671, 34, 0, '2023-11-06 14:46:49', '2023-11-06 14:46:49', 0);
INSERT INTO `user_video` VALUES (33, 1672489513671, 35, 0, '2023-11-06 14:48:38', '2023-11-06 14:48:38', 0);
INSERT INTO `user_video` VALUES (34, 1672489513671, 36, 0, '2023-11-06 14:48:48', '2023-11-06 14:48:48', 0);
INSERT INTO `user_video` VALUES (35, 1672489513671, 37, 0, '2023-11-06 14:48:59', '2023-11-06 14:48:59', 0);
INSERT INTO `user_video` VALUES (36, 1672489513671, 38, 0, '2023-11-06 14:49:23', '2023-11-06 14:49:23', 0);
INSERT INTO `user_video` VALUES (37, 1672489513671, 39, 0, '2023-11-06 14:49:37', '2023-11-06 14:49:37', 0);
INSERT INTO `user_video` VALUES (38, 1672489513671, 40, 0, '2023-11-06 14:49:48', '2023-11-06 14:49:48', 0);
INSERT INTO `user_video` VALUES (39, 1672489513671, 41, 0, '2023-11-06 14:50:34', '2023-11-06 14:50:34', 0);
INSERT INTO `user_video` VALUES (40, 1672489513671, 42, 0, '2023-11-06 14:50:46', '2023-11-06 14:50:46', 0);
INSERT INTO `user_video` VALUES (41, 1672489513671, 43, 0, '2023-11-06 14:51:09', '2023-11-06 14:51:09', 0);
INSERT INTO `user_video` VALUES (42, 1672489513671, 44, 0, '2023-11-06 14:51:18', '2023-11-06 14:51:18', 0);
INSERT INTO `user_video` VALUES (43, 1672489513671, 45, 0, '2023-11-06 14:51:26', '2023-11-06 14:51:26', 0);
INSERT INTO `user_video` VALUES (44, 1672489513671, 46, 0, '2023-11-06 14:51:38', '2023-11-06 14:51:38', 0);

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `author_id` bigint NOT NULL COMMENT '视频上传的作者id',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '视频简介',
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频播放地址',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频封面地址',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `favorite_count` int NOT NULL DEFAULT 0 COMMENT '点赞量',
  `play_counts` int NOT NULL DEFAULT 0 COMMENT '播放次数',
  `comment_count` int NOT NULL DEFAULT 0 COMMENT '评论量',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除(0-未删, 1-已删)',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频类别',
  `collect_count` int NOT NULL DEFAULT 0 COMMENT '收藏次数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '`video`' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, 1672489513669, '这是一个关羽视频，他无所不能', 'http://s38oif2dm.hn-bkt.clouddn.com/86fdfcd94f621fb29ef8172dd9ec68f5.mp4', '。', '关羽启动', 1, 0, 0, '2023-01-28 16:49:38', '2023-11-06 14:12:27', 0, 'hot', 1);
INSERT INTO `video` VALUES (2, 1672489513671, '这是一个好看动人的视频', 'http://s38oif2dm.hn-bkt.clouddn.com/7621a726f3df74c6b9bec19aefdc7407.mp4', '.', '02女大学生', 100, 0, 0, '2023-10-29 21:52:05', '2023-11-01 11:32:55', 0, 'sports', 1);
INSERT INTO `video` VALUES (3, 1672489513669, '鸡头，鸡头，可爱的鸡头', 'http://s38oif2dm.hn-bkt.clouddn.com/547fe704fa852affa69b06a9535ecd7e.mp4', '。', '神奇的鸡头', 20, 0, 0, '2023-10-04 22:13:46', '2023-11-06 14:12:46', 0, 'hot', 1);
INSERT INTO `video` VALUES (4, 1672489513671, '快乐的soplay，欢迎大家来看哦', 'http://s38oif2dm.hn-bkt.clouddn.com/b868152d39f68221162ed58d112bbf76.mp4', '。', 'cospl', 33, 10, 21, '2023-10-03 11:29:10', '2023-11-06 14:13:19', 0, 'sports', 0);
INSERT INTO `video` VALUES (5, 1672489513671, '苦逼打工人的上班一天全记录', 'http://s38oif2dm.hn-bkt.clouddn.com/b34f19fc522fc7ef9f86d28ffa7cb0d7.mp4', '。', '上班的一天', 4532, 12, 1, '2023-10-05 11:31:03', '2023-11-06 14:13:38', 0, 'sports', 2);
INSERT INTO `video` VALUES (34, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-a2cfdfe4cbc7dbfa59ef1149d4407dbf.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-a2cfdfe4cbc7dbfa59ef1149d4407dbf.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:46:49', '2023-11-06 14:51:48', 0, 'sports', 0);
INSERT INTO `video` VALUES (35, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-8eed54c2db9ccf9f8d2e02937eed049e.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-8eed54c2db9ccf9f8d2e02937eed049e.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:48:38', '2023-11-06 14:51:48', 0, 'sports', 0);
INSERT INTO `video` VALUES (36, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-f1c1f3af8bf48697895613ce20411474.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-f1c1f3af8bf48697895613ce20411474.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:48:48', '2023-11-06 14:51:49', 0, 'sports', 0);
INSERT INTO `video` VALUES (37, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-7c667b66a2e13b2f4ffab002239a88d5.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-7c667b66a2e13b2f4ffab002239a88d5.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:48:59', '2023-11-06 14:51:49', 0, 'sports', 0);
INSERT INTO `video` VALUES (38, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-9ee91a170bf3dd7580acd0a5dbc7fdd6.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-9ee91a170bf3dd7580acd0a5dbc7fdd6.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:49:23', '2023-11-06 14:51:50', 0, 'sports', 0);
INSERT INTO `video` VALUES (39, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-fe815711003c6775feee0de6f2aa56f9.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-fe815711003c6775feee0de6f2aa56f9.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:49:37', '2023-11-06 14:51:50', 0, 'sports', 0);
INSERT INTO `video` VALUES (40, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-f166a00316ba3a00e0d00e43ae3b39bf.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-f166a00316ba3a00e0d00e43ae3b39bf.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:49:48', '2023-11-06 14:51:51', 0, 'sports', 0);
INSERT INTO `video` VALUES (41, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-f07e2d9b9f937be5694242532fdb2e89.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-f07e2d9b9f937be5694242532fdb2e89.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:50:34', '2023-11-06 14:51:52', 0, 'sports', 0);
INSERT INTO `video` VALUES (42, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-c049ce8d08d2131ef6b226603d66029e.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-c049ce8d08d2131ef6b226603d66029e.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:50:46', '2023-11-06 14:51:53', 0, 'hot', 0);
INSERT INTO `video` VALUES (43, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-0c26eb186af4e742ddffb410a25d0952.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-0c26eb186af4e742ddffb410a25d0952.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:51:09', '2023-11-06 14:51:54', 0, 'hot', 0);
INSERT INTO `video` VALUES (44, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-547fe704fa852affa69b06a9535ecd7e.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-547fe704fa852affa69b06a9535ecd7e.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:51:18', '2023-11-06 14:51:55', 0, 'hot', 0);
INSERT INTO `video` VALUES (45, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-b1f195a0eb539bbf837404decdeb7375.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-b1f195a0eb539bbf837404decdeb7375.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:51:26', '2023-11-06 14:51:55', 0, 'hot', 0);
INSERT INTO `video` VALUES (46, 1672489513671, '', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-c89e4100e4e6f462278af715efb61a49.mp4', 'http://s38oif2dm.hn-bkt.clouddn.com/1672489513671-c89e4100e4e6f462278af715efb61a49.jpg', '我的dy', 0, 0, 0, '2023-11-06 14:51:38', '2023-11-06 14:51:58', 0, 'hot', 0);

SET FOREIGN_KEY_CHECKS = 1;
