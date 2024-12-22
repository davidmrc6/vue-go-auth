<script setup lang="ts">
import router from '@/router';
import { ref } from 'vue';

const props = defineProps({
  onSubmit: {
    type: Function,
    required: true,
  },
  error: {
    type: [String, null],
    default: '',
  }
})

const form = ref({
  name: '',
  email: '',
  password: '',
  role: 'user',
})

const handleSubmit = () => {
  props.onSubmit(form.value)
}
</script>

<template>
  <div class="flex flex-col bg-gray-200 rounded-md w-[26rem] h-[20rem] justify-center items-center shadow-lg">
    <form class="flex flex-col space-y-1">
      <span v-if="error">
        {{ error }}
      </span>
      <input
        v-model="form.name"
        type="text"
        placeholder="name"
        class="px-1"
      >
      <input
        v-model="form.email"
        type="text"
        placeholder="email"
        class="px-1"
      >
      <input
        v-model="form.password"
        type="password"
        placeholder="password"
        class="px-1"
      >
      <button
        @click.prevent="handleSubmit"
      >
        Register
      </button>
    </form>
    <div>
      <span class="text-sm">
        Already have an account? <button @click="() => router.push('/login')">Log in.</button>
      </span>
    </div>
  </div>
</template>
