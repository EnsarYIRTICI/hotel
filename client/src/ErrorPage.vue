<template>
    <div class="flex flex-col items-center justify-center min-h-screen bg-red-100">
      <h1 class="text-4xl font-bold text-red-600">Sunucuya Bağlanılamadı</h1>
      <p class="text-lg text-gray-700 mt-4">Lütfen internet bağlantınızı kontrol edin veya daha sonra tekrar deneyin.</p>
      <p class="text-sm text-gray-500 mt-2">Hata: Sunucuya bağlanırken bir problem oluştu.</p>
    </div>
  </template>
  
  <script>
  export default {
    name: 'ErrorPage',
  };
  </script>
  
  <style scoped>
  h1 {
    text-align: center;
    margin-top: 50px;
  }
  p {
    text-align: center;
    font-size: 18px;
    margin-top: 20px;
  }
  </style>
  