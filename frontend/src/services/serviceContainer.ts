import AuthService from '@/services/AuthService/AuthService';
import FilterService from '@/services/FilterService/FilterService';
import { AxiosInstance } from 'axios';

// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
export default function serviceContainer(axios: AxiosInstance) {
  return {
    get authService(): AuthService {
      return new AuthService(axios);
    },

    get filterService(): FilterService {
      return new FilterService(axios);
    },
  };
}
