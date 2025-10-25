<script setup>
import { ref } from 'vue'
import { Icon } from '@iconify/vue'

const error = ref(false)
const loginMessage = ref('')

const showPassword = ref(false)

const loginForm = ref({
    username: '',
    password: '',
})

const handleLogin = (e) => {
    e.preventDefault()
    if (!loginForm.value.username || !loginForm.value.password) {
        error.value = true
        loginMessage.value = 'Complete todos los campos.'
    }
}

const togglePasswordVisibility = () => {
    showPassword.value = !showPassword.value
}
</script>

<template>
    <div class="min-h-screen flex items-center justify-center bg-green-50 p-4">
        <div class="grid grid-cols-2 gap-4">
            <div class="w-150 bg-white rounded-2xl shadow-lg max-w-sm p-8 border border-green-100">
                <h1 class="text-center mb-4 font-bold text-green-600 text-xl">Inicio de Sesion</h1>
                
                <form @submit="handleLogin" class="space-y-4">
                    <div v-if="error" class="flex items-center justify-center text-[15px] font-semibold border border-red-200 bg-red-100 p-2 rounded-lg text-center text-red-500">
                        <Icon icon="mdi:error" width="25" class="mr-2" />
                        {{ loginMessage }}
                    </div>

                    <div class="mb-2">
                        <label class="block text-sm text-gray-600 mb-1">Nombre de Usuario</label>
                        <input
                            type="text"
                            v-model="loginForm.username"
                            placeholder="Usuario"
                            class="w-full border border-gray-200 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-green-400"
                        >
                    </div>

                    <div class="mb-2">
                        <label class="block text-sm text-gray-600 mb-1">Contraseña</label>
                        <div class="relative">
                            <input 
                                :type="showPassword ? 'text' : 'password'"
                                v-model="loginForm.password"
                                placeholder="Contraseña"
                                class="w-full border border-gray-200 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-green-400"
                            >
                            <button
                                type="button"
                                @click.prevent="togglePasswordVisibility"
                                class="rounded-lg absolute right-3 top-1/2 transform -translate-y-1/2 focus:outline-none hover:text-green-500 text-shadow-lg transition"
                            >
                                <Icon :icon="showPassword ? 'el:eye-close' : 'mdi:eye'" width="25" />
                            </button>
                        </div>
                    </div>

                    <button type="submit" class="cursor-pointer mt-2 w-full bg-green-500 hover:bg-green-600 text-white font-semibold py-2 rounded-lg transition-colors">
                        Ingresar
                    </button>
                </form>
            </div>
            <div class="bg-white rounded-2xl shadow-lg w-full max-w-sm p-8 border border-green-100">
                
            </div>
        </div>
    </div>
</template>