import Vue from 'vue'
import Vuex from 'vuex'
import userInfo from './modules/userInfo'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

export default new Vuex.Store({
	// 模块化,
  modules: {
  	userInfo
  },
	
	getters: {
		isLogined: state => {
			return state.userInfo.isLogined
		}
	},

	// 插件
	// 持久化vuex的数据
	plugins: [createPersistedState({
	    key: 'store',
	    storage: window.localStorage,
	})]
})
