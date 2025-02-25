<template>
  <div class="signup-container">
    <div class="signup-card">
      <div class="header-container">
        <h1 class="gradient-title">Create Account</h1>
        <p class="subtitle">Join BlueNote today</p>
      </div>
      
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label for="email">Email</label>
          <input 
            type="email" 
            id="email" 
            v-model="email" 
            required
            placeholder="Enter your email"
            class="form-input"
          >
        </div>
        
        <div class="form-group">
          <label for="password">Password</label>
          <input 
            type="password" 
            id="password" 
            v-model="password" 
            required
            placeholder="Enter your password"
            class="form-input"
          >
          <small>Password must contain at least 8 characters including letters, numbers, and special characters</small>
        </div>
        
        <div class="form-group">
          <label for="confirmPassword">Confirm Password</label>
          <input 
            type="password" 
            id="confirmPassword" 
            v-model="confirmPassword" 
            required
            placeholder="Confirm your password"
            class="form-input"
          >
        </div>

        <div class="form-group">
          <button type="submit" :disabled="loading" class="signup-button">
            {{ loading ? 'Creating Account...' : 'Sign Up' }}
          </button>
        </div>

        <div v-if="error" class="error">{{ error }}</div>
        <div v-if="success" class="success">{{ success }}</div>
        
        <div class="footer-links">
          <p>
            Already have an account? 
            <router-link to="/login" class="highlight-link">Login</router-link>
          </p>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import api from '../services/api'

export default {
  name: 'SignupView',
  data() {
    return {
      email: '',
      password: '',
      confirmPassword: '',
      loading: false,
      error: null,
      success: null
    }
  },
  methods: {
    async handleSubmit() {
      this.loading = true
      this.error = null
      this.success = null
      
      if (this.password !== this.confirmPassword) {
        this.error = 'Passwords do not match'
        this.loading = false
        return
      }
      
      try {
        await api.signup({
          email: this.email,
          password: this.password,
          confirmPassword: this.confirmPassword
        })
        
        this.success = 'Account created successfully! You can now login.'
        
        // Reset form
        this.email = ''
        this.password = ''
        this.confirmPassword = ''
        
        // Redirect to login after a short delay
        setTimeout(() => {
          this.$router.push('/login')
        }, 2000)
        
      } catch (error) {
        console.error('Signup error:', error)
        this.error = error.response?.data || 'Failed to create account. Please try again.'
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
.signup-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #e6f7ff 0%, #ffffff 100%);
  padding: 20px;
}

.signup-card {
  width: 100%;
  max-width: 400px;
  background: #ffffff;
  border: none;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(52, 152, 219, 0.1);
  padding: 24px;
  transition: transform 0.3s ease;
}

.signup-card:hover {
  transform: translateY(-5px);
}

.header-container {
  text-align: center;
  margin-bottom: 24px;
}

.gradient-title {
  font-size: 32px;
  font-weight: 600;
  background: linear-gradient(135deg, #2980b9 0%, #3498db 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin: 0;
}

.subtitle {
  color: #666;
  margin-top: 8px;
  font-size: 15px;
}

.form-input {
  border-radius: 8px;
  padding: 12px;
  border: 1px solid #eaeaea;
  transition: all 0.3s ease;
  width: 100%;
}

.form-input:hover, .form-input:focus {
  border-color: #3498db;
  box-shadow: 0 0 0 2px rgba(52, 152, 219, 0.1);
}

.signup-button {
  background: linear-gradient(135deg, #2980b9 0%, #3498db 100%);
  border: none;
  width: 100%;
  height: 48px;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  margin-top: 12px;
  transition: all 0.3s ease;
  color: white;
  cursor: pointer;
}

.signup-button:hover, .signup-button:focus {
  opacity: 0.9;
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(52, 152, 219, 0.2);
}

.footer-links {
  text-align: center;
  margin-top: 24px;
  border-top: 1px solid #f0f0f0;
  padding-top: 20px;
}

.highlight-link {
  color: #3498db;
  font-weight: 500;
  transition: all 0.3s ease;
  text-decoration: none;
}

.highlight-link:hover {
  opacity: 0.8;
  text-decoration: underline;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
  color: #555;
}

.error {
  color: #e74c3c;
  margin-top: 10px;
  text-align: center;
}

.success {
  color: #27ae60;
  margin-top: 10px;
  text-align: center;
}

small {
  display: block;
  color: #666;
  margin-top: 5px;
  font-size: 0.8rem;
}
</style>
