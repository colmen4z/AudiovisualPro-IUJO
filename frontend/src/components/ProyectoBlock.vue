<script setup>

</script>

<template>
    <div class="bg-white p-3 rounded-lg shadow-md border border-gray-300 text-center">
        <h1 class="text-xl text-green-700 font-bold mb-2 underline underline-offset-4">Nombre del proyecto</h1>
        <div class="mb-3">
            <h3 class="mb-2 font-semibold">Cliente: nombre del cliente</h3>
            <div class="grid grid-cols-2 gap-3 mb-2 font-semibold">
                <span>Estado del proyecto</span>
                <span>Tipo del Proyecto</span>
                <span>Fecha de inicio: fecha</span>
                <span>Fecha de finalizacion: fecha</span>
            </div>
            <p class="mb-2 font-semibold">Locacion</p>
            <span class="mb-3 text-green-500 font-semibold text-lg">Presupuesto: 0.00</span>
        </div>
        <div class="grid grid-cols-2 gap-2">
            <button class="bg-blue-100 rounded rounded-lg font-bold p-3 text-md text-blue-500 hover:scale-101 transition hover:bg-blue-500 hover:outline-0 hover:text-white hover:shadow-lg active:scale-[0.98] ease-in-out cursor-pointer">Editar Proyecto</button>
            <button class="bg-pink-100 rounded rounded-lg font-bold p-3 text-md text-pink-500 hover:scale-101 transition hover:bg-pink-500 hover:outline-0 hover:text-white hover:shadow-lg active:scale-[0.98] ease-in-out cursor-pointer">Eliminar Proyecto</button>
        </div>
    </div>
</template>