<template>
  <div id="app">
    <header v-if="isLoggedIn" class="app-header">
      <div class="container">
        <div class="header-content">
          <div class="logo">BlueNote</div>
          <nav class="main-nav">
            <router-link to="/profile">Profile</router-link>
            <a href="#" @click.prevent="logout">Logout</a>
          </nav>
        </div>
      </div>
    </header>
    <main class="app-main">
      <router-view />
    </main>
  </div>
</template>

<script>
export default {
  name: 'App',
  computed: {
    isLoggedIn() {
      return this.$store.getters.isAuthenticated
    }
  },
  methods: {
    logout() {
      this.$store.dispatch('logout')
      this.$router.push('/login')
    }
  }
}
</script>

<style>
/* Reset and base styles */
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: #f9fafb;
  color: #333;
  line-height: 1.6;
  min-height: 100vh;
}

#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

/* Header styles */
.app-header {
  background-color: #fff;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  padding: 15px 0;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  font-size: 24px;
  font-weight: 600;
  background: linear-gradient(135deg, #2980b9 0%, #3498db 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.main-nav {
  display: flex;
  gap: 20px;
}

.main-nav a {
  color: #3498db;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s;
}

.main-nav a:hover {
  color: #2980b9;
}

/* Main content area */
.app-main {
  flex: 1;
  padding-top: 0px; /* Account for fixed header */
  margin-bottom: 20px;
}

/* Form styles */
.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #555;
}

input, textarea, select {
  width: 100%;
  padding: 12px;
  border: 1px solid #dfe3e8;
  border-radius: 8px;
  font-size: 16px;
  transition: all 0.3s;
}

input:hover, textarea:hover, select:hover {
  border-color: #3498db;
}

input:focus, textarea:focus, select:focus {
  border-color: #3498db;
  box-shadow: 0 0 0 2px rgba(52, 152, 219, 0.2);
  outline: none;
}

button {
  background: linear-gradient(135deg, #2980b9 0%, #3498db 100%);
  color: white;
  border: none;
  padding: 12px 24px;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s;
}

button:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(52, 152, 219, 0.2);
}

button:disabled {
  background: #c5d1d9;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}
</style>
