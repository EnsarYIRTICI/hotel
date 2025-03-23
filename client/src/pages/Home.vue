<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-100">
    <div class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full text-center">
      <h2 class="text-2xl font-semibold text-gray-800 mb-4">Hoşgeldiniz, Ana Sayfadasınız!</h2>
      <button 
        @click="logout" 
        class="mt-4 w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 transition duration-300"
      >
        Çıkış Yap
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter();

const logout = () => {
  localStorage.removeItem('token'); // Token'ı temizle
  router.push('/'); // Login sayfasına yönlendir
};
</script>

<style scoped>
/* Burada stil ekleyebilirsiniz */
</style>
