DROP TABLE IF EXISTS manage_email;

CREATE TABLE manage_email (
    id serial NOT NULL,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    create_by bigint NOT NULL,
    update_by bigint NOT NULL,
    "from" text NOT NULL,
    host text NOT NULL,
    port bigint NOT NULL,
    username text NOT NULL,
    password text NOT NULL,
    enable_ssl smallint NOT NULL,
    is_verify smallint NOT NULL,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS manage_menu;

CREATE TABLE manage_menu (
    id serial NOT NULL,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    create_by bigint NOT NULL,
    update_by bigint NOT NULL,
    status varchar(1) NOT NULL,
    parent_id bigint NOT NULL,
    menu_type varchar(1) NOT NULL,
    menu_name varchar(50) NOT NULL,
    hide_in_menu smallint NOT NULL,
    active_menu varchar(50) NOT NULL,
    "order" bigint NOT NULL,
    route_name varchar(255) NOT NULL,
    route_path varchar(255) NOT NULL,
    component varchar(255) NOT NULL,
    icon varchar(255) NOT NULL,
    icon_type varchar(1) NOT NULL,
    i18n_key varchar(255) NOT NULL,
    keep_alive smallint NOT NULL,
    href text NOT NULL,
    multi_tab smallint NOT NULL,
    fixed_index_in_tab bigint NOT NULL,
    query text NOT NULL,
    permissions text NOT NULL,
    constant smallint NOT NULL,
    button_code text NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO manage_menu (id, create_time, update_time, create_by, update_by, status, parent_id, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    (1,'2024-11-10 21:12:35','2024-11-10 21:12:35',0,0,'1',0,'2','关于',0,'',6,'about','/about','layout.base$view.about','fluent:book-information-24-regular','1','route.about',0,'',0,0,'[]','[]',0,''),
    (2,'2024-11-10 21:43:39','2024-11-10 21:43:39',0,0,'1',0,'2','首页',0,'',1,'home','/home','layout.base$view.home','mdi:monitor-dashboard','1','route.home',0,'',0,0,'[]','[]',0,''),
    (3,'2024-11-11 09:19:18','2024-11-11 09:19:18',0,0,'1',0,'1','系统管理',0,'',5,'manage','/manage','layout.base','carbon:cloud-service-management','1','route.manage',0,'',0,0,'[]','[]',0,''),
    (4,'2024-11-11 09:20:26','2024-11-11 09:20:26',0,0,'1',3,'2','用户管理',0,'',1,'manage_user','/manage/user','view.manage_user','ic:round-manage-accounts','1','route.manage_user',0,'',0,0,'[]','[{"code":"manage:user:list","desc":"用户列表"}]',0,''),
    (7,'2024-11-11 10:29:03','2024-11-11 10:29:03',0,0,'1',3,'2','角色管理',0,'',6,'manage_role','/manage/role','view.manage_role','carbon:user-role','1','route.manage_role',0,'',0,0,'[]','[{"code":"manage:role:list","desc":"角色列表"}]',0,''),
    (8,'2024-11-11 10:29:33','2024-11-11 10:29:33',0,0,'1',3,'2','菜单管理',0,'',5,'manage_menu','/manage/menu','view.manage_menu','material-symbols:route','1','route.manage_menu',0,'',0,0,'[]','[{"code":"manage:menu:list","desc":"菜单列表"}]',0,''),
    (13,'2024-11-25 05:23:25','2024-11-25 05:23:25',0,0,'1',0,'2','403',1,'',0,'403','/403','layout.blank$view.403','no-permission','2','route.403',0,'',0,0,'[]','[]',1,''),
    (14,'2024-11-25 05:24:40','2024-11-25 05:24:40',0,0,'1',0,'2','404',1,'',0,'404','/404','layout.blank$view.404','not-found','2','route.404',0,'',0,0,'[]','[]',1,''),
    (15,'2024-11-25 05:25:14','2024-11-25 05:25:14',0,0,'1',0,'2','500',1,'',0,'500','/500','layout.blank$view.500','service-error','2','route.500',0,'',0,0,'[]','[]',1,''),
    (17,'2024-11-25 05:52:32','2024-11-25 05:52:32',0,0,'1',0,'2','login',1,'',0,'login','/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?','layout.blank$view.login','no-icon','2','route.login',0,'',0,0,'[]','[]',1,''),
    (21,'2024-11-26 16:04:56','2024-11-26 16:04:56',0,0,'1',4,'3','创建用户',0,'',0,'','','','','1','button.manage.user.add',0,'',0,0,'[]','[{"code":"manage:user:add","desc":"新增用户"},{"code":"manage:role:getAll","desc":"获取所有角色"}]',0,'manage:user:add'),
    (22,'2024-11-28 11:13:13','2024-11-28 11:13:13',0,0,'1',4,'3','删除用户',0,'',1,'','','','','1','button.manage.user.delete',0,'',0,0,'[]','[{"code":"manage:user:delete","desc":"删除用户"}]',0,'manage:user:delete'),
    (23,'2024-11-28 11:16:09','2024-11-28 11:16:09',0,0,'1',4,'3','编辑用户',0,'',2,'','','','','1','button.manage.user.edit',0,'',0,0,'[]','[{"code":"manage:user:edit","desc":"编辑用户"},{"code":"manage:role:getAll","desc":"获取所有角色"}]',0,'manage:user:edit'),
    (25,'2024-11-28 11:25:25','2024-11-28 11:25:25',0,0,'1',7,'3','新增角色',0,'',0,'','','','','1','button.manage.role.add',0,'',0,0,'[]','[{"code":"manage:role:add","desc":"新增角色"}]',0,'manage:role:add'),
    (26,'2024-11-28 11:26:03','2024-11-28 11:26:03',0,0,'1',7,'3','编辑角色',0,'',1,'','','','','1','button.manage.role.edit',0,'',0,0,'[]','[{"code":"manage:role:edit","desc":"编辑角色"},{"code":"manage:role:getMenus","desc":"获取菜单权限"},{"code":"manage:role:setMenus","desc":"设置菜单权限"},{"code":"manage:menu:getAllPages","desc":"获取菜单页面"},{"code":"manage:menu:tree","desc":"获取菜单树"},{"code":"manage:role:updateHome","desc":"更新角色首页"},{"code":"manage:role:getHome","desc":"获取角色首页"}]',0,'manage:role:edit'),
    (27,'2024-11-28 11:27:10','2024-11-28 11:27:10',0,0,'1',7,'3','删除用户',0,'',2,'','','','','1','button.manage.role.delete',0,'',0,0,'[]','[{"code":"manage:role:delete","desc":"删除用户"}]',0,'manage:role:delete'),
    (29,'2024-11-28 12:22:58','2024-11-28 12:22:58',0,0,'1',8,'3','添加下一级菜单',0,'',0,'','','','','1','button.manage.menu.add',0,'',0,0,'[]','[{"code":"manage:menu:add","desc":"添加下一级菜单"}]',0,'manage:menu:add'),
    (30,'2024-11-28 12:23:59','2024-11-28 12:23:59',0,0,'1',8,'3','编辑菜单',0,'',1,'','','','','1','button.manage.menu.edit',0,'',0,0,'[]','[{"code":"manage:menu:edit","desc":"编辑菜单"}]',0,'manage:menu:edit'),
    (31,'2024-11-28 12:24:38','2024-11-28 12:24:38',0,0,'1',8,'3','删除菜单',0,'',2,'','','','','1','button.manage.menu.delete',0,'',0,0,'[]','[{"code":"manage:menu:delete","desc":"删除菜单"}]',0,'manage:menu:delete'),
    (32,'2024-11-28 15:59:02','2024-11-28 15:59:02',0,0,'1',0,'2','个人中心',1,'/manage/user',0,'user-center','/user-center','layout.base$view.user-center','','1','route.user-center',0,'',0,0,'[]','[]',0,''),
    (36,'2024-11-30 03:08:35','2024-11-30 03:08:35',0,0,'1',0,'1','文档',0,'',2,'document','/document','layout.base','mdi:file-document-multiple-outline','1','route.document',0,'',0,0,'[]','[]',0,''),
    (37,'2024-11-30 03:10:37','2024-11-30 03:10:37',0,0,'1',36,'2','项目文档',0,'',0,'document_project','/document/project','view.iframe-page','logo','2','route.document_project',0,'https://jzero.jaronnie.com',0,0,'[]','[]',0,''),
    (38,'2024-11-30 03:20:13','2024-11-30 03:20:13',0,0,'1',3,'2','用户详情',1,'',0,'manage_user-detail','/manage/user-detail/:id','view.manage_user-detail','','1','route.manage_user-detail',0,'',0,0,'[]','[]',0,'');

SELECT setval('manage_menu_id_seq', COALESCE((SELECT MAX(id)+1 FROM manage_menu), 1));

DROP TABLE IF EXISTS manage_role;

CREATE TABLE manage_role (
    id serial NOT NULL,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    create_by bigint NOT NULL,
    update_by bigint NOT NULL,
    name varchar(50) NOT NULL,
    status varchar(1) NOT NULL,
    code varchar(255) NOT NULL,
    "desc" text NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO manage_role (id, create_time, update_time, create_by, update_by, name, status, code, "desc")
VALUES
    (1,'2024-10-29 22:23:50','2024-10-29 22:23:50',1,0,'超级管理员','1','R_SUPER','超级管理员');

SELECT setval('manage_role_id_seq', COALESCE((SELECT MAX(id)+1 FROM manage_role), 1));

DROP TABLE IF EXISTS manage_role_menu;

CREATE TABLE manage_role_menu (
    id serial NOT NULL,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    create_by bigint NOT NULL,
    update_by bigint NOT NULL,
    role_id bigint NOT NULL,
    menu_id bigint NOT NULL,
    is_home smallint NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO manage_role_menu (id, create_time, update_time, create_by, update_by, role_id, menu_id, is_home)
VALUES
    (772,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,1,0),
    (773,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,8,0),
    (774,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,7,0),
    (775,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,4,0),
    (776,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,20,0),
    (777,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,2,1),
    (778,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,3,0),
    (779,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,21,0),
    (780,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,22,0),
    (781,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,23,0),
    (782,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,24,0),
    (783,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,25,0),
    (784,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,26,0),
    (785,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,27,0),
    (786,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,28,0),
    (787,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,29,0),
    (788,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,30,0),
    (789,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,31,0),
    (790,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,32,0),
    (791,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,35,0),
    (792,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,38,0),
    (793,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,36,0),
    (794,'2024-12-04 01:33:23','2024-12-04 01:33:23',0,0,1,37,0);

SELECT setval('manage_role_menu_id_seq', COALESCE((SELECT MAX(id)+1 FROM manage_role_menu), 1));

DROP TABLE IF EXISTS manage_user;

CREATE TABLE manage_user (
    id serial NOT NULL,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    create_by bigint NOT NULL,
    update_by bigint NOT NULL,
    username varchar(30) NOT NULL UNIQUE,
    password varchar(100) NOT NULL,
    nickname varchar(30) NOT NULL,
    gender varchar(1) NOT NULL,
    phone varchar(20) NOT NULL,
    status varchar(1) NOT NULL,
    email varchar(100) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO manage_user (id, create_time, update_time, create_by, update_by, username, password, nickname, gender, phone, status, email)
VALUES
    (1,'2024-10-24 09:45:00','2024-10-31 09:40:13',0,0,'admin','123456','超级管理员','1','','1','');

SELECT setval('manage_user_id_seq', COALESCE((SELECT MAX(id)+1 FROM manage_user), 1));

DROP TABLE IF EXISTS manage_user_role;

CREATE TABLE manage_user_role (
    id serial NOT NULL,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    create_by bigint NOT NULL,
    update_by bigint NOT NULL,
    user_id bigint NOT NULL,
    role_id bigint NOT NULL,
    PRIMARY KEY (id)
);


INSERT INTO manage_user_role (id, create_time, update_time, create_by, update_by, user_id, role_id)
VALUES
    (1,'2024-10-31 09:40:13','2024-10-31 09:40:13',0,0,1,1),
    (4,'2024-11-28 12:49:54','2024-11-28 12:49:54',0,0,2,2),
    (5,'2024-11-28 12:50:29','2024-11-28 12:50:29',0,0,3,2);

SELECT setval('manage_user_role_id_seq', COALESCE((SELECT MAX(id)+1 FROM manage_user_role), 1));