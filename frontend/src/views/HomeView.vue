<script setup lang="ts">
import { storeToRefs } from 'pinia';

import { useAuthStore } from '@/stores/useAuth';

const authStore = useAuthStore()
const { user, loading } = storeToRefs(authStore)
</script>

<template>
  <div v-if="loading" class="flex min-h-screen w-full max-w-4xl mx-auto justify-center items-center">
    <span>Loading...</span>
  </div>
  <div v-else class="flex min-h-screen w-full max-w-4xl mx-auto justify-center">
    <span>Welcome {{ user?.name }}</span>
  </div>
</template>
