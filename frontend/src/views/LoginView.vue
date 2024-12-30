<script setup lang="ts">
import { useAuthStore } from '@/stores/useAuth'

import LoginForm from '@/components/LoginForm.vue'
import { storeToRefs } from 'pinia'
import type { LoginCredentials } from '@/models/User'
import router from '@/router';

const authStore = useAuthStore()
const { error, loading } = storeToRefs(authStore)

const handleLogin = async (credentials: LoginCredentials) => {
  const success = await authStore.login(credentials)
  if (success) {
    router.push('/')
  }
}
</script>

<template>
  <div class="w-full bg-gray-100">
    <div class="flex min-h-screen w-full max-w-4xl mx-auto items-center justify-center">
      <LoginForm
        :on-submit="handleLogin"
        :error="error"
        :loading="loading"
      />
    </div>
  </div>
</template>
