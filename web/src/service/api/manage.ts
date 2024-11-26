import { request } from '../request';

/** get role list */
export function GetRoleList(params?: Api.Manage.RoleSearchParams) {
  return request<Api.Manage.RoleList>({
    url: '/manage/getRoleList',
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
    url: '/manage/getAllRoles',
    method: 'get'
  });
}

/** add role */
export function AddRole(req: Api.Manage.AddRoleRequest) {
  return request<Api.Manage.AddRoleResponse>({
    url: '/manage/addRole',
    method: 'post',
    data: req
  });
}

/** edit role */
export function EditRole(req: Api.Manage.EditRoleRequest) {
  return request<Api.Common.Empty>({
    url: '/manage/editRole',
    method: 'post',
    data: req
  });
}

export function GetRoleMenus(params: Api.Manage.GetRoleMenusRequest) {
  return request<number[]>({
    url: '/manage/getRoleMenus',
    method: 'get',
    params
  });
}

export function SetRoleMenus(req: Api.Manage.SetRoleMenusRequest) {
  return request<Api.Common.Empty>({
    url: '/manage/setRoleMenus',
    method: 'post',
    data: req
  });
}

/* delete role */
export function DeleteRole(req: number[]) {
  return request<Api.Common.Empty>({
    url: '/manage/deleteRole',
    method: 'post',
    data: {
      ids: req
    }
  });
}

/** add user */
export function AddUser(req: Api.Manage.AddUserRequest) {
  return request<Api.Manage.AddUserResponse>({
    url: '/manage/addUser',
    method: 'post',
    data: req
  });
}

/** edit user */
export function EditUser(req: Api.Manage.EditUserRequest) {
  return request<Api.Manage.EditUserResponse>({
    url: '/manage/editUser',
    method: 'post',
    data: req
  });
}

export function DeleteUser(req: number[]) {
  return request<Api.Common.Empty>({
    url: '/manage/deleteUser',
    method: 'post',
    data: {
      ids: req
    }
  });
}

/** get user list */
export function GetUserList(params?: Api.Manage.UserSearchParams) {
  return request<Api.Manage.UserList>({
    url: '/manage/getUserList',
    method: 'get',
    params
  });
}

export function AddMenu(req: Api.Manage.AddMenuRequest) {
  return request<Api.Common.Empty>({
    url: '/manage/addMenu',
    method: 'post',
    data: req
  });
}

export function EditMenu(req: Api.Manage.EditMenuRequest) {
  return request<Api.Common.Empty>({
    url: '/manage/editMenu',
    method: 'post',
    data: req
  });
}

export function DeleteMenu(req: number[]) {
  return request<Api.Common.Empty>({
    url: '/manage/deleteMenu',
    method: 'post',
    data: {
      ids: req
    }
  });
}

/** get menu list */
export function GetMenuList() {
  return request<Api.Manage.MenuList>({
    url: '/manage/getMenuList/v2',
    method: 'get'
  });
}

/** get all pages */
export function GetAllPages() {
  return request<string[]>({
    url: '/manage/getAllPages',
    method: 'get'
  });
}

/** get menu tree */
export function GetMenuTree() {
  return request<Api.Manage.MenuTree[]>({
    url: '/manage/getMenuTree',
    method: 'get'
  });
}

export function GetAllButtons() {}
