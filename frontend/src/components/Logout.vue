<script setup>
import { defineProps, defineEmits } from 'vue'

import { Icon } from "@iconify/vue"

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
	<div v-if="isOpen" @click="cerrarBackdrop" class="fixed inset-0 overflow-y-auto flex justify-center items-center">
		<div class="bg-white p-4 rounded-lg shadow-xl w-75 border-2 border-gray-300">
			<div class="flex justify-center items-center">
				<Icon icon="solar:exit-bold" width="100" class="text-pink-600 mb-3" />
			</div>
			<div class="mb-3">
				<p class="text-xl font-semibold text-center text-black">Esta seguro/a que desea salir del sistema?</p>
			</div>
			<div class="grid grid-cols-2 gap-3">
				<button @click="emit('confirm')" class="rounded rounded-lg font-bold p-3 text-md text-pink-500 hover:scale-102 transition hover:bg-pink-500 hover:outline-0 hover:text-white hover:shadow-lg active:scale-[0.98] ease-in-out">Cerrar Sesion</button>
				<button @click="emit('cancel')" class="rounded rounded-lg font-bold p-3 text-md text-green-600 hover:scale-102 transition hover:bg-green-500 hover:outline-0 hover:text-white hover:shadow-lg active:scale-[0.98] ease-in-out">Cancelar</button>
			</div>
		</div>
	</div>
</template>