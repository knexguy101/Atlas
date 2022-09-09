import { createRouter, createWebHistory } from 'vue-router'
import Create from '/src/components/Create.vue'
import Tasks from '/src/components/Tasks.vue'

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
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router