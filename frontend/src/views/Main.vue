<script setup>
import { RouterView } from 'vue-router'
import SideNavbar from '../components/SideNavbar.vue';

</script>

<template>
    <div class="flex h-screen w-full">
        <SideNavbar />

        <main class="flex-grow p-6">
            <router-view />
        </main>
    </div>
</template>