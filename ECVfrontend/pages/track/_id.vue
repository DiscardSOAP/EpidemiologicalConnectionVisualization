<template>
  <div style="width:100%;height:100%;">
    <h1 style="width:100%;text-align:center;font-weight:normal">
      {{nameMap[$route.params.id]}}病例行动轨迹跟踪
      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn
            icon
            color="primary"
            :to="`/topology/${$route.params.id}/`"
            v-on="on"
          >
            <v-icon x-large>mdi-share-variant</v-icon>
          </v-btn>
        </template>
        <span>切换至行动轨迹图</span>
      </v-tooltip>

    </h1>
    <div
      id="demo"
      style="width:100%;height:90%;position:relative"
    >
      <el-amap
        :amap-manager="amapManager"
        :events="mapEvents"
        class="amap-demo"
        style="width:100%;height:100%"
      >
      </el-amap>
      <v-date-picker
        style="position:absolute;right:0%;top:0%;z-index:999;"
        v-model="date"
        no-title
        reactive
        :show-current="false"
        width="200"
      ></v-date-picker>
      <v-speed-dial
        v-model="fab"
        bottom
        right
        absolute
        direction="left"
        open-on-hover
      >
        <template v-slot:activator>
          <v-btn
            color="blue darken-2"
            :disabled="fab"
            dark
            x-large
            fab
          >
            <v-icon v-if="fab">mdi-calendar-multiple</v-icon>
            <v-icon v-else>mdi-calendar</v-icon>
          </v-btn>
        </template>
        <v-btn
          fab
          dark
          color="purple"
          @click.stop="toNextDay()"
        >
          <v-icon>mdi-chevron-double-right</v-icon>
        </v-btn>
        <v-btn
          fab
          dark
          color="indigo"
          @click.stop="toPreDay()"
        >
          <v-icon style="transform: rotateX(180deg);">mdi-chevron-double-left</v-icon>
        </v-btn>
      </v-speed-dial>
    </div>
  </div>
</template>

<script>
import { AMapManager } from 'vue-amap';
import moment from "moment"
import VueAMap from 'vue-amap'
let amapManager = new VueAMap.AMapManager();
import * as _ from 'lodash';

const nameMap = {
  'tianjin': '天津',
  'beijing': '北京',
  'xinjiang': '新疆',
  'heilongjiang': '黑龙江',
  'jilin': '吉林',
  'liaoning': '辽宁',
  'neimenggu': '内蒙古',
  'hebei': '河北',
  'shanxi': '山西',
  'shaanxi': '陕西',
  'ningxia': '宁夏',
  'gansu': '甘肃',
  'qinghai': '青海',
  'xizang': '西藏',
  'sichuan': '四川',
  'yunnan': '云南',
  'guangxi': '广西',
  'guizhou': '贵州',
  'chongqing': '重庆',
  'henan': '河南',
  'shandong': '山东',
  'hubei': '湖北',
  'anhui': '安徽',
  'jiangsu': '江苏',
  'hunan': '湖南',
  'jiangxi': '江西',
  'zhejiang': '浙江',
  'guangdong': '广东',
  'fujian': '福建',
  'aomen': '澳门',
  'xianggang': '香港',
  'taiwan': '台湾',
  'shanghai': '上海',
  'nanhaizhudao': '南海诸岛',
  'hainan': '海南'
}


export default {
  auth: false,
  data () {
    let self = this
    return {
      nameMap: {
        'tianjin': '天津',
        'hainan': '海南',
        'beijing': '北京',
        'guangdong': '广东',
        'xinjiang': '新疆',
        'heilongjiang': '黑龙江',
        'jilin': '吉林',
        'liaoning': '辽宁',
        'neimenggu': '内蒙古',
        'hebei': '河北',
        'shanxi': '山西',
        'shaanxi': '陕西',
        'ningxia': '宁夏',
        'gansu': '甘肃',
        'qinghai': '青海',
        'xizang': '西藏',
        'sichuan': '四川',
        'yunnan': '云南',
        'guangxi': '广西',
        'guizhou': '贵州',
        'chongqing': '重庆',
        'henan': '河南',
        'shandong': '山东',
        'hubei': '湖北',
        'anhui': '安徽',
        'jiangsu': '江苏',
        'hunan': '湖南',
        'jiangxi': '江西',
        'zhejiang': '浙江',
        'fujian': '福建',
        'aomen': '澳门',
        'xianggang': '香港',
        'taiwan': '台湾',
        'shanghai': '上海',
        'nanhaizhudao': '南海诸岛',
      },
      date: new Date().toISOString().substr(0, 10),
      amapManager,
      map: null,
      routes: [],
      humans: [],
      readyNum: 0,
      ready: false,
      fab: false,

      mapEvents: {
        init (o) {
          self.map = o
          o.setMapStyle('amap://styles/dark')
          o.plugin(["AMap.Driving"], function () {
            for (let idx in self.news) {
              let patient = self.news[idx]
              //console.log("patient:", patient)
              let keyAddrs = []
              let route = []
              for (let event of patient.events)
                keyAddrs.push({
                  keyword: event.addr,
                  city: nameMap[self.$route.params.id]
                })
              //console.log("key addrs:", keyAddrs)
              let searchedPathNum = 0
              for (let i = 0; i < keyAddrs.length - 1; i++) {
                let stAddr = keyAddrs[i]
                let edAddr = keyAddrs[i + 1]
                var driving = new AMap.Driving({})
                //console.log(_.cloneDeep([stAddr, edAddr]))
                driving.search(_.cloneDeep([stAddr, edAddr]), function (status, result) {
                  //console.log(result)
                  route.push({
                    R: result.origin.R,
                    Q: result.origin.Q,
                    pathID: i,
                    pathIndex: 0
                  })
                  let steps = result.routes[0].steps
                  for (let j = 0; j < steps.length; j++) {
                    route.push({
                      R: steps[j].start_location.R,
                      Q: steps[j].start_location.Q,
                      pathID: i,
                      pathIndex: j + 1
                    })
                  }
                  route.push({
                    R: steps[steps.length - 1].end_location.R,
                    Q: steps[steps.length - 1].end_location.Q,
                    pathID: i,
                    pathIndex: steps.length + 1
                  })
                  route.push({
                    R: result.destination.R,
                    Q: result.destination.Q,
                    pathID: i,
                    pathIndex: steps.length + 2
                  })
                  searchedPathNum++
                  if (searchedPathNum == keyAddrs.length - 1) {
                    route.sort((a, b) => {
                      if (a.pathID < b.pathID || (a.pathID == b.pathID && a.pathIndex < b.pathIndex))
                        return -1;
                      else if (a.pathID == b.pathID && a.pathIndex == b.pathIndex)
                        return 0;
                      else
                        return 1;
                    })
                    //console.log(route)
                    let patientRoute = []
                    let patientTrack = []
                    route.forEach(x => {
                      patientRoute.push([x.R, x.Q])
                      patientTrack.push({
                        pos: [x.R, x.Q],
                        date: patient.events[x.pathID].endDate
                      })
                    })
                    //console.log(patientRoute)
                    self.routes.push(patientTrack)
                    //console.log("???")
                    let human = new AMap.Marker({
                      map: o,
                      position: _.cloneDeep(patientRoute[0]),
                      content: `<div style="width:25px;height:34px;background-image:url(https://i.loli.net/2020/05/28/HaQerMxyFq2KnsI.png);background-repeat:no-repeat;background-size:25px 34px;text-align:center;opacity: 0.8;"> <div>${patient.id}</div> </div>`
                    })
                    self.humans.push(human)
                    AMap.event.addListener(human, 'click', function name () {
                      const { href } = self.$router.resolve({
                        path: `/news/${patient.id}/`
                      })
                      window.open(href, '_blank')
                    })
                    self.readyNum++
                    //human.moveAlong(path, 600)
                  }
                })
                o.addControl(driving)
              }
            }
          })
          //console.log(self.routes)
          o.setFitView()
        }
      }
    }
  },
  asyncData (context) {
    return context.app.$axios.post(`/api/track/`, {
      category: nameMap[context.route.params.id]
    }).then(
      res => {
        //console.log(res.data)
        return res.data
      },
      () => {
        context.error({ statusCode: 400, message: "error" })
        return {}
      }
    )
  },
  mounted () {
    //this.track()
  },
  watch: {
    "readyNum": function (newVal) {
      if (newVal == this.news.length) {
        this.ready = true
      }
    },
    "ready": function (newVal) {
      this.track()
    },
    "date": function (newDate) {
      this.track()
    }
  },
  methods: {
    toNextDay: function () {
      this.date = moment(this.date).add(1, 'days').format('YYYY-MM-DD')
    },
    toPreDay: function () {
      this.date = moment(this.date).subtract(1, 'days').format('YYYY-MM-DD')
    },
    track: function () {
      for (let id in this.humans) {
        let human = this.humans[id]
        human.stopMove()
        let route = this.routes[id]
        let path = []
        let pos = route[0].pos
        let minR = 1e9, minQ = 1e9, maxR = -1e9, maxQ = -1e9
        route.forEach(x => {
          if (x.date == this.date) {
            path.push(x.pos)
            minR = Math.min(minR, x.pos[0])
            minQ = Math.min(minQ, x.pos[1])
            maxR = Math.max(maxR, x.pos[0])
            maxQ = Math.max(maxQ, x.pos[1])
          }
          if (moment(x.date) < moment(this.date))
            pos = x.pos
        })
        console.log(route)
        console.log(path)
        console.log(pos)
        human.setPosition(_.cloneDeep(pos))
        if (path.length > 0)
          human.moveAlong(_.cloneDeep(path), Math.max(maxR - minR, maxQ - minQ) * 60000)
      }
    }
  }
}
</script>
