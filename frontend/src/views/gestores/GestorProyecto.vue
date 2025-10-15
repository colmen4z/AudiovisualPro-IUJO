<script setup>
import { ref } from "vue"
import { Icon } from "@iconify/vue"

import AddProyectoView from '../../components/AddProyecto.vue'
import ProyectoBlock from '../../components/ProyectoBlock.vue'

const addProyectoModalVisible = ref(false)

const mostrarAddProyecto = () => {
	addProyectoModalVisible.value = true
}

const cerrarAddProyecto = () => {
	addProyectoModalVisible.value = false
}
</script>

<template>
	<div class="bg-white shadow-lg rounded outline-2 outline-gray-300 h-full w-full p-4">
        <h1 class="text-xl font-bold mb-3">Gestion de Proyectos</h1>

		<div>
			<div class="mb-3">
				<button @click="mostrarAddProyecto" class="bg-green-500 text-white p-3 rounded font-semibold text-md transition ease-in-out shadow-md hover:bg-green-600 hover:scale-102 hover:shadow-lg active:scale-[0.98] flex items-center justify-center cursor-pointer">
					<Icon icon="ix:project-new" class="mr-2" width="25" />
					Agregar Proyecto
				</button>
			</div>

			<div class="border-2 border-gray-300 rounded-lg max-h-130 overflow-y-auto">
				<div class="bg-white shadow-xl rounded-2xl p-3">
					<h2 class="text-xl font-semibold mb-3">Proyectos registrados</h2>
					<div class="grid grid-cols-2 gap-2">
						<ProyectoBlock />
						<ProyectoBlock />
						<ProyectoBlock />
						<ProyectoBlock />
						<ProyectoBlock />
						<ProyectoBlock />
						<ProyectoBlock />
					</div>
				</div>
			</div>
		</div>
	</div>
	
	<AddProyectoView :isOpen="addProyectoModalVisible" @cancel="cerrarAddProyecto" />
</template>