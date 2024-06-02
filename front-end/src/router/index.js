import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

// 解决冗余路由的问题
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}


const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: {
      title: '登录/注册'
    }
  },
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/Home.vue'),
    meta: {
      title: '首页',
      requireAuth: true
    }
  },
  {
    path: '/dataScreen',
    name: 'DataScreen',
    component: () => import('@/views/DataScreen.vue'),
    meta: {
      title: '数据大屏',
    }
  },
  {
    path: '/dataSquare',
    name: 'DataSquare',
    component: () => import('@/views/DataSquare.vue'),
    meta: {
      title: '数据广场',
    }
  },
  {
    path: '/uploadFile',
    name: 'UploadFile',
    component: () => import('@/views/UploadFile.vue'),
    meta: {
      title: '谣言检测',
    }
  },
  {
    path: '/analysisResult',
    name: 'AnalysisResult',
    component: () => import('@/views/AnalysisResult.vue'),
    meta: {
      title: '分析结果',
    }
  },
  {
    path: '/personalSpace',
    name: 'PersonalSpace',
    component: () => import('@/views/PersonalSpace.vue'),
    meta: {
      title: '个人中心',
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
