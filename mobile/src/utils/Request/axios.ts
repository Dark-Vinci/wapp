import axios, { AxiosInstance } from 'axios';
import Constants from 'expo-constants';

export const instantiate = (token?: string): AxiosInstance => {
  const instance = axios.create({
    baseURL: Constants.expoConfig?.extra?.apiURL,
    timeout: 10000,
    headers: {
      'Content-Type': 'application/json',
    },
  });

  instance.interceptors.request.use(
    config => {
      config.headers.Authorization = `Bearer ${token}`;
      return config;
    },
    error => {
      // Handle request error
      return Promise.reject(error);
    }
  )

  instance.interceptors.response.use(
    response => {
      // You can modify response data here
      return response.data;
    },
    error => {
      // Handle response error
      return Promise.reject(error);
    }
  )

  return instance;
};

