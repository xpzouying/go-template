import path from "path"
import react from "@vitejs/plugin-react"
import { defineConfig } from "vite"

import { viteMockServe } from 'vite-plugin-mock';

export default defineConfig({
  plugins: [
    react(),
    viteMockServe({
      supportTs: false,
      mockPath: './src/mock/',
      enable: true,
    })
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  server: {
    proxy: {
      "/status": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
})
