<template>
  <div class="edit-profile-container">
    <div class="edit-profile-card">
      <h1 class="gradient-title">Edit Profile</h1>
      
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label for="nickname">Nickname</label>
          <input 
            type="text" 
            id="nickname" 
            v-model="formData.nickname" 
            placeholder="Enter your nickname (2-20 characters)"
            required
            class="form-input"
          >
        </div>
        
        <div class="form-group">
          <label for="birthday">Birthday</label>
          <input 
            type="date" 
            id="birthday" 
            v-model="formData.birthday" 
            placeholder="Select your birthday"
            required
            class="form-input"
          >
        </div>
        
        <div class="form-group">
          <label for="aboutMe">About Me</label>
          <textarea 
            id="aboutMe" 
            v-model="formData.aboutMe" 
            placeholder="Tell us about yourself (up to 150 characters)"
            rows="4"
            class="form-input"
          ></textarea>
          <small>{{ formData.aboutMe?.length || 0 }}/150 characters</small>
        </div>
        
        <div class="form-group">
          <button type="submit" :disabled="loading" class="submit-btn">
            {{ loading ? 'Saving...' : 'Save Changes' }}
          </button>
          <router-link to="/profile" class="cancel-btn">Cancel</router-link>
        </div>
        
        <div v-if="error" class="error">{{ error }}</div>
        <div v-if="success" class="success">{{ success }}</div>
      </form>
    </div>
  </div>
</template>

<script>
import api from '../services/api'

export default {
  name: 'EditProfileView',
  data() {
    return {
      formData: {
        nickname: '',
        birthday: '',
        aboutMe: ''
      },
      loading: false,
      error: null,
      success: null
    }
  },
  async created() {
    // Pre-fill form with current user data
    await this.loadUserData()
  },
  methods: {
    async loadUserData() {
      // Try to get from store first
      const user = this.$store.getters.currentUser
      
      if (user) {
        this.formData.nickname = user.Nickname || ''
        this.formData.birthday = user.Birthday || ''
        this.formData.aboutMe = user.AboutMe || ''
        return
      }
      
      // Fetch from API if not in store
      try {
        const response = await api.fetchProfile()
        const profile = response.data
        
        this.formData.nickname = profile.Nickname || ''
        this.formData.birthday = profile.Birthday || ''
        this.formData.aboutMe = profile.AboutMe || ''
      } catch (error) {
        console.error('Error loading user data:', error)
        this.error = 'Failed to load your profile data.'
      }
    },
    
    async handleSubmit() {
      this.loading = true
      this.error = null
      this.success = null
      
      // Validate form
      if (this.formData.nickname.length < 2 || this.formData.nickname.length > 20) {
        this.error = 'Nickname must be between 2 and 20 characters'
        this.loading = false
        return
      }
      
      if (this.formData.aboutMe && this.formData.aboutMe.length > 150) {
        this.error = 'About Me text must not exceed 150 characters'
        this.loading = false
        return
      }
      
      try {
        await api.updateProfile(this.formData)
        this.success = 'Profile updated successfully!'
        
        // Refresh profile data in store
        await this.refreshProfile()
        
        // Redirect back to profile after short delay
        setTimeout(() => {
          this.$router.push('/profile')
        }, 1500)
      } catch (error) {
        console.error('Profile update error:', error)
        this.error = error.response?.data || 'Failed to update profile. Please try again.'
      } finally {
        this.loading = false
      }
    },
    
    async refreshProfile() {
      try {
        const response = await api.fetchProfile()
        this.$store.dispatch('setUserProfile', response.data)
      } catch (error) {
        console.error('Error refreshing profile:', error)
      }
    }
  }
}
</script>

<style scoped>
.edit-profile-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f4f8;
}

.edit-profile-card {
  background-color: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  max-width: 500px;
  width: 100%;
}

.gradient-title {
  text-align: center;
  margin-bottom: 30px;
  background: linear-gradient(to right, #4facfe, #00f2fe);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.form-group {
  margin-bottom: 20px;
}

.form-input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

.submit-btn {
  background-color: #4facfe;
  color: #fff;
  padding: 10px 15px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.submit-btn:disabled {
  background-color: #ccc;
}

.error {
  color: red;
  margin-top: 10px;
}

.success {
  color: green;
  margin-top: 10px;
}

.cancel-btn {
  display: inline-block;
  margin-left: 10px;
  padding: 10px 15px;
  background-color: #ccc;
  color: #333;
  text-decoration: none;
  border-radius: 4px;
}

.cancel-btn:hover {
  background-color: #bbb;
}

small {
  display: block;
  color: #666;
  text-align: right;
  margin-top: 5px;
}
</style>
