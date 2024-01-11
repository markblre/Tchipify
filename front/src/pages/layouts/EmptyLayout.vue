<template>
  <div>
    <NavigationBar/>
    <RouterView />
  </div>
</template>

<script setup>
import { RouterView } from 'vue-router'
import NavigationBar from '@/components/shared/NavigationBar.vue'
</script>

<style scoped>

</style>