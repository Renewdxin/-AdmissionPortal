import { defineStore } from 'pinia' //创建仓库
import { ref } from 'vue'
export const useUserStore = defineStore(
  '3g-user',
  () => {
    const token = ref('Bear abcdefg')
    const setToken = (newToken) => {
      token.value = newToken
    }
    return {
      token,
      setToken
    }
  },
  {
    persist: true
  }
)
