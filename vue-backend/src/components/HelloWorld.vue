<script setup lang="ts">
import { ref } from 'vue'

defineProps<{ msg: string }>()

const count = ref(0)
</script>

<template>  
  <div class="flex h-screen w-full justify-center items-center">
      <h1 class="text-2xl text-red-600">{{ msg }} Empty Message</h1>
      <router-link :to="{name: 'admin.login'}" >Login</router-link>
  </div>
</template>
