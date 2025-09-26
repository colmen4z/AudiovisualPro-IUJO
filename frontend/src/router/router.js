import { createRouter, createWebHistory } from 'vue-router'
//PAGINAS MAIN
import Login from '../views/Login.vue'
import Main from '../views/Main.vue'
import Home from '../views/Home.vue'
//PAGINAS GESTORES
import GestorCliente from '../views/gestores/GestorCliente.vue'
import GestorProyecto from '../views/gestores/GestorProyecto.vue'
import GestorPersonal from '../views/gestores/GestorPersonal.vue'
import GestorRecursos from '../views/gestores/GestorRecursos.vue'
import GestorLocaciones from '../views/gestores/GestorLocaciones.vue'
//PAGINA 404
import FourZeroFour from '../views/FourZeroFour.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/login",
            name: "Login",
            component: Login
        },
        {
            path: "/",
            name: "MainLayout",
            component: Main,
            children: [
                {
                    path: "",
                    name: "Home",
                    component: Home,
                    redirect: ""
                },
                {
                    path: "gestor-proyecto",
                    name: "GestorProyecto",
                    component: GestorProyecto
                },
                {
                    path: "gestor-cliente",
                    name: "GestorCliente",
                    component: GestorCliente
                },
                {
                    path: "gestor-personal",
                    name: "GestorPersonal",
                    component: GestorPersonal
                },
                {
                    path: "gestor-recursos",
                    name: "GestorRecursos",
                    component: GestorRecursos
                },
                {
                    path: "gestor-locaciones",
                    name: "GestorLocaciones",
                    component: GestorLocaciones
                }
            ]
        },
        {
            path: '/:pathMatch(.*)*',
            name: 'NotFound',
            component: FourZeroFour
        }
    ]
})

export default router
