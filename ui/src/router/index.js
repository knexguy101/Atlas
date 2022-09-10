import { createRouter, createWebHistory } from 'vue-router'
import Create from '/src/components/Create.vue'
import Tasks from '/src/components/Tasks.vue'
import Profiles from '/src/components/Profiles.vue'
import Proxies from '/src/components/Proxies.vue'
import Accounts from '/src/components/Accounts.vue'

const routes = [
    {
        path: '/',
        name: 'Create',
        component: Create,
    },
    {
        path: '/tasks',
        name: 'Tasks',
        component: Tasks,
    },
    {
        path: '/profiles',
        name: 'Profiles',
        component: Profiles,
    },
    {
        path: '/proxies',
        name: 'Proxies',
        component: Proxies,
    },
    {
        path: '/accounts',
        name: 'Accounts',
        component: Accounts,
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router