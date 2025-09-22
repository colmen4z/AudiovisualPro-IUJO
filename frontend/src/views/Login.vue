<script setup>

</script>

<template>
    <div class="p-4 bg-white outline-3 outline-gray-200 shadow-lg rounded w-70 h-115">
        <div>
            <img src="../assets/audiovisualpro_logo/AudiovisualPro128.png" alt="AudiovisualPro Logo" class="mx-auto mb-2">
            <h2 class="text-2xl font-bold mb-5 text-green-500 text-center">Inicio de Sesion</h2>
            <div class="mb-5">
                <h4 class="font-semibold mb-1">Usuario</h4>
                <input type="text" name="username" id="username" placeholder="Usuario" class="font-[450] outline outline-gray-300 rounded p-2 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition w-full">
            </div>
            <div class="mb-5">
                <h4 class="font-semibold mb-1">Contraseña</h4>
                <input type="password" name="password" id="password" placeholder="Contraseña" class="font-[450] outline outline-gray-300 rounded p-2 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 transition w-full">
            </div>
            <button class="bg-green-400 p-3 rounded font-semibold hover:bg-green-500 focus:outline-2 focus:outline-offset-2 focus:outline-green-500 w-full hover:scale-105 transition active:scale-100 hover:text-white hover:text-shadow-2xs">Ingresar</button>
        </div>
    </div>
</template>