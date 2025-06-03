import { request } from '../request';

export function GetVersion() {
  return request({
    url: '/api/v1/version',
    method: 'get'
  });
}
