<script setup lang="ts">
import { storeToRefs } from 'pinia';

import { useAuthStore } from '@/stores/useAuth';
import router from '@/router'

const authStore = useAuthStore()
const { user, loading, isAuthenticated } = storeToRefs(authStore)

const handleLogout = async () => {
  try { await authStore.logout()
    router.push('/login')
  } catch (error) {
    console.error(error)
  }
}

</script>

<template>
  <div class="flex flex-col min-h-screen w-full max-w-4xl mx-auto justify-center items-center">
    <div v-if="loading">
      <span>Loading...</span>
    </div>
    <div v-else>
      <span>Welcome {{ user?.name }}</span>
    </div>
    <div>
      <button
        v-if="isAuthenticated"
        @click="handleLogout"
        class="underline hover:opacity-50 transition-all"
      > Log out </button>
    </div>
  </div>
</template>
