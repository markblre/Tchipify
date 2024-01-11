import {createRouter, createWebHistory} from 'vue-router';
import EmptyLayout from "@/pages/layouts/EmptyLayout.vue";
import Login from "@/pages/Login.vue";
import Register from "@/pages/Register.vue";
import Account from "@/pages/Account.vue";
import Profile from "@/components/Profile.vue";
import {useAuthStore} from "@/stores/auth_store.js";
import Home from "@/pages/Home.vue";
import Ratings from "@/pages/Ratings.vue";

const router = createRouter({
    history: createWebHistory('/'),
    routes: [
        {
            path: '/',
            name: 'home',
            component: EmptyLayout,
            children: [
                {
                  path: '',
                  name: 'home2',
                  component: Home,
                },
                {
                    path: 'login',
                    name: 'login',
                    component: Login,
                    meta: {guestAllowed: true},
                },
                {
                    path: 'register',
                    name: 'register',
                    component: Register,
                    meta: {guestAllowed: true},
                },
                {
                    path: 'account',
                    name: 'account',
                    component: Account,
                    children: [
                        {
                            path: 'profile',
                            name: 'profile',
                            component: Profile,
                        },
                        {
                            path: 'ratings',
                            name: 'ratings',
                            component: Ratings,
                        },
                    ]
                },
            ]
        },
    ]
})



router.beforeEach(async (to, from) => {
    const auth = useAuthStore();

    await auth.silent_login()
    // if path requires authentication

    if (!!!to.meta.guestAllowed) {
        if (!auth.isAuthenticated){
            to.path = 'login'
            return { name: 'login'}
        }
    }
})

export default router