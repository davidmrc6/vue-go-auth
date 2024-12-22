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
  email: '',
  password: '',
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
        Log in
      </button>
    </form>
    <div>
      <span class="text-sm">
        New? <button @click="() => router.push('/register')">Register here.</button>
      </span>
    </div>
  </div>
</template>
