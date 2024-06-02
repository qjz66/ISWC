<template>
  <div>
    <header class="header">
      <div class="container">
        <div class="header-inner">
          <div class="logo" @click="RouteToHome">
            <span>
              <img src="../picture/icon.png" alt="logo" /> RUMOUR ROBOT
            </span>
          </div>
          <nav class="main-nav">
            <ul class="menu">
              <li class="menu-item">
                <span
                  @click="HandleClass(1)"
                  :class="{ highlight: currentIndex === 1 }"
                  >首页</span
                >
              </li>
              <li class="menu-item">
                <span
                  @click="HandleClass(2)"
                  :class="{ highlight: currentIndex === 2 }"
                  >数据大屏</span
                >
              </li>
              <li class="menu-item">
                <span
                  @click="HandleClass(3)"
                  :class="{ highlight: currentIndex === 3 }"
                  >数据广场</span
                >
              </li>
              <li class="menu-item">
                <span
                  @click="HandleClass(4)"
                  :class="{ highlight: currentIndex === 4 }"
                  >个人中心</span
                >
              </li>
            </ul>
          </nav>
          <span class="header-btn" @click="RouteToUploadFile"
            ><i class="fa fa-rocket"></i>检测谣言</span
          >
        </div>
      </div>
    </header>
  </div>
</template>

<script>
export default {
  name: 'Header',
  data() {
    return {
      currentIndex: null
    }
  },
  methods: {
    HandleClass(index) {
      this.currentIndex = index
      console.log(index, this.currentIndex)
      if (index === 1) this.$router.push('/')
      else if (index === 2) this.$router.push('/dataScreen')
      else if (index === 3) this.$router.push('/dataSquare')
      else if (index === 4) this.$router.push('/personalSpace')
    },
    RouteToHome() {
      this.$router.push('/')
    },
    RouteToUploadFile() {
      this.$router.push('/uploadFile')
    }
  },
  mounted() {
    console.log(this.$route.path)
    if (this.$route.path === '/') this.currentIndex = 1
    else if (this.$route.path === '/dataScreen') this.currentIndex = 2
    else if (this.$route.path === '/dataSquare') this.currentIndex = 3
    else if (this.$route.path === '/personalSpace') this.currentIndex = 4
  }
}
</script>

<style lang="scss" scoped>
.header {
  position: absolute;
  width: 100%;
  height: 106px;
  background-color: transparent;
  display: flex;
  align-items: center;
  justify-content: space-between;
  z-index: 99999;

  .header-inner {
    display: flex;
    align-items: center;
    justify-content: flex-start;

    span {
      cursor: pointer;
    }

    .logo span {
      font-size: 24px;
      line-height: 1;
      font-weight: 500;
      display: flex;
      align-items: center;
      color: #fff;
      font-family: 'Michroma', sans-serif;

      img {
        height: 48px;
        margin-right: 3px;
        width: 49px;
      }
    }

    .main-nav {
      position: absolute;
      top: 50%;
      -webkit-transform: translateY(-50%);
      -ms-transform: translateY(-50%);
      -o-transform: translateY(-50%);
      transform: translateY(-50%);
      left: 45%;
      padding-left: 5px;

      .menu {
        display: flex;
        align-items: center;
        justify-content: flex-start;

        li {
          position: relative;
          margin-right: 23px;
          letter-spacing: 0.6px;

          span {
            cursor: pointer;
            font-family: 'Chakra Petch', sans-serif;
            position: relative;
            display: block;
            color: #fff;
            font-size: 24px;
            font-weight: 700;
            line-height: 24px;
            text-transform: uppercase;
            padding-right: 21px;
            padding-top: 20px;
            padding-bottom: 20px;
          }

          span:hover {
            background-image: linear-gradient(
              267deg,
              #641ff9 27.93%,
              #02c5f7 99.79%,
              #c6bafd -31.69%,
              #b7a2fd -22.78%
            );
            background-clip: text;
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
          }
        }

        li.menu-item.current > span::after {
          background-image: linear-gradient(
            267deg,
            #641ff9 27.93%,
            #02c5f7 99.79%,
            #c6bafd -31.69%,
            #b7a2fd -22.78%
          );
          background-clip: text;
          -webkit-background-clip: text;
          -webkit-text-fill-color: transparent;
        }
      }
    }

    .header-btn {
      font-family: 'Chakra Petch', sans-serif;
      font-size: 22px;
      display: flex;
      align-items: center;
      position: relative;
      background: linear-gradient(
        267deg,
        #641ff9 27.93%,
        #02c5f7 99.79%,
        #c6bafd -31.69%,
        #b7a2fd -22.78%
      );
      color: #fff;
      font-weight: 700;
      text-transform: uppercase;
      padding: 8px 30px;
      margin-left: auto;
      height: 48px;
      border: none;
      border-radius: 30px;
      -webkit-transition: all 0.3s ease;
      -moz-transition: all 0.3s ease;
      -ms-transition: all 0.3s ease;
      -o-transition: all 0.3s ease;
      transition: all 0.3s ease;

      .fa-rocket {
        margin-left: -8px;
        margin-right: 5px;
        margin-top: 3px;
      }
    }

    .header-btn:hover {
      background: linear-gradient(
        -267deg,
        #641ff9 27.93%,
        #02c5f7 99.79%,
        #c6bafd -31.69%,
        #b7a2fd -22.78%
      );
    }

    .mobile-button {
      position: absolute;
      top: 50%;
      -webkit-transform: translateY(-50%);
      -ms-transform: translateY(-50%);
      -o-transform: translateY(-50%);
      transform: translateY(-50%);
      right: 15px;
      width: 26px;
      height: 26px;
      display: none;
      background-color: transparent;
      cursor: pointer;
      -webkit-transition: all 0.3s ease;
      -moz-transition: all 0.3s ease;
      -ms-transition: all 0.3s ease;
      -o-transition: all 0.3s ease;
      transition: all 0.3s ease;

      span {
        position: absolute;
        top: 50%;
        left: 0;
        width: 100%;
        height: 3px;
      }
    }

    .mobile-button:before,
    .mobile-button:after {
      content: '';
      position: absolute;
      top: 50%;
      left: 0;
      width: 100%;
      height: 3px;
    }

    .mobile-button:before,
    .mobile-button:after,
    .mobile-button span {
      background-color: #fff;
      -webkit-transition: all ease 0.3s;
      -moz-transition: all ease 0.3s;
      transition: all ease 0.3s;
    }

    .mobile-button:before {
      -webkit-transform: translateY(-7px);
      -moz-transform: translateY(-7px);
      transform: translateY(-7px);
    }

    .mobile-button:after {
      -webkit-transform: translateY(7px);
      -moz-transform: translateY(7px);
      transform: translateY(7px);
    }

    .mobile-button.active:before {
      -webkit-transform: rotate3d(0, 0, 1, 45deg);
      -moz-transform: rotate3d(0, 0, 1, 45deg);
      transform: rotate3d(0, 0, 1, 45deg);
    }

    .mobile-button.active:after {
      -webkit-transform: rotate3d(0, 0, 1, -45deg);
      -moz-transform: rotate3d(0, 0, 1, -45deg);
      transform: rotate3d(0, 0, 1, -45deg);
    }

    .mobile-button.active span {
      opacity: 0;
    }
  }
}

.highlight {
  background-image: linear-gradient(
    267deg,
    #641ff9 27.93%,
    #02c5f7 99.79%,
    #c6bafd -31.69%,
    #b7a2fd -22.78%
  );
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
</style>