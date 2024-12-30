<script setup lang="ts">
import router from '@/router';
import { ref, computed } from 'vue';

const props = defineProps({
  onSubmit: {
    type: Function,
    required: true,
  },
  error: {
    type: [String, null],
    default: '',
  },
  loading: {
    type: Boolean,
    default: false,
  },
})

const form = ref({
  email: '',
  password: '',
})
const passwordVisible = ref(false)

// Regex pattern to validate email
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

// Check email validity
const emailError = computed(() => {
  return !emailRegex.test(form.value.email) && form.value.email !== ''
    ? 'Invalid email address.'
    : ''
})

const changePasswordVisibility = () => {
  passwordVisible.value = !passwordVisible.value
}

const passwordFieldType = computed(() => {
  return passwordVisible.value ? 'text' : 'password'
})

const handleSubmit = () => {
  props.onSubmit(form.value)
}
</script>

<template>
  <div class="flex flex-col bg-gray-200 rounded-md w-[26rem] h-[20rem] justify-center items-center shadow-lg">
    <form class="flex flex-col space-y-1 w-full max-w-[60%]">
      <span
        v-if="error"
        class="text-red-600 text-sm block h-[1.25rem]"
      >
        {{ error }}
      </span>
      <span v-else-if="emailError">
        {{ emailError }}
      </span>
      <input
        v-model="form.email"
        type="text"
        placeholder="email"
        class="px-1"
      >
      <span class="flex">
        <input
          v-model="form.password"
          :type="passwordFieldType"
          placeholder="password"
          class="px-1"
        >
        <button @click="changePasswordVisibility" type="button">
          {{ passwordVisible ? "Hide" : "Show" }}
        </button>
      </span>
      <button
        @click.prevent="handleSubmit"
        :disabled="loading"
      >
        {{ loading ? "Loading" : "Log in" }}
      </button>
    </form>
    <div>
      <span class="text-sm">
        New? <button @click="() => router.push('/register')">Register here.</button>
      </span>
    </div>
  </div>
</template>
