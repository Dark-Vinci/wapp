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

  private handleError(error: AxiosError): void {
    if (error.request) {
      console.log({ error });
      return;
      // return Promise.reject(error);
    }

    if ((error as AxiosError).response) {
      console.log({ error });
      return;
      // return Promise.reject(error);
    }
  }

  public async createAccount(userData: any): Promise<any> {
    const signupURL = `${this.backendBaseURL}/signup`;

    try {
      const response = await axios.post(signupURL, userData);

      return response as unknown as { token: string };
    } catch (error: unknown) {
      this.handleError(error as AxiosError);
    }
  }

  public async loginUser(userData: any): Promise<any> {
    const loginURL = `${this.backendBaseURL}/login`;

    try {
      const response = await axios.post(loginURL, userData);

      return response as unknown as { token: string };
    } catch (error: unknown) {
      this.handleError(error as AxiosError);
    }
  }
}
