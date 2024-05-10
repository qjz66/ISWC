import Vue from 'vue'
import App from './App.vue'
// 路由
import router from './router'
// vuex
import store from './store'
// 引入iView
import ViewUI from 'view-design'
// 引入正则表达式的工具
import Valid from './utils/valid'

import '@/styles/base.scss'
import 'view-design/dist/styles/iview.css'

// 使用iView
Vue.use(ViewUI);

// 把工具注入到vue实例中
Vue.prototype.$Valid = Valid;

Vue.config.productionTip = false

// 实现全局路由守卫
router.beforeEach((to, from, next) => {
	if (to.meta.title) {
	  document.title = to.meta.title;
	}

	if (to.meta.requireAuth) {
		if (store.state.userInfo.data.token) {
			if (to.path == '/login') {
        next('/');
      } else {
        next();
      }
		} else {
			next('/login');
		}
	} else {
		if (store.state.userInfo.data.token) {
			next('/');
		} else {
			next();
		}
	}

})

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
