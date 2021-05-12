import Vue from 'vue';
import Axios, { AxiosRequestConfig } from 'axios';
import VueAxios from 'vue-axios';
import store from './store';

Vue.config.productionTip = false;

Vue.use(VueAxios, Axios);

Vue.axios.defaults.baseURL = `${window.location.origin}/api`;
Vue.axios.defaults.responseType = 'json';

Vue.axios.interceptors.request.use((config: AxiosRequestConfig) => {
  const requestConfig = config;

  if (store.getters.jwt) {
    requestConfig.headers.common.Authorization = store.getters.jwt;
  }

  return requestConfig;
});
