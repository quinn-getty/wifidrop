import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  base:"./",
  build:{
    outDir:'../server/dist'
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:27149/',
        changeOrigin: true,
      },
    }
  }
})
