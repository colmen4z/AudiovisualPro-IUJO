<script setup>
import { Icon } from "@iconify/vue"
</script>

<template>
	<div class="bg-white shadow-lg rounded outline-2 outline-gray-300 h-full w-full p-4">
        <h1 class="text-xl font-bold mb-3">Gestion de Clientes</h1>

        <div>
        	<div class="mb-3">
        		<button class="bg-green-500 text-white p-3 rounded font-semibold text-md transition ease-in-out shadow-md hover:bg-green-600 hover:scale-102 hover:shadow-lg active:scale-[0.98] flex items-center justify-center cursor-pointer">
        			<Icon icon="mdi:account-add" width="25" heigth="25" class="mr-2" />
	        		Agregar Cliente
	        	</button>
        	</div>

        	<div class="border-2 border-gray-300 rounded">
	        	<div class="bg-white shadow-xl rounded-2xl p-3">
			        <h2 class="text-xl font-semibold mb-3">Clientes registrados</h2>
			        <table class="w-full text-sm text-left text-gray-700">
			            <thead class="text-xs text-gray-500 uppercase border-b border-gray-200">
			                <tr>
			                    <th scope="col" class="py-3 px-4">
			                        Nombre
			                    </th>
			                    <th scope="col" class="py-3 px-4 hidden sm:table-cell">
			                        RIF
			                    </th>
			                    <th scope="col" class="py-3 px-4">
			                        Email
			                    </th>
			                    <th scope="col" class="py-3 px-4 hidden md:table-cell">
			                        Teléfono
			                    </th>
			                </tr>
			            </thead>
			            <tbody>
			                <tr class="border-b hover:bg-gray-50 transition duration-150">
			                    <td class="py-4 px-4 font-medium text-gray-900">
			                        María Rodríguez
			                    </td>
			                    <td class="py-4 px-4 hidden sm:table-cell">
			                        J-12345678-9
			                    </td>
			                    <td class="py-4 px-4">
			                        maria.r@ejemplo.com
			                    </td>
			                    <td class="py-4 px-4 hidden md:table-cell">
			                        +58 412-1234567
			                    </td>
			                </tr>
			                <tr class="bg-gray-50 border-b hover:bg-gray-100 transition duration-150">
			                    <td class="py-4 px-4 font-medium text-gray-900">
			                        Pedro Gómez
			                    </td>
			                    <td class="py-4 px-4 hidden sm:table-cell">
			                        V-98765432-1
			                    </td>
			                    <td class="py-4 px-4">
			                        pedro.g@ejemplo.com
			                    </td>
			                    <td class="py-4 px-4 hidden md:table-cell">
			                        +58 414-7654321
			                    </td>
			                </tr>
			                <tr class="border-b hover:bg-gray-50 transition duration-150">
			                    <td class="py-4 px-4 font-medium text-gray-900">
			                        Ana Martínez
			                    </td>
			                    <td class="py-4 px-4 hidden sm:table-cell">
			                        G-11223344-5
			                    </td>
			                    <td class="py-4 px-4">
			                        ana.m@ejemplo.com
			                    </td>
			                    <td class="py-4 px-4 hidden md:table-cell">
			                        +58 426-9988776
			                    </td>
			                </tr>
			            </tbody>
			        </table>
			    </div>
		    </div>
        </div>
	</div>
</template>