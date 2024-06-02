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
    </div>
    <!-- 新闻标题 -->
    <div class="news-title">
      <dv-decoration-11
        >新闻标题：<strong>人工智能现在有多智能</strong></dv-decoration-11
      >
    </div>
    <!-- 分析的结论 -->
    <div class="news-result">
      <span>此新闻为谣言的概率为:</span>
      <dv-decoration-9 style="width: 150px; height: 150px">66%</dv-decoration-9>
    </div>
    <!-- 分析的过程 -->
    <div class="news-analysis">
      <div class="news-analysis-title">分析过程如下：</div>
      <p>
        近日，一组照片在网上火了：皮肤黝黑的少年站在网球场上对着镜头露出腼腆的笑，他背着一个竹编的背篓，里面装着一双网球拍。这名叫王发的少年来自云南临沧沧源佤族自治县，这里延绵崎岖的山路阻隔着少年与外界的联系，却没有阻拦住少年成为职业球员的梦想。
      </p>
      <p>
        为了实现自己的梦想，王发连续六年不到六点起床，挥拍7000余次，这才取得了参加国家巡回赛的机会。与王发相同的还有另一个少年，名为石学念，他因在摆摊时练习篮球而走红，网友们看到他围着围裙打球的视频后深受触动，其中一个高赞评论说道：“身上穿的是生活，手上打的是梦想”。
      </p>
      <p>
        这样的少年还有很多很多，他们或许来自不同的家庭，来自不同的地区，但他们都有同一个名字：追梦者。曾经我认为，梦想是个很宏大的词汇，它不仅要立意高远，在实践中更要惊天动地，似乎只有每一步都做出了巨大成就才能称之为“正在向梦想迈进”。但事实上，真正的梦想应该是浸润在日常生活之中的“微不足道”的坚持，是平凡日子里的“微小习惯”。
      </p>
      <p>
        如果说逐梦是奔向星辰大海，那么奔赴的前提，一定是踏上平凡之路。
        踏上平凡之路，让我们沉淀内心，明确方向。在新媒体技术高速发展的今天，纷繁复杂的信息充斥着我们的屏幕，使我们可以足不出户而知天下，可以体验他人的人生、看到不同世界的生活。因此，我们不停地向外张望，却忘记了审视自己的内心，也忘记了被掩盖的梦想。踏上平凡之路，即回归平凡、回归生活，将目光收回到自己身上，审视当下的生活状态，明确自己的生活目标和前进方向，为自己点亮一个路标。
      </p>
    </div>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
export default {
  name: 'analysisResult',
  components: {
    Header
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
  margin-top: -20px;
  display: flex;
  justify-content: center;
  align-items: center;

  .dv-decoration-7 {
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
}

.news-title {
  display: flex;
  justify-content: center;
  align-items: center;

  .dv-decoration-11 {
    width: 75%;
    height: 80px;
    font-size: 30px;
    color: white;
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