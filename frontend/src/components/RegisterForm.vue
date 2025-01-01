<script setup lang="ts">
import router from '@/router';
import { ref, computed } from 'vue';
import type { PropType } from 'vue';

const props = defineProps({
  onSubmit: {
    type: Function as PropType<(
      form: {
        name: string;
        email: string;
        password: string;
        role: string
      }) => void
    >,
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
const touched = ref({
  name: false,
  email: false,
  password: false,
})

// Validate name by checking if it's empty
const validateName = (name: string) => {
  if (!name && touched.value.name) return 'Name is required'
  return ''
}

// Validate email by testing against a Regex expression
const validateEmail = (email: string) => {
  if (!email && touched.value.email) return 'Email is required'
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email) && touched.value.email) {
    return 'Please enter a valid email address'
  }
  return ''
}

// Validate password by checking if it's empty
const validatePassword = (password: string) => {
  if (!password && touched.value.password) return 'Password is required'
  return ''
}

const nameError = computed(() => validateName(form.value.name))
const emailError = computed(() => validateEmail(form.value.email))
const passwordError = computed(() => validatePassword(form.value.password))
const formError = computed(() => {
  if (nameError.value) return nameError.value
  if (emailError.value) return emailError.value
  if (passwordError.value) return passwordError.value
  return props.error
})

// Check which fields have been touched
const handleFieldFocus = (field: 'name' | 'email' | 'password') => {
  touched.value[field] = true
}

const handleSubmit = () => {
  touched.value.name = true
  touched.value.email = true
  touched.value.password = true

  // Submit only if inputs are correct
  if ( !nameError.value && !emailError.value && !passwordError.value) {
    props.onSubmit(form.value)
  }
}
</script>

<template>
  <div class="flex flex-col bg-gray-200 rounded-md w-[26rem] h-[20rem] justify-center items-center shadow-lg">
    <form
      class="flex flex-col space-y-1 w-full max-w-[60%]"
      @submit.prevent="handleSubmit"
    >
      <!-- Errors -->
      <div v-if="formError" class="text-red-600 text-sm text-center" role="alert">
        {{ formError }}
      </div>

      <!-- Name input box -->
      <div class="space-y-1">
        <input
          v-model="form.name"
          type="username"
          placeholder="Name"
          class="w-full px-3 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="{ 'border-red-500': nameError }"
          @focus="handleFieldFocus('name')"
        >
      </div>

      <!-- Email input box -->
      <div class="space-y-1">
        <input
          v-model="form.email"
          type="email"
          placeholder="Email"
          class="w-full px-3 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="{ 'border-red-500': emailError }"
          @focus="handleFieldFocus('email')"
        >
      </div>

      <!-- Password input box -->
      <div class="space-y-1">
        <div class="relative">
          <input
            v-model="form.password"
            :type="passwordVisible ? 'text' : 'password'"
            placeholder="Password"
            class="w-full px-3 pr-12 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-blue-500"
            :class="{ 'border-red-500': passwordError }"
            @focus="handleFieldFocus('password')"
          >
          <button
            type="button"
            class="absolute right-2 top-1/2 -translate-y-1/2 text-sm text-gray-600 hover:text-gray-800"
            @click="passwordVisible = !passwordVisible"
          >
            {{ passwordVisible ? 'Hide' : 'Show' }}
          </button>
        </div>
      </div>

      <!-- Register button -->
      <button
        type="submit"
        class="w-full py-2 bg-blue-500 text-white rounded hover:bg-blue-600 disabled:opacity-50"
        :disabled="loading"
      >
        {{ loading ? 'Loading...' : 'Register' }}
      </button>

      <p class="text-center text-sm">
        Already have an account? <button type="button" class="text-blue-500 hover:underline" @click="() => router.push('/login')">Log in</button>
      </p>

    </form>
  </div>
</template>
