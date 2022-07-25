/*
Navicat MySQL Data Transfer

Source Server         : 本地测试
Source Server Version : 50726
Source Host           : 127.0.0.1:3306
Source Database       : ginstudy

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2022-07-22 08:35:14
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `auth_rule`;
CREATE TABLE `auth_rule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `pathname` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `ismenu` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否启用 默认1 菜单 0 文件',
  `created` int(11) DEFAULT NULL,
  `updated` int(11) DEFAULT NULL,
  `deletetime` int(11) DEFAULT NULL,
  `weigh` int(11) DEFAULT NULL,
  `status` varchar(40) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `children` text,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=257 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of auth_rule
-- ----------------------------
INSERT INTO `auth_rule` VALUES ('37', '36', 'file', 'el-icon-notebook-2', '/system/admin', '用户列表', '用户列表', '0', '1614759419', null, null, '1', 'normal', 'system/admin/index', null);
INSERT INTO `auth_rule` VALUES ('38', '36', 'file', 'el-icon-notebook-2', '/system/rules', '前端菜单', '前端菜单', '0', '1614759461', null, null, '2', 'normal', 'system/rules/index', null);
INSERT INTO `auth_rule` VALUES ('39', '36', 'file', 'el-icon-notebook-2', '/achievements/generalityGoal/departmentAssessment', '职能部门考核', '职能部门考核', '0', '1614759515', null, null, '3', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('36', '0', 'file', 'el-icon-menu', '/system', '系统管理', '用户管理', '1', '1621413412', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('225', '0', 'file', 'el-icon-menu', '/personalPerformance', '个人绩效管理', '个人绩效管理', '1', '1619409666', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('226', '225', 'file', 'el-icon-notebook-2', '/personalPerformance/allStaffEvaluation', '全员民主测评', '全员民主测评', '0', '1619415264', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('227', '225', 'file', 'el-icon-notebook-2', '/personalPerformance/quarterAchievements', '季度绩效测评', '季度绩效测评', '0', '1619415355', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('228', '225', 'file', 'el-icon-notebook-2', '/personalPerformance/yearsAchievements', '年度绩效测评', '年度绩效测评', '0', '1619415408', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('229', '225', 'file', 'el-icon-notebook-2', '/personalPerformance/personnelClass', '人员分类管理', '人员分类管理', '0', '1619416533', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('235', '0', 'file', 'el-icon-menu', '/judicature', '司法绩效管理', '司法绩效管理', '1', '1620270606', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('236', '235', 'file', 'el-icon-tickets', '/judicature/petitionInvolLitigation', '涉诉信访统计', '涉诉信访统计', '0', '1620270690', null, null, '60', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('237', '235', 'file', 'el-icon-tickets', '/judicature/caseQuality', '案件质量评查', '案件质量评查', '0', '1620270741', null, null, '50', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('238', '235', 'file', 'el-icon-tickets', '/judicature/caseFiling', '案件归档登记', '案件归档登记', '0', '1620270798', null, null, '40', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('239', '235', 'file', 'el-icon-tickets', '/judicature/performanceIndex', '绩效指数统计', '绩效指数统计', '0', '1620270850', null, null, '30', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('240', '235', 'file', 'el-icon-tickets', '/judicature/departmentPerformance', '部门绩效评估', '部门绩效评估', '0', '1620270903', null, null, '20', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('241', '235', 'file', 'el-icon-tickets', '/judicature/personalPerformance', '个人绩效评估', '个人绩效评估', '0', '1620270951', null, null, '10', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('243', '0', 'file', 'el-icon-menu', '/postGoal', '岗位目标管理', '岗位目标管理', '1', '1620442073', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('245', '243', 'file', 'el-icon-notebook-2', '/postGoal/jobTarget', '岗位目标', '岗位目标', '0', '1620442243', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('246', '243', 'file', 'el-icon-notebook-2', '/postGoal/personnelTasks', '人员任职', '人员任职', '0', '1620442339', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('247', '243', 'file', 'el-icon-notebook-2', '/postGoal/spotCheck', '平时抽查考核', '平时抽查考核', '0', '1620442389', null, null, '0', 'normal', null, null);
INSERT INTO `auth_rule` VALUES ('248', '243', 'file', 'el-icon-notebook-2', '/postGoal/quarterlyTarget', '季度目标考核', '季度目标考核', '0', '1620442420', null, null, '0', 'normal', null, null);
