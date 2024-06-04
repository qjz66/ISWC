const userInfo = {
  // 为模块开启独立的命名空间
  namespaced: true,
  
  // 数据
  state: {
    data: {
      userName:'玖月'
    },
    isLogined: false
  },

  getters: {
    userInfo: state => {
      return state.data
    }
  },

  mutations: {
    // 设置用户信息
    setUserInfo(state, userInfo) {
      state.data = userInfo
      state.isLogined = true
    },
    // 清除用户信息
    clearUserInfo(state,info) {
      state.data = info
      state.isLogined = false
    },
    // 修改用户信息
    // modifyUserInfo(state, newInfo) {
    //   state.data = Object.assign(state.data, newInfo)
    // },

  },

  actions: {
    // 保存用户信息
    saveInfo({ commit }, result) {
      commit('setUserInfo', result)
    },
    // 退出登录
    logout({commit}) {
      commit('clearUserInfo', {})
      // 跳转到登录页面
      location.href = '/login'
    }
  }

}

export default userInfo
