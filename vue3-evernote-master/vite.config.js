// vite.config.js

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import * as path from 'path'
// https://vitejs.dev/config/
export default defineConfig (({
  command,
  mode
}) => {
  return {
  plugins: [vue()],
  resolve: {
    extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json', '.vue'],
    alias : {
      '@': path.resolve(__dirname, './src'),
      'vue$': 'vue/dist/vue.runtime.esm-bundler.js',
    },
  },
}
})