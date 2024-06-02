<template>
  <div style="height: 100%">
    <div class="header">
      <Header />
    </div>
    <div class="page">
      <DataScreenChartVue :selectRangeDate="selectRangeDate" />
    </div>
  </div>
</template>

<script>
import DataScreenChartVue from './DataScreenChart.vue'
import Header from '@/components/Header.vue'

export default {
  name: 'DataScreen',
  data() {
    return {
      activeName: 'month', // 默认显示近一月
      modal: false,
      flag: true,
      selectRangeDate: [],
      startTime: '',
      endTime: '',
      optionStart: {
        disabledDate(date) {
          // 禁止选择今天之后的日期
          return date && date.valueOf() > Date.now() - 86400000
        }
      },
      optionEnd: {},
      resizeFn: null
    }
  },
  components: {
    DataScreenChartVue,
    Header
  },
  mounted() {
    window.addEventListener('resize', this.resizeFn)
    this.handleSelect(this.activeName) // 默认显示近一个月
  },
  methods: {
    pickStartDate(date) {
      // 选择开始时间的回调
      this.startTime = date
      this.optionEnd = {
        disabledDate(d) {
          // 禁止选择开始时间之前的日期
          return d && d.valueOf() < new Date(date).valueOf() - 86400000
        }
      }
    },
    pickEndDate(date) {
      // 选择结束时间的回调
      this.endTime = date
    },
    getMonthBetween(start, end) {
      // 获取开始时间和结束时间之内的所有月份
      this.selectRangeDate = []
      let s = start.split('-') // 字符串装换数组
      let e = end.split('-')
      let date = new Date()
      let min = date.setFullYear(s[0], s[1] - 1) // 设置最小时间
      let max = date.setFullYear(e[0], e[1] - 1) // 设置最大时间
      let curr = min
      while (curr <= max) {
        // 循环添加月份
        var month = curr.getMonth()
        var arr = [curr.getFullYear(), month + 1]
        this.selectRangeDate.push(arr)
        curr.setMonth(month + 1)
      }
    },
    getDays(day) {
      // 获取天数
      let arr = []
      for (let i = -day; i < 0; i++) {
        // 循环添加天数
        let today = new Date() // 获取今天
        let targetday_milliseconds = today.getTime() + 1000 * 60 * 60 * 24 * i
        today.setTime(targetday_milliseconds) //设置i天前的时间
        let tYear = today.getFullYear()
        let tMonth = today.getMonth()
        let tDate = today.getDate()
        let date = [tYear, tMonth + 1, tDate]
        arr.push(date)
      }
      return arr
    },
    handleSelect(name) {
      switch (name) {
        case 'day':
          break
        case 'week':
          this.selectRangeDate = this.getDays(7) // 获取近一周的天数
          this.flag = true

          break
        case 'month':
          this.selectRangeDate = this.getDays(30) // 获取近一个月的天数
          this.flag = true
          break
        case 'filter':
          this.modal = true
          break
        default:
          break
      }
    }
  }
}
</script>

<style lang="less" scoped>
.header {
  height: 106px;
  background: #050828;
  justify-content: space-between;
  align-items: center;
}
.page {
  height: calc(~'100% - 106px');
}
</style>
