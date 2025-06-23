import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  server: {
    proxy: {
      '/get-messages': 'http://localhost:8080'
    },
    host: '0.0.0.0', // Разрешает доступ с любых IP в локальной сети
    allowedHosts: [
      'boriskobzarev.tplinkdns.com' // Ваш новый домен
    ]
  }
})