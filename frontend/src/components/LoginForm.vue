<script setup lang="ts">
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
  <div>
    <form class="flex flex-col space-y-1">
      <span v-if="error">
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
  </div>
</template>
