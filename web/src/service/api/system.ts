import { request } from '../request';

/** get role list */
export function GetRoleList(params?: Api.System.RoleSearchParams) {
  return request<Api.System.RoleList>({
    url: '/systemManage/getRoleList',
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
  return request<Api.System.AllRole[]>({
    url: '/systemManage/getAllRoles',
    method: 'get'
  });
}

/** add user */
export function AddUser(req: Api.System.AddUserRequest) {
  return request<Api.System.AddUserRequest>({
    url: '/system/addUser',
    method: 'post',
    data: req
  });
}

/** edit user */
export function EditUser(req: Api.System.EditUserRequest) {
  return request<Api.System.EditUserResponse>({
    url: '/system/editUser',
    method: 'post',
    data: req
  });
}

export function DeleteUser(req: number[]) {
  return request<Api.Common.Empty>({
    url: '/system/deleteUser',
    method: 'post',
    data: {
      ids: req
    }
  });
}

/** get user list */
export function GetUserList(params?: Api.System.UserSearchParams) {
  return request<Api.System.UserList>({
    url: '/system/getUserList',
    method: 'get',
    params
  });
}

/** get menu list */
export function GetMenuList() {
  return request<Api.System.MenuList>({
    url: '/systemManage/getMenuList/v2',
    method: 'get'
  });
}

/** get all pages */
export function GetAllPages() {
  return request<string[]>({
    url: '/systemManage/getAllPages',
    method: 'get'
  });
}

/** get menu tree */
export function GetMenuTree() {
  return request<Api.System.MenuTree[]>({
    url: '/systemManage/getMenuTree',
    method: 'get'
  });
}
