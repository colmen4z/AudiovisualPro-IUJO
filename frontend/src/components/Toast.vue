<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
    modelValue: Boolean,
    message: String,
    type: { type: String, default: 'Loading' }
})

const emit = defineEmits(['update:modelValue'])
const visible = ref(props.modelValue)

watch(
    () => props.modelValue,
    (val) => (visible.value = val)
)

watch(visible, (val) => emit('update:modelValue', val))
</script>

<template>
    <transition name="fade">
        <div
            v-if="visible"
            class="fixed top-5 left-1/2 transform -translate-x-1/2 bg-white border border-green-100 shadow-lg rounded-xl px-4 py-3 flex items-center gap-2 z-50"
        >
            <p class="text-sm font-semibold">
                {{ message }}
            </p>
        </div>
    </transition>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.3s easy;
}
.fade-enter-from,
.fade.leave-to {
    opacity: 0;
}
</style>