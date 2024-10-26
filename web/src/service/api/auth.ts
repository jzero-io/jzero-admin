import { request } from '../request';

/**
 * PwdLogin
 *
 * @param username Username
 * @param password Password
 */
export function fetchPwdLogin(username: string, password: string) {
  return request<Api.Auth.LoginToken>({
    url: '/auth/pwd-login',
    method: 'post',
    data: {
      username,
      password
    }
  });
}

/**
 * CodeLogin
 *
 * @param email Email
 * @param verificationUuid
 * @pparam verificationCode
 */
export function fetchCodeLogin(req: Api.Auth.CodeLoginRequest) {
  return request<Api.Auth.LoginToken>({
    url: '/auth/code-login',
    method: 'post',
    data: req
  });
}

/** Register */
export function fetchRegister(req: Api.Auth.RegisterRequest) {
  return request<Api.Auth.RegisterResponse>({
    url: '/auth/register',
    method: 'post',
    data: req
  });
}

/** ResetPassword */
export function resetPassword(req: Api.Auth.ResetPasswordRequest) {
  return request<Api.Auth.ResetPasswordResponse>({
    url: '/auth/resetPassword',
    method: 'post',
    data: req
  });
}

/** SendVerificationCode */
export function SendVerificationCode(params: Api.Auth.SendVerificationCodeRequest) {
  return request<Api.Auth.SendVerificationCodeResponse>({
    url: '/auth/sendVerificationCode',
    method: 'get',
    params
  });
}

/** Get user info */
export function fetchGetUserInfo() {
  return request<Api.Auth.UserInfo>({ url: '/auth/getUserInfo' });
}

/**
 * Refresh token
 *
 * @param refreshToken Refresh token
 */
export function fetchRefreshToken(refreshToken: string) {
  return request<Api.Auth.LoginToken>({
    url: '/auth/refreshToken',
    method: 'post',
    data: {
      refreshToken
    }
  });
}

/**
 * return custom backend error
 *
 * @param code error code
 * @param msg error message
 */
export function fetchCustomBackendError(code: string, msg: string) {
  return request({ url: '/auth/error', params: { code, msg } });
}
