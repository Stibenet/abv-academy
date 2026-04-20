import { createRouter, createWebHistory } from 'vue-router'

import HomePage from '@/pages/HomePage.vue'
import InfoOrgPage from '@/pages/InfoOrgPage.vue'

const routes = [
    {
        path: '/',
        component: HomePage
    },
    {
        path: '/info',
        component: InfoOrgPage
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router