<template>
  <div
    id="demo"
    style="width:100%;height:100%;position:relative"
  >
    <el-amap
      :amap-manager="amapManager"
      :events="mapEvents"
      class="amap-demo"
      style="width:100%;height:100%"
    >
    </el-amap>
    <!--div style="background:red;height:100%;width:100%;position:absolute;top:0%;left:0%;">
    </div-->
    <v-date-picker
      style="position:absolute;right:0%;top:0%;z-index:999;"
      v-model="date"
      no-title
      reactive
      width="200"
      readonly
    ></v-date-picker>
    <v-speed-dial
      v-model="fab"
      bottom
      :show-current="false"
      right
      absolute
      direction="left"
      open-on-hover
    >
      <template v-slot:activator>
        <v-btn
          v-model="fab"
          color="blue darken-2"
          :disabled="fab"
          dark
          x-large
          fab
        >
          <v-icon v-if="fab">mdi-close</v-icon>
          <v-icon v-else>mdi-play-box-outline</v-icon>
        </v-btn>
      </template>
      <v-btn
        fab
        dark
        color="purple"
        @click="human.stopMove()"
      >
        <v-icon>mdi-stop</v-icon>
      </v-btn>
      <v-btn
        fab
        dark
        color="indigo"
        @click="human.resumeMove()"
      >
        <v-icon>mdi-play</v-icon>
      </v-btn>
      <v-btn
        fab
        dark
        color="red"
        @click="human.pauseMove()"
      >
        <v-icon>mdi-pause</v-icon>
      </v-btn>
      <v-btn
        fab
        dark
        color="green"
        @click="human.moveAlong(path,velocity)"
      >
        <v-icon>mdi-google-play</v-icon>
      </v-btn>
    </v-speed-dial>
  </div>
</template>

<script>
import { AMapManager } from 'vue-amap';
import moment, { max, min } from "moment"
import VueAMap from 'vue-amap'
let amapManager = new VueAMap.AMapManager();
import * as _ from 'lodash';
export default {
  auth: false,
  data () {
    let self = this
    return {
      date: new Date().toISOString().substr(0, 10),
      map: null,
      interval: null,
      amapManager,
      human: null,
      path: [],
      fab: false,
      velocity: 600,
      mapEvents: {
        init (o) {
          console.log(self.date)
          self.map = o
          o.setMapStyle('amap://styles/dark');
          o.plugin(["AMap.Driving"], function () {
            let points = []
            for (let event of self.events) {
              points.push({
                keyword: event.addr,
                city: self.category
              })
            }
            //console.log(points)
            let finish_num = 0
            let routes = []
            let mks = []
            for (let i = 0; i < points.length - 1; i++) {
              //while (finish_num != i);
              //console.log(points.length - 1)
              let st = points[i]
              let ed = points[i + 1]
              var driving = new AMap.Driving({})
              //console.log(_.cloneDeep([st, ed]))
              driving.search(_.cloneDeep([st, ed]), function (status, result) {
                console.log(result)
                finish_num += 1
                routes.push({
                  R: result.origin.R,
                  Q: result.origin.Q,
                  id: i,
                  idx: 0
                })
                //console.log(i, [result.origin.R, result.origin.Q])
                let steps = result.routes[0].steps
                for (let j = 0; j < steps.length; j++) {
                  routes.push({
                    R: steps[j].start_location.R,
                    Q: steps[j].start_location.Q,
                    id: i,
                    idx: j + 1
                  })
                  //console.log(i, [steps[j].start_location.R, steps[j].start_location.Q])
                }
                routes.push({
                  R: steps[steps.length - 1].end_location.R,
                  Q: steps[steps.length - 1].end_location.Q,
                  id: i,
                  idx: steps.length + 1
                })
                //console.log(i, [steps[steps.length - 1].end_location.R, steps[steps.length - 1].end_location.Q])
                routes.push({
                  R: result.destination.R,
                  Q: result.destination.Q,
                  id: i,
                  idx: steps.length + 2
                })

                mks.push({
                  R: result.origin.R,
                  Q: result.origin.Q,
                  id: i
                })
                if (i == points.length - 2)
                  mks.push({
                    R: result.destination.R,
                    Q: result.destination.Q,
                    id: points.length - 1
                  })

                if (finish_num == points.length - 1) {
                  routes.sort((a, b) => {
                    if (a.id < b.id || (a.id == b.id && a.idx < b.idx))
                      return -1;
                    else if (a.id == b.id && a.idx == b.idx)
                      return 0;
                    else
                      return 1;
                  })
                  let path = []
                  let minR = 1e9, maxR = -1e9, minQ = 1e9, maxQ = -1e9
                  routes.forEach(x => {
                    path.push([x.R, x.Q])
                    minR = Math.min(minR, x.R)
                    maxR = Math.max(maxR, x.R)
                    minQ = Math.min(minQ, x.Q)
                    maxQ = Math.max(maxQ, x.Q)
                  })
                  self.velocity = Math.max(maxR - minR, maxQ - minQ) * 60000
                  console.log(Math.max(maxR - minR, maxQ - minQ))
                  console.log(path)
                  let humanContent = `<div style="width:25px;height:34px;background-image:url(https://i.loli.net/2020/05/28/HaQerMxyFq2KnsI.png);background-repeat:no-repeat;background-size:25px 34px;text-align:center;opacity: 0.8;"> </div>`

                  var polyline = new AMap.Polyline({
                    map: o,
                    path: _.cloneDeep(path),
                    showDir: true,
                    strokeColor: "#28F",  //线颜色
                    // strokeOpacity: 1,     //线透明度
                    strokeWeight: 6,      //线宽
                    // strokeStyle: "solid"  //线样式
                  });

                  var passedPolyline = new AMap.Polyline({
                    map: o,
                    // path: lineArr,
                    strokeColor: "#AF5",  //线颜色
                    // strokeOpacity: 1,     //线透明度
                    strokeWeight: 6,      //线宽
                    // strokeStyle: "solid"  //线样式
                  });
                  self.path = path

                  mks.sort((a, b) => {
                    if (a.id < b.id)
                      return -1;
                    else if (a.id == b.id)
                      return 0;
                    else
                      return 1;
                  })
                  //console.log(mks)
                  let markers = []
                  let infoWindows = []
                  mks.forEach((mk, idx) => {
                    let markerContent = `<div style="width:25px;height:34px;background-image:url(https://i.loli.net/2020/05/28/d39PwI27fBvz1Fe.png);background-repeat:no-repeat;background-size:25px 34px;text-align:center;opacity: 0.8;"> <div>${idx + 1}</div> </div>`
                    //console.log('!!!!')
                    var marker = new AMap.Marker({
                      map: o,
                      content: markerContent,
                      position: [mk.R + (Math.random() - 0.5) * 3e-4, mk.Q + (Math.random() - 0.5) * 3e-4]
                    });
                    markers.push(marker)
                    //console.log(mk.R, mk.Q)
                    let text = ` \
                      <div style="color:black"> \
                        <span style="font-weight:bold"> 地点 </span>\
                        <span> ${self.events[idx].addr}</span> \
                      </div> \
                      <div style="color:black"> \
                        <span style="font-weight:bold"> 时间 </span>\
                        <span>${moment(self.events[idx].startDate).format('MM-DD')} ${self.events[idx].startTime}~${moment(self.events[idx].endDate).format('MM-DD')} ${self.events[idx].endTime}</span> \
                      </div> \
                      <div style="color:black"> \
                        <span style="font-weight:bold"> 描述 </span>\
                        <span> ${self.events[idx].description}</span> \
                      </div> `
                    let infoWindow = new AMap.InfoWindow({
                      content: text,
                      anchor: 'top-center',
                    });
                    infoWindows.push(infoWindow)
                    AMap.event.addListener(marker, 'click', function () {
                      infoWindow.open(o, marker.getPosition());
                    });
                  })
                  let human = new AMap.Marker({
                    map: o,
                    position: _.cloneDeep(path[0]),
                    // icon: "https://i.loli.net/2020/05/28/HaQerMxyFq2KnsI.png",
                    // offset: new AMap.Pixel(-26, -13),
                    content: humanContent,
                    //autoRotation: true,
                    //angle: -90,
                  });
                  self.human = human
                  //human.moveAlong(_.cloneDeep(path), 200);
                  o.setFitView();

                  console.log(self.events)
                  let current = 0
                  console.log("velocity", self.velocity)
                  human.on('moving', function (e) {

                    let currentPos = e.passedPath[e.passedPath.length - 1]
                    //console.log(currentPos.R, currentPos.Q)

                    if (current < mks.length && Math.abs(currentPos.R - mks[current].R) + Math.abs(currentPos.Q - mks[current].Q) < self.velocity * 6e-7) {
                      //console.log(current)
                      self.date = self.events[current].startDate
                      if (self.events[current].startDate != self.events[current].endDate) {
                        //console.log(self.events[current].startDate, self.events[current].endDate)
                        human.pauseMove()
                        if (self.interval != null)
                          clearInterval(self.interval)
                        let cur = moment(self.events[current].startDate)
                        self.interval = setInterval(() => {
                          let current2 = current - 1
                          //console.log("Here!!", current)
                          cur.add(1, 'days')
                          self.date = cur.format('YYYY-MM-DD')
                          console.log(self.date, self.events[current2].endDate)
                          if (current2 >= self.events.length || self.date == self.events[current2].endDate) {
                            clearInterval(self.interval)
                            self.interval = null
                            self.human.resumeMove()
                          }
                        }, 500);
                      }
                      infoWindows[current].open(o, markers[current].getPosition())
                      current += 1
                    }

                    passedPolyline.setPath(e.passedPath);
                  });

                }

                if (status === 'complete') {
                  //console.log('绘制驾车路线完成');
                } else {
                  console.log('获取驾车数据失败：' + result);
                }
              });
              o.addControl(driving);
            }
          })
        }
      }
    }
  },
  asyncData (context) {
    return context.app.$axios.get(`/api/news/${context.route.params.id}/`).then(
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
}
</script>
