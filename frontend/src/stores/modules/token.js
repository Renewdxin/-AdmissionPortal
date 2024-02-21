import { defineStore } from 'pinia' //创建仓库
import { ref } from 'vue'
export const useTokenStore = defineStore(
  'xiyouMobile-token',
  () => {
    const token = ref('')
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
