import { request } from '../request';

/**
 * LoginByPwd
 *
 * @param username Username
 * @param password Password
 */
export function LoginByPwd(username: string, password: string) {
  return request<Api.Auth.LoginResponse>({
    url: '/api/auth/pwd-login',
    method: 'post',
    data: {
      username,
      password
    }
  });
}

/**
 * LoginByCode
 *
 * @param email Email
 * @param verificationUuid
 * @pparam verificationCode
 */
export function LoginByCode(req: Api.Auth.CodeLoginRequest) {
  return request<Api.Auth.LoginResponse>({
    url: '/api/auth/code-login',
    method: 'post',
    data: req
  });
}

/** Register */
export function Register(req: Api.Auth.RegisterRequest) {
  return request<Api.Auth.RegisterResponse>({
    url: '/api/auth/register',
    method: 'post',
    data: req
  });
}

/** ResetPassword */
export function ResetPassword(req: Api.Auth.ResetPasswordRequest) {
  return request<Api.Auth.ResetPasswordResponse>({
    url: '/api/auth/resetPassword',
    method: 'post',
    data: req
  });
}

/** SendVerificationCode */
export function SendVerificationCode(params: Api.Auth.SendVerificationCodeRequest) {
  return request<Api.Auth.SendVerificationCodeResponse>({
    url: '/api/auth/sendVerificationCode',
    method: 'get',
    params
  });
}

/** Get user info */
export function GetUserInfo() {
  return request<Api.Auth.GetUserInfoResponse>({ url: '/api/auth/getUserInfo' });
}

/**
 * Refresh token
 *
 * @param refreshToken Refresh token
 */
export function RefreshToken(refreshToken: string) {
  return request<Api.Auth.LoginResponse>({
    url: '/api/auth/refreshToken',
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
export function GetCustomBackendError(code: string, msg: string) {
  return request({ url: '/api/auth/error', params: { code, msg } });
}
