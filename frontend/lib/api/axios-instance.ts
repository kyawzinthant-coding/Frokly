import axios, { AxiosError, AxiosRequestConfig } from "axios";

export const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  // The API is on another origin; without this the browser won't store or send the auth cookie.
  withCredentials: true,
});

// Orval calls this for every generated request/hook — one place to add
// auth headers, logging, or error handling later without touching generated code.
export const customInstance = <T>(config: AxiosRequestConfig): Promise<T> => {
  return axiosInstance(config).then((response) => response.data);
};

// Orval reads this to type hook errors: a failed request rejects with an AxiosError
// whose response.data is the API's error body.
export type ErrorType<Error> = AxiosError<Error>;
