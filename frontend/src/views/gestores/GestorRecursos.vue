<script setup>
import { ref } from 'vue'
import { Icon } from "@iconify/vue"

import AddRecursoView from '../../components/AddRecurso.vue'

const addRecursoModalVisible = ref(false)

const mostrarAddRecurso = () => {
	addRecursoModalVisible.value = true
}

const cerrarAddRecurso = () => {
	addRecursoModalVisible.value = false
}
</script>

<template>
	<div class="bg-white shadow-lg rounded outline-2 outline-gray-300 h-full w-full p-4">
        <h1 class="text-xl font-bold mb-3">Gestion de Recursos Tecnicos</h1>

		<div>
			<div class="mb-3">
				<button @click="mostrarAddRecurso" class="bg-green-500 text-white p-3 rounded font-semibold text-md transition ease-in-out shadow-md hover:bg-green-600 hover:scale-102 hover:shadow-lg active:scale-[0.98] flex items-center justify-center cursor-pointer">
					<Icon icon="material-symbols:camera" width="25" class="mr-2" />
					Agregar Recurso
				</button>
			</div>
		</div>
	</div>

	<AddRecursoView :isOpen="addRecursoModalVisible" @cancel="cerrarAddRecurso" />
</template>