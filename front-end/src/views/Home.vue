<template>
  <div class="body-container">
    <Header />

    <section class="home" id="home">
      <div class="container">
        <div class="row d-flex align-items-center justify-content-between">
          <div class="col-lg-6">
            <div class="home-content">
              <h1>
                <span>RUMOUR BUZZ</span> <br />面向用户的自动化谣言检测系统
              </h1>
              <p>
                我们谣言检测模型基于TextCNN模型和MLP技术的多视图特征提取器、基于领域记忆库的新闻隐性标签发掘和丰富模块以及基于领域适配器的领域视图权重聚合模块,可以准确识别出您给出的新闻是否为一则谣言
              </p>
              <div class="my-btn">
                <span class="main-btn" @click="RouteToUploadFile"
                  ><i class="fa fa-rocket"></i> <span>检测谣言</span></span
                >
              </div>
            </div>
          </div>
          <div class="col-lg-6">
            <div class="right-box">
              <div class="home-image">
                <img
                  class="img-fluid"
                  src="../picture/home.png"
                  alt="home image"
                />
                <img
                  class="img-fluid shape"
                  src="../picture/shape.png"
                  alt="shape"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="features">
      <div class="container">
        <div class="row">
          <div class="col-12">
            <h3 class="main-title">为什么选择我们这个RUMOUR BUZZ</h3>
          </div>
        </div>
        <div class="row">
          <div class="col-lg-4 col-md-6 item">
            <div class="feature-item">
              <img src="../picture/11.png" alt="feature image" />
              <h3 class="feature-title">
                基于TextCNN模型和多层感知机的多视图特征提取
              </h3>
              <p>多维度提取语义、情感、风格特征</p>
            </div>
          </div>
          <div class="col-lg-4 col-md-6 item">
            <div class="feature-item">
              <img src="../picture/21.png" alt="feature image" />
              <h3 class="feature-title">基于领域记忆库的隐性标签发掘</h3>
              <p>自动化地发掘、捕捉和储存隐性领域特征信息</p>
            </div>
          </div>
          <div class="col-lg-4 col-md-6 item">
            <div class="feature-item">
              <img src="../picture/31.png" alt="feature image" />
              <h3 class="feature-title">基于领域适配器的跨视图权重聚合</h3>
              <p>建模和吸收领域差异，修正特征预测，实现准确预测</p>
            </div>
          </div>
        </div>
      </div>
    </section>
    <Footer />
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import Footer from '../components/Footer.vue'

export default {
  name: 'home',
  components: {
    Header,
    Footer
  },
  methods: {
    RouteToUploadFile() {
      this.$router.push('/uploadFile')
    }
  }
}
</script>

<style scope>
.body-container {
  background: url('../assets/footer-animation-bg.svg') center center;
  background-color: #050828;
}

.home {
  position: relative;
  overflow: hidden;
  background-size: cover;
  padding-top: 115px;
}

.home .home-content {
  margin-top: 40px;
}

.home .home-content h1 {
  position: relative;
  color: #fff;
  font-size: 45px;
  font-weight: 700;
  text-transform: capitalize;
  margin-bottom: 25px;
  line-height: 1.2;
  text-align: left;
  width: 700px;
}

.home .home-content h1:before {
  content: '';
  width: 200px;
  height: 200px;
  background: linear-gradient(
    180deg,
    #109b9b -43.08%,
    #052a2a -7.55%,
    #2b7a7a 33.87%,
    #27c793 82.91%
  );
  filter: blur(40px);
  display: inline-block;
  position: absolute;
  top: -63px;
  left: -35px;
  border-radius: 100%;
  opacity: 0.5;
  z-index: 1;
}

.home .home-content h1 span {
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

.home .home-content p {
  color: #d1cccc;
  font-size: 17px;
  font-weight: 400;
  line-height: 1.8;
  margin-bottom: 45px;
}

.home .home-content .main-btn {
  font-size: 20px;
  padding: 14px 25px;
}

.home .home-content .custom-btn {
  margin-right: 0;
  background: transparent;
  border: none;
  color: #d7d4d4;
}

.home .home-content .custom-btn:hover {
  color: #fff;
}

.home .home-facts {
  margin-top: 45px;
}

.home .home-facts .facts-list {
  display: flex;
}

.home .home-facts .facts-list li {
  position: relative;
  margin: 0 45px;
}

.home .home-facts .facts-list li:first-of-type {
  margin-left: 0;
}

.home .home-facts .facts-list li span {
  display: block;
  color: #d1cccc;
  font-weight: 600;
  font-size: 15px;
  font-family: 'Montserrat', sans-serif;
}

.home .home-facts .facts-list li span.number {
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
  font-size: 20px;
  font-family: 'Montserrat', sans-serif;
  margin-bottom: 10px;
}

.home .right-box {
  position: relative;
  padding-left: 30px;
}

.home .home-image {
  position: relative;
}

.home .home-image img {
  margin-left: 50px;
  display: block;
}

.home .home-image img.shape {
  position: absolute;
  top: -40px;
  left: 66px;
  opacity: 0.4;
  z-index: 1;
  -webkit-animation: moving 1s infinite alternate;
  animation: moving 1s infinite alternate;
}

@-webkit-keyframes moving {
  0% {
    transform: translateX(0);
  }
  100% {
    transform: translateX(-30px);
  }
}

/* --------------------------------------
  Features
-----------------------------------------*/

.features {
  padding-top: 120px;
  padding-bottom: 50px;
}

.features .item:nth-of-type(n + 4) .feature-item {
  margin-top: 30px;
}

.feature-item {
  background: rgba(255, 255, 255, 0.03);
  border: 2px solid rgba(255, 255, 255, 0.1);
  padding: 20px;
  border-radius: 24px;
  -webkit-transition: all 0.3s ease;
  -moz-transition: all 0.3s ease;
  -ms-transition: all 0.3s ease;
  -o-transition: all 0.3s ease;
  transition: all 0.3s ease;
}

.feature-item:hover {
  border: 2px solid #02c5f7;
  transform: translateY(-10px);
}

.features .feature-item img {
  margin-bottom: 25px;
}

.features .feature-item h3 {
  font-size: 18px;
  color: #fff;
  line-height: 1.5;
  margin-bottom: 15px;
}

.features .feature-item p {
  color: #d1cccc;
  font-size: 17px;
  font-weight: 400;
  line-height: 1.8;
  margin-bottom: 0;
}
</style>