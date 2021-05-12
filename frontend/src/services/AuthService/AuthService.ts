import { AxiosInstance } from 'axios';
import { LoginDetails, RegisterDetails } from '@/services/AuthService/interfaces';

export default class AuthService {
  constructor(
    private readonly axios: AxiosInstance,
  ) {
  }

  public async login(loginDetails: LoginDetails): Promise<string> {
    const resp = await this.axios.post('/auth/login', loginDetails);

    return resp.data.token;
  }

  public async register(registerDetails: RegisterDetails): Promise<void> {
    await this.axios.post('/users', registerDetails);
  }
}
