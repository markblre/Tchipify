<template>
  <nav class="navbar navbar-dark bg-black sticky-top py-4">
    <a class="navbar-brand d-flex align-items-center" v-on:click="$router.push('/')">
      <img class="d-inline-block align-top mx-3" src="/src/assets/img/logo_purple.jpg" width="50" height="50" alt="">
      Tchipify
    </a>
    <div v-if="isAuthenticated" class="d-flex">
      <div class="d-flex">
        <router-link to="account/profile"><span class="clickable fa fa-user text-primary fs-4 mx-3" title="account"/> </router-link>
        <span class="fs-6 text-secondary">{{ currentUser.name }}</span>
      </div>
      <span v-on:click="logout" class="clickable fa fa-power-off text-primary fs-4 mx-3" title="logout"/>
    </div>
  </nav>
</template>

<script setup>
import {useAuthStore} from "@/stores/auth_store.js";
import {storeToRefs} from "pinia";
import {useRouter} from "vue-router";

const authStore = useAuthStore()
const {currentUser, isAuthenticated} = storeToRefs(authStore)
const router = useRouter()

async function logout() {
  authStore.logout()
      .then(() => {
        router.push('/login')
      })
}
</script>

<style>
.clickable:hover {
  cursor: pointer;
  opacity: 50%;
}
</style>