<template>
  <div class="loginContainer">
    <div class="loginCard">
      <div class="headerContainer">
        <h1 class="gradientTitle">Welcome to BlueNote</h1>
        <p class="subtitle">Sign in to your account</p>
      </div>
      
      <!-- Login method tabs -->
      <div class="loginTabs">
        <div 
          class="loginTab" 
          :class="{ active: activeTab === 'email' }"
          @click="activeTab = 'email'"
        >
          <i class="tabIcon">✉️</i>
          <span>Email</span>
        </div>
        <div 
          class="loginTab" 
          :class="{ active: activeTab === 'wechat' }"
          @click="activeTab = 'wechat'"
        >
          <i class="tabIcon">💬</i>
          <span>WeChat</span>
        </div>
        <div 
          class="loginTab" 
          :class="{ active: activeTab === 'phone' }"
          @click="activeTab = 'phone'"
        >
          <i class="tabIcon">📱</i>
          <span>Phone</span>
        </div>
      </div>
      
      <!-- Email login form -->
      <form v-if="activeTab === 'email'" @submit.prevent="handleSubmit" class="loginForm">
        <div class="form-group">
          <label for="email">Email</label>
          <input 
            type="email" 
            id="email" 
            v-model="email" 
            required
            placeholder="Enter your email"
            class="formInput"
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
            class="formInput"
          >
        </div>

        <div class="form-group">
          <button type="submit" :disabled="loading" class="loginButton">
            {{ loading ? 'Logging in...' : 'Login' }}
          </button>
        </div>
      </form>
      
      <!-- WeChat QR login -->
      <div v-else-if="activeTab === 'wechat'" class="wechatLoginContainer">
        <div class="qrCodePlaceholder">
          <div class="qrCodeBox">
            <div class="qrCodeImage"></div>
            <p>Scan with WeChat to login</p>
          </div>
        </div>
        <p class="infoText">Open WeChat app and scan the QR code to log in</p>
      </div>
      
      <!-- Phone login -->
      <div v-else-if="activeTab === 'phone'" class="phoneLoginContainer">
        <div class="form-group">
          <label for="phoneNumber">Phone Number</label>
          <div class="phoneInputGroup">
            <select class="countryCode">
              <option value="+86">+86</option>
              <option value="+1">+1</option>
              <option value="+44">+44</option>
              <option value="+61">+61</option>
              <option value="+81">+81</option>
            </select>
            <input 
              type="tel" 
              id="phoneNumber" 
              v-model="phoneNumber" 
              required
              placeholder="Enter your phone number"
              class="formInput phoneInput"
            >
          </div>
        </div>
        
        <div class="form-group">
          <label for="verificationCode">Verification Code</label>
          <div class="verificationGroup">
            <input 
              type="text" 
              id="verificationCode" 
              v-model="verificationCode" 
              required
              placeholder="Enter verification code"
              class="formInput"
            >
            <button 
              type="button" 
              class="sendCodeBtn" 
              :disabled="codeSent" 
              @click="sendVerificationCode"
            >
              {{ codeSent ? `Resend in ${countdown}s` : 'Get Code' }}
            </button>
          </div>
        </div>
        
        <div class="form-group">
          <button 
            type="button" 
            @click="handlePhoneLogin" 
            :disabled="loading" 
            class="loginButton"
          >
            {{ loading ? 'Logging in...' : 'Login' }}
          </button>
        </div>
      </div>

      <div v-if="error" class="error">{{ error }}</div>
      
      <div class="footerLinks">
        <p>
          Don't have an account? 
          <router-link to="/signup" class="highlightLink">Sign up</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import api from '../services/api'

export default {
  name: 'LoginView',
  data() {
    return {
      activeTab: 'email', // Default active tab
      email: '',
      password: '',
      phoneNumber: '',
      verificationCode: '',
      loading: false,
      error: null,
      codeSent: false,
      countdown: 60
    }
  },
  methods: {
    async handleSubmit() {
      this.loading = true
      this.error = null
      
      try {
        const response = await api.login({
          email: this.email,
          password: this.password
        })
        
        // JWT token is handled by axios interceptor in api.js
        // Also checking if we got a successful response
        if (response.status === 200) {
          // Fetch user profile after successful login
          await this.fetchUserProfile()
          
          // Redirect to profile page
          this.$router.push('/profile')
        }
      } catch (error) {
        console.error('Login error:', error)
        this.error = error.response?.data?.message || error.response?.data || 'Failed to login. Please try again.'
      } finally {
        this.loading = false
      }
    },
    
    async fetchUserProfile() {
      try {
        const response = await api.fetchProfile()
        this.$store.dispatch('setUserProfile', response.data)
      } catch (error) {
        console.error('Error fetching user profile:', error)
      }
    },
    
    sendVerificationCode() {
      // Simulate sending verification code
      this.codeSent = true
      this.countdown = 60
      
      const timer = setInterval(() => {
        this.countdown--
        if (this.countdown <= 0) {
          clearInterval(timer)
          this.codeSent = false
        }
      }, 1000)
      
      // In a real app, you would call an API here
      console.log("Sending verification code to:", this.phoneNumber)
    },
    
    handlePhoneLogin() {
      this.loading = true
      this.error = null
      
      // Validate phone number and verification code
      if (!this.phoneNumber) {
        this.error = "Please enter your phone number"
        this.loading = false
        return
      }
      
      if (!this.verificationCode) {
        this.error = "Please enter the verification code"
        this.loading = false
        return
      }
      
      // Simulate API call for phone login
      setTimeout(() => {
        // For demo purposes, just redirect
        this.loading = false
        this.$router.push('/profile')
      }, 1500)
    }
  }
}
</script>

<style scoped>
.loginContainer {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #e6f7ff 0%, #ffffff 100%);
  padding: 20px;
}

.loginCard {
  width: 100%;
  max-width: 400px;
  background: #ffffff;
  border: none;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(52, 152, 219, 0.1);
  padding: 24px;
  transition: transform 0.3s ease;
}

.loginCard:hover {
  transform: translateY(-5px);
}

.headerContainer {
  text-align: center;
  margin-bottom: 24px;
}

.gradientTitle {
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

/* Login tabs styling */
.loginTabs {
  display: flex;
  border-bottom: 1px solid #eee;
  margin-bottom: 24px;
}

.loginTab {
  flex: 1;
  text-align: center;
  padding: 12px;
  cursor: pointer;
  color: #666;
  transition: all 0.3s ease;
  border-bottom: 2px solid transparent;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.loginTab:hover {
  color: #3498db;
}

.loginTab.active {
  color: #3498db;
  border-bottom: 2px solid #3498db;
}

.tabIcon {
  font-size: 20px;
  margin-bottom: 4px;
}

.formInput {
  border-radius: 8px;
  padding: 12px;
  border: 1px solid #eaeaea;
  transition: all 0.3s ease;
}

.formInput:hover, .formInput:focus {
  border-color: #3498db;
  box-shadow: 0 0 0 2px rgba(52, 152, 219, 0.1);
}

.loginButton {
  background: linear-gradient(135deg, #2980b9 0%, #3498db 100%);
  border: none;
  width: 100%;
  height: 48px;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  margin-top: 12px;
  transition: all 0.3s ease;
}

.loginButton:hover, .loginButton:focus {
  opacity: 0.9;
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(52, 152, 219, 0.2);
}

.footerLinks {
  text-align: center;
  margin-top: 24px;
  border-top: 1px solid #f0f0f0;
  padding-top: 20px;
}

.highlightLink {
  color: #3498db;
  font-weight: 500;
  transition: all 0.3s ease;
  text-decoration: none;
}

.highlightLink:hover {
  opacity: 0.8;
  text-decoration: underline;
}

/* WeChat QR code login styling */
.wechatLoginContainer {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px 0;
}

.qrCodePlaceholder {
  width: 100%;
  display: flex;
  justify-content: center;
  margin: 20px 0;
}

.qrCodeBox {
  width: 200px;
  height: 240px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.qrCodeImage {
  width: 180px;
  height: 180px;
  background: linear-gradient(45deg, #f3f3f3 25%, #e6e6e6 25%, #e6e6e6 50%, #f3f3f3 50%, #f3f3f3 75%, #e6e6e6 75%, #e6e6e6);
  background-size: 10px 10px;
  border-radius: 8px;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.qrCodeImage::after {
  content: "QR Code";
  color: #3498db;
  font-weight: 500;
}

.qrCodeBox p {
  margin-top: 15px;
  color: #666;
  font-size: 14px;
}

.infoText {
  color: #999;
  font-size: 14px;
  text-align: center;
}

/* Phone login styling */
.phoneInputGroup {
  display: flex;
  gap: 10px;
}

.countryCode {
  width: 80px;
  padding: 12px;
  border-radius: 8px;
  border: 1px solid #eaeaea;
  background-color: #f9f9f9;
}

.verificationGroup {
  display: flex;
  gap: 10px;
}

.sendCodeBtn {
  white-space: nowrap;
  padding: 0 15px;
  border-radius: 8px;
  background-color: #3498db;
  color: #fff;
  border: none;
  cursor: pointer;
  font-size: 14px;
}

.sendCodeBtn:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.error {
  color: #e74c3c;
  margin-top: 15px;
  text-align: center;
}

.loginForm {
  display: flex;
  flex-direction: column;
  gap: 15px;
}
.phoneInput {
  flex: 1;
}
</style>
