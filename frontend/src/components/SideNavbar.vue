<script setup>
import { Icon } from "@iconify/vue";
</script>

<template>
    <div class="bg-white p-3 shadow-lg h-screen w-60">
        <div class="flex justify-center">
            <img src="../assets/audiovisualpro_logo/AudiovisualPro128.png" alt="AudiovisualPro Logo">
        </div>
        <div class="flex flex-col gap-4 p-4">
            <p class="font-semibold">Gestiones</p>
            <button class="rounded outline-2 outline-gray-200 hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition">Proyectos</button>
            <button class="rounded outline-2 outline-gray-200 hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition">Clientes</button>
            <button class="rounded outline-2 outline-gray-200 hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition">Personal</button>
            <button class="rounded outline-2 outline-gray-200 hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition">Recursos Tecnicos</button>
            <button class="rounded outline-2 outline-gray-200 hover:outline-0 p-2 hover:bg-green-500 hover:text-white font-semibold hover:scale-105 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition">Locaciones</button>
        </div>
    </div>
</template>