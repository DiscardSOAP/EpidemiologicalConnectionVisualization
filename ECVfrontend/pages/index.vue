<template>
  <div style="width:100%">
    <div
      style="height:600px;width:100%"
      ref="myEchart"
    ></div>
    <div
      style="height:600px;width:100%"
      ref="myEchart2"
    ></div>
  </div>
</template>
<script>
import echarts from "echarts"
import '../node_modules/echarts/map/js/china.js'
export default {
  auth: false,
  data () {
    return {
      chart: null,
      chart2: null,
      data2: {},
      data: {},
      maxNum: 0,
      minNum: 0,
      category: '',
    }
  },
  mounted () {
    this.chinaConfigure()
    this.chartConfigure()
  },
  asyncData (context) {
    return context.app.$axios
      .get(`/api/topology/`)
      .then(
        res => {
          let data = res.data
          console.log(res.data)
          let maxNum = 0
          for (let city in data) {
            maxNum = Math.max(maxNum, data[city])
          }
          if (maxNum == 0)
            maxNum = 1
          //console.log(users)
          return {
            data: data,
            maxNum: maxNum,
          }
        },
        () => {
          context.error({ statusCode: 400, message: "error!" })
          return {}
        }
      )
  },
  beforeDestroy () {
    if (this.chart) {
      this.chart.dispose()
      this.chart = null
    }
    if (this.chart2) {
      this.chart2.dispose()
      this.chart2 = null
    }
  },
  methods: {
    chartConfigure () {
      if (this.chart2) {
        this.chart2.dispose()
        this.chart2 = null
      }
      var xAxisData = []
      var activeNumbers = []
      var lockedNumbers = []
      for (var i = 0; i < 100; i++) {
        xAxisData.push('' + i);
        activeNumbers.push((Math.sin(i / 5) * (i / 5 - 10) + i / 6) * 5);
        lockedNumbers.push((Math.sin(i / 5) * (i / 5 - 10) + i / 3) * 2);
      }
      let option = {
        title: {
          text: this.category ? this.category : '全国' + '病例变化',
          textStyle: {
            color: "white",
            fontSize: 20
          },
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'shadow',
            label: {
              show: true
            }
          }
        },
        toolbox: {
          show: true,
          feature: {
            mark: { show: true },
            dataView: { show: true, readOnly: false },
            magicType: { show: true, type: ['line', 'bar'] },
            restore: { show: true },
            saveAsImage: { show: true }
          },
        },
        calculable: true,
        legend: {
          data: ['游荡病例', '隔离病例'],
          itemGap: 5,
          textStyle: {
            color: "white",
          },
        },
        grid: {
          top: '12%',
          left: '1%',
          right: '10%',
          containLabel: true
        },
        xAxis: [
          {
            name: '日期',
            nameTextStyle: {
              color: "white"
            },
            type: 'category',
            data: xAxisData,
            axisLabel: {
              color: "white"
            }
          }
        ],
        yAxis: [
          {
            type: 'value',
            name: '人数',
            nameTextStyle: {
              color: "white"
            },
            axisLabel: {
              formatter: function (a) {
                a = +a;
                return isFinite(a)
                  ? echarts.format.addCommas(+a / 1000)
                  : '';
              },
              color: "white"
            }
          }
        ],
        dataZoom: [
          {
            show: true,
            start: 94,
            end: 100
          },
          {
            type: 'inside',
            start: 94,
            end: 100
          },
          {
            show: true,
            yAxisIndex: 0,
            filterMode: 'empty',
            width: 30,
            height: '80%',
            showDataShadow: false,
            left: '95%'
          }
        ],
        series: [
          {
            name: '隔离病例',
            type: 'bar',
            stack: '感染病例',
            data: lockedNumbers
          },
          {
            name: '游荡病例',
            type: 'bar',
            stack: '感染病例',
            data: activeNumbers
          },

        ]
      };
      this.chart2 = echarts.init(this.$refs.myEchart2)
      window.onresize = this.chart2.resize
      this.chart2.setOption(option)
    },
    chinaConfigure () {
      let nameMap = {
        '天津': 'tianjin',
        '北京': 'beijing',
        '新疆': 'xinjiang',
        '黑龙江': 'heilongjiang',
        '吉林': 'jilin',
        '辽宁': 'liaoning',
        '内蒙古': 'neimenggu',
        '河北': 'hebei',
        '山西': 'shanxi',
        '陕西': 'shaanxi',
        '宁夏': 'ningxia',
        '甘肃': 'gansu',
        '青海': 'qinghai',
        '西藏': 'xizang',
        '四川': 'sichuan',
        '云南': 'yunnan',
        '广西': 'guangxi',
        '贵州': 'guizhou',
        '重庆': 'chongqing',
        '河南': 'henan',
        '山东': 'shandong',
        '湖北': 'hubei',
        '安徽': 'anhui',
        '江苏': 'jiangsu',
        '湖南': 'hunan',
        '江西': 'jiangxi',
        '浙江': 'zhejiang',
        '广东': 'guangdong',
        '福建': 'fujian',
        '澳门': 'aomen',
        '香港': 'xianggang',
        '台湾': 'taiwan',
        '上海': 'shanghai',
        '南海诸岛': 'nanhaizhudao',
        "海南": "hainan"
      }

      console.log(this.data)
      let data = []
      for (let city in this.data) {
        data.push({
          "name": city,
          "value": this.data[city],
        })
      }
      data.push({
        "name": "贵州",
        "value": 100,
      })
      console.log(data)
      this.chart = echarts.init(this.$refs.myEchart) //这里是为了获得容器所在位置    
      window.onresize = this.chart.resize
      var convertData = function (data) {
        var res = [];
        for (var i = 0; i < data.length; i++) {
          var geoCoord = geoCoordMap[data[i].name];
          if (geoCoord) {
            res.push({
              name: data[i].name,
              value: geoCoord.concat(data[i].value)
            });
          }
        }
        return res;
      };
      var geoCoordMap = {
        '西藏': [91.11, 29.97],
        '上海': [121.48, 31.22],
        '福建': [118.3, 26.08],
        '广东': [113.23, 23.16],
        '广西': [108.83, 23.56],
        '海南': [110.05, 19.02],
        '辽宁': [123.38, 40.8],
        '吉林': [125.95, 43.48],
        '江西': [115.69, 27.68],
        '内蒙古': [111.65, 42.82],
        '四川': [102.36, 30.67],
        '陕西': [108.95, 34.27],
        '重庆': [107.54, 29.59],
        '江苏': [119.98, 33.34],
        '贵州': [106.71, 26.57],
        '北京': [116.46, 40.00],
        '新疆': [87.68, 41.77],
        '山东': [118, 36.25],
        '天津': [117.5, 39.13],
        '云南': [101.71, 24.57],
        '黑龙江': [127.63, 46.75],
        '河北': [115.48, 38.53],
        '湖南': [112, 27.71],
        '安徽': [117.27, 31.86],
        '湖北': [113.01, 30.92],
        '浙江': [120.01, 29.32],
        '河南': [113.95, 33.57],
        '青海': [97.11, 35.97],
        '甘肃': [102.1, 38.57],
        '宁夏': [106.1, 37.02],
        '山西': [112.55, 37.57],
        '台湾': [121.1, 24.02],
        '南海诸岛': [127.1, 24.42],
        '香港': [114.13, 22.36],
        '澳门': [113.63, 22.12],
      };

      this.chart.setOption({ // 进行相关配置
        title: {
          text: '中国疫情分布地图',
          subtext: '数据录入来自DiscarpSOAP',
          textStyle: {
            color: "white",
            fontSize: 30
          },
          subtextStyle: {
            fontSize: 15
          },
          x: "center"
        },
        tooltip: {
          trigger: 'item'
        },
        geo: {
          roam: true,
          zoom: 1.2,
          map: 'china',
          label: {
            normal: {
              show: true, // 是否显示对应地名
              textStyle: {
                color: 'rgba(127,127,127)'
              }
            },
            emphasis: {
              color: 'rgba(0,0,0,0)'
            }
          },
          itemStyle: {
            normal: {
              borderColor: 'rgba(0, 0, 0, 0.2)'
            },
            emphasis: {
              color: 'rgba(255,255,255,1)',
              areaColor: null,
              shadowOffsetX: 0,
              shadowOffsetY: 0,
              shadowBlur: 20,
              borderWidth: 0,
              shadowColor: 'rgba(0, 0, 0, 0.5)'
            }
          }

        },
        series: [
          {
            name: 'pm2.5',
            type: 'scatter',
            coordinateSystem: 'geo',
            data: convertData(data),
            symbolSize: function (val) {
              return val[2] / 10;
            },
            encode: {
              value: 2
            },
            label: {
              formatter: '{b}',
              position: 'right',
              show: false
            },
            itemStyle: {
              color: '#F44'
            },
            emphasis: {
              label: {
                show: false
              },
            }
          },
          {
            name: 'Top 5',
            type: 'effectScatter',
            coordinateSystem: 'geo',
            data: convertData(data.sort(function (a, b) {
              return b.value - a.value;
            }).slice(0, 6)),
            symbolSize: function (val) {
              return val[2] / 10;
            },
            encode: {
              value: 2
            },
            showEffectOn: 'render',
            rippleEffect: {
              brushType: 'stroke'
            },
            hoverAnimation: true,
            itemStyle: {
              color: 'red',
              shadowBlur: 10,
              shadowColor: '#333'
            },
            zlevel: 1
          }
        ]
      })
      var _this = this;
      this.chart.on('click', function (params) {
        _this.category = params.name
        _this.chartConfigure()
      });
    }
  }
}
</script>
