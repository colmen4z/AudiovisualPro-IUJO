<script setup>
import { ref } from 'vue'
import { Icon } from "@iconify/vue"

import AddLocacionView from '../../components/AddLocacion.vue'

const addLocacionModalVisible = ref(false)

const mostrarAddLocacion = () => {
	addLocacionModalVisible.value = true
}

const cerrarAddLocacion = () => {
	addLocacionModalVisible.value = false
}
</script>

<template>
	<div class="bg-white shadow-lg rounded outline-2 outline-gray-300 h-full w-full p-4">
        <h1 class="text-xl font-bold mb-3">Gestion de Locaciones</h1>

        <div>
        	<div class="mb-3">
        		<button @click="mostrarAddLocacion" class="bg-green-500 text-white p-3 rounded font-semibold text-md transition ease-in-out shadow-md hover:bg-green-600 hover:scale-102 hover:shadow-lg active:scale-[0.98] flex items-center justify-center cursor-pointer">
        			<Icon icon="mdi:add-location" width="25" heigth="25" class="mr-2" />
	        		Agregar Locacion
	        	</button>
        	</div>

        	<div class="mb-3 flex items-center border-2 border-gray-300 rounded-lg shadow-md">
        		<Icon icon="material-symbols:search" width="25" class="m-2 text-gray-500" />
        		<input type="text" name="buscador" placeholder="Buscar Locacion" class="p-2 text-md w-full focus:outline-0">
        	</div>

        	<div class="border-2 border-gray-300 rounded-lg max-h-100 overflow-y-auto">
	        	<div class="bg-white shadow-xl rounded-2xl p-3">
			        <h2 class="text-xl font-semibold mb-3">Clientes registrados</h2>
			        <table class="w-full text-sm table-auto text-left text-gray-700">
			            <thead class="text-xs text-gray-500 uppercase border-b border-gray-200">
			                <tr>
			                    <th scope="col" class="py-3 px-4">
			                        Nombre
			                    </th>
			                    <th scope="col" class="py-3 px-4 hidden sm:table-cell">
			                        Direccion
			                    </th>
			                    <th scope="col" class="py-3 px-4">
			                        Descripcion
			                    </th>
			                </tr>
			            </thead>
			            <tbody>
			                <tr class="border-b hover:bg-gray-50 transition duration-150">
			                    <td class="py-4 px-4 font-medium text-gray-900">
			                    	Catedral de Barquisimeto
			                    </td>
			                    <td class="py-4 px-4 hidden sm:table-cell">
			                    	Avenida Venezuela
			                    </td>
			                    <td class="py-4 px-4">
			                    	Avenida Venezuela entre avenida Simon Rodriguez y calle 30
			                    </td>
			                </tr>
			                <tr class="border-b hover:bg-gray-50 transition duration-150">
			                    <td class="py-4 px-4 font-medium text-gray-900">
			                    	Catedral de Barquisimeto
			                    </td>
			                    <td class="py-4 px-4 hidden sm:table-cell">
			                    	Avenida Venezuela
			                    </td>
			                    <td class="py-4 px-4">
			                    	Avenida Venezuela entre avenida Simon Rodriguez y calle 30
			                    </td>
			                </tr>
			                <tr class="border-b hover:bg-gray-50 transition duration-150">
			                    <td class="py-4 px-4 font-medium text-gray-900">
			                    	Catedral de Barquisimeto
			                    </td>
			                    <td class="py-4 px-4 hidden sm:table-cell">
			                    	Avenida Venezuela
			                    </td>
			                    <td class="py-4 px-4">
			                    	Avenida Venezuela entre avenida Simon Rodriguez y calle 30
			                    </td>
			                </tr>
			                <tr class="border-b hover:bg-gray-50 transition duration-150">
			                    <td class="py-4 px-4 font-medium text-gray-900">
			                    	Catedral de Barquisimeto
			                    </td>
			                    <td class="py-4 px-4 hidden sm:table-cell">
			                    	Avenida Venezuela
			                    </td>
			                    <td class="py-4 px-4">
			                    	Avenida Venezuela entre avenida Simon Rodriguez y calle 30
			                    </td>
			                </tr>
			                <tr class="border-b hover:bg-gray-50 transition duration-150">
			                    <td class="py-4 px-4 font-medium text-gray-900">
			                    	Catedral de Barquisimeto
			                    </td>
			                    <td class="py-4 px-4 hidden sm:table-cell">
			                    	Avenida Venezuela
			                    </td>
			                    <td class="py-4 px-4">
			                    	Avenida Venezuela entre avenida Simon Rodriguez y calle 30
			                    </td>
			                </tr>
			                <tr class="border-b hover:bg-gray-50 transition duration-150">
			                    <td class="py-4 px-4 font-medium text-gray-900">
			                    	Catedral de Barquisimeto
			                    </td>
			                    <td class="py-4 px-4 hidden sm:table-cell">
			                    	Avenida Venezuela
			                    </td>
			                    <td class="py-4 px-4">
			                    	Avenida Venezuela entre avenida Simon Rodriguez y calle 30
			                    </td>
			                </tr>
			            </tbody>
			        </table>
			    </div>
		    </div>
        </div>
	</div>

	<AddLocacionView :isOpen="addLocacionModalVisible" @cancel="cerrarAddLocacion" />
</template>