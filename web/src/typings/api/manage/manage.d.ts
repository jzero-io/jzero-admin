declare namespace Api {
  /**
   * namespace Manage
   *
   * backend api module: "Manage"
   */
  namespace Manage {
    type CommonSearchParams = Pick<Common.PaginatingCommonParams, 'current' | 'size'>;

    /** role */
    type Role = Common.CommonRecord<{
      /** role name */
      roleName: string;
      /** role code */
      roleCode: string;
      /** role description */
      roleDesc: string;
    }>;

    type AddRoleRequest = {
      /** role name */
      roleName: string;
      /** role code */
      roleCode: string;
      /** role description */
      roleDesc: string;
      status: string | null;
    };

    type EditRoleRequest = {
      uuid: string | undefined;
      /** role name */
      roleName: string;
      /** role code */
      roleCode: string;
      /** role description */
      roleDesc: string;
      status: string | null;
    };

    // eslint-disable-next-line @typescript-eslint/no-empty-object-type
    type EditRoleResponse = {};

    type SetRoleMenusRequest = {
      roleUuid: string;
      menuUuids: string[];
    };

    type GetRoleMenusRequest = {
      roleUuid: string;
    };

    type UpdateRoleHomeRequest = {
      roleUuid: string;
      home: string;
    };

    // eslint-disable-next-line @typescript-eslint/no-empty-object-type
    type AddRoleResponse = {};

    /** role search params */
    type RoleSearchParams = CommonType.RecordNullable<
      Pick<Api.Manage.Role, 'roleName' | 'roleCode' | 'status'> & CommonSearchParams
    >;

    /** role list */
    type RoleList = Common.PaginatingQueryRecord<Role>;

    /** all role */
    type AllRole = Pick<Role, 'uuid' | 'roleName' | 'roleCode'>;

    /**
     * user gender
     *
     * - "1": "male"
     * - "2": "female"
     */
    type UserGender = '1' | '2';

    /**
     * user status
     *
     * - "1": "Enable"
     * - "2": "Disable"
     */
    type EnableStatus = '0' | '1';

    /** user */
    type User = Common.CommonRecord<{
      /** username */
      username: string;
      /** user gender */
      userGender: UserGender | null;
      /** user nick name */
      nickName: string;
      /** user phone */
      userPhone: string;
      /** user email */
      userEmail: string;
      /** user role code collection */
      userRoles: string[];
    }>;

    /** add user request */
    type AddUserRequest = {
      /** username */
      username: string;
      /** user gender */
      userGender: UserGender | null;
      /** user nick name */
      nickName: string;
      /** user phone */
      userPhone: string;
      /** user email */
      userEmail: string;
      /** user role code collection */
      userRoles: string[];
      /** password */
      password: string;
      status: EnableStatus | null;
    };

    // eslint-disable-next-line @typescript-eslint/no-empty-object-type
    type AddUserResponse = {};

    /** edit user request */
    type EditUserRequest = {
      uuid: string | undefined;
      /** username */
      username: string;
      /** user gender */
      userGender: UserGender | null;
      /** user nick name */
      nickName: string;
      /** user phone */
      userPhone: string;
      /** user email */
      userEmail: string;
      /** user role code collection */
      userRoles: string[];
      status: EnableStatus | null;
    };

    // eslint-disable-next-line @typescript-eslint/no-empty-object-type
    type EditUserResponse = {};

    /** user search params */
    type UserSearchParams = CommonType.RecordNullable<
      Pick<Api.Manage.User, 'username' | 'userGender' | 'nickName' | 'userPhone' | 'userEmail' | 'status'> &
        CommonSearchParams
    >;

    /** user list */
    type UserList = Common.PaginatingQueryRecord<User>;

    /**
     * menu type
     *
     * - "1": directory
     * - "2": menu
     * - "3": button
     */
    type MenuType = '1' | '2' | '3';

    type MenuPermission = {
      /**
       * button code
       *
       * it can be used to control the menu permission
       */
      code: string;
      /** button description */
      desc: string;
    };

    /**
     * icon type
     *
     * - "1": iconify icon
     * - "2": local icon
     */
    type IconType = '1' | '2';

    type MenuPropsOfRoute = Pick<
      import('vue-router').RouteMeta,
      | 'i18nKey'
      | 'keepAlive'
      | 'constant'
      | 'order'
      | 'href'
      | 'hideInMenu'
      | 'activeMenu'
      | 'multiTab'
      | 'fixedIndexInTab'
      | 'query'
    >;

    type Menu = Common.CommonRecord<{
      /** parent menu uuid */
      parentUuid: string;
      /** menu type */
      menuType: MenuType;
      /** menu name */
      menuName: string;
      /** route name */
      routeName: string;
      /** route path */
      routePath: string;
      /** component */
      component?: string;
      /** iconify icon name or local icon name */
      icon: string;
      /** icon type */
      iconType: IconType;
      /** button code */
      buttonCode: string;
      /** permissions */
      permissions?: MenuPermission[] | null;
      /** children menu */
      children?: Menu[] | null;
    }> &
      MenuPropsOfRoute;

    /** menu list */
    type MenuList = Common.PaginatingQueryRecord<Menu>;

    type MenuTree = {
      uuid: string;
      label: string;
      pUuid: string;
      children?: MenuTree[];
    };

    type AddMenuRequest = {
      menuType: MenuType;
      menuName: string;
      routeName: string;
      routePath: string;
      component?: string;
      i18nKey?: App.I18n.I18nKey | null;
      icon: string;
      iconType: IconType;
      parentUuid: string;
      status: string | null;
      keepAlive?: boolean | null;
      constant?: boolean | null;
      order?: number | null;
      href?: string | null;
      hideInMenu?: boolean | null;
      activeMenu?: import('@elegant-router/types').RouteKey | null;
      mutiTab?: boolean | null;
      fixedIndexInTab?: number | null;
      query?: { key: string; value: string }[] | null;
      buttonCode: string;
      permissions?: MenuPermission[] | null;
    };

    type EditMenuRequest = {
      uuid?: string;
      menuType: MenuType;
      menuName: string;
      routeName: string;
      routePath: string;
      component?: string;
      i18nKey?: App.I18n.I18nKey | null;
      icon: string;
      iconType: IconType;
      parentUuid: string;
      status: string | null;
      keepAlive?: boolean | null;
      constant?: boolean | null;
      order?: number | null;
      href?: string | null;
      hideInMenu?: boolean | null;
      activeMenu?: import('@elegant-router/types').RouteKey | null;
      mutiTab?: boolean | null;
      fixedIndexInTab?: number | null;
      query?: { key: string; value: string }[] | null;
      buttonCode: string;
      permissions?: MenuPermission[] | null;
    };
  }
}
