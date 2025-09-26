<script setup>
import { Icon } from "@iconify/vue";
</script>

<template>
    <div class="bg-white p-3 shadow-lg h-screen w-60 outline-2 outline-gray-300">
        <div class="flex justify-center">
            <img src="../assets/audiovisualpro_logo/AudiovisualPro128.png" alt="AudiovisualPro Logo">
        </div>

        <div class="flex flex-col mb-2">
            <router-link to="/" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition"><Icon icon="mdi:home" :height="20" /> Home</router-link>
        </div>

        <div class="flex flex-col p-2">

            <p class="font-semibold mb-2 text-green-600 text-xl">Gestiones</p>

            <router-link to="/gestor-proyecto" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition"><Icon icon="ix:project" :height="20" /> Proyectos</router-link>

            <router-link to="/gestor-cliente" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition"><Icon icon="mdi:account" :height="20" /> Clientes</router-link>

            <router-link to="/gestor-personal" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition"><Icon icon="mdi:account-hard-hat" :height="20" /> Personal</router-link>

            <router-link to="/gestor-recursos" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition"><Icon icon="mdi:audio-video" :height="20" /> Recursos Tecnicos</router-link>

            <router-link to="/gestor-locaciones" class="rounded hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition"><Icon icon="mdi:location" :height="20" /> Locaciones</router-link>

        </div>
    </div>
</template>