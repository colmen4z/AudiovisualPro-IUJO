<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const username = ref('')
const password = ref('')
const message = ref('')
const isError = ref(false)
const isLoading = ref(false)

const backendUrl = 'http://localhost:3000'

const handleLogin = async () => {
  message.value = ''
  isError.value = false
  isLoading.value = true

  try {
    const response = await fetch(`${backendUrl}/api/auth/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        usuario: username.value,
        contrasena: password.value,
      }),
    });

    const data = await response.json()
    
    if (response.ok) {
      message.value = data.message || 'Ingreso exitoso'
      isError.value = false

      localStorage.setItem('jwt_token', data.token)

      console.log('Token JWT almacenado:', data.token)
      
      router.push({ name: 'Home' })
    } else {
      message.value = 'Error desconocido al iniciar sesión.'
      isError.value = true
      console.error('Error de login:', data.error)
    }
    
  } catch (error) {
    message.value = 'No se pudo conectar con el servidor. Verifique que el backend esté corriendo en el puerto 3000.'
    isError.value = true
    console.error('Error de red:', error)
  } finally {
    isLoading.value = false
  }
};
</script>

<template>
    <div class="w-300 p-6 bg-white shadow-2xl rounded-xl max-w-sm mx-auto my-12 outline-2 outline-gray-300">
        <div>
            <img src="../assets/audiovisualpro_logo/AudiovisualPro128.png" alt="AudiovisualPro Logo" class="mx-auto mb-4 w-20 h-20 object-contain">
            <h2 class="text-3xl font-bold mb-6 text-green-600 text-center tracking-tight">Inicio de Sesion</h2>

            <div v-if="message" :class="[
                'p-3 mb-4 rounded-lg font-medium text-sm transition duration-300 ease-in-out',
                isError ? 'bg-pink-100 text-pink-700 border border-pink-300' : 'bg-green-100 text-green-700 border border-green-300'
            ]">
                {{ message }}
            </div>

            <form @submit.prevent="handleLogin">
                <div class="mb-5">
                    <label for="username" class="block text-sm font-semibold mb-1 text-gray-700">Usuario</label>
                    <input 
                        type="text" 
                        id="username" 
                        v-model="username"
                        placeholder="Nombre de usuario" 
                        required
                        class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-102 transition duration-150 ease-in-out disabled:bg-gray-50"
                        :disabled="isLoading"
                    >
                </div>

                <div class="mb-6">
                    <label for="password" class="block text-sm font-semibold mb-1 text-gray-700">Contraseña</label>
                    <input 
                        type="password" 
                        id="password" 
                        v-model="password"
                        placeholder="Contraseña" 
                        required
                        class="font-[450] border border-gray-300 rounded-lg p-3 w-full focus:outline-0 focus:ring-2 focus:ring-green-500 focus:border-green-500 hover:scale-102 transition duration-150 ease-in-out disabled:bg-gray-50"
                        :disabled="isLoading"
                    >
                </div>
                
                <button 
                    type="submit"
                    class="bg-green-500 text-white p-3 rounded-xl font-semibold text-lg w-full transition duration-300 ease-in-out 
                           shadow-md hover:bg-green-600 hover:scale-102 hover:shadow-lg active:scale-[0.98] disabled:bg-green-300 flex items-center justify-center"
                    :disabled="isLoading"
                >
                    {{ isLoading ? 'Ingresando...' : 'Ingresar' }}
                </button>
            </form>
        </div>
    </div>
</template>
