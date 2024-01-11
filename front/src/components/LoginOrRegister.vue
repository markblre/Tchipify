<template>
  <div>
    <div class="d-flex flex-column justify-content-center bg-black bg-opacity-50 mx-auto mt-5 w-40 rounded-2">
      <h1 class="text-secondary text-center my-4">{{ capitalizedAction }}</h1>
      <form class="w-75 mx-auto mt-5" @submit.prevent="execute_action()">
        <div v-if="action === 'register'" class="form-group row">
          <label class="col-lg-2 col-form-label col-form-label-lg">Name</label>
          <div class="col-lg-10">
            <input v-model="user.name" type="text" class="form-control form-control-lg" placeholder="Name">
          </div>
        </div>
        <div class="form-group row mt-3">
          <label class="col-lg-2 col-form-label col-form-label-lg">Username</label>
          <div class="col-lg-10">
            <input v-model="user.username" type="text" class="form-control form-control-lg" placeholder="Username">
          </div>
        </div>
        <div class="form-group row mt-3">
          <label class="col-lg-2 col-form-label col-form-label-lg">Password</label>
          <div class="col-lg-10">
            <input v-model="user.password" type="password" class="form-control form-control-lg" placeholder="Password">
          </div>
        </div>
        <div class="form-group row d-flex justify-content-center my-5">
          <div class="col-lg-10 d-flex justify-content-center">
            <button type="submit" class="btn btn-lg btn-primary">{{ capitalizedAction }}</button>
          </div>
        </div>
      </form>
      <router-link class="text-end fs-6 m-2" :to="otherAction">Or {{ otherAction }}</router-link>
    </div>
  </div>
</template>

<script setup lang="js">
import {computed, onMounted, reactive} from "vue";
import {useAuthStore} from "@/stores/auth_store.js";
import {useRouter} from "vue-router";
import {storeToRefs} from "pinia";

const props = defineProps({
  action: String
})

const authStore = useAuthStore()
const router = useRouter()

const capitalizedAction = computed(() => {
  return props.action.charAt(0).toUpperCase() + props.action.slice(1)
})
const otherAction = computed(() => {
  return props.action === 'login' ? 'register' : 'login'
})

const user = reactive({})


async function execute_action() {
  if (this.props.action === "login") {
    authStore.login(this.user)
        .then(() => {
          router.push('/')
        })
        .catch(() => {
        })
  } else {
    authStore.register(this.user)
        .then(() => {
          router.push('/')
        })
        .catch(() => {
        })
  }
}
</script>

<style scoped>
.w-40 {
  width: 40%;
}
</style>