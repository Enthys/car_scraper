import Vue from 'vue';
import Vuex from 'vuex';
import createPersistedState from 'vuex-persistedstate';

Vue.use(Vuex);

export default new Vuex.Store({
  plugins: [createPersistedState()],
  state: {
    jwt: '',
  },
  mutations: {
    USER_LOGIN(state, payload: {token: string}) {
      state.jwt = payload.token;
    },
    USER_LOGOUT(state) {
      state.jwt = '';
    },
  },
  getters: {
    jwt: (state) => `Bearer ${state.jwt}`,
  },
});
