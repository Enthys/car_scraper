<template>
  <div class="center container">
    <div class="half" @keypress.enter="login">
      <div class="error" v-if="error !== ''">ERROR: {{ error }}</div>
      Email: <input type="email" v-model="email">
      <br>
      Password: <input type="password" v-model="password">
      <br>
      <button @click="login">Login</button>
    </div>
  </div>
</template>

<script lang="ts">
import AuthService from '@/services/AuthService/AuthService';
import { Component, Vue } from 'vue-property-decorator';

@Component({
  inject: ['authService'],
})
export default class Login extends Vue {
  private readonly authService!: AuthService;

  private email = '';

  private password = '';

  private error = '';

  private async login(): Promise<void> {
    try {
      const token = await this.authService.login({
        email: this.email,
        password: this.password,
      });

      await this.$store.commit('USER_LOGIN', { token });
      await this.$router.push('/');
    } catch (err) {
      this.error = 'Invalid Email or Password';
    }
  }
}
</script>

<style scoped>

</style>
