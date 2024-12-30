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
  name: '',
  email: '',
  password: '',
  role: 'user',
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
  if (!emailError.value) {
    props.onSubmit(form.value);
  }
}
</script>

<template>
  <div class="flex flex-col bg-gray-200 rounded-md w-[26rem] h-[20rem] justify-center items-center shadow-lg">
    <form class="flex flex-col space-y-1">
      <span v-if="error">
        {{ error }}
      </span>
      <span v-else-if="emailError">
        {{ emailError }}
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
        {{ loading ? 'Loading' : 'Register' }}
      </button>
    </form>
    <div>
      <span class="text-sm">
        Already have an account? <button @click="() => router.push('/login')">Log in.</button>
      </span>
    </div>
  </div>
</template>
