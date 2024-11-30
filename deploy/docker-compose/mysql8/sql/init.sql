SET NAMES utf8mb4;

# 转储表 casbin_rule
# ------------------------------------------------------------

DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;

INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES
	(80,'p','2','manage:menu:list','','','',''),
	(98,'p','2','manage:role:list','','','',''),
	(117,'p','2','manage:user:list','','','',''),
	(137,'p','1','manage:user:list','','','',''),
	(138,'p','1','manage:role:list','','','',''),
	(139,'p','1','manage:menu:list','','','',''),
	(140,'p','1','manage:user:add','','','',''),
	(141,'p','1','manage:role:getAll','','','',''),
	(142,'p','1','manage:user:delete','','','',''),
	(143,'p','1','manage:user:edit','','','',''),
	(144,'p','1','manage:role:getAll','','','',''),
	(145,'p','1','manage:role:add','','','',''),
	(146,'p','1','manage:role:edit','','','',''),
	(147,'p','1','manage:role:getMenus','','','',''),
	(148,'p','1','manage:role:setMenus','','','',''),
	(149,'p','1','manage:menu:getAllPages','','','',''),
	(150,'p','1','manage:menu:tree','','','',''),
	(151,'p','1','manage:role:delete','','','',''),
	(152,'p','1','manage:menu:add','','','',''),
	(153,'p','1','manage:menu:edit','','','',''),
	(154,'p','1','manage:menu:delete','','','','');

/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 manage_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_menu`;

CREATE TABLE `manage_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `status` varchar(1) NOT NULL,
  `parent_id` bigint NOT NULL,
  `menu_type` varchar(1) NOT NULL,
  `menu_name` varchar(50) NOT NULL,
  `hide_in_menu` tinyint(1) NOT NULL,
  `active_menu` varchar(50) DEFAULT NULL,
  `order` bigint NOT NULL,
  `route_name` varchar(255) NOT NULL,
  `route_path` varchar(255) NOT NULL,
  `component` varchar(255) NOT NULL,
  `icon` varchar(255) NOT NULL,
  `icon_type` varchar(1) NOT NULL,
  `i18n_key` varchar(255) NOT NULL,
  `keep_alive` tinyint(1) NOT NULL,
  `href` longtext,
  `multi_tab` tinyint(1) DEFAULT NULL,
  `fixed_index_in_tab` bigint DEFAULT NULL,
  `query` longtext,
  `permissions` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `constant` tinyint(1) NOT NULL,
  `button_code` longtext,
  `iframe_page_url` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_menu` WRITE;
/*!40000 ALTER TABLE `manage_menu` DISABLE KEYS */;

INSERT INTO `manage_menu` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `status`, `parent_id`, `menu_type`, `menu_name`, `hide_in_menu`, `active_menu`, `order`, `route_name`, `route_path`, `component`, `icon`, `icon_type`, `i18n_key`, `keep_alive`, `href`, `multi_tab`, `fixed_index_in_tab`, `query`, `permissions`, `constant`, `button_code`, `iframe_page_url`)
VALUES
	(1,'2024-11-10 21:12:35','2024-11-10 21:12:35',NULL,NULL,'1',0,'2','关于',0,'',6,'about','/about','layout.base$view.about','fluent:book-information-24-regular','1','route.about',0,'',0,NULL,'[]','[]',0,NULL,NULL),
	(2,'2024-11-10 21:43:39','2024-11-10 21:43:39',NULL,NULL,'1',0,'2','首页',0,'',1,'home','/home','layout.base$view.home','mdi:monitor-dashboard','1','route.home',0,'',0,NULL,'[]','[]',0,NULL,NULL),
	(3,'2024-11-11 09:19:18','2024-11-11 09:19:18',NULL,NULL,'1',0,'1','系统管理',0,'',5,'manage','/manage','layout.base','carbon:cloud-service-management','1','route.manage',0,'',0,NULL,'[]','[]',0,NULL,NULL),
	(4,'2024-11-11 09:20:26','2024-11-11 09:20:26',NULL,NULL,'1',3,'2','用户管理',0,'',1,'manage_user','/manage/user','view.manage_user','ic:round-manage-accounts','1','route.manage_user',0,'',0,NULL,'[]','[{\"code\":\"manage:user:list\",\"desc\":\"用户列表\"}]',0,'',NULL),
	(7,'2024-11-11 10:29:03','2024-11-11 10:29:03',NULL,NULL,'1',3,'2','角色管理',0,'',2,'manage_role','/manage/role','view.manage_role','carbon:user-role','1','route.manage_role',0,'',0,NULL,'[]','[{\"code\":\"manage:role:list\",\"desc\":\"角色列表\"}]',0,'',NULL),
	(8,'2024-11-11 10:29:33','2024-11-11 10:29:33',NULL,NULL,'1',3,'2','菜单管理',0,'',5,'manage_menu','/manage/menu','view.manage_menu','material-symbols:route','1','route.manage_menu',0,'',0,NULL,'[]','[{\"code\":\"manage:menu:list\",\"desc\":\"菜单列表\"}]',0,'',NULL),
	(13,'2024-11-25 05:23:25','2024-11-25 05:23:25',NULL,NULL,'1',0,'2','403',1,'',0,'403','/403','layout.blank$view.403','no-permission','2','route.403',0,'',0,NULL,'[]','[]',1,NULL,NULL),
	(14,'2024-11-25 05:24:40','2024-11-25 05:24:40',NULL,NULL,'1',0,'2','404',1,'',0,'404','/404','layout.blank$view.404','not-found','2','route.404',0,'',0,NULL,'[]','[]',1,NULL,NULL),
	(15,'2024-11-25 05:25:14','2024-11-25 05:25:14',NULL,NULL,'1',0,'2','500',1,'',0,'500','/500','layout.blank$view.500','service-error','2','route.500',0,'',0,NULL,'[]','[]',1,NULL,NULL),
	(17,'2024-11-25 05:52:32','2024-11-25 05:52:32',NULL,NULL,'1',0,'2','login',1,'',0,'login','/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?','layout.blank$view.login','no-icon','2','route.login',0,'',0,NULL,'[]','[]',1,NULL,NULL),
	(21,'2024-11-26 16:04:56','2024-11-26 16:04:56',NULL,NULL,'1',4,'3','创建用户',0,'',0,'','','','','1','button.manage.user.add',0,'',0,NULL,'[]','[{\"code\":\"manage:user:add\",\"desc\":\"新增用户\"},{\"code\":\"manage:role:getAll\",\"desc\":\"获取所有角色\"}]',0,'manage:user:add',NULL),
	(22,'2024-11-28 11:13:13','2024-11-28 11:13:13',NULL,NULL,'1',4,'3','删除用户',0,'',1,'','','','','1','button.manage.user.delete',0,'',0,NULL,'[]','[{\"code\":\"manage:user:delete\",\"desc\":\"删除用户\"}]',0,'manage:user:delete',NULL),
	(23,'2024-11-28 11:16:09','2024-11-28 11:16:09',NULL,NULL,'1',4,'3','编辑用户',0,'',2,'','','','','1','button.manage.user.edit',0,'',0,NULL,'[]','[{\"code\":\"manage:user:edit\",\"desc\":\"编辑用户\"},{\"code\":\"manage:role:getAll\",\"desc\":\"获取所有角色\"}]',0,'manage:user:edit',NULL),
	(25,'2024-11-28 11:25:25','2024-11-28 11:25:25',NULL,NULL,'1',7,'3','新增角色',0,'',0,'','','','','1','button.manage.role.add',0,'',0,NULL,'[]','[{\"code\":\"manage:role:add\",\"desc\":\"新增角色\"}]',0,'manage:role:add',NULL),
	(26,'2024-11-28 11:26:03','2024-11-28 11:26:03',NULL,NULL,'1',7,'3','编辑角色',0,'',1,'','','','','1','button.manage.role.edit',0,'',0,NULL,'[]','[{\"code\":\"manage:role:edit\",\"desc\":\"编辑角色\"},{\"code\":\"manage:role:getMenus\",\"desc\":\"获取菜单权限\"},{\"code\":\"manage:role:setMenus\",\"desc\":\"设置菜单权限\"},{\"code\":\"manage:menu:getAllPages\",\"desc\":\"获取菜单页面\"},{\"code\":\"manage:menu:tree\",\"desc\":\"获取菜单树\"}]',0,'manage:role:edit',NULL),
	(27,'2024-11-28 11:27:10','2024-11-28 11:27:10',NULL,NULL,'1',7,'3','删除用户',0,'',2,'','','','','1','button.manage.role.delete',0,'',0,NULL,'[]','[{\"code\":\"manage:role:delete\",\"desc\":\"删除用户\"}]',0,'manage:role:delete',NULL),
	(29,'2024-11-28 12:22:58','2024-11-28 12:22:58',NULL,NULL,'1',8,'3','添加下一级菜单',0,'',0,'','','','','1','button.manage.menu.add',0,'',0,NULL,'[]','[{\"code\":\"manage:menu:add\",\"desc\":\"添加下一级菜单\"}]',0,'manage:menu:add',NULL),
	(30,'2024-11-28 12:23:59','2024-11-28 12:23:59',NULL,NULL,'1',8,'3','编辑菜单',0,'',1,'','','','','1','button.manage.menu.edit',0,'',0,NULL,'[]','[{\"code\":\"manage:menu:edit\",\"desc\":\"编辑菜单\"}]',0,'manage:menu:edit',NULL),
	(31,'2024-11-28 12:24:38','2024-11-28 12:24:38',NULL,NULL,'1',8,'3','删除菜单',0,'',2,'','','','','1','button.manage.menu.delete',0,'',0,NULL,'[]','[{\"code\":\"manage:menu:delete\",\"desc\":\"删除菜单\"}]',0,'manage:menu:delete',NULL),
	(32,'2024-11-28 15:59:02','2024-11-28 15:59:02',NULL,NULL,'1',0,'2','个人中心',1,'/manage/user',0,'user-center','/user-center','layout.base$view.user-center','','1','route.user-center',0,'',0,NULL,'[]','[]',0,'',NULL),
	(36,'2024-11-30 03:08:35','2024-11-30 03:08:35',NULL,NULL,'1',0,'1','文档',0,'',2,'document','/document','layout.base','mdi:file-document-multiple-outline','1','route.document',0,'',0,NULL,'[]','[]',0,'',NULL),
	(37,'2024-11-30 03:10:37','2024-11-30 03:10:37',NULL,NULL,'1',36,'2','项目文档',0,'',0,'document_project','/document/project','view.iframe-page','logo','2','route.document_project',0,'https://jzero.jaronnie.com',0,NULL,'[]','[]',0,'',NULL),
	(38,'2024-11-30 03:20:13','2024-11-30 03:20:13',NULL,NULL,'1',3,'2','用户详情',1,'',0,'manage_user-detail','/manage/user-detail/:id','view.manage_user-detail','','1','route.manage_user-detail',0,'',0,NULL,'[]','[]',0,'',NULL);

/*!40000 ALTER TABLE `manage_menu` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 manage_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_role`;

CREATE TABLE `manage_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `name` varchar(50) NOT NULL,
  `status` varchar(1) NOT NULL,
  `code` varchar(255) NOT NULL,
  `desc` longtext NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_role` WRITE;
/*!40000 ALTER TABLE `manage_role` DISABLE KEYS */;

INSERT INTO `manage_role` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `name`, `status`, `code`, `desc`)
VALUES
	(1,'2024-10-29 22:23:50','2024-10-29 22:23:50',1,NULL,'超级管理员','1','R_SUPER','超级管理员'),
	(2,'2024-11-28 12:46:13','2024-11-28 12:46:13',1,NULL,'测试用户','1','R_TEST','测试用户');

/*!40000 ALTER TABLE `manage_role` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 manage_role_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_role_menu`;

CREATE TABLE `manage_role_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `role_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_role_menu` WRITE;
/*!40000 ALTER TABLE `manage_role_menu` DISABLE KEYS */;

INSERT INTO `manage_role_menu` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `role_id`, `menu_id`)
VALUES
	(298,'2024-11-29 14:18:09','2024-11-29 14:18:09',NULL,NULL,2,2),
	(299,'2024-11-29 14:18:09','2024-11-29 14:18:09',NULL,NULL,2,1),
	(300,'2024-11-29 14:18:09','2024-11-29 14:18:09',NULL,NULL,2,4),
	(301,'2024-11-29 14:18:09','2024-11-29 14:18:09',NULL,NULL,2,7),
	(302,'2024-11-29 14:18:09','2024-11-29 14:18:09',NULL,NULL,2,8),
	(303,'2024-11-29 14:18:09','2024-11-29 14:18:09',NULL,NULL,2,20),
	(304,'2024-11-29 14:18:09','2024-11-29 14:18:09',NULL,NULL,2,24),
	(305,'2024-11-29 14:18:09','2024-11-29 14:18:09',NULL,NULL,2,28),
	(306,'2024-11-29 14:18:09','2024-11-29 14:18:09',NULL,NULL,2,32),
	(369,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,1),
	(370,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,8),
	(371,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,7),
	(372,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,4),
	(373,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,20),
	(374,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,2),
	(375,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,3),
	(376,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,21),
	(377,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,22),
	(378,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,23),
	(379,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,24),
	(380,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,25),
	(381,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,26),
	(382,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,27),
	(383,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,28),
	(384,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,29),
	(385,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,30),
	(386,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,31),
	(387,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,32),
	(388,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,35),
	(389,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,36),
	(390,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,37),
	(391,'2024-11-30 03:22:18','2024-11-30 03:22:18',NULL,NULL,1,38);

/*!40000 ALTER TABLE `manage_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 manage_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_user`;

CREATE TABLE `manage_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `username` varchar(30) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nickname` varchar(30) NOT NULL,
  `gender` varchar(1) NOT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `status` varchar(1) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_manage_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_user` WRITE;
/*!40000 ALTER TABLE `manage_user` DISABLE KEYS */;

INSERT INTO `manage_user` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `username`, `password`, `nickname`, `gender`, `phone`, `status`, `email`)
VALUES
	(1,'2024-10-24 09:45:00','2024-10-31 09:40:13',NULL,NULL,'jzero','123456','jzero','1','','1','jaron@jaronnie.com'),
	(3,'2024-11-28 12:50:29','2024-11-28 12:50:29',NULL,NULL,'test','123456','测试用户','1','1234567890','1','test@qq.com');

/*!40000 ALTER TABLE `manage_user` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 manage_user_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_user_role`;

CREATE TABLE `manage_user_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `user_id` bigint NOT NULL,
  `role_id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_user_role` WRITE;
/*!40000 ALTER TABLE `manage_user_role` DISABLE KEYS */;

INSERT INTO `manage_user_role` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `user_id`, `role_id`)
VALUES
	(1,'2024-10-31 09:40:13','2024-10-31 09:40:13',NULL,NULL,1,1),
	(4,'2024-11-28 12:49:54','2024-11-28 12:49:54',NULL,NULL,2,2),
	(5,'2024-11-28 12:50:29','2024-11-28 12:50:29',NULL,NULL,3,2);

/*!40000 ALTER TABLE `manage_user_role` ENABLE KEYS */;
UNLOCK TABLES;
