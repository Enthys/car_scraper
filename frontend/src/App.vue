<template>
  <div id="app">
    <div id="nav" v-if="!$store.state.jwt">
      <router-link to="/login">Login</router-link> &nbsp;|&nbsp;
      <router-link to="/register">Register</router-link>
    </div>
    <div id="nav" v-else>
      <router-link to="/">Create filter</router-link>
      &nbsp;|&nbsp;
      <router-link to="/filters">Current filters</router-link>
      &nbsp;|&nbsp;
      <a href="#" @click="logout">Logout</a>
    </div>
    <router-view/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class App extends Vue {
  private async logout(): Promise<void> {
    await this.$store.commit('USER_LOGOUT');
    await this.$router.push('/login');
  }
}

</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}

.error {
  border: 1px solid #a60808;
  background-color: #ff9f95;
  border-radius: 5px;
}

.center {
  margin: 0 auto;
  justify-items: center;
}

.container-h {
  display: flex;
  flex-direction: row;
}

.container-v {
  display: flex;
  flex-direction: column;
}

.half {
  width: 50%;
}
</style>
