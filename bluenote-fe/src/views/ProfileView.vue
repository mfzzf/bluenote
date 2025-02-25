<template>
  <div class="profile-container">
    <div class="profile-card">
      <h1 class="gradient-title">User Profile</h1>
      
      <div v-if="loading" class="loading">
        Loading profile...
      </div>
      
      <div v-else-if="error" class="error">
        {{ error }}
      </div>
      
      <div v-else class="profile-details">
        <div class="profile-field">
          <strong>Nickname:</strong>
          <span>{{ profile.Nickname || 'Not set' }}</span>
        </div>
        
        <div class="profile-field">
          <strong>Email:</strong>
          <span>{{ profile.Email }}</span>
        </div>
        
        <div class="profile-field">
          <strong>Phone:</strong>
          <span>{{ profile.Phone || 'Not set' }}</span>
        </div>
        
        <div class="profile-field">
          <strong>Birthday:</strong>
          <span>{{ profile.Birthday || 'Not set' }}</span>
        </div>
        
        <div class="profile-field">
          <strong>About Me:</strong>
          <p>{{ profile.AboutMe || 'No information provided' }}</p>
        </div>
        
        <div class="actions">
          <router-link to="/edit-profile" class="edit-btn">Edit Profile</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import api from '../services/api'

export default {
  name: 'ProfileView',
  data() {
    return {
      profile: {},
      loading: true,
      error: null
    }
  },
  async created() {
    await this.fetchProfile()
  },
  methods: {
    async fetchProfile() {
      this.loading = true
      this.error = null
      
      try {
        const response = await api.fetchProfile()
        this.profile = response.data
        this.$store.dispatch('setUserProfile', response.data)
      } catch (error) {
        console.error('Error fetching profile:', error)
        this.error = 'Failed to load profile. Please try again.'
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
.profile-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #e6f7ff 0%, #ffffff 100%);
  padding: 20px;
}

.profile-card {
  width: 100%;
  max-width: 600px;
  background: #ffffff;
  border: none;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(52, 152, 219, 0.1);
  padding: 24px;
  transition: transform 0.3s ease;
}

.profile-card:hover {
  transform: translateY(-5px);
}

.gradient-title {
  font-size: 32px;
  font-weight: 600;
  background: linear-gradient(135deg, #2980b9 0%, #3498db 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin: 0;
  text-align: center;
  margin-bottom: 24px;
}

.loading {
  text-align: center;
  margin: 20px 0;
  font-style: italic;
  color: #666;
}

.error {
  color: #e74c3c;
  text-align: center;
  margin: 20px 0;
}

.profile-field {
  margin-bottom: 15px;
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
}

.profile-field strong {
  display: block;
  margin-bottom: 5px;
  color: #555;
}

.actions {
  margin-top: 25px;
  text-align: center;
}

.edit-btn {
  display: inline-block;
  padding: 12px 24px;
  background: linear-gradient(135deg, #2980b9 0%, #3498db 100%);
  color: white;
  text-decoration: none;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.edit-btn:hover {
  opacity: 0.9;
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(52, 152, 219, 0.2);
}
</style>
