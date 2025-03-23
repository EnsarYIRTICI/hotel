<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-lg w-full sm:w-96">
      <h2 class="text-2xl font-semibold text-center mb-6">Giriş Yap</h2>
      <form @submit.prevent="submitForm">
        <div class="mb-4">
          <label for="identity_number" class="block text-sm font-medium text-gray-700">Kimlik Numarası</label>
          <input 
            type="text" 
            id="identity_number" 
            v-model="identityNumber" 
            class="w-full p-3 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="Kimlik numaranızı girin" 
            required
          />
        </div>
        <div class="mb-6">
          <label for="password" class="block text-sm font-medium text-gray-700">Şifre</label>
          <input 
            type="password" 
            id="password" 
            v-model="password" 
            class="w-full p-3 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="Şifrenizi girin" 
            required
          />
        </div>
        <div class="mb-4">
          <button 
            type="submit" 
            class="w-full py-3 px-4 bg-blue-600 text-white rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            Giriş Yap
          </button>
        </div>
        <p v-if="error" class="text-red-500 text-sm text-center">Hatalı kimlik numarası veya şifre!</p>
      </form>
      
      <div class="mt-4 text-center">
        <p class="text-sm text-gray-600">Hesabınız yok mu? 
          <router-link to="/accounts/register" class="text-blue-600 hover:text-blue-800">Kayıt ol</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      identityNumber: '',
      password: '',
      error: false,
    };
  },
  methods: {
    async submitForm() {
      try {
        const response = await axios.post('http://localhost:3000/api/auth/login', {
          identity_number: this.identityNumber,
          password: this.password,
        });

        // Başarılı giriş
        const token = response.data.token;
        const customer = response.data.customer;

        // Burada token'ı bir depolama alanında saklayabilirsiniz (örneğin, Vuex veya localStorage)
        localStorage.setItem('authToken', token);

        // Giriş başarılıysa yönlendir
        this.$router.push({ name: 'dashboard' });
      } catch (error) {
        // Hatalı giriş
        this.error = true;
      }
    },
  },
};
</script>
