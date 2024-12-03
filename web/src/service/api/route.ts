import { request } from '../request';

/** get constant routes */
export function GetConstantRoutes() {
  return request<Api.Route.MenuRoute[]>({ url: '/api/route/getConstantRoutes' });
}

/** get user routes */
export function GetUserRoutes() {
  return request<Api.Route.UserRoute>({ url: '/api/route/getUserRoutes' });
}

/**
 * whether the route is exist
 *
 * @param routeName route name
 */
export function IsRouteExist(routeName: string) {
  return request<boolean>({ url: '/api/route/isRouteExist', params: { routeName } });
}
