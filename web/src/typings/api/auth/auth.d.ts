/* eslint-disable @typescript-eslint/no-empty-interface */
declare namespace Api {
  /**
   * namespace Auth
   *
   * backend api module: "auth"
   */
  namespace Auth {
    interface LoginResponse {
      token: string;
      refreshToken: string;
    }

    interface GetUserInfoResponse {
      userId: string;
      username: string;
      roles: string[];
      buttons: string[];
    }

    interface CodeLoginRequest {
      email: string;
      verificationCode: string;
      verificationUuid: string;
    }

    interface RegisterRequest {
      email: string;
      verificationCode: string;
      verificationUuid: string;
      username: string;
      password: string;
    }

    // eslint-disable-next-line @typescript-eslint/no-empty-object-type
    interface RegisterResponse {}

    interface SendVerificationCodeRequest {
      verificationType: string;
      email: string;
    }

    interface SendVerificationCodeResponse {
      verificationUuid: string;
    }

    interface ResetPasswordRequest {
      email: string;
      verificationUuid: string;
      verificationCode: string;
      password: string;
    }

    // eslint-disable-next-line @typescript-eslint/no-empty-object-type
    interface ResetPasswordResponse {}
  }
}
