<template>
    <div class="flex items-center justify-center min-h-screen bg-gray-100">
      <div class="bg-white p-8 rounded-lg shadow-lg w-full sm:w-96">
        <h2 class="text-2xl font-semibold text-center mb-6">Kayıt Ol</h2>
        <form @submit.prevent="submitForm">
          <div class="mb-4">
            <label for="full_name" class="block text-sm font-medium text-gray-700">Ad Soyad</label>
            <input 
              type="text" 
              id="full_name" 
              v-model="fullName" 
              class="w-full p-3 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="Adınızı ve soyadınızı girin" 
              required
            />
          </div>
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
          <div class="mb-4">
            <label for="email" class="block text-sm font-medium text-gray-700">E-Posta</label>
            <input 
              type="email" 
              id="email" 
              v-model="email" 
              class="w-full p-3 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="E-posta adresinizi girin" 
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
              class="w-full py-3 px-4 bg-green-600 text-white rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              Kayıt Ol
            </button>
          </div>
          <p v-if="error" class="text-red-500 text-sm text-center">{{ errorMessage }}</p>
        </form>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        fullName: '',
        identityNumber: '',
        email: '',
        password: '',
        error: false,
        errorMessage: '',
      };
    },
    methods: {
      async submitForm() {
        try {
          const response = await axios.post('http://localhost:3000/api/auth/register', {
            full_name: this.fullName,
            identity_number: this.identityNumber,
            email: this.email,
            password: this.password,
          });
  
          // Başarılı kayıt
          const token = response.data.token;
          localStorage.setItem('authToken', token);
  
          // Kayıt başarılıysa giriş ekranına yönlendir
          this.$router.push({ name: 'login' });
        } catch (error) {
          this.error = true;
          this.errorMessage = error.response?.data?.error || 'Kayıt işlemi başarısız oldu!';
        }
      },
    },
  };
  </script>
  