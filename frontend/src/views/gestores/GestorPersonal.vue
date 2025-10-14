<script setup>
import { ref } from 'vue'
import { Icon } from "@iconify/vue"

import AddPersonalView from '../../components/AddPersonal.vue'
import AddRolPersonalView from '../../components/AddRolPersonal.vue'

const addPersonalModalVisible = ref(false)
const addRolesModalVisible = ref(false)

//ABRIR MODALES
const mostrarAddPersonal = () => {
	addPersonalModalVisible.value = true
}
const mostrarAddRol = () => {
	addRolesModalVisible.value = true
}

//CERRAR MODALES
const cerrarAddPersonal = () => {
	addPersonalModalVisible.value = false
}
const cerrarAddRol = () => {
	addRolesModalVisible.value = false
}
</script>

<template>
	<div class="bg-white shadow-lg rounded outline-2 outline-gray-300 h-full w-full p-4">
        <h1 class="text-xl font-bold mb-3">Gestion de Personal</h1>

        <div>
        	<div class="mb-3 flex items-center gap-2">
        		<button @click="mostrarAddPersonal" class="bg-green-500 text-white p-3 rounded font-semibold text-md transition ease-in-out shadow-md hover:bg-green-600 hover:scale-102 hover:shadow-lg active:scale-[0.98] flex items-center justify-center cursor-pointer">
        			<Icon icon="mdi:add-bold" width="25" heigth="25" class="mr-2" />
	        		Agregar Personal
	        	</button>
	        	<button @click="mostrarAddRol" class="bg-green-500 text-white p-3 rounded font-semibold text-md transition ease-in-out shadow-md hover:bg-green-600 hover:scale-102 hover:shadow-lg active:scale-[0.98] flex items-center justify-center cursor-pointer">
	        		<Icon icon="mdi:add-bold" width="25" heigth="25" class="mr-2" />
	        		Agregar Rol
	        	</button>
	        	<button class="bg-green-500 text-white p-3 rounded font-semibold text-md transition ease-in-out shadow-md hover:bg-green-600 hover:scale-102 hover:shadow-lg active:scale-[0.98] flex items-center justify-center cursor-pointer">
	        		<Icon icon="eos-icons:role-binding" width="25" class="mr-2" />
	        		Gestionar Roles
	        	</button>
        	</div>
        </div>
	</div>

	<AddPersonalView :isOpen="addPersonalModalVisible" @cancel="cerrarAddPersonal" />
	<AddRolPersonalView :isOpen="addRolesModalVisible" @cancel="cerrarAddRol" />
</template>