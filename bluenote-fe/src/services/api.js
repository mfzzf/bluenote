import axios from 'axios'
import store from '../store'

const api = axios.create({
  baseURL: 'http://192.168.31.97:8080',
  headers: {
    'Content-Type': 'application/json'
  },
  withCredentials: true // Enable sending cookies with requests
})

// Add token to requests if user is authenticated
api.interceptors.request.use(config => {
  const token = store.state.token
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Handle authentication header from responses
api.interceptors.response.use(response => {
  // Check for JWT token in headers (could be 'x-jwt-token', 'Authorization' or another custom header)
  const token = response.headers['x-jwt-token'] || response.headers['authorization']
  if (token) {
    // Extract token if it's in 'Bearer token' format
    const tokenValue = token.startsWith('Bearer ') ? token.slice(7) : token
    store.dispatch('login', tokenValue)
  }
  return response
}, error => {
  // Handle 401 errors (unauthorized)
  if (error.response && error.response.status === 401) {
    store.dispatch('logout')
  }
  return Promise.reject(error)
})

export default {
  // Auth endpoints
  login(credentials) {
    return api.post('/users/login', credentials)
  },
  
  signup(userData) {
    return api.post('/users/signup', userData)
  },
  
  // Profile endpoints
  fetchProfile() {
    return api.get('/users/profile')
  },
  
  updateProfile(profileData) {
    return api.post('/users/edit', profileData)
  },
  
  logout() {
    return api.post('/users/logout')
  }
}
