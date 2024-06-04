<template>
  <div class="body-container">
    <!-- Header -->
    <div class="header">
      <Header />
    </div>
    <!-- 分析报告 -->
    <div class="analysis-result">
      <dv-decoration-7 style="width: 275px; height: 150px"
        >分析报告</dv-decoration-7
      >
      <span class="share-report" @click="shareResult"
        >分享该报告 <img src="../picture/zhuanfa.png" alt=""
      /></span>
    </div>
    <!-- 新闻标题 -->
    <div class="news-title">
      <el-col :span="16">
        <el-card shadow="always">
          {{ content }}
        </el-card>
      </el-col>
    </div>
    <!-- 分析的结论 -->
    <div class="news-result">
      <span>置信概率:</span>
      <dv-decoration-9 style="width: 150px; height: 150px"
        >{{ possibility }}%</dv-decoration-9
      >
    </div>
    <!-- 分析的过程 -->
    <!-- <div class="news-analysis">
      <div class="news-analysis-title">分析：</div>
      <p>新闻隐性标签提取：科技、健康、生活</p>
      <p>关键词（热点分析）：屏 手机 三星</p>
      <p>视图权重参数：语义：0.82;情感：0.71;风格：0.93</p>
      <p>分析参考：科技：0.93;健康：0.76;生活：0.86</p>
    </div> -->
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import { shareAnalysisResult, getNowTime } from '../utils/api'
export default {
  name: 'analysisResult',
  components: {
    Header
  },
  data() {
    return {
      content: '',
      possibility: 0
    }
  },
  mounted() {
    this.content = this.$route.params.content
    this.possibility = (this.$route.params.possibility * 100).toFixed(1)
  },
  methods: {
    // 分享新的动态
    shareResult() {
      let params = {
        id: this.$store.state.userInfo.data.id
      }

      let data = new FormData()
      data.append('content', this.content)
      data.append('date', getNowTime())
      data.append('fromName', this.$store.state.userInfo.data.userName)

      shareAnalysisResult(params, data).then((res) => {
        console.log('分享动态:', res)
        if (res.status == 200) {
          console.log('分享动态成功')
        } else {
          console.log('分享动态失败')
        }
      })
    }
  }
}
</script>

<style lang="less">
.body-container {
  background: url('../assets/footer-animation-bg.svg') center center;
  background-color: #050828;
}
.header {
  height: 106px;
  justify-content: space-between;
  align-items: center;
}

.analysis-result {
  top: 80px;
  width: 100%;
  margin-top: -20px;
  position: relative;

  .dv-decoration-7 {
    position: relative;
    left: 50%;
    transform: translate(-50%, -50%);
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
    font-size: 45px;
  }

  .share-report {
    cursor: pointer;
    position: relative;
    top: -163px;
    left: 74%;
    font-size: 20px;
    color: #1296db;
    right: 0%;

    img {
      margin-bottom: 5px;
      width: 25px;
      height: 25px;
    }
  }
}

.news-title {
  display: flex;
  justify-content: center;
  align-items: center;

  .el-col-16 {
    margin-left: 20px;
    padding-left: 20px;
    padding-right: 20px;
  }
  .el-card {
    font-size: 20px;
    color: white;
    background-color: #363b79;
    border: none;
    border-radius: 20px;
    margin-top: 20px;
  }

  .el-card:hover {
    background-image: linear-gradient(to right, #25aae1, #40e495);
    box-shadow: 0 7px 15px 0 rgba(49, 196, 190, 0.75);
  }
}

.news-result {
  margin-top: 40px;
  display: flex;
  justify-content: center;
  align-items: center;

  span {
    font-size: 30px;
    color: #e65e5e;
  }

  .dv-decoration-9 {
    font-size: 30px;
    color: #ff0000;
  }
}

.news-analysis {
  margin-top: 40px;
  margin-left: 20%;
  padding-bottom: 20px;
  width: 60%;
  font-size: 20px;
  color: white;

  .news-analysis-title {
    margin-top: 20px;
    margin-bottom: 50px;
    font-size: 24px;
    color: skyblue;
  }

  p {
    text-indent: 2em;
  }
}
</style>