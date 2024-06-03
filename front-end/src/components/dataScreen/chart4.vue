<template>
  <div class="doubleLine"></div>
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
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          left: '11%',
          top: '10%',
          itemWidth: 7,
          itemHeight: 7,
          textStyle: {
            color: '#5CB1C1',
            fontSize: 10
          }
        },
        grid: {
          top: '12%',
          left: '10%',
          right: '10%',
          bottom: '10%',
          containLabel: false
        },

        xAxis: {
          type: 'category',
          boundaryGap: false,
          axisLine: {
            symbol: ['none', 'arrow'],
            symbolSize: [6, 6],
            symbolOffset: [0, 10],
            lineStyle: {
              color: '#122C49'
            }
          },
          axisTick: { show: false },
          axisLabel: {
            color: '#61B9C8',
            fontSize: 9
          },
          data: [
            '2023.11',
            '2023.12',
            '2024.1',
            '2024.2',
            '2024.3',
            '2024.4',
            '2024.5'
          ]
        },
        yAxis: [
          {
            type: 'value',
            scale: true,
            max: 400,
            min: 0,
            interval: 50,
            axisLine: {
              symbol: ['none', 'arrow'],
              symbolSize: [6, 6],
              lineStyle: {
                color: '#122C49'
              }
            },
            axisLabel: {
              color: '#61B9C8',
              showMaxLabel: false,
              fontSize: 9
            },
            name: '(条)',
            nameGap: -10,
            nameTextStyle: {
              color: '#61B9C8',
              fontSize: 9,
              align: 'right',
              padding: [0, 6, 0, 0]
            },
            splitLine: {
              show: false
            }
          },
          {
            type: 'value',
            scale: true,
            max: 400,
            min: 0,
            axisLine: {
              symbol: ['none', 'arrow'],
              symbolSize: [6, 6],
              lineStyle: {
                color: '#122C49'
              }
            },
            axisLabel: {
              color: '#61B9C8',
              showMaxLabel: false,
              fontSize: 9
            },
            name: '(条)',
            nameGap: -10,
            nameTextStyle: {
              color: '#61B9C8',
              fontSize: 9,
              align: 'left',
              padding: [0, 0, 0, 6]
            },
            interval: 50,
            splitLine: {
              show: false
            }
          }
        ],
        series: [
          {
            name: '微博',
            type: 'line',
            smooth: true,
            symbol: 'none',
            lineStyle: {
              color: '#F39800'
            },
            itemStyle: {
              color: '#F39800'
            },
            data: [50, 132, 100, 300, 90, 230, 210]
          },
          {
            name: 'b站',
            yAxisIndex: 1,
            type: 'line',
            smooth: true,
            symbol: 'none',
            lineStyle: {
              color: '#BF232A'
            },
            itemStyle: {
              color: '#BF232A'
            },
            data: [330, 310, 132, 100, 300, 90, 230]
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
.doubleLine {
  height: 100%;
  width: 100%;
}
</style>
