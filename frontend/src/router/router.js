import { createRouter, createWebHistory } from 'vue-router'
import PageOneTest from '../views/PageOneTest.vue'
import PageTwoTest from '../views/PageTwoTest.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            name: "PageOneTest",
            component: PageOneTest
        },
        {
            path: "/PageTwoTest",
            name: "PageTwoTest",
            component: PageTwoTest
        }
    ]
})

export default router