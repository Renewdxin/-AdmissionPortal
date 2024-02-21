import { defineStore } from 'pinia' //创建仓库
import { ref } from 'vue'
export const useIdStore = defineStore(
  'xiyouMobile-id',
  () => {
    const id = ref('')
    const setId = (newId) => {
      id.value = newId
    }
    return {
      id,
      setId
    }
  },
  {
    persist: true
  }
)
