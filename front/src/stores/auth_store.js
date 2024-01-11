import {computed, reactive, ref} from "vue";
import {defineStore} from "pinia";
import { useAxios } from "@vueuse/integrations/useAxios";
import {useGeneralResponses} from "@/composables/general_responses.js";


export const useAuthStore = defineStore('auth', () => {
    //----- Notifications
    const generalResponses = useGeneralResponses()
    //----- Store variables
    const currentUser = ref({})
    const isAuthenticated = computed(() => {
        return currentUser.value !== null && currentUser.value !== undefined && currentUser.value.name !== null && currentUser.value.name !== undefined
    })
    const localstorageName = "tchipify_is_authenticated"

    let authBaseUrl = "/api/"
    const authAxiosConfig = computed(() => { // so it's reactive
        return {
            headers: {
                Accept: "application/json"
            }
        }
    })

    async function silent_login(){
        if (localStorage[localstorageName] === "true") {
            try {
                await introspect()
            } catch (e) { // not logged in
                localStorage[localstorageName] = false
            }
        } else {
            localStorage[localstorageName] = false
        }

    }

    async function login(user){
        const config = {
            headers: authAxiosConfig,
            method: 'POST',
            data: user,
        }
        const { error } = await useAxios(authBaseUrl+'login', config)
        if (!error.value) {
            localStorage[localstorageName] = true
            return introspect()
        } else {
            generalResponses.manageError(error.value)
            // manage error and let the component display it however it wants to
            return Promise.reject(error.value)
        }
    }

    async function logout() {
        const config = {
            headers: authAxiosConfig,
            method: 'POST',
        }
        const {error} = await useAxios(authBaseUrl+'logout', config)
        if (error.value) {
            generalResponses.manageError(error)
            // manage error and let the component display it however it wants to
            return Promise.reject(error.value)
        }
        localStorage[localstorageName] = false
        currentUser.value = null
    }

    async function introspect(){
        const {data, error} = await useAxios(authBaseUrl+'introspect', authAxiosConfig.value)
        if (!error.value) {
            currentUser.value = data.value
        } else {
            // manage error and let the component display it however it wants to
            generalResponses.manageError(error)
            return Promise.reject(error.value)
        }
    }

    async function register(user){
        const config = {
            headers: authAxiosConfig,
            method: 'POST',
            data: user,
        }
        const {error} = await useAxios(authBaseUrl+'register', config)
        if (!error.value) {
            return login({username: user.username, password: user.password})
        } else {
            generalResponses.manageError(error)
            return Promise.reject(error.value)
        }
    }



    return {authBaseUrl, authAxiosConfig, isAuthenticated, currentUser, login, logout, register, silent_login}
})