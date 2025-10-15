<script setup>
import { ref, computed } from "vue"
import { Icon } from "@iconify/vue"

const nameProyectoField = ref('')
const tipoProyectoField = ref('Selecciona un tipo...')
const estadoProyectoField = ref('Selecciona un estado...')
const presupuestoProyecto = ref(0)

const props = defineProps({
    isOpen: {
        type: Boolean,
        required: true
    }
})

const emits = defineEmits(['confirm', 'cancel'])

const cerrarBackdrop = (event) => {
    if (event.target === event.currentTarget) {
        emits('cancel')
    }
}

const limpiarCampos = () => {
    nameProyectoField.value = ''
    tipoProyectoField.value = 'Selecciona un tipo...'
    estadoProyectoField.value = 'Selecciona un estado...'
}

const formateoPresupuesto = computed(() => {
    if (!presupuestoProyecto.value) return '';
    return new Intl.NumberFormat('es-VE', {
        style: 'currency',
        currency: 'VES'
    }).format(presupuestoProyecto.value)
})
</script>

<template>
    <div @click="cerrarBackdrop" v-if="isOpen" class="backdrop-blur-[2px] bg-gray-200/50 fixed inset-0 overflow-y-auto flex justify-center items-center">
        <div class="bg-white p-4 rounded-lg shadow-xl w-175 border-2 border-gray-300">
            <div class="mb-3 flex items-center justify-between">
				<span class="text-xl font-semibold">Agregar Personal</span>
				<button @click="emit('cancel')" class="bg-pink-500 rounded-lg rounded-2 text-white shadow-md cursor-pointer hover:scale-105 transition active:scale-[0.98] ">
					<Icon icon="material-symbols:close" width="35" />
				</button>
			</div>
            <div class="mb-3">
                <div class="mb-2">
                    <label for="nombre_proyecto" class="uppercase font-semibold text-gray-500 text-xs ml-1">Nombre de Proyecto</label>
                    <input v-model="nameProyectoField" type="text" name="nombre_proyecto" placeholder="Nombre" class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-101 transition ease-in-out disabled:bg-gray-50">
                </div>
                <div class="grid grid-cols-2 gap-2">
                    <div>
                        <label for="tipo_proyecto" class="uppercase font-semibold text-gray-500 text-xs ml-1">Tipo</label>
                        <select v-model="tipoProyectoField" name="tipo_proyecto" class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-101 transition ease-in-out disabled:bg-gray-50">
                            <option disabled selected>Selecciona un tipo...</option>
                            <option value="">Marketing</option>
                            <option value="">Television</option>
                            <option value="">Cine</option>
                        </select>
                    </div>
                    <div>
                        <label for="estado_proyecto" class="uppercase font-semibold text-gray-500 text-xs ml-1">Estado</label>
                        <select v-model="estadoProyectoField" name="estado_proyecto" class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-101 transition ease-in-out disabled:bg-gray-50">
                            <option disabled selected>Selecciona un estado...</option>
                            <option value="">Iniciado</option>
                            <option value="">En Curso</option>
                            <option value="">Finalizado</option>
                        </select>
                    </div>
                    <div>
                        <label for="fecha_inicio" class="uppercase font-semibold text-gray-500 text-xs ml-1">Fecha de Inicio</label>
                        <input type="date" class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-101 transition ease-in-out disabled:bg-gray-50">
                    </div>
                    <div>
                        <label for="fecha_fin_estimada" class="uppercase font-semibold text-gray-500 text-xs ml-1">Fecha de Finalizacion Estimada</label>
                        <input type="date" class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-101 transition ease-in-out disabled:bg-gray-50">
                    </div>
                </div>
                <div class="mb-2">
                    <label for="cliente_proyecto" class="uppercase font-semibold text-gray-500 text-xs ml-1">Cliente</label>
                    <select name="cliente_proyecto" class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-101 transition ease-in-out disabled:bg-gray-50">
                        <option disabled selected>Selecciona un cliente...</option>
                        <option value="">Sair Colmenarez - J-123123123</option>
                        <option value="">Sara Ramos - J-123456789</option>
                        <option value="">Samuel Peña - J-768594745</option>
                        <option value="">Thairiana Mendez - J-349745000</option>
                        <option value="">Jose Daniel Colmenarez - J-463534123</option>
                    </select>
                </div>
                <div class="grid grid-cols-2 gap-2">
                    <div class="mb-2">
                        <label for="locacion_proyecto" class="uppercase font-semibold text-gray-500 text-xs ml-1">Locacion</label>
                        <select name="locacion_proyecto" class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-101 transition ease-in-out disabled:bg-gray-50">
                            <option disabled selected>Selecciona una locacion...</option>
                            <option value="">La Catedral</option>
                            <option value="">Sambil</option>
                            <option value="">Avenida Libertador</option>
                        </select>
                    </div>
                    <div class="mb-2">
                        <label for="presupuesto_proyecto" class="uppercase font-semibold text-gray-500 text-xs ml-1">Presupuesto</label>
                        <input v-model="presupuestoProyecto" min="0" step="0.01" placeholder="0.00" type="number" class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-101 transition ease-in-out disabled:bg-gray-50">
                    </div>
                </div>

                <div class="mb-2 text-center mt-1">
                    <div v-if="presupuestoProyecto > 0">
                        <span class="text-lg text-green-500 font-semibold">Presupuesto del proyecto: {{ formateoPresupuesto }} </span>
                    </div>
                    <div v-else>
                        <span class="text-lg text-gray-500 fomt-semibold">No hay presupuesto...</span>
                    </div>
                </div>
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