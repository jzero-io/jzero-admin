-- Create function to automatically update update_time
CREATE OR REPLACE FUNCTION update_update_time_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

DROP TABLE IF EXISTS "manage_email";

CREATE TABLE "manage_email" (
    id serial NOT NULL,
    uuid varchar(36) NOT NULL UNIQUE,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "from" text NOT NULL,
    host text NOT NULL,
    port bigint NOT NULL,
    username text NOT NULL,
    password text NOT NULL,
    enable_ssl smallint NOT NULL,
    is_verify smallint NOT NULL,
    PRIMARY KEY (id)
);

CREATE TRIGGER update_manage_email_update_time BEFORE UPDATE ON "manage_email"
FOR EACH ROW EXECUTE FUNCTION update_update_time_column();

DROP TABLE IF EXISTS "manage_menu";

CREATE TABLE "manage_menu" (
    id serial NOT NULL,
    uuid varchar(36) NOT NULL UNIQUE,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status varchar(1) NOT NULL,
    parent_uuid varchar(36) NOT NULL,
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

CREATE TRIGGER update_manage_menu_update_time BEFORE UPDATE ON "manage_menu"
FOR EACH ROW EXECUTE FUNCTION update_update_time_column();

INSERT INTO "manage_menu" (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('93b45888-417e-4963-a08c-50ccdf2460e3','2024-11-10 21:12:35','2024-11-10 21:12:35','1','','2','关于',0,'',6,'about','/about','layout.base$view.about','fluent:book-information-24-regular','1','route.about',0,'',0,0,'[]','[]',0,''),
    ('f7e8d9c6-b5a4-4382-8271-605f4e3d2c1b','2024-11-10 21:43:39','2024-11-10 21:43:39','1','','2','首页',0,'',1,'home','/home','layout.base$view.home','mdi:monitor-dashboard','1','route.home',0,'',0,0,'[]','[]',0,''),
    ('a1b2c3d4-e5f6-4782-91a0-b9c8d7e6f5a4','2024-11-11 09:19:18','2024-11-11 09:19:18','1','','1','系统管理',0,'',5,'manage','/manage','layout.base','carbon:cloud-service-management','1','route.manage',0,'',0,0,'[]','[]',0,''),
    ('b3c4d5e6-f7a8-4912-83b4-c5d6e7f8a9b0','2024-11-11 09:20:26','2024-11-11 09:20:26','1','a1b2c3d4-e5f6-4782-91a0-b9c8d7e6f5a4','2','用户管理',0,'',1,'manage_user','/manage/user','view.manage_user','ic:round-manage-accounts','1','route.manage_user',0,'',0,0,'[]','[{"code":"manage:user:list","desc":"用户列表"}]',0,''),
    ('c5d6e7f8-a9b0-4123-95c6-d7e8f9a0b1c2','2024-11-11 10:29:03','2024-11-11 10:29:03','1','a1b2c3d4-e5f6-4782-91a0-b9c8d7e6f5a4','2','角色管理',0,'',6,'manage_role','/manage/role','view.manage_role','carbon:user-role','1','route.manage_role',0,'',0,0,'[]','[{"code":"manage:role:list","desc":"角色列表"}]',0,''),
    ('d7e8f9a0-b1c2-4334-87d8-e9f0a1b2c3d4','2024-11-11 10:29:33','2024-11-11 10:29:33','1','a1b2c3d4-e5f6-4782-91a0-b9c8d7e6f5a4','2','菜单管理',0,'',5,'manage_menu','/manage/menu','view.manage_menu','material-symbols:route','1','route.manage_menu',0,'',0,0,'[]','[{"code":"manage:menu:list","desc":"菜单列表"}]',0,''),
    ('e9f0a1b2-c3d4-4556-79e0-f1a2b3c4d5e6','2024-11-25 05:23:25','2024-11-25 05:23:25','1','','2','403',1,'',0,'403','/403','layout.blank$view.403','no-permission','2','route.403',0,'',0,0,'[]','[]',1,''),
    ('f1a2b3c4-d5e6-4767-8af2-a3b4c5d6e7f8','2024-11-25 05:24:40','2024-11-25 05:24:40','1','','2','404',1,'',0,'404','/404','layout.blank$view.404','not-found','2','route.404',0,'',0,0,'[]','[]',1,''),
    ('a3b4c5d6-e7f8-4990-81a4-b5c6d7e8f9a0','2024-11-25 05:25:14','2024-11-25 05:25:14','1','','2','500',1,'',0,'500','/500','layout.blank$view.500','service-error','2','route.500',0,'',0,0,'[]','[]',1,''),
    ('b5c6d7e8-f9a0-41a2-93b6-c7d8e9f0a1b2','2024-11-25 05:52:32','2024-11-25 05:52:32','1','','2','login',1,'',0,'login','/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?','layout.blank$view.login','no-icon','2','route.login',0,'',0,0,'[]','[]',1,''),
    ('c7d8e9f0-a1b2-43b4-85c8-d9e0f1a2b3c4','2024-11-26 16:04:56','2024-11-26 16:04:56','1','b3c4d5e6-f7a8-4912-83b4-c5d6e7f8a9b0','3','创建用户',0,'',0,'','','','','1','button.manage.user.add',0,'',0,0,'[]','[{"code":"manage:user:add","desc":"新增用户"},{"code":"manage:role:getAll","desc":"获取所有角色"}]',0,'manage:user:add'),
    ('d9e0f1a2-b3c4-65c6-97d0-e1f2a3b4c5d6','2024-11-28 11:13:13','2024-11-28 11:13:13','1','b3c4d5e6-f7a8-4912-83b4-c5d6e7f8a9b0','3','删除用户',0,'',1,'','','','','1','button.manage.user.delete',0,'',0,0,'[]','[{"code":"manage:user:delete","desc":"删除用户"}]',0,'manage:user:delete'),
    ('e1f2a3b4-c5d6-87d8-a9e2-f3a4b5c6d7e8','2024-11-28 11:16:09','2024-11-28 11:16:09','1','b3c4d5e6-f7a8-4912-83b4-c5d6e7f8a9b0','3','编辑用户',0,'',2,'','','','','1','button.manage.user.edit',0,'',0,0,'[]','[{"code":"manage:user:edit","desc":"编辑用户"},{"code":"manage:role:getAll","desc":"获取所有角色"}]',0,'manage:user:edit'),
    ('f3a4b5c6-d7e8-a9e0-b1f4-a5b6c7d8e9f0','2024-11-28 11:25:25','2024-11-28 11:25:25','1','c5d6e7f8-a9b0-4123-95c6-d7e8f9a0b1c2','3','新增角色',0,'',0,'','','','','1','button.manage.role.add',0,'',0,0,'[]','[{"code":"manage:role:add","desc":"新增角色"}]',0,'manage:role:add'),
    ('a5b6c7d8-e9f0-b1f2-c3a6-b7c8d9e0f1a2','2024-11-28 11:26:03','2024-11-28 11:26:03','1','c5d6e7f8-a9b0-4123-95c6-d7e8f9a0b1c2','3','编辑角色',0,'',1,'','','','','1','button.manage.role.edit',0,'',0,0,'[]','[{"code":"manage:role:edit","desc":"编辑角色"},{"code":"manage:role:getMenus","desc":"获取菜单权限"},{"code":"manage:role:setMenus","desc":"设置菜单权限"},{"code":"manage:menu:getAllPages","desc":"获取菜单页面"},{"code":"manage:menu:tree","desc":"获取菜单树"},{"code":"manage:role:updateHome","desc":"更新角色首页"},{"code":"manage:role:getHome","desc":"获取角色首页"}]',0,'manage:role:edit'),
    ('b7c8d9e0-f1a2-c3d4-e5b8-c9d0e1f2a3b4','2024-11-28 11:27:10','2024-11-28 11:27:10','1','c5d6e7f8-a9b0-4123-95c6-d7e8f9a0b1c2','3','删除用户',0,'',2,'','','','','1','button.manage.role.delete',0,'',0,0,'[]','[{"code":"manage:role:delete","desc":"删除用户"}]',0,'manage:role:delete'),
    ('c9d0e1f2-a3b4-d5e6-f7c0-d1e2f3a4b5c6','2024-11-28 12:22:58','2024-11-28 12:22:58','1','d7e8f9a0-b1c2-4334-87d8-e9f0a1b2c3d4','3','添加下一级菜单',0,'',0,'','','','','1','button.manage.menu.add',0,'',0,0,'[]','[{"code":"manage:menu:add","desc":"添加下一级菜单"}]',0,'manage:menu:add'),
    ('d1e2f3a4-b5c6-e7f8-a9d2-e3f4a5b6c7d8','2024-11-28 12:23:59','2024-11-28 12:23:59','1','d7e8f9a0-b1c2-4334-87d8-e9f0a1b2c3d4','3','编辑菜单',0,'',1,'','','','','1','button.manage.menu.edit',0,'',0,0,'[]','[{"code":"manage:menu:edit","desc":"编辑菜单"}]',0,'manage:menu:edit'),
    ('e3f4a5b6-c7d8-f9a0-b1e4-f5a6b7c8d9e0','2024-11-28 12:24:38','2024-11-28 12:24:38','1','d7e8f9a0-b1c2-4334-87d8-e9f0a1b2c3d4','3','删除菜单',0,'',2,'','','','','1','button.manage.menu.delete',0,'',0,0,'[]','[{"code":"manage:menu:delete","desc":"删除菜单"}]',0,'manage:menu:delete'),
    ('f5a6b7c8-d9e0-a1b2-c3f6-a7b8c9d0e1f2','2024-11-28 15:59:02','2024-11-28 15:59:02','1','','2','个人中心',1,'/manage/user',0,'user-center','/user-center','layout.base$view.user-center','','1','route.user-center',0,'',0,0,'[]','[]',0,''),
    ('a7b8c9d0-e1f2-b3c4-d5a8-b9c0d1e2f3a4','2024-11-30 03:08:35','2024-11-30 03:08:35','1','','1','文档',0,'',2,'document','/document','layout.base','mdi:file-document-multiple-outline','1','route.document',0,'',0,0,'[]','[]',0,''),
    ('b9c0d1e2-f3a4-c5d6-e7b0-c1d2e3f4a5b6','2024-11-30 03:10:37','2024-11-30 03:10:37','1','a7b8c9d0-e1f2-b3c4-d5a8-b9c0d1e2f3a4','2','项目文档',0,'',0,'document_project','/document/project','view.iframe-page','logo','2','route.document_project',0,'https://jzero.jaronnie.com',0,0,'[]','[]',0,''),
    ('c1d2e3f4-a5b6-d7e8-f9c2-d3e4f5a6b7c8','2024-11-30 03:20:13','2024-11-30 03:20:13','1','a1b2c3d4-e5f6-4782-91a0-b9c8d7e6f5a4','2','用户详情',1,'',0,'manage_user-detail','/manage/user-detail/:id','view.manage_user-detail','','1','route.manage_user-detail',0,'',0,0,'[]','[]',0,'');

DROP TABLE IF EXISTS "manage_role";

CREATE TABLE "manage_role" (
    id serial NOT NULL,
    uuid varchar(36) NOT NULL UNIQUE,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(50) NOT NULL,
    status varchar(1) NOT NULL,
    code varchar(255) NOT NULL,
    "desc" text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TRIGGER update_manage_role_update_time BEFORE UPDATE ON "manage_role"
FOR EACH ROW EXECUTE FUNCTION update_update_time_column();

INSERT INTO "manage_role" (uuid, create_time, update_time, name, status, code, "desc")
VALUES
    ('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','2024-10-29 22:23:50','2024-10-29 22:23:50','超级管理员','1','R_SUPER','超级管理员');


DROP TABLE IF EXISTS "manage_role_menu";

CREATE TABLE "manage_role_menu" (
    id serial NOT NULL,
    uuid varchar(36) NOT NULL UNIQUE,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    role_uuid varchar(36) NOT NULL,
    menu_uuid varchar(36) NOT NULL,
    is_home smallint NOT NULL,
    PRIMARY KEY (id)
);

CREATE TRIGGER update_manage_role_menu_update_time BEFORE UPDATE ON "manage_role_menu"
FOR EACH ROW EXECUTE FUNCTION update_update_time_column();

INSERT INTO "manage_role_menu" (uuid, create_time, update_time, role_uuid, menu_uuid, is_home)
VALUES
    ('2b3c4d5e-6f7a-8b9c-0d1e-2f3a4b5c6d7e','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','93b45888-417e-4963-a08c-50ccdf2460e3',0),
    ('3c4d5e6f-7a8b-9c0d-1e2f-3a4b5c6d7e8f','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','d7e8f9a0-b1c2-4334-87d8-e9f0a1b2c3d4',0),
    ('4d5e6f7a-8b9c-0d1e-2f3a-4b5c6d7e8f9a','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','c5d6e7f8-a9b0-4123-95c6-d7e8f9a0b1c2',0),
    ('5e6f7a8b-9c0d-1e2f-3a4b-5c6d7e8f9a0b','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','b3c4d5e6-f7a8-4912-83b4-c5d6e7f8a9b0',0),
    ('6f7a8b9c-0d1e-2f3a-4b5c-6d7e8f9a0b1c','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','f7e8d9c6-b5a4-4382-8271-605f4e3d2c1b',1),
    ('7a8b9c0d-1e2f-3a4b-5c6d-7e8f9a0b1c2d','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a1b2c3d4-e5f6-4782-91a0-b9c8d7e6f5a4',0),
    ('8b9c0d1e-2f3a-4b5c-6d7e-8f9a0b1c2d3e','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','c7d8e9f0-a1b2-43b4-85c8-d9e0f1a2b3c4',0),
    ('9c0d1e2f-3a4b-5c6d-7e8f-9a0b1c2d3e4f','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','d9e0f1a2-b3c4-65c6-97d0-e1f2a3b4c5d6',0),
    ('0d1e2f3a-4b5c-6d7e-8f9a-0b1c2d3e4f5a','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','e1f2a3b4-c5d6-87d8-a9e2-f3a4b5c6d7e8',0),
    ('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a6b','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','f3a4b5c6-d7e8-a9e0-b1f4-a5b6c7d8e9f0',0),
    ('2f3a4b5c-6d7e-8f9a-0b1c-2d3e4f5a6b7c','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a5b6c7d8-e9f0-b1f2-c3a6-b7c8d9e0f1a2',0),
    ('3a4b5c6d-7e8f-9a0b-1c2d-3e4f5a6b7c8d','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','b7c8d9e0-f1a2-c3d4-e5b8-c9d0e1f2a3b4',0),
    ('4b5c6d7e-8f9a-0b1c-2d3e-4f5a6b7c8d9e','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','c9d0e1f2-a3b4-d5e6-f7c0-d1e2f3a4b5c6',0),
    ('5c6d7e8f-9a0b-1c2d-3e4f-5a6b7c8d9e0f','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','d1e2f3a4-b5c6-e7f8-a9d2-e3f4a5b6c7d8',0),
    ('6d7e8f9a-0b1c-2d3e-4f5a-6b7c8d9e0f1a','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','e3f4a5b6-c7d8-f9a0-b1e4-f5a6b7c8d9e0',0),
    ('7e8f9a0b-1c2d-3e4f-5a6b-7c8d9e0f1a2b','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','f5a6b7c8-d9e0-a1b2-c3f6-a7b8c9d0e1f2',0),
    ('8f9a0b1c-2d3e-4f5a-6b7c-8d9e0f1a2b3c','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','c1d2e3f4-a5b6-d7e8-f9c2-d3e4f5a6b7c8',0),
    ('9a0b1c2d-3e4f-5a6b-7c8d-9e0f1a2b3c4d','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a7b8c9d0-e1f2-b3c4-d5a8-b9c0d1e2f3a4',0),
    ('0b1c2d3e-4f5a-6b7c-8d9e-0f1a2b3c4d5e','2024-12-04 01:33:23','2024-12-04 01:33:23','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','b9c0d1e2-f3a4-c5d6-e7b0-c1d2e3f4a5b6',0);

DROP TABLE IF EXISTS "manage_user";

CREATE TABLE "manage_user" (
    id serial NOT NULL,
    uuid varchar(36) NOT NULL UNIQUE,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    username varchar(30) NOT NULL UNIQUE,
    password varchar(100) NOT NULL,
    nickname varchar(30) NOT NULL,
    gender varchar(1) NOT NULL,
    phone varchar(20) NOT NULL,
    status varchar(1) NOT NULL,
    email varchar(100) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TRIGGER update_manage_user_update_time BEFORE UPDATE ON "manage_user"
FOR EACH ROW EXECUTE FUNCTION update_update_time_column();

INSERT INTO "manage_user" (uuid, create_time, update_time, username, password, nickname, gender, phone, status, email)
VALUES
    ('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e6f','2024-10-24 09:45:00','2024-10-31 09:40:13','admin','123456','超级管理员','1','','1','');

DROP TABLE IF EXISTS "manage_user_role";

CREATE TABLE "manage_user_role" (
    id serial NOT NULL,
    uuid varchar(36) NOT NULL UNIQUE,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_uuid varchar(36) NOT NULL,
    role_uuid varchar(36) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TRIGGER update_manage_user_role_update_time BEFORE UPDATE ON "manage_user_role"
FOR EACH ROW EXECUTE FUNCTION update_update_time_column();

INSERT INTO "manage_user_role" (uuid, create_time, update_time, user_uuid, role_uuid)
VALUES
    ('2d3e4f5a-6b7c-8d9e-0f1a-2b3c4d5e6f7a','2024-10-31 09:40:13','2024-10-31 09:40:13','1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e6f','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d');
