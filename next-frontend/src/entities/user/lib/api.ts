import {
  SMS_NOT_SENT,
  SMS_SENT,
  SMS_VERIFICATION_FAILED,
  SMS_VERIFICATION_SUCCED,
} from "./constants";

const BASE_AUTH_URL = process.env.NEXT_PUBLIC_API_URL + "/auth";
const BASE_USER_URL = process.env.NEXT_PUBLIC_API_URL + "/users";

type SMSAuth = {
  expiresIn: number;
  sent: boolean;
};

type SMSAuthResponse = {
  message: string;
};

type CodeVerification = {
  expiresIn: number;
  tokenType: string;
  verified: boolean;
};

type RefreshTokenResponse  = {
  refreshed: boolean;
  expiresIn: number;
}

type IsAuthResponse = {
  id: number;
  phone: string;
  email: string;
  name: string;
  authenticated: boolean;
}

const checkResponse = <T>(res: Response): Promise<T> =>
  res.ok ? res.json() : res.json().then((err) => Promise.reject(err));

export const refreshToken = async () => {
  return await fetch(`${BASE_AUTH_URL}/refresh`, {
    method: "POST",
    credentials: "include"
  })
  .then((res) => checkResponse<RefreshTokenResponse>(res))
  .then((data) => {
    if (data.refreshed) {
      return { message: "Success!" } as SMSAuthResponse
    }

    return { message: "Failed to refresh." } as SMSAuthResponse
  })
}

export const fetchWithRefresh = async <T>(
  url: RequestInfo,
  options: RequestInit
) => {
  try {
    const res = await fetch(url, options);
    return await checkResponse<T>(res);
  } catch (err) {
      const refreshData = await refreshToken();

      if (refreshData.message === "Success!") {
        const res = await fetch(url, options);
        return await checkResponse<T>(res);
      } else {
        return Promise.reject(err);
      }
   
  }
};

export const sendCode = async (phone: string) => {
  return await fetch(`${BASE_AUTH_URL}/sms/send`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify({ phone: phone }),
  })
    .then((res) => checkResponse<SMSAuth>(res))
    .then((data) => {
      if (data.sent) {
        return {
          message: SMS_SENT,
        } as SMSAuthResponse;
      }

      return {
        message: SMS_NOT_SENT,
      } as SMSAuthResponse;
    });
};

export const verifyCode = async (code: string) => {
  return await fetch(`${BASE_AUTH_URL}/sms/verify`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify({
      phone: sessionStorage.getItem("phone-number"),
      code: code,
    }),
  })
    .then((res) => checkResponse<CodeVerification>(res))
    .then((data) => {
      if (data.verified) {
        return {
          message: SMS_VERIFICATION_SUCCED,
          verified: true,
        } as SMSAuthResponse & { verified: boolean };
      }

      return {
        message: SMS_VERIFICATION_FAILED,
        verified: false,
      } as SMSAuthResponse & { verified: boolean };
    });
};

export const checkAuth = async () => {
  return await fetchWithRefresh<IsAuthResponse>(`${BASE_AUTH_URL}/check`, {
    method: "GET",
    credentials: "include"
  });
};

export const logout = async () => {
  return await fetch(`${BASE_AUTH_URL}/logout`, {
    method: "POST",
    credentials: "include"
  });
};

export const changeName = async (id: number, name: string) => {
  return await fetch(`${BASE_USER_URL}/${id}/name`, {
    method: "PATCH",
    credentials: "include",
    body: JSON.stringify({
      name: name
    })
  })
}

export const changeEmail = async (id: number, email: string) => {
  return await fetch(`${BASE_USER_URL}/${id}/email`, {
    method: "PATCH",
    credentials: "include",
    body: JSON.stringify({
      email: email,
    })
  })
}