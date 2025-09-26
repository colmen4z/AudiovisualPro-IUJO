import { createRouter, createWebHistory } from 'vue-router'
//PAGINAS MAIN
import Login from '../views/Login.vue'
import Main from '../views/Main.vue'
//PAGINA 404
import FourZeroFour from '../views/FourZeroFour.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            name: "Login",
            component: Login
        },
        {
            path: "/main",
            name: "Main",
            component: Main
        },
        {
            path: '/:pathMatch(.*)*',
            redirect: '/'
        },
        {
            path: '/:pathMatch(.*)*',
            name: 'NotFound',
            component: FourZeroFour
        }
    ]
})

export default router