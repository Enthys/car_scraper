import { AxiosInstance } from 'axios';

export default class FilterService {
  private readonly axios: AxiosInstance;

  constructor(axios: AxiosInstance) {
    this.axios = axios;
  }

  public async createFilter(type: string, filter: unknown): Promise<void> {
    await this.axios.post('/filters', {
      type,
      filter: JSON.stringify(filter),
    });
  }
}
