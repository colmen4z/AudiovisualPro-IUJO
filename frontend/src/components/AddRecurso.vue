<script setup>
import { Icon } from '@iconify/vue'
import { ref } from 'vue'

const props = defineProps({
    isOpen: {
        type: Boolean,
        required: true
    }
})

const emit = defineEmits(['confirm', 'cancel'])

const cerrarBackdrop = (event) => {
    if (event.target === event.currentTarget) {
        emit('cancel')
    }
}
</script>

<template>
    <div v-if="isOpen" @click="cerrarBackdrop" class="backdrop-blur-[2px] bg-gray-200/50 fixed inset-0 overflow-y-auto flex justify-center items-center">
        <div class="bg-white p-4 rounded-lg shadow-xl w-175 border-2 border-gray-300">
            <div class="mb-3 flex items-center justify-between">
				<span class="text-xl font-semibold">Agregar Personal</span>
				<button @click="emit('cancel')" class="bg-pink-500 rounded-lg rounded-2 text-white shadow-md cursor-pointer hover:scale-105 transition active:scale-[0.98] ">
					<Icon icon="material-symbols:close" width="35" />
				</button>
			</div>
            <div class="mb-3">

            </div>
            <div>
				<div class="grid grid-cols-2 gap-2">
					<button @click="limpiarCampos" class="rounded rounded-lg font-bold p-3 text-md text-gray-500 hover:scale-101 transition hover:bg-gray-500 hover:outline-0 hover:text-white hover:shadow-lg active:scale-[0.98] ease-in-out cursor-pointer">Limpiar Datos</button>
					<button class="rounded rounded-lg font-bold p-3 text-md text-green-500 hover:scale-101 transition hover:bg-green-500 hover:outline-0 hover:text-white hover:shadow-lg active:scale-[0.98] ease-in-out cursor-pointer">Agregar Personal</button>
				</div>
			</div>
        </div>
    </div>
</template>