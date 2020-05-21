<template>
  <div style="width:100%;height:100%">
    <h1 style="width:100%;text-align:center;font-weight:normal">
      {{nameMap[$route.params.id]}}疫情传播图谱
      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn
            icon
            color="yellow"
            :to="`/track/${$route.params.id}/`"
            v-on="on"
          >
            <v-icon x-large>mdi-google-maps</v-icon>
          </v-btn>
        </template>
        <span>切换至行动轨迹图</span>
      </v-tooltip>

    </h1>

    <network
      style="width:100%;height:90%;"
      ref="network"
      :nodes="nodes"
      :edges="edges"
      :options="options"
      @click="onItemSelected"
    >
    </network>
    <v-bottom-sheet v-model="sheet.open">
      <v-sheet
        height="200px"
        style="text-align:center"
      >
        <v-btn
          class="mt-6"
          text
          large
          :color="sheet.error?'red':'purple'"
          :disabled="sheet.loading"
          :loading="sheet.loading"
          :to="`/news/${sheet.id}/`"
        >{{sheet.error?'ERROR':'Learn More'}}</v-btn>
        <!-- <v-card
          color="#385F73"
          dark
          style="width:70%"
          
        >
          <v-card-title class="headline">Unlimited music now</v-card-title>

          <v-card-subtitle>Listen to your favorite artists and albums whenever and wherever, online and offline.</v-card-subtitle>

          <v-card-actions>
            <v-btn text>Listen Now</v-btn>
          </v-card-actions>
        </v-card> -->
        <div
          class="py-3"
          style="margin:0 auto;width:70%;text-align:justify;"
          v-html="sheet.error?'Connection Error!':sheet.news.text"
        ></div>
      </v-sheet>
    </v-bottom-sheet>
  </div>
</template>

<script>
import { Network } from "vue-vis-network"

const nameMap = {
  'tianjin': '天津',
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
}

export default {
  auth: false,
  components: {
    Network
  },
  data: function () {
    return {
      nameMap: {
        'tianjin': '天津',
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
      sheet: {
        open: false,
        loading: false,
        error: false,
        id: 0,
        news: {
          title: '',
          publishTime: '',
          markdownText: '',
          text: '',
          events: [],
        },
      },
      nodes: [],
      edges: [],
      options: {
        nodes: {
          shape: 'dot',
          size: 30,
          font: {
            size: 32,
            color: '#ffffff'
          },
          borderWidth: 2
        },
        edges: {
          smooth: {
            type: 'continuous',
          }
        }
      }
    }
  },
  methods: {
    onNodeSelected: function (node) {
      this.sheet.open = true
      this.sheet.id = node
      this.sheet.loading = true
      this.sheet.error = false
      this.sheet.news = {
        title: '',
        publishTime: '',
        markdownText: '',
        events: [],
      }
      this.$axios.get(`/api/news/${node}/`).then(
        res => {
          this.sheet.loading = false
          this.sheet.news = res.data
          this.sheet.error = false
        },
        () => {
          this.sheet.error = true
          this.sheet.loading = false
        }
      )
    },
    onEdgeSelected: function (edge) {
      for (let idx in this.edges)
        if (this.edges[idx].id == edge) {
          this.edges[idx].label == "father and son"
        }
    },
    onItemSelected: function ({ nodes, edges }) {
      let eventType = null
      if (nodes.length > 0) {
        eventType = 'node'
      } else if (edges.length > 0) {
        eventType = 'edge'
      } else {
        eventType = 'unknown'
      }
      switch (eventType) {
        case 'node':
          this.onNodeSelected(nodes[0])
          break
        case 'edge':
          this.onEdgeSelected(edges[0])
          break
        default:
          break
      }
    }
  },
  asyncData (context) {
    return context.app.$axios.post(`/api/topology/`, {
      category: nameMap[context.route.params.id]
    }).then(
      res => {
        let graph = res.data
        console.log(graph)
        let nodes = []
        let outDegree = {}
        for (let idx in graph) {
          let nd = graph[idx]
          nodes.push({ id: nd.id, label: nd.id.toString(), color: 'rgb(255,127,127)' })
          outDegree[nd.id] = 0
        }
        if (nodes.length == 0)
          return {}
        let edges = []
        for (let idx in graph) {
          let nd = graph[idx]
          for (let idx2 in nd.relationships) {
            let r = nd.relationships[idx2]
            if (r.relationship == 'From' && r.id in outDegree) {
              outDegree[r.id]++
              edges.push({
                from: r.id,
                to: nd.id,
                id: edges.length,
                arrows: 'to',
                color: { color: 'rgb(255,127,127)' }
              })
            }
            else {
              edges.push({
                from: nd.id,
                to: r.id,
                id: edges.length,
                label: r.relationship,
                color: { color: 'white' },
                dashes: true
              })
            }
          }
        }
        let minOut = 1e9
        let maxOut = 0
        for (let nd in outDegree) {
          minOut = Math.min(minOut, outDegree[nd])
          maxOut = Math.max(maxOut, outDegree[nd])
        }
        for (let idx in nodes) {
          let nd = nodes[idx]
          let out = outDegree[nd.id]
          if (minOut == maxOut)
            break
          else {
            let v = 191 - 1.0 * (out - minOut) / (maxOut - minOut) * 191
            v = parseInt(v)
            v = Math.min(Math.max(v, 0), 191)
            nodes[idx].color = `rgb(255,${v},${v})`
          }
        }
        console.log(nodes)
        console.log(edges)
        return {
          nodes: nodes,
          edges: edges
        }
      },
      () => {
        context.error({ statusCode: 400, message: "error" })
        return {}
      }
    )
  },
  mounted () {
    /*this.nodes = [
      { id: 0, label: "0", group: 0 },
      { id: 1, label: "1", group: 0 },
      { id: 2, label: "2", group: 0 },
      { id: 3, label: "3", group: 1 },
      { id: 4, label: "4", group: 1 },
      { id: 5, label: "5", group: 1 },
      { id: 6, label: "6", group: 2 },
      { id: 7, label: "7", group: 2 },
      { id: 8, label: "8", group: 2 },
      { id: 9, label: "9", group: 3 },
      { id: 10, label: "10", group: 3 },
      { id: 11, label: "11", group: 3 },
      { id: 12, label: "12", group: 4 },
      { id: 13, label: "13", group: 4 },
      { id: 14, label: "14", group: 4 },
      { id: 15, label: "15", group: 5 },
      { id: 16, label: "16", group: 5 },
      { id: 17, label: "17", group: 5 },
      { id: 18, label: "18", group: 6 },
      { id: 19, label: "19", group: 6 },
      { id: 20, label: "20", group: 6 },
      { id: 21, label: "21", group: 7 },
      { id: 22, label: "22", group: 7 },
      { id: 23, label: "23", group: 7 },
      { id: 24, label: "24", group: 8 },
      { id: 25, label: "25", group: 8 },
      { id: 26, label: "26", group: 8 },
      { id: 27, label: "27", group: 9 },
      { id: 28, label: "28", group: 9 },
      { id: 29, label: "29", group: 9 }
    ]

    // create an array with edges
    this.edges = [
      { from: 1, to: 0, id: 0, arrows: 'from', label: 'hello' },
      { from: 2, to: 0, id: 1, arrows: 'from' },
      { from: 4, to: 3, id: 2, arrows: 'from' },
      { from: 5, to: 4, id: 3, arrows: 'from' },
      { from: 4, to: 0, id: 4, arrows: 'from' },
      { from: 7, to: 6, id: 5, arrows: 'from' },
      { from: 8, to: 7, id: 6, arrows: 'from' },
      { from: 7, to: 0, id: 7, arrows: 'from' },
      { from: 10, to: 9, id: 8, arrows: 'from' },
      { from: 11, to: 10, id: 9, arrows: 'from' },
      { from: 10, to: 4, id: 10, arrows: 'from' },
      { from: 13, to: 12, id: 11, arrows: 'from' },
      { from: 14, to: 13, id: 12, arrows: 'from' },
      { from: 13, to: 0, id: 13, arrows: 'from' },
      { from: 16, to: 15, id: 14, arrows: 'from' },
      { from: 17, to: 15, id: 15, arrows: 'from' },
      { from: 15, to: 10, id: 16, arrows: 'from' },
      { from: 19, to: 18, id: 17, arrows: 'from' },
      { from: 20, to: 19, id: 18, arrows: 'from' },
      { from: 19, to: 4, id: 19, arrows: 'from' },
      { from: 22, to: 21, id: 20, arrows: 'from' },
      { from: 23, to: 22, id: 21, arrows: 'from' },
      { from: 22, to: 13, id: 22, arrows: 'from' },
      { from: 25, to: 24, id: 23, arrows: 'from' },
      { from: 26, to: 25, id: 24, arrows: 'from' },
      { from: 25, to: 7, id: 25, arrows: 'from' },
      { from: 28, to: 27, id: 26, arrows: 'from' },
      { from: 29, to: 28, id: 27, arrows: 'from' },
      { from: 28, to: 0, id: 28, arrows: 'from' }
    ]*/
  }
}
</script>

<style>
</style>