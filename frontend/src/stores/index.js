import { createPinia } from 'pinia' //创建状态管理库
import persist from 'pinia-plugin-persistedstate' //实现数据持久化

const pinia = createPinia().use(persist)

export default pinia

export * from '@/stores/modules/token'
export * from '@/stores/modules/id'
