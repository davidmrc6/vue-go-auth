import {
  ref,
  computed,
  type Ref,
  type ComputedRef } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'

import type { LoginCredentials, User } from '@/models/User'
import type { ApiError } from '@/models/Error'

export const useAuthStore = defineStore('auth', () => {
  const user: Ref<User | null> = ref(null)
  const loading: Ref<boolean> = ref(false)
  const error: Ref<string | null> = ref(null)

  const isAuthenticated: ComputedRef<boolean> = computed(() => !!user.value)

  const clearError = () => {
    error.value = null
  }

  axios.defaults.baseURL = 'http://localhost:8080/api/v1'
  axios.defaults.withCredentials = true

  const register = async (userData: User) => {
    try {
      clearError()
      loading.value = true
      await axios.post('/register', {
        name: userData.name,
        email: userData.email,
        password: userData.password,
        role: userData.role ?? 'user',
      })
      // After succesful registration, log in
      return await login({
        email: userData.email,
        password: userData.password,
      })
    } catch (err) {
      if (axios.isAxiosError(err) && err.response?.data) {
        const apiError = err.response.data as ApiError
        error.value = apiError.error
      } else {
        error.value = 'An unexpected error occurred during registration'
      }
      throw err
    } finally {
      loading.value = false
    }
  }

  const login = async (credentials: LoginCredentials) => {
    try {
      clearError()
      loading.value = true
      const response = await axios.post('/login', credentials)
      // Fetch user details after successful login
      await fetchUser()
      return true
    } catch (err) {
      if (axios.isAxiosError(err) && err.response?.data) {
        const apiError = err.response.data as ApiError
        error.value = apiError.error
      } else {
        error.value = 'An unexpected error occurred during login'
      }
    } finally {
      loading.value = false
    }
  }

  const fetchUser = async () => {
    try {
      clearError()
      const response = await axios.get('/me')
      user.value = response.data
    } catch (err) {
      if (axios.isAxiosError(err)) {
        if (err.response?.status === 401) {
          logout()
          error.value = 'Session expired. Please log in again.'
        } else {
          error.value = 'Failed to fetch user data'
        }
      }
      throw err
    }
  }

  const logout = async () => {
    try {
      clearError()
      await axios.get('/logout')
      user.value = null
    } catch (err) {
      error.value = 'Failed to logout properly'
      console.error(err)
    }
  }

  const initialize = async () => {
    try {
      await fetchUser()
    } catch (err) {
      console.debug('Not authenticated')
    }
  }

  return {
    user,
    loading,
    error,
    isAuthenticated,
    register,
    login,
    logout,
    fetchUser,
    initialize,
    clearError
  }
})
