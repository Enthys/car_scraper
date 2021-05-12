<template>
  <div>
    Email: <input type="email" v-model="email">
    <br>
    Password: <input type="password" v-model="password">
    <br>
    Confirm Password: <input type="password" v-model="passwordConfirm">
    <br>
    <button @click="register">Login</button>

  </div>
</template>

<script lang="ts">
import AuthService from '@/services/AuthService/AuthService';
import { Component, Vue } from 'vue-property-decorator';

@Component({
  inject: [
    'authService',
  ],
})
export default class Register extends Vue {
  private readonly authService!: AuthService;

  private email = '';

  private password = '';

  private passwordConfirm = '';

  private async register(): Promise<void> {
    await this.authService.register({
      email: this.email,
      password: this.password,
      passwordConfirm: this.passwordConfirm,
    });

    await this.$router.push('/login');
  }
}
</script>

<style scoped>

</style>
