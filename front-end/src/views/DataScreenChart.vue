<template>
  <div class="dataScreen">
    <Row class="content">
      <Col span="8">
        <!-- 微博热榜分析 -->
        <div class="list">
          <div class="left">
            <span class="title"
              ><span class="title-gradient">微博热榜分析</span></span
            >
            <div class="left-info">
              <el-row :gutter="12">
                <el-col :span="24">
                  <el-card shadow="always"
                    >#泰坦号深海内爆一年后又有富豪要挑战#美国俄亥俄州的一位亿万富翁计划乘坐深海潜水器抵达“泰坦号”事故发生地，以证明该行业的安全性。<span
                      style="color: red"
                      >置信率:80%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    【#一老板称可用30万赎回被卖缅甸夫妻#】据海报新闻报道：近日，重庆一对夫妻易先生和王女士被骗到泰
                    国卖到缅甸的事情引发关注。5月27日，王女士的表姐刘女士告诉海报新闻记者，一位自称是缅甸妙瓦底的老板向他们表示，可以用30万赎回两人。<span
                      style="color: red"
                      >置信率:78%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    #20岁女生留学新加坡坠亡#，家属：离毕业不到10天，财物遗失。<span
                      style="color: red"
                      >置信率:91%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    重庆21岁女孩常州被害后遭分尸，凶手以被害人语气向家属要钱。<span
                      style="color: red"
                      >置信率:87%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    【#法国镇长挪用公款供养女巫#】据英国广播公司（BBC）27日报道，法国知名旅游景点阿格德镇镇长吉尔·德·埃托雷因挪用公款于今年3月被捕。最新调查发现，被挪用的公款，很大一部分花在了当地一个自称能“通灵”的女巫身上。<span
                      style="color: red"
                      >置信率:68%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    【#遮挡富士山网红打卡点的黑网已破洞#
                    日媒：“没想到这么快”】据日本共同社报道，日本山梨县富士河口湖町一处网红打卡点因游客太多且随之带来的不文明行为增多而竖起黑色屏障网，然而当地政府28日表示，用来遮挡风景的黑色屏障网目前已经出现了多个破洞。<span
                      style="color: red"
                      >置信率:77%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    【#飞机滑出后女子疑因情感问题强行要求下机#】据看看新闻KNEWS报道：据网友爆料，5月27日深夜，由浦东飞往深圳的DZ6208航班，飞机滑出后，一名女子疑因情感问题，强行要求下飞机。最终机上所有乘客也只得下机。机组根据规定对飞机进行二次清舱后，旅客重新登机起飞，航班因此延误2小时。<span
                      style="color: red"
                      >置信率:84%</span
                    >
                  </el-card>
                </el-col>
              </el-row>
            </div>
          </div>
        </div>
      </Col>
      <Col span="8">
        <div class="circlePie" id="circlePie">
          <canvas id="main" width="500" height="500"></canvas>
          <canvas id="dot"></canvas>
        </div>
      </Col>
      <Col span="8">
        <!-- b站热榜分析 -->
        <div class="list">
          <div class="right">
            <span class="title"
              ><span class="title-gradient">b站热榜分析</span></span
            >
            <div class="right-info">
              <el-row :gutter="12">
                <el-col :span="24">
                  <el-card shadow="always">
                    【#蛐蛐你是什么意思#？】蛐蛐一词多用于形容小声议论他人，尤其指背后说人闲话、坏话的行为。近年来因短视频流行，该词成网络热梗。<span
                      style="color: red"
                      >置信率:79%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    #牛排价格都不如猪排贵了#最近一段时间，牛肉价格“大跳水”，引发关注。以往上好的牛腱子肉通常需要一斤花费百来元，现在仅需几十元就可以拿下。曾经高高在上的牛肉为啥会跌价，
                    #可以实现牛肉自由了吗#？<span style="color: red"
                      >置信率:70%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    #知名配音演员公开维权#被侵权配音演员林景称“真正让人在意的不是AI，而是侵权。”#云中君钟离配音演员遭AI侵权<span
                      style="color: red"
                      >置信率:91%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    【#韩称朝鲜再向韩放飞600多个带垃圾的气球#】据央视军事报道：当地时间6月2日上午，据韩国联合参谋本部消息，朝鲜从1日起再次向韩国放飞600多个携带垃圾等污染物的气球。1日晚间起，韩国京畿道以及首尔等地区陆续发现朝鲜放飞的气球，韩国多地政府向民众紧急推送应急手机警报，呼吁民众减少外出并小心空中飘浮物落下。韩国军方提醒市民发现可疑物体时不要触碰并立刻就近向军方或警方报告。<span
                      style="color: red"
                      >置信率:82%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    【#总被闹钟吓醒心脏承受不住#】几乎每个需要早起的日子，闹钟都会派上用场。有些人怕起不来，会定多个闹钟或换一个急促的铃声，殊不知这样反而会开启“糟糕”的一天。戳图，看#这样定闹钟可以无痛起床#↓↓#为什么有的人一天只睡4小时也很精神#<span
                      style="color: red"
                      >置信率:83%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    #白开水在4小时内饮用活性最佳#水对人体很重要，人可以三天无饭，但绝不能三天无水。想要健康，必须要正确认识水，选对水，喝对水。#矿泉水苏打水喝了真的健康吗#？<span
                      style="color: red"
                      >置信率:79%</span
                    >
                  </el-card>
                </el-col>
                <el-col :span="24">
                  <el-card shadow="always">
                    【安徽#芜湖再通报黄金回收店鬼秤事件#：拟对涉事店铺罚款并列入严重违法失信名单】6月1日，针对“网络博主打假称安徽芜湖一金店惊现‘五两秤’”一事，安徽芜湖市镜湖区市场监督管理局再度发布情况通报。通报称，经查明，该店使用加装芯片改造破坏其准确度的计量器具从事贵金属回收贸易结算；在收购贵金属经营活动中未保证商品量的量值准确。<span
                      style="color: red"
                      >置信率:92%</span
                    >
                  </el-card>
                </el-col>
              </el-row>
            </div>
          </div>
        </div>
      </Col>
    </Row>
    <Row class="bottom-list">
      <Col span="16">
        <!-- 近几天微博和b站的谣言数 -->
        <div class="list">
          <div class="bottom">
            <span class="title"
              ><span class="title-gradient">近几月各大平台的谣言数</span></span
            >
            <span class="angle1"></span>
            <span class="angle2"></span>
            <span class="angle3"></span>
            <span class="angle4"></span>
            <chart4 ref="chart9" id="bottom_1"></chart4>
          </div>
        </div>
      </Col>
      <Col span="8">
        <!-- 各大平台谣言占比 -->
        <div class="list right-bottom">
          <div class="bottom bottom1">
            <span class="title"
              ><span class="title-gradient">各大平台谣言占比</span></span
            >
            <span class="angle1"></span>
            <span class="angle2"></span>
            <span class="angle3"></span>
            <span class="angle4"></span>
            <chart7 ref="chart12" id="bottom_4"></chart7>
          </div>
        </div>
      </Col>
    </Row>
  </div>
</template>

<script>
const chart1 = () => import('../components/dataScreen/chart1')
const chart2 = () => import('../components/dataScreen/chart2')
const chart3 = () => import('../components/dataScreen/chart3')
const chart4 = () => import('../components/dataScreen/chart4')
const chart5 = () => import('../components/dataScreen/chart5')
const chart6 = () => import('../components/dataScreen/chart6')
const chart7 = () => import('../components/dataScreen/chart7')
const areaChart = () => import('../components/dataScreen/areaChart')
const radarChart = () => import('../components/dataScreen/radar')
import Header from '@/components/Header.vue'

export default {
  name: 'dataScreenChart',
  props: ['selectRangeDate'],
  components: {
    areaChart,
    radarChart,
    chart1,
    chart2,
    chart3,
    chart4,
    chart5,
    chart6,
    chart7,
    Header
  },
  data() {
    return {
      everyPer: 0,
      xOffset: 0,
      circle: {
        x: 250,
        y: 250,
        radius: 218
      },
      title: [
        // '数据1:456,789',
        // '数据2:123,12',
        // '数据3:134,23',
        // '数据4:234,234',
        // '数据5:123,123',
        // '数据6:678,123'
      ],
      // 模块一面积图数据
      cnfigData1: {
        color: '#5CB1C1',
        name: ['（次）', '（人）'],
        data: [
          {
            name: '次数',
            color: ['#9e70ff', '#6e5eff'],
            data: [200, 12, 21, 54, 260, 130, 210]
          },
          {
            name: '人数',
            color: ['#48cefd', '#5356f1'],
            data: [50, 182, 234, 191, 190, 30, 10]
          }
        ]
      },
      configData2: {
        data: [213, 190, 137, 99, 63, 196, 248, 212, 248, 112]
      },
      // 模块一雷达图数据
      chatRadarData: {
        title: '标题',
        position: ['center', '85%'],
        center: ['50%', '50%'],
        indicator: [
          { text: '分类1' },
          { text: '分类2' },
          { text: '分类3' },
          { text: '分类4' },
          { text: '分类5' },
          { text: '分类6' }
        ],
        data: [
          {
            name: '次数',
            color: '#0DF5F8',
            value: [100, 8, 0.4, -80, 2000, 332]
          },
          {
            name: '人数',
            color: '#7921AD',
            value: [60, 5, 0.3, -100, 1500, 221]
          }
        ]
      },
      // 模块二雷达图数据
      officeRadarData: {
        title: '标题',
        position: ['center', '85%'],
        center: ['50%', '50%'],
        indicator: [
          { text: '分类1' },
          { text: '分类2' },
          { text: '分类3' },
          { text: '分类4' },
          { text: '分类5' },
          { text: '分类6' },
          { text: '分类7' },
          { text: '分类8' },
          { text: '分类9' },
          { text: '分类10' }
        ],
        data: [
          {
            name: '数据',
            color: '#55D35B',
            value: [213, 190, 137, 99, 63, 196, 248, 212, 248, 112]
          }
        ]
      },
      // 模块三雷达图数据
      friendRadarData: {
        title: '分类标题',
        position: ['center', '85%'],
        center: ['50%', '50%'],
        indicator: [
          { text: '分类1' },
          { text: '分类2' },
          { text: '分类3' },
          { text: '分类4' },
          { text: '分类5' },
          { text: '分类6' }
        ],
        data: [
          {
            name: '数据',
            color: '#FA8486',
            value: [100, 8, 0.4, -80, 2000, 332]
          }
        ]
      },
      // 模块四雷达图数据
      momentsRadarData: {
        title: '标题',
        position: ['center', '85%'],
        center: ['50%', '50%'],
        indicator: [
          { text: '分类1' },
          { text: '分类2' },
          { text: '分类3' },
          { text: '分类4' },
          { text: '分类5' },
          { text: '分类6' }
        ],
        data: [
          {
            name: '个数',
            color: '#D91748',
            value: [100, 8, 0.4, -80, 2000, 332]
          }
        ]
      },
      warea: { x: 250, y: 250, max: 700 },
      dots: [],
      resizeFn: null,
      animationFrame1: null,
      animationFrame2: null,
      centerBox: {
        width: 0,
        height: 0
      }
    }
  },
  methods: {
    drawDot() {
      let canvas = document.getElementById('dot')
      canvas.width = document.querySelector('#circlePie').offsetWidth
      canvas.height = document.querySelector('#circlePie').offsetHeight
      let ctx = canvas.getContext('2d')
      ctx.clearRect(0, 0, canvas.width, canvas.height)
      // 将鼠标坐标添加进去，产生一个用于比对距离的点数组
      let ndots = [this.warea].concat(this.dots)
      this.dots.forEach((dot) => {
        // 粒子位移
        dot.x += dot.xa
        dot.y += dot.ya
        // 遇到边界将加速度反向
        dot.xa *= dot.x > canvas.width || dot.x < 0 ? -1 : 1
        dot.ya *= dot.y > canvas.height || dot.y < 0 ? -1 : 1
        // 绘制点
        ctx.fillStyle = '#ffffff'
        ctx.beginPath()
        ctx.arc(dot.x - 0.5, dot.y - 0.5, 2, 0, 2 * Math.PI, true)
        ctx.closePath()
        ctx.fill()
        // 循环比对粒子间的距离
        for (let i = 0; i < ndots.length; i++) {
          let d2 = ndots[i]
          if (dot === d2 || d2.x === null || d2.y === null) continue
          let xc = dot.x - d2.x
          let yc = dot.y - d2.y
          // 两个粒子之间的距离
          let dis = Math.sqrt(xc * xc + yc * yc)
          // 距离比
          let ratio
          // 如果两个粒子之间的距离小于粒子对象的max值，则在两个粒子间画线
          if (dis < d2.max) {
            // 计算距离比
            ratio = (d2.max - dis) / d2.max
            // 画线
            ctx.beginPath()
            ctx.lineWidth = ratio / 2
            if (d2 === this.warea) {
              ctx.strokeStyle = 'rgba(255,255,255,0)'
            } else {
              // 距离变大 连线颜色变浅
              ctx.strokeStyle = 'rgba(255,255,255,' + (ratio + 0.2) + ')'
            }
            ctx.moveTo(dot.x, dot.y)
            ctx.lineTo(d2.x, d2.y)
            ctx.stroke()
          }
        }
        // 将已经计算过的粒子从数组中删除
        ndots.splice(ndots.indexOf(dot), 1)
      })
      this.animationFrame1 = window.requestAnimationFrame(this.drawDot)
    },
    rads(x) {
      // 弧度转换
      return (Math.PI * x) / 180
    },
    act() {
      //清空画布
      const canvas = document.querySelector('#main')
      canvas.style.width = this.centerBox.height + 'px'
      canvas.style.height = this.centerBox.height + 'px'
      const context = canvas.getContext('2d')
      context.clearRect(0, 0, canvas.width, canvas.height)
      this.drawPie(this.everyPer, context)
      this.animationFrame2 = window.requestAnimationFrame(this.act)
      this.everyPer += Math.PI / 180
      let speed = 0.07 //波浪速度，数越大速度越快
      this.xOffset += speed
    },
    drawPie(everyPer, context) {
      context.save()
      context.fillStyle = 'rgba(18,55,88,.2)'
      context.beginPath()
      context.arc(this.circle.x, this.circle.y, 245, 0, 2 * Math.PI, true)
      context.closePath()
      context.fill()
      context.restore()

      //外圆
      context.save()
      context.shadowBlur = 50
      context.shadowColor = '#123959'
      context.fillStyle = '#080D27'
      context.beginPath()
      context.arc(this.circle.x, this.circle.y, 235, 0, 2 * Math.PI, true)
      context.closePath()
      context.fill()
      context.restore()
      for (let i = 0; i < this.title.length; i++) {
        //绘制文字。
        context.save()
        // 画文字
        this.drawCircularText(
          this.circle,
          this.title[i],
          this.rads(i * 60 - 110),
          this.rads(i * 60 - 65),
          i,
          context
        )
        context.restore()
      }
      // 旋转小球
      let x = 240 * Math.cos(everyPer)
      let y = 240 * Math.sin(everyPer)
      context.save()
      context.fillStyle = 'rgb(56,252,253)'
      context.shadowBlur = 80
      context.shadowColor = '#39E9EE'
      context.translate(this.circle.x, this.circle.y)
      context.beginPath()
      context.arc(x, y, 5, 0, 2 * Math.PI)
      context.arc(-x, -y, 5, 0, 2 * Math.PI)
      context.closePath()
      context.fill()
      context.restore()
      //
      context.save()
      context.fillStyle = '#153776'
      context.beginPath()
      context.arc(this.circle.x, this.circle.y, 200, 0, 2 * Math.PI, true)
      context.closePath()
      context.fill()

      context.fillStyle = '#121535'
      context.beginPath()
      context.arc(this.circle.x, this.circle.y, 190, 0, 2 * Math.PI, true)
      context.closePath()
      context.fill()
      //内圆
      let nowRange = 74
      context.save()
      this.drawCircle(context)
      this.drawSin(this.xOffset, context, nowRange)
      this.drawText(context, nowRange)
      context.restore()
      for (let i = 0; i < 6; i++) {
        //绘制刻度。
        context.save()
        context.translate(this.circle.x, this.circle.y)
        context.rotate(-Math.PI / 2 + Math.PI / 6 + (i * Math.PI) / 3) //旋转坐标轴。坐标轴x的正方形从 向上开始算起
        context.beginPath()
        context.moveTo(190, 0)
        context.lineTo(200, 0)
        context.lineWidth = 4
        context.strokeStyle = '#0A122D'
        context.stroke()
        context.closePath()
        context.restore()
      }
    },
    drawCircle(ctx) {
      // 画圆 中心圆
      ctx.beginPath()
      ctx.fillStyle = '#209ADF'
      ctx.arc(this.circle.x, this.circle.y, 120, 0, 2 * Math.PI)
      ctx.fill()
      ctx.beginPath()
      ctx.arc(this.circle.x, this.circle.y, 120, 0, 2 * Math.PI)
      ctx.clip()
    },
    drawSin(xOffset, ctx, nowRange) {
      //画sin 曲线函数
      let mW = 240
      let mH = 240
      let sX = 0
      let axisLength = mW //轴长
      let waveWidth = 0.04 //波浪宽度,数越小越宽
      let waveHeight = 12 //波浪高度,数越大越高
      ctx.save()
      ctx.translate(130, 130)
      let points = [] //用于存放绘制Sin曲线的点
      ctx.beginPath()
      //在整个轴长上取点
      for (let x = sX; x < sX + axisLength; x += 20 / axisLength) {
        //此处坐标(x,y)的取点，依靠公式 “振幅高*sin(x*振幅宽 + 振幅偏移量)”
        let y = -Math.sin((sX + x) * waveWidth + xOffset)
        let dY = mH * (1 - nowRange / 100)
        points.push([x, dY, dY + y * waveHeight])
        ctx.lineTo(x, dY + y * waveHeight)
      }
      //封闭路径
      ctx.lineTo(axisLength, mH)
      ctx.lineTo(sX, mH)
      ctx.lineTo(points[0][0], points[0][1])
      ctx.fillStyle = '#2C50B1'
      ctx.fill()

      ctx.restore()
    },
    // 中心显示文字
    drawText(ctx, nowRange) {
      ctx.save()
      ctx.translate(130, 130)
      let size = 50
      ctx.font = size + 'px Microsoft Yahei'
      ctx.textAlign = 'center'
      ctx.fillStyle = '#95EFFF'
      ctx.fillText(nowRange + '%', 120, 120 - size / 2)
      ctx.restore()
      ctx.save()
      size = 25
      ctx.translate(130, 130)
      ctx.font = size + 'px Microsoft Yahei'
      ctx.textAlign = 'center'
      ctx.fillStyle = '#95EFFF'
      ctx.fillText('平均置信概率', 120, 120 + size)
      ctx.restore()
    },
    // 旋转的文字
    drawCircularText(s, string, startAngle, endAngle, n, context) {
      let radius = s.radius, // 文字环绕的中心圆半径
        angleDecrement, // 一个文字所占的角度
        angle = parseFloat(startAngle), // 文字的起始角度
        index = 0, // 文字的索引值
        character // 当前要画的文字
      let arr = string.split(':')
      context.save()
      context.fillStyle = '#fff'
      context.font = '12px 微软雅黑 '
      context.textAlign = 'center'
      context.textBaseline = 'middle'
      if (n < 2 || n === 5) {
        // 上三个不需要反转的文字
        while (index < string.length) {
          character = string.charAt(index)
          if (arr[0].indexOf(character) >= 0) {
            if (arr[0].length > 6) {
              angleDecrement = (startAngle - endAngle) / (string.length - 3)
            } else {
              angleDecrement = (startAngle - endAngle) / (string.length - 1)
            }
          } else {
            angleDecrement = (startAngle - endAngle) / (string.length + 6)
          }
          context.save()
          context.beginPath()
          context.translate(
            s.x + Math.cos(angle) * radius,
            s.y + Math.sin(angle) * radius
          )
          context.rotate(Math.PI / 2 + angle)
          context.fillText(character, 0, 0)
          angle -= angleDecrement
          index++
          context.restore()
        }
      } else {
        // 下面三个需要反转的文字
        while (index < string.length) {
          character = string.split('').reverse().join('').charAt(index) // 字符串反转
          if (arr[1].indexOf(character) >= 0) {
            angleDecrement = (startAngle - endAngle) / (string.length + 6)
          } else {
            if (arr[0].length > 6) {
              angleDecrement = (startAngle - endAngle) / (string.length - 3)
            } else {
              angleDecrement = (startAngle - endAngle) / (string.length - 1)
            }
          }
          context.save()
          context.beginPath()
          context.translate(
            s.x + Math.cos(angle) * radius,
            s.y + Math.sin(angle) * radius
          )
          context.rotate(-Math.PI / 2 + angle) // 旋转文字
          context.fillText(character, 0, 0)
          angle -= angleDecrement
          index++
          context.restore()
        }
      }
      context.restore()
    }
  },
  mounted() {
    this.centerBox = {
      width: document.querySelector('#circlePie').offsetWidth,
      height: document.querySelector('#circlePie').offsetHeight
    }
    for (let i = 0; i < 200; i++) {
      // 随机200个运动的圆点
      let x = Math.random() * this.centerBox.width // 随机的x偏移量
      let y = Math.random() * this.centerBox.height // 随机y轴偏移量
      let xa = Math.random() * 2 - 1 // x轴运动速度
      let ya = Math.random() * 2 - 1 // y轴运动速度
      this.dots.push({
        x: x,
        y: y,
        xa: xa,
        ya: ya,
        // 两个圆点之间需要连线的距离
        max: 40
      })
    }
    this.act()
    this.drawDot()
    this.resizeFn = this.$debounce(() => {
      // 通过捕获系统的onresize事件触发我们需要执行的事件
      this.centerBox = {
        width: document.querySelector('#circlePie').offsetWidth,
        height: document.querySelector('#circlePie').offsetHeight
      }
      for (let i = 1; i < 13; i++) {
        this.$refs['chart' + i].setChart()
      }
    }, 500)
    window.addEventListener('resize', this.resizeFn)
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.resizeFn)
    window.cancelAnimationFrame(this.animationFrame1)
    window.cancelAnimationFrame(this.animationFrame2)
  }
}
</script>

<style lang="less">
.dataScreen {
  height: 100%;
  width: 100%;
  padding: 14px 20px 20px;
  background: #050828;
  overflow: hidden;

  .content {
    height: 65%;

    .ivu-col {
      height: 100%;
    }

    .circlePie {
      height: 100%;
      padding: 0 0 40px;
      text-align: center;
      position: relative;

      canvas {
        position: absolute;
        left: 50%;
        top: 0;
        transform: translate(-50%, 0);
      }

      #dot {
        background: rgba(0, 0, 0, 0);
      }
    }

    .list {
      height: 100%;
      .left {
        .left-info {
          margin-top: 10px;
          height: 440px;
          overflow-y: auto;
          overflow-x: hidden;

          .el-col-24 {
            margin-left: 20px;
            padding-left: 20px;
            padding-right: 20px;
            width: calc(100% - 40px);
          }
          .el-card {
            font-size: 15px;
            color: white;
            background-color: #363b79;
            border: none;
            border-radius: 20px;
            margin-top: 20px;

            .el-card__body {
              padding: 0;
              padding-left: 20px;
              padding-bottom: 10px;
              padding-top: 5px;
            }
          }
          .el-card:hover {
            box-shadow: 0 7px 15px 0 rgba(49, 196, 190, 0.75);
          }
        }
        .left-info::-webkit-scrollbar {
          width: 10px;
          height: 10px;
          /**/
        }
        .left-info::-webkit-scrollbar-track {
          background: rgb(239, 239, 239);
          border-radius: 2px;
        }
        .left-info::-webkit-scrollbar-thumb {
          background: #bfbfbf;
          border-radius: 10px;
        }
        .left-info::-webkit-scrollbar-thumb:hover {
          background: #333;
        }
        .left-info::-webkit-scrollbar-corner {
          background: #179a16;
        }
      }

      .right {
        .right-info {
          margin-top: 10px;
          height: 440px;
          overflow-y: auto;
          overflow-x: hidden;

          .el-col-24 {
            margin-left: 20px;
            padding-left: 20px;
            padding-right: 20px;
            width: calc(100% - 40px);
          }
          .el-card {
            font-size: 15px;
            color: white;
            background-color: #363b79;
            border: none;
            border-radius: 20px;
            margin-top: 20px;

            .el-card__body {
              padding: 0;
              padding-left: 20px;
              padding-bottom: 10px;
              padding-top: 5px;
            }
          }
          .el-card:hover {
            box-shadow: 0 7px 15px 0 rgba(49, 196, 190, 0.75);
          }
        }
        .right-info::-webkit-scrollbar {
          width: 10px;
          height: 10px;
          /**/
        }
        .right-info::-webkit-scrollbar-track {
          background: rgb(239, 239, 239);
          border-radius: 2px;
        }
        .right-info::-webkit-scrollbar-thumb {
          background: #bfbfbf;
          border-radius: 10px;
        }
        .right-info::-webkit-scrollbar-thumb:hover {
          background: #333;
        }
        .right-info::-webkit-scrollbar-corner {
          background: #179a16;
        }
      }

      .left,
      .right {
        background: #0d1341;
      }

      .left,
      .right,
      .center {
        width: 100%;
        height: 90%;
        border: 1px solid #0d2451;
        position: relative;

        #left1,
        #left2,
        #left3,
        #right1,
        #right2,
        #right3,
        #center2 {
          height: 100%;
        }

        .chart-68 {
          width: 68%;
          height: 100%;
          float: left;
        }

        .chart-32 {
          width: 32%;
          height: 100%;
          float: left;
        }
      }
    }
  }

  .bottom-list {
    height: 35%;

    .ivu-col {
      height: 100%;

      .list {
        height: 100%;
        width: 33.3333%;
        padding-right: 1.5%;
        float: left;
        margin-left: 400px;

        #bottom_4 {
          margin-left: 60px;
          padding: 0 20px;
        }

        .bottom {
          width: 100%;
          height: 100%;
          border: 1px solid #0d2451;
          position: relative;
        }
      }

      .right-bottom {
        margin-left: -200px;
        width: 397px;

        .bottom1 {
          width: 100%;
        }
      }
    }
  }
}
</style>
