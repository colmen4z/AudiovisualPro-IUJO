<script setup>
import { RouterView } from 'vue-router';
</script>

<template>
  <router-link class="link" to="/">page one test</router-link>
  <router-link class="link" to="/PageTwoTest">page two test</router-link>
  <RouterView />
</template>
