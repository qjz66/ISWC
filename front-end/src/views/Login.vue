<template>
  <div class="login-container" style="height: 100%">
    <div class="about-image">
      <img class="img-fluid" src="../picture/about.png" alt="about image" />
    </div>
    <div class="login-box">
      <!-- 登录/注册选择栏 -->
      <div class="login-text">
        <a
          href="javascript:;"
          :class="typeView == 0 ? 'active' : ''"
          @click="handleTab(0)"
          >登录</a
        >
        <b>·</b>
        <a
          href="javascript:;"
          :class="typeView == 1 ? 'active' : ''"
          @click="handleTab(1)"
          >注册</a
        >
      </div>

      <!-- 登录模块 -->
      <div class="right-content" v-show="typeView == 0">
        <!-- 账号和密码的输入框 -->
        <div class="input-box">
          <!-- 不开启基于之前输入内容的补全 -->
          <input
            autocomplete="off"
            type="text"
            class="input"
            v-model="formLogin.userName"
            placeholder="请输入登录邮箱/手机号"
          />
          <!-- 密码长度最多20个字符 -->
          <input
            autocomplete="off"
            type="password"
            class="input"
            v-model="formLogin.userPwd"
            maxlength="20"
            @keyup.enter="login"
            placeholder="请输入登录密码"
          />
        </div>

        <!-- 登录按钮,主按钮 -->
        <Button
          class="loginBtn"
          type="primary"
          :disabled="isDisabled"
          :loading="isLoading"
          @click.stop="login"
        >
          立即登录
        </Button>

        <!-- 记住密码,忘记密码选项 -->
        <div class="option">
          <Checkbox class="remember" v-model="checked">
            <span class="checked">记住我</span>
          </Checkbox>
          <span class="forget-pwd" @click.stop="forgetPwd">忘记密码?</span>
        </div>
      </div>

      <!-- 注册模块 -->
      <div class="right_content" v-show="typeView == 1">
        <!-- 账号和密码的输入框 -->
        <div class="input-box">
          <input
            autocomplete="off"
            type="text"
            class="input"
            v-model="formRegister.userName"
            placeholder="请输入注册邮箱/手机号"
          />
          <input
            autocomplete="off"
            type="password"
            class="input"
            v-model="formRegister.userPwd"
            maxlength="20"
            @keyup.enter="register"
            placeholder="请输入密码"
          />
          <input
            autocomplete="off"
            type="password"
            class="input"
            v-model="formRegister.userPwd2"
            maxlength="20"
            @keyup.enter="register"
            placeholder="请再次确认密码"
          />
        </div>

        <!-- 注册按钮,主按钮 -->
        <Button
          class="loginBtn"
          type="primary"
          :disabled="isRegAble"
          :loading="isLoading"
          @click.stop="register"
        >
          立即注册
        </Button>
      </div>
    </div>
  </div>
</template>

<script>
import { login, register } from '@/utils/api'

export default {
  name: 'login',
  components: {},
  data() {
    return {
      // 登录信息
      formLogin: {
        userName: '',
        userPwd: ''
      },
      // 注册信息
      formRegister: {
        userName: '',
        userPwd2: '',
        userPwd: ''
      },
      typeView: 0, //显示不同的view,即登录界面还是注册界面
      checked: false, // 记住登录
      isLoading: false
    }
  },
  computed: {
    // 登陆按钮状态
    // 如果账号和密码都没有输入,则禁用按钮,isDisabled=true
    isDisabled() {
      return !(this.formLogin.userName && this.formLogin.userPwd)
    },
    // 注册按钮状态
    isRegAble() {
      return !(
        this.formRegister.userName &&
        this.formRegister.userPwd &&
        this.formRegister.userPwd2
      )
    }
  },
  mounted() {
    this.getCookie()
  },
  methods: {
    // 登录/注册tab切换
    handleTab(type) {
      this.typeView = type
      this.clearInput() // 清除输入框
    },

    // 忘记密码界面
    forgetPwd() {
      this.$Message.info('忘记密码，请联系客服')
    },

    // 立即登录
    login() {
      // 禁用按钮或者加载中不可登录
      if (this.isDisabled || this.isLoading) {
        return false
      }

      // 正则表达式判断邮箱/手机号
      if (!this.$Valid.validUserName(this.formLogin.userName)) {
        this.$Message.error('请输入正确的邮箱/手机号')
        return false
      }

      // 正则表达式判断密码
      if (!this.$Valid.validPass(this.formLogin.userPwd)) {
        this.$Message.error('密码应为8到20位字母或数字!')
        return false
      }

      // 判断复选框是否被勾选,勾选则调用配置cookie方法
      if (this.checked) {
        // 传入账号名,密码,保存天数3个参数
        this.setCookie(this.formLogin.userName, this.formLogin.userPwd, 7)
      } else {
        // 清空Cookie
        this.clearCookie()
      }

      this.isLoading = true

      // 构造发送数据
      let form = {
        username: this.formLogin.userName,
        password: this.formLogin.userPwd
      }

      // 发送登录的网络请求
      login(form)
        .then((res) => {
          console.log('登录===', res)
          console.log('登录=', res.data)
          console.log('状态码', res.status)
          this.isLoading = false
          if (res.status == 200) {
            // 清除输入框
            this.clearInput()
            this.$Message.success('登录成功')
            // 保存用户信息
            this.$store.dispatch('userInfo/saveInfo', res.data)
            console.log('成功')
            console.log('11:', store.state.userInfo.data)
            // 路由跳转
            this.$router.push('/')
          } else {
            console.log('失败')
            this.$Message.error(res.msg)
          }
        })
        .catch(() => {
          this.isLoading = false
        })
    },

    // 立即注册
    register() {
      if (this.isRegAble || this.isLoading) {
        return false
      }

      // 校验
      if (!this.$Valid.validUserName(this.formRegister.userName)) {
        this.$Message.error('请输入正确的邮箱/手机号')
        return false
      } else if (!this.$Valid.validPass(this.formRegister.userPwd)) {
        this.$Message.error('密码应为8到20位字母或数字!')
        return false
      } else if (!this.$Valid.validPass(this.formRegister.userPwd2)) {
        this.$Message.error('确认密码有误')
        return false
      } else if (this.formRegister.userPwd2 !== this.formRegister.userPwd) {
        this.$Message.error('两次密码不一致')
        return false
      }

      this.isLoading = true

      let data = {
        username: this.formRegister.userName,
        password: this.formRegister.userPwd2
      }

      register(data)
        .then((res) => {
          this.isLoading = false
          console.log('注册===', res)
          if (res.code == 0) {
            // 清除输入框
            this.clearInput()
            this.$Message.success('注册成功')
            // 保存用户信息
            this.$store.dispatch('userInfo/saveInfo', res.data)
            // 路由跳转
            this.$router.push('/')
          } else {
            this.$Message.error(res.msg)
          }
        })
        .catch(() => {
          this.isLoading = false
        })
    },

    // 设置cookie
    setCookie(user_name, user_pwd, exdays) {
      // 获取时间
      let exdate = new Date()
      // 保存的天数
      exdate.setTime(exdate.getTime() + 24 * 60 * 60 * 1000 * exdays)
      // 字符串拼接cookie
      window.document.cookie =
        'userName' + '=' + user_name + ';path=/;expires=' + exdate.toUTCString()
      window.document.cookie =
        'userPwd' + '=' + user_pwd + ';path=/;expires=' + exdate.toUTCString()
    },

    // 读取cookie
    getCookie() {
      if (document.cookie.length > 0) {
        // 这里显示的格式需要切割一下自己可输出看下
        let arr = document.cookie.split('; ')
        console.log(arr)
        for (let i = 0; i < arr.length; i++) {
          // 再次切割
          let arr2 = arr[i].split('=')
          // 判断查找相对应的值
          if (arr2[0] == 'userName') {
            // 保存数据并赋值
            this.formLogin.userName = arr2[1]
          } else if (arr2[0] == 'userPwd') {
            this.formLogin.userPwd = arr2[1]
          }
        }
      }
    },

    //清除cookie
    clearCookie() {
      // 修改前2个值都为空，天数为负1天就好了
      this.setCookie('', '', -1)
    },

    // 清空输入框
    clearInput() {
      this.formLogin = {
        userName: '',
        userPwd: ''
      }
      this.formRegister = {
        userName: '',
        userPwd2: '',
        userPwd: ''
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.login-container {
  background: url('../assets/footer-animation-bg.svg') center center;
  background-color: #050828;
  background-position: center;
  background-size: cover;
  position: relative;
  width: 100%;
  height: 100%;

  .about-image {
    padding-top: 3%;
    padding-left: 13%;

    .img-fluid {
      height: 800px;
    }
  }

  .login-box {
    position: absolute;
    left: 60%;
    top: 50%;
    -webkit-transform: translateY(-50%);
    transform: translateY(-50%);
    box-sizing: border-box;
    text-align: center;
    box-shadow: 0 1px 11px rgba(0, 0, 0, 0.3);
    border-radius: 2px;
    /*margin: 100px auto 0;*/
    width: 420px;
    background: #fff;
    padding: 45px 35px;
    .option {
      text-align: left;
      margin-top: 18px;
      .checked {
        padding-left: 5px;
      }
      .forget-pwd,
      .goback {
        float: right;
        font-size: 14px;
        font-weight: 400;
        color: #4981f2;
        line-height: 20px;
        cursor: pointer;
      }
      .protocol {
        color: #4981f2;
        cursor: pointer;
      }
    }

    .login-text {
      width: 100%;
      text-align: center;
      padding: 0 0 30px;
      font-size: 24px;
      letter-spacing: 1px;
      a {
        padding: 10px;
        color: #969696;
        &.active {
          font-weight: 600;
          color: rgba(73, 129, 242, 1);
          border-bottom: 2px solid rgba(73, 129, 242, 1);
        }
        &:hover {
          border-bottom: 2px solid rgba(73, 129, 242, 1);
        }
      }
      b {
        padding: 10px;
      }
    }
    .title {
      font-weight: 600;
      padding: 0 0 30px;
      font-size: 24px;
      letter-spacing: 1px;
      color: rgba(73, 129, 242, 1);
    }

    .input-box {
      .input {
        &:nth-child(1) {
          /*margin-top: 10px;*/
        }
        &:nth-child(2),
        &:nth-child(3) {
          margin-top: 20px;
        }
      }
    }

    .loginBtn {
      width: 100%;
      height: 45px;
      margin-top: 40px;
      font-size: 15px;
    }

    .input {
      padding: 10px 0px;
      font-size: 15px;
      width: 350px;
      color: #2c2e33;
      outline: none; // 去除选中状态边框
      border: 1px solid #fff;
      border-bottom-color: #e7e7e7;
      background-color: rgba(0, 0, 0, 0); // 透明背景
    }

    input:focus {
      border-bottom-color: #0f52e0;
      outline: none;
    }
    .button {
      line-height: 45px;
      cursor: pointer;
      margin-top: 50px;
      border: none;
      outline: none;
      height: 45px;
      width: 350px;
      background: rgba(216, 216, 216, 1);
      border-radius: 2px;
      color: white;
      font-size: 15px;
    }
  }

  // 滚动条样式
  ::-webkit-scrollbar {
    width: 10px;
  }

  ::-webkit-scrollbar-track {
    -webkit-box-shadow: inset006pxrgba(0, 0, 0, 0.3);
    border-radius: 8px;
  }

  ::-webkit-scrollbar-thumb {
    border-radius: 10px;
    background: rgba(0, 0, 0, 0.2);
    -webkit-box-shadow: inset006pxrgba(0, 0, 0, 0.5);
  }

  ::-webkit-scrollbar-thumb:window-inactive {
    background: rgba(0, 0, 0, 0.4);
  }

  //设置更改input 默认颜色
  ::-webkit-input-placeholder {
    /* WebKit browsers */
    color: #bebebe;
    font-size: 16px;
  }

  ::-moz-placeholder {
    /* Mozilla Firefox 19+ */
    color: #bebebe;
    font-size: 16px;
  }

  :-ms-input-placeholder {
    /* Internet Explorer 10+ */
    color: #bebebe;
    font-size: 16px;
  }

  input:-webkit-autofill {
    box-shadow: 0 0 0px 1000px rgba(255, 255, 255, 1) inset;
    -webkit-box-shadow: 0 0 0px 1000px rgba(255, 255, 255, 1) inset;
    -webkit-text-fill-color: #2c2e33;
  }

  .ivu-checkbox-wrapper {
    margin-right: 0;
  }
}
</style>
