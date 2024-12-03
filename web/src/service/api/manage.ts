import { request } from '../request';

/** get role list */
export function GetRoleList(params?: Api.Manage.RoleSearchParams) {
  return request<Api.Manage.RoleList>({
    url: '/api/manage/getRoleList',
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
    url: '/api/manage/getAllRoles',
    method: 'get'
  });
}

/** add role */
export function AddRole(req: Api.Manage.AddRoleRequest) {
  return request<Api.Manage.AddRoleResponse>({
    url: '/api/manage/addRole',
    method: 'post',
    data: req
  });
}

/** edit role */
export function EditRole(req: Api.Manage.EditRoleRequest) {
  return request<Api.Common.Empty>({
    url: '/api/manage/editRole',
    method: 'post',
    data: req
  });
}

export function GetRoleMenus(params: Api.Manage.GetRoleMenusRequest) {
  return request<number[]>({
    url: '/api/manage/getRoleMenus',
    method: 'get',
    params
  });
}

export function SetRoleMenus(req: Api.Manage.SetRoleMenusRequest) {
  return request<Api.Common.Empty>({
    url: '/api/manage/setRoleMenus',
    method: 'post',
    data: req
  });
}

export function UpdateRoleHome(req: Api.Manage.UpdateRoleHomeRequest) {
  return request<Api.Common.Empty>({
    url: '/api/manage/updateRoleHome',
    method: 'post',
    data: req
  });
}

/* delete role */
export function DeleteRole(req: number[]) {
  return request<Api.Common.Empty>({
    url: '/api/manage/deleteRole',
    method: 'post',
    data: {
      ids: req
    }
  });
}

/** add user */
export function AddUser(req: Api.Manage.AddUserRequest) {
  return request<Api.Manage.AddUserResponse>({
    url: '/api/manage/addUser',
    method: 'post',
    data: req
  });
}

/** edit user */
export function EditUser(req: Api.Manage.EditUserRequest) {
  return request<Api.Manage.EditUserResponse>({
    url: '/api/manage/editUser',
    method: 'post',
    data: req
  });
}

export function DeleteUser(req: number[]) {
  return request<Api.Common.Empty>({
    url: '/api/manage/deleteUser',
    method: 'post',
    data: {
      ids: req
    }
  });
}

/** get user list */
export function GetUserList(params?: Api.Manage.UserSearchParams) {
  return request<Api.Manage.UserList>({
    url: '/api/manage/getUserList',
    method: 'get',
    params
  });
}

export function AddMenu(req: Api.Manage.AddMenuRequest) {
  return request<Api.Common.Empty>({
    url: '/api/manage/addMenu',
    method: 'post',
    data: req
  });
}

export function EditMenu(req: Api.Manage.EditMenuRequest) {
  return request<Api.Common.Empty>({
    url: '/api/manage/editMenu',
    method: 'post',
    data: req
  });
}

export function DeleteMenu(req: number[]) {
  return request<Api.Common.Empty>({
    url: '/api/manage/deleteMenu',
    method: 'post',
    data: {
      ids: req
    }
  });
}

/** get menu list */
export function GetMenuList() {
  return request<Api.Manage.MenuList>({
    url: '/api/manage/getMenuList/v2',
    method: 'get'
  });
}

/** get all pages */
export function GetAllPages(roleId: number) {
  return request<string[]>({
    url: '/api/manage/getAllPages',
    method: 'get',
    params: {
      roleId
    }
  });
}

/** get menu tree */
export function GetMenuTree() {
  return request<Api.Manage.MenuTree[]>({
    url: '/api/manage/getMenuTree',
    method: 'get'
  });
}

export function GetAllButtons() {}
