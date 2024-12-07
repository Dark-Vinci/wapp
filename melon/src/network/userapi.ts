import axios, { AxiosError, AxiosInstance } from 'axios';

export class UserApi {
  private readonly backendBaseURL: string;
  private readonly axiosInstance: AxiosInstance;

  public constructor(backendBaseURL: string, a: AxiosInstance) {
    this.backendBaseURL = backendBaseURL;
    this.axiosInstance = a;
  }

  public async getUserApi(): Promise<void> {
    try {
      const response = await this.axiosInstance.get('/user');

      console.log(response);
    } catch (error) {
      console.log(error);
    }
  }

  public async loginUser(userData: any) {
    const loginURL = `${this.backendBaseURL}/login`;

    try {
      const response = await axios.post(loginURL, userData);

      return response as unknown as { token: string };
    } catch (error: unknown) {
      if ((error as AxiosError).request) {
        console.log({ error });
        return Promise.reject(error);
      }

      if ((error as AxiosError).response) {
        console.log({ error });
        return Promise.reject(error);
      }

      return Promise.reject(error);
    }
  }
}
