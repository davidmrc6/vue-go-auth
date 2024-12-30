<script setup lang="ts">
import { useAuthStore } from '@/stores/useAuth'
import RegisterForm from '@/components/RegisterForm.vue';
import type { User } from '@/models/User'
import router from '@/router';

import { storeToRefs } from 'pinia'

const authStore = useAuthStore()
const { error, loading } = storeToRefs(authStore)

const handleRegister = async (credentials: User) => {
  const success = await authStore.register(credentials)
  if (success) {
    router.push('/')
  }
}
</script>

<template>
  <div class="w-full bg-gray-100">
    <div class="flex min-h-screen w-full max-w-4xl mx-auto items-center justify-center">
      <RegisterForm
        :on-submit="handleRegister"
        :error="error"
        :loading="loading"
      />
    </div>
  </div>
</template>
