import { request } from '../request';

export function GetVersion() {
  return request({
    url: '/api/version',
    method: 'get'
  });
}
