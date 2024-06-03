<template>
  <div class="pieChart"></div>
</template>

<script>
export default {
  name: '',
  data() {
    return {
      myChart: null
    }
  },
  methods: {
    setChart() {
      let option = {
        title: [],
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        legend: {
          data: [
            { name: '微博', icon: 'circle' },
            { name: 'b站', icon: 'circle' },
            { name: '今日头条', icon: 'circle' },
            { name: '抖音', icon: 'circle' }
          ],
          left: '8%',
          top: '10%',
          itemWidth: 7,
          itemHeight: 7,
          textStyle: {
            color: '#00CCFF',
            fontSize: 10
          }
        },
        series: [
          {
            name: '谣言占比',
            type: 'pie',
            radius: '40%',
            center: ['17%', '60%'],
            data: [
              { value: 335, name: '微博' },
              { value: 310, name: 'b站' },
              { value: 234, name: '今日头条' },
              { value: 135, name: '抖音' }
            ],
            label: {
              fontSize: 8,
              color: '#00CCFF'
            },
            labelLine: {
              length: 15,
              length2: 10,
              lineStyle: {
                color: '#3F3F5C'
              }
            },
            itemStyle: {
              color: function (params) {
                var colorList = [
                  '#F74F64',
                  '#00CCFF',
                  '#315371',
                  '#142AFE',
                  '#9814FE'
                ]
                return colorList[params.dataIndex]
              }
            },
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ]
      }
      if (!this.myChart) this.myChart = this.$echarts(this.$el)

      this.myChart.clear()
      this.myChart.resize()
      this.myChart.setOption(option)
    }
  },
  mounted() {
    this.setChart()
  }
}
</script>

<style lang="less" scoped>
.pieChart {
  height: 100%;
  width: 100%;
  padding: 0 20px;
}
</style>
