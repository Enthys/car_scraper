import './bootstrap';
import Vue from 'vue';
import serviceContainer from '@/services/serviceContainer';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';

new Vue({
  router,
  store,
  render: (h) => h(App),
  provide: serviceContainer(Vue.axios),
}).$mount('#app');
