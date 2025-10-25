import { request } from '../request';

/** get role list */
export function GetRoleList(params?: Api.Manage.RoleSearchParams) {
  return request<Api.Manage.RoleList>({
    url: '/api/v1/manage/getRoleList',
    method: 'get',
    params
  });
}

/**
 * get all roles
 *
 * these roles are all enabled
 */
export function GetAllRoles() {
  return request<Api.Manage.AllRole[]>({
    url: '/api/v1/manage/getAllRoles',
    method: 'get'
  });
}

/** add role */
export function AddRole(req: Api.Manage.AddRoleRequest) {
  return request<Api.Manage.AddRoleResponse>({
    url: '/api/v1/manage/addRole',
    method: 'post',
    data: req
  });
}

/** edit role */
export function EditRole(req: Api.Manage.EditRoleRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/editRole',
    method: 'post',
    data: req
  });
}

export function GetRoleMenus(params: Api.Manage.GetRoleMenusRequest) {
  return request<{ menuUuids: string[] }>({
    url: '/api/v1/manage/getRoleMenus',
    method: 'get',
    params
  });
}

export function SetRoleMenus(req: Api.Manage.SetRoleMenusRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/setRoleMenus',
    method: 'post',
    data: req
  });
}

export function GetRoleHome(roleUuid: string) {
  return request<string>({
    url: '/api/v1/manage/getRoleHome',
    method: 'get',
    params: {
      roleUuid
    }
  });
}

export function UpdateRoleHome(req: Api.Manage.UpdateRoleHomeRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/updateRoleHome',
    method: 'post',
    data: req
  });
}

/* delete role */
export function DeleteRole(req: string[]) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/deleteRole',
    method: 'post',
    data: {
      uuids: req
    }
  });
}

/** add user */
export function AddUser(req: Api.Manage.AddUserRequest) {
  return request<Api.Manage.AddUserResponse>({
    url: '/api/v1/manage/addUser',
    method: 'post',
    data: req
  });
}

/** edit user */
export function EditUser(req: Api.Manage.EditUserRequest) {
  return request<Api.Manage.EditUserResponse>({
    url: '/api/v1/manage/editUser',
    method: 'post',
    data: req
  });
}

export function DeleteUser(req: string[]) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/deleteUser',
    method: 'post',
    data: {
      uuids: req
    }
  });
}

/** get user list */
export function GetUserList(params?: Api.Manage.UserSearchParams) {
  return request<Api.Manage.UserList>({
    url: '/api/v1/manage/getUserList',
    method: 'get',
    params
  });
}

export function AddMenu(req: Api.Manage.AddMenuRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/addMenu',
    method: 'post',
    data: req
  });
}

export function EditMenu(req: Api.Manage.EditMenuRequest) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/editMenu',
    method: 'post',
    data: req
  });
}

export function DeleteMenu(req: string[]) {
  return request<Api.Common.Empty>({
    url: '/api/v1/manage/deleteMenu',
    method: 'post',
    data: {
      uuids: req
    }
  });
}

/** get menu list */
export function GetMenuList() {
  return request<Api.Manage.MenuList>({
    url: '/api/v1/manage/getMenuList',
    method: 'get'
  });
}

/** get all pages */
export function GetAllPages(roleUuid: string) {
  return request<string[]>({
    url: '/api/v1/manage/getAllPages',
    method: 'get',
    params: {
      roleUuid
    }
  });
}

/** get menu tree */
export function GetMenuTree() {
  return request<Api.Manage.MenuTree[]>({
    url: '/api/v1/manage/getMenuTree',
    method: 'get'
  });
}

export function GetAllButtons() {}
