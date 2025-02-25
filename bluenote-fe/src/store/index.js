import { createStore } from 'vuex'

export default createStore({
  state: {
    token: localStorage.getItem('token') || '',
    user: JSON.parse(localStorage.getItem('user')) || null
  },
  getters: {
    isAuthenticated: state => !!state.token,
    currentUser: state => state.user
  },
  mutations: {
    setToken(state, token) {
      state.token = token
    },
    setUser(state, user) {
      state.user = user
    },
    clearAuth(state) {
      state.token = ''
      state.user = null
    }
  },
  actions: {
    login({ commit }, token) {
      // Store token in localStorage
      localStorage.setItem('token', token)
      commit('setToken', token)
      
      // We don't need to manually set cookies because the browser will
      // automatically handle cookies sent from the server
    },
    setUserProfile({ commit }, user) {
      localStorage.setItem('user', JSON.stringify(user))
      commit('setUser', user)
    },
    logout({ commit }) {
      // Remove stored token and user data
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      commit('clearAuth')
    }
  }
})
