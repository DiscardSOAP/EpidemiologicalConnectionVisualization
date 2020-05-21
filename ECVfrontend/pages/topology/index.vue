<template>
  <div
    class="echarts"
    style="height:100%;width:100%"
  >
    <div
      style="height:100%;width:100%"
      ref="myEchart"
    ></div>
  </div>
</template>
<script>
import echarts from "echarts";
import '../../node_modules/echarts/map/js/china.js' // 引入中国地图数据
export default {
  auth: false,
  data () {
    return {
      chart: null,
      data: {},
      maxNum: 0,
      minNum: 0
    };
  },
  mounted () {
    this.chinaConfigure();
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
    if (!this.chart) {
      return;
    }
    this.chart.dispose();
    this.chart = null;
  },
  methods: {
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
        '南海诸岛': 'nanhaizhudao'
      }

      console.log(this.data)
      let data = []
      for (let city in this.data) {
        data.push({
          "name": city,
          "value": this.data[city],
        })
      }
      console.log(data)
      let myChart = echarts.init(this.$refs.myEchart); //这里是为了获得容器所在位置    
      window.onresize = myChart.resize;
      myChart.setOption({ // 进行相关配置
        //backgroundColor: "#02AFDB",
        tooltip: {}, // 鼠标移到图里面的浮动提示框
        title: {
          text: '中国疫情地图',
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
        toolbox: {
        },
        /*dataRange: {
          show: false,
          min: 0,
          max: 1000,
          text: ['High', 'Low'],
          realtime: false,
          calculable: true,
          color: ['orangered', 'yellow', 'lightskyblue']
        },*/
        visualMap: {
          min: this.minNum,
          max: this.maxNum,
          text: ['High', 'Low'],
          realtime: false,
          calculable: true,
          inRange: {
            color: ['lightskyblue', 'yellow', 'orangered']
          },
          textStyle: {
            color: "white"
          },
        },
        geo: { // 这个是重点配置区
          map: 'china', // 表示中国地图
          roam: true,
          label: {
            normal: {
              show: true, // 是否显示对应地名
              textStyle: {
                color: 'rgb(127,127,127)'
              }
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
        series: [{
          type: 'scatter',
          coordinateSystem: 'geo' // 对应上方配置
        },
        {
          name: '确诊人数', // 浮动框的标题
          type: 'map',
          geoIndex: 0,
          data: data
        }
        ]
      })
      var _this = this;
      myChart.on('click', function (params) {
        console.log(nameMap[params.name])
        _this.$router.push(`/topology/${nameMap[params.name]}/`)

      });
    }
  }
}
</script>