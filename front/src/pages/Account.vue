<template>
  <div class="d-flex">
    <div class="flex-shrink-0 p-3 text-white bg-dark w-10" style="height: 90vh;">
      <span class="fs-4 text-primary">Account</span>
      <hr>
      <ul class="nav nav-pills flex-column mb-auto">
        <li class="nav-item">
          <router-link to="/account/profile" class="nav-link" aria-current="page">
            Profile
          </router-link>
        </li>
        <li>
          <router-link to="/account/ratings" class="nav-link">
            Ratings
          </router-link>
        </li>
      </ul>
    </div>
    <RouterView/>
  </div>
</template>

<script setup>
import {RouterView} from "vue-router";</script>

<style scoped>
.w-10 {
  width: 10%;
}
</style>