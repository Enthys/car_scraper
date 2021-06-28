import { AxiosInstance } from 'axios';

// eslint-disable-next-line no-shadow
export enum FilterType {
  MOBILEBG_CARS = 'MobileBGCar',
  MOBILEBG_BIKES = 'MobileBGBike',
  MOBILEBG_BUSES = 'MobileBGBus',
  CARSBG_CARS = 'CarsBGCar',
  CARSBG_BIKES = 'CarsBGBike',
  CARSBG_BUSES = 'CarsBGBus',
}

export interface FilterInterface {
  id: number
  type: FilterType
  search: string
}

export interface FilterSearchInterface {
  id: number,
  type: FilterType,
  brand?: string,
  model?: string|string[]
  yearFrom?: string
  yearTo?: string
  priceFrom?: string
  priceTo?: string
}

export default class FilterService {
  private readonly axios: AxiosInstance;

  constructor(axios: AxiosInstance) {
    this.axios = axios;
  }

  public async getFilters(): Promise<FilterInterface[]> {
    const resp = await this.axios.get('/filters');

    return resp.data;
  }

  public async createFilter(type: string, filter: unknown): Promise<void> {
    await this.axios.post('/filters', {
      type,
      filter: JSON.stringify(filter),
    });
  }

  async deleteFilter(filterId: number): Promise<void> {
    await this.axios.delete(`/filters/${filterId}`);
  }
}
