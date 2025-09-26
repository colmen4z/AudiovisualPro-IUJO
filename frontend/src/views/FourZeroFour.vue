<script setup>
import { ref } from 'vue'

const contador = ref(0)

function increment() {
    contador.value++
}

function redirectPreviousPage() {
    window.history.back()
}
</script>

<template>
    <div class="flex flex-col">
        <div class="flex align-items justify-center">
            <img src="../assets/audiovisualpro_logo/AudiovisualPro128.png">
        </div>
        <div class="text-center mb-6">
            <h1 class="text-green-500 text-4xl font-bold mb-6">Pagina no encontrada</h1>
            <button @click="redirectPreviousPage" class="bg-white font-semibold rounded outline-2 outline-gray-300 p-3 shadow-lg hover:bg-green-500 hover:text-white hover:outline-0 active:outline-2 active:outline-offset-2 active:outline-green-500 transition">Volver a la pagina anterior</button>
        </div>
        <div class="text-center">
            <button @click="increment" class="bg-white font-semibold rounded outline-2 outline-gray-300 p-3 shadow-lg hover:bg-green-500 hover:text-white hover:outline-0 active:outline-2 active:outline-offset-2 active:outline-green-500 transition">Juega con el contador: {{ contador }}</button>
        </div>
    </div>
</template>