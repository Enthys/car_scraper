import Vue from 'vue';
import VueRouter, { NavigationGuardNext, Route, RouteConfig } from 'vue-router';
import Store from '../store';

Vue.use(VueRouter);

function IsLoggedIn(): boolean {
  return !!Store.state.jwt;
}

function HasLoggedInStatus(status: boolean, redirectName: string) {
  return (to: Route, from: Route, next: NavigationGuardNext) => {
    if (IsLoggedIn() === status) {
      next();
    } else {
      next({ name: redirectName });
    }
  };
}

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: () => import('./../views/Filter/CreateFilter.vue'),
    beforeEnter: HasLoggedInStatus(true, 'Login'),
  },
  {
    path: '/filters',
    name: 'Filters',
    component: () => import('./../views/Filter/GetFilters.vue'),
    beforeEnter: HasLoggedInStatus(true, 'Login'),
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Authentication/Login.vue'),
    beforeEnter: HasLoggedInStatus(false, 'Home'),
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Authentication/Register.vue'),
    beforeEnter: HasLoggedInStatus(false, 'Home'),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
