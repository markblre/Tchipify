<template>
  <div class="w-100 d-flex flex-column text-center m-4">
    <h2 class="mx-auto mt-5 text-secondary shadow-lg">Update profile</h2>
    <form class="w-50 m-auto" @submit.prevent="update">
      <div class="form-group row">
        <label class="col-lg-3 col-form-label col-form-label-lg">Name</label>
        <div class="col-lg-8">
          <input v-model="user.name" type="text" class="form-control form-control-lg" placeholder="Name">
        </div>
      </div>
      <div class="form-group row mt-4">
        <label class="col-lg-3 col-form-label col-form-label-lg">Username</label>
        <div class="col-lg-8">
          <input v-model="user.username" type="text" class="form-control form-control-lg" placeholder="Username">
        </div>
      </div>
      <div class="form-group row mt-4">
        <label class="col-lg-3 col-form-label col-form-label-lg">New password</label>
        <div class="col-lg-8">
          <input v-model="user.password" type="password" class="form-control form-control-lg" placeholder="Password">
        </div>
      </div>
      <div class="form-group row d-flex justify-content-center my-5">
        <div class="col-lg-10 d-flex justify-content-center">
          <button type="submit" class="btn btn-lg btn-primary">Update</button>
        </div>
      </div>
    </form>
  </div>
</template>

<script setup>
import {onMounted, reactive} from "vue";
import {useAuthStore} from "@/stores/auth_store.js";
import {storeToRefs} from "pinia";
import {useAxios} from "@vueuse/integrations/useAxios";
import {useGeneralResponses} from "@/composables/general_responses.js";
import {useToast} from "vue-toast-notification";
import router from "@/router/index.js";

const authStore = useAuthStore()
const { currentUser } = storeToRefs(authStore)

const user = reactive({})
const generalResponses = useGeneralResponses()
const toast = useToast();


onMounted(() => {
  user.value = {
    name: currentUser.value.name,
    username: currentUser.value.username
  }
})

async function update(){
  let userToSend = user
  delete userToSend.value
  const config = {
    headers: authStore.authAxiosConfig,
    method: 'PUT',
    data: userToSend,
  }
  const { error } = await useAxios(authStore.authBaseUrl+'users/'+authStore.currentUser.id, config)
  if (error.value) {
    generalResponses.manageError(error.value)
    // manage error and let the component display it however it wants to
    return Promise.reject(error.value)
  } else {
    toast.success("Update succeeded, please login again.")
    await authStore.logout()
    router.push("/login")

  }
}
</script>

<style scoped>

</style>