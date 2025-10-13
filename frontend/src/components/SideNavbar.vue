<script setup>
import { ref } from 'vue'
import { Icon } from "@iconify/vue"
import { useRouter } from 'vue-router'

import LogoutView from '../components/Logout.vue'

const router = useRouter()

const logoutModalVisible = ref(false)

const mostrarLogout = () => {
    logoutModalVisible.value = true
}

const logout = () => {
    localStorage.removeItem('jwt_token')
    console.log('Sesion cerrada')

    logoutModalVisible.value = false
    router.push({ name: 'Login' })
}

const cancelarLogout = () => {
    logoutModalVisible.value = false
}
</script>

<template>
    <div class="bg-white p-3 shadow-lg h-screen w-60 outline-2 outline-gray-300">
        <div class="flex justify-center">
            <img src="../assets/audiovisualpro_logo/AudiovisualPro128.png" alt="AudiovisualPro Logo">
        </div>

        <div class="flex flex-col mb-2">
            <router-link to="/" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition flex items-center">
                <Icon icon="material-symbols-light:home" width="25" class="mr-1" />
                Home
            </router-link>
        </div>

        <div class="flex flex-col p-2 gap-1">

            <p class="font-semibold mb-2 text-green-600 text-xl">Gestiones</p>

            <router-link to="/gestor-proyecto" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition flex items-center">
                <Icon icon="ix:project" width="25" class="mr-1" />
                Proyectos
            </router-link>

            <router-link to="/gestor-cliente" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition flex items-center">
                <Icon icon="mdi:account" width="25" class="mr-1" />
                Clientes
            </router-link>

            <router-link to="/gestor-personal" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition flex items-center">
                <Icon icon="mdi:worker" width="25" class="mr-1" />
                Personal
            </router-link>

            <router-link to="/gestor-recursos" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition flex items-center">
                <Icon icon="material-symbols:camera" width="25" class="mr-1" />
                Recursos Tecnicos
            </router-link>

            <router-link to="/gestor-locaciones" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition flex items-center">
                <Icon icon="mdi:location" width="25" class="mr-1" />
                Locaciones
            </router-link>

            <button @click="mostrarLogout" class="rounded hover:outline-0 p-2 hover:bg-pink-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-pink-500 transition text-left flex items-center cursor-pointer">
                <Icon icon="material-symbols:logout" width="25" class="mr-1" />
                Cerrar Sesion
            </button>
        </div>
    </div>

    <LogoutView :isOpen="logoutModalVisible" @confirm="logout" @cancel="cancelarLogout" />
</template>