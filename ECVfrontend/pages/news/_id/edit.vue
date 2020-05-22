<template>
  <div>
    <v-form
      ref="form"
      v-model="newsForm.valid"
    >
      <v-container>
        <v-row>
          <v-col>
            <v-text-field
              v-model="newsForm.form.title"
              :counter="50"
              :rules="[v=>v.length > 0 || '新闻标题不能为空']"
              label="新闻标题"
              required
            ></v-text-field>
          </v-col>
          <v-col cols="2">
            <v-select
              v-model="newsForm.form.category"
              :items="newsForm.categories"
              :rules="[v => v.length > 0 || '区域不能为空']"
              label="区域"
              required
            ></v-select>
          </v-col>
          <v-col cols="2">
            <v-btn
              color='primary'
              block
              :disabled='!newsForm.valid'
              @click='submit()'
            >
              Upload
              <v-icon
                right
                dark
              >mdi-cloud-upload</v-icon>
            </v-btn>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <client-only>
              <mavon-editor
                :disabled="newsForm.loading||newsForm.ok"
                style="z-index:0"
                v-model="newsForm.form.markdownText"
                @change="changeText"
              />
            </client-only>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-combobox
              v-model="comboBox.rawRelationships"
              :filter="comboBoxFilter"
              :hide-no-data="!comboBox.search"
              :items="comboBox.items"
              :search-input.sync="comboBox.search"
              :rules="comboBox.rules"
              hide-selected
              label="建立病例联系"
              multiple
              outlined
              small-chips
            >
              <template v-slot:no-data>
                <v-list-item>
                  <v-chip
                    :color="`${comboBox.colors[comboBox.colorCnt % comboBox.colors.length]}`"
                    label
                  >
                    {{ comboBox.search }}
                  </v-chip>
                </v-list-item>
              </template>
              <template v-slot:selection="{ attrs, item, parent, selected }">
                <v-chip
                  v-if="item === Object(item)"
                  v-bind="attrs"
                  :color="`${item.color}`"
                  :input-value="selected"
                  label
                >
                  <span class="pr-2">
                    {{ item.text }}
                  </span>
                  <v-icon @click="parent.selectItem(item)">mdi-close</v-icon>
                </v-chip>
              </template>
              <template v-slot:item="{ index, item }">
                <v-text-field
                  v-if="comboBox.editing === item"
                  v-model="comboBox.editing.text"
                  autofocus
                  flat
                  background-color="transparent"
                  solo
                  persistent-hint
                  hint="输入的格式应为 '联系的文字说明 : 另一患者id'"
                  @keyup.enter="comboBoxEdit(index, item)"
                ></v-text-field>
                <v-chip
                  v-else
                  :color="`${item.color}`"
                  dark
                  label
                >
                  {{ item.text }}
                </v-chip>
                <v-spacer></v-spacer>
                <v-list-item-action @click.stop>
                  <v-btn
                    icon
                    @click.stop.prevent="comboBoxEdit(index, item)"
                  >
                    <v-icon>{{ comboBox.editing !== item ? 'mdi-pencil' : 'mdi-check' }}</v-icon>
                  </v-btn>
                </v-list-item-action>
              </template>
            </v-combobox>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-timeline v-if="newsForm.form.events.length>0">
              <v-timeline-item
                v-for="(event, i) in newsForm.form.events"
                :key="i"
                :color="newsForm.colors[i%newsForm.colors.length]"
                small
                right
              >
                <template v-slot:opposite>
                  <span
                    :class="`headline font-weight-bold ${newsForm.colors[i%newsForm.colors.length]}--text`"
                    v-text="event.startDate+' '+event.startTime+' ~ '+event.endDate+' '+event.endTime"
                  ></span>
                </template>
                <div class="py-2">
                  <v-row>
                    <h2 :class="`headline ${newsForm.colors[i%newsForm.colors.length]}--text`">
                      {{event.addr}}
                      <v-btn
                        icon
                        x-large
                        color="grey"
                        @click="deleteEvent(event)"
                      >
                        <v-icon>mdi-delete</v-icon>
                      </v-btn>
                    </h2>

                  </v-row>
                  <v-row>
                    {{event.description}}
                  </v-row>
                </div>
              </v-timeline-item>
            </v-timeline>
          </v-col>
        </v-row>

      </v-container>
    </v-form>
    <v-dialog
      max-width="60%"
      scrollable
      ref="dialog"
      v-model="dialog"
    >
      <template v-slot:activator="{ on }">
        <v-btn
          dark
          fab
          bottom
          right
          fixed
          x-large
          color="primary"
          v-on="on"
        >
          <v-icon>mdi-plus</v-icon>
        </v-btn>

      </template>
      <v-card>
        <v-card-title>添加患者行动轨迹</v-card-title>
        <v-divider></v-divider>
        <v-card-text>
          <v-form
            ref="form"
            v-model='eventForm.valid'
          >
            <v-row>
              <v-col>
                <v-combobox
                  v-model="eventForm.form.addr"
                  :filter="addrBoxFilter"
                  :hide-no-data="!addrBox.search"
                  :items="addrBox.items"
                  :search-input.sync="addrBox.search"
                  hide-selected
                  label="地址"
                  outlined
                  small-chips
                >
                  <template v-slot:no-data>
                    <v-list-item>
                      {{ addrBox.search }}
                    </v-list-item>
                  </template>
                  <template v-slot:selection="{ attrs, item}">
                    {{ item }}
                  </template>
                  <template v-slot:item="{ index, item }">
                    <!--v-text-field
                      v-if="comboBox.editing === item"
                      v-model="comboBox.editing.text"
                      autofocus
                      flat
                      background-color="transparent"
                      solo
                      persistent-hint
                      hint="输入的格式应为 '联系的文字说明 : 另一患者id'"
                      @keyup.enter="comboBoxEdit(index, item)"
                    ></v-text-field>
                    <v-chip
                      v-else
                      :color="`${item.color}`"
                      dark
                      label
                    -->
                    {{ item }}
                    <!--/v-chip>
                    <v-spacer></v-spacer>
                    <v-list-item-action @click.stop>
                      <v-btn
                        icon
                        @click.stop.prevent="comboBoxEdit(index, item)"
                      >
                        <v-icon>{{ comboBox.editing !== item ? 'mdi-pencil' : 'mdi-check' }}</v-icon>
                      </v-btn>
                    </v-list-item-action-->
                  </template>
                </v-combobox>
                <!--v-text-field
                  v-model="eventForm.form.addr"
                  :counter="50"
                  :rules="[v=>v.length > 0 || '地址不能为空']"
                  label="地址"
                  required
                ></v-text-field-->
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-menu
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  ref="menu1"
                  v-model="menu1"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="eventForm.form.startDate"
                      class="mt-1"
                      label="开始日期"
                      prepend-icon="mdi-calendar"
                      dense
                      readonly
                      outlined
                      v-on="on"
                      required
                      :rules="[v=>v.length > 0 || '开始日期不能为空']"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="eventForm.form.startDate"
                    scrollable
                    no-title
                    @input="menu1=false"
                    light
                  >
                  </v-date-picker>
                </v-menu>
              </v-col>
              <v-col>
                <v-menu
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  ref="menu2"
                  v-model="menu2"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="eventForm.form.startTime"
                      class="mt-1"
                      label="开始时刻"
                      dense
                      readonly
                      outlined
                      v-on="on"
                      required
                      :rules="[v=>v.length > 0 || '开始时刻不能为空']"
                    ></v-text-field>
                  </template>
                  <v-time-picker
                    v-model="eventForm.form.startTime"
                    format="24hr"
                    @input="menu2=false"
                    light
                  >
                  </v-time-picker>
                </v-menu>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-menu
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  ref="menu3"
                  v-model="menu3"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="eventForm.form.endDate"
                      class="mt-1"
                      label="结束日期"
                      prepend-icon="mdi-calendar"
                      dense
                      readonly
                      outlined
                      v-on="on"
                      required
                      :rules="[v=>v.length > 0 || '结束日期不能为空']"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="eventForm.form.endDate"
                    scrollable
                    no-title
                    @input="menu3=false"
                    light
                  >
                  </v-date-picker>
                </v-menu>

              </v-col>
              <v-col>
                <v-menu
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  ref="menu4"
                  v-model="menu4"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="eventForm.form.endTime"
                      class="mt-1"
                      label="结束时刻"
                      dense
                      readonly
                      outlined
                      v-on="on"
                      required
                      :rules="[v=>v.length > 0 || '结束时刻不能为空']"
                    ></v-text-field>
                  </template>
                  <v-time-picker
                    v-model="eventForm.form.endTime"
                    format="24hr"
                    @input="menu4=false"
                    light
                  >
                  </v-time-picker>
                </v-menu>
              </v-col>
            </v-row>
            <v-row v-if="eventForm.form.startTime.length >0 && eventForm.form.startDate.length >0 && eventForm.form.endTime.length >0 && eventForm.form.endDate.length >0 && !validateTime()">
              <v-col>
                <v-alert
                  type="error"
                  outlined
                >
                  结束时刻必须晚于开始时刻
                </v-alert>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-textarea
                  outlined
                  name="description"
                  label="事件描述"
                  v-model="eventForm.form.description"
                ></v-textarea>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer />

          <v-btn
            text
            color='primary'
            x-large
            :disabled="!eventForm.valid || (eventForm.form.startTime.length >0 && eventForm.form.startDate.length >0 && eventForm.form.endTime.length >0 && eventForm.form.endDate.length >0 && !validateTime())"
            @click="submitEventForm()"
          >SUBMIT</v-btn>
          <v-btn
            text
            color='red'
            x-large
            @click="clearEventForm()"
          >CANCEL</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
import Vue from 'vue'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import moment from 'moment'
import { AMapManager } from 'vue-amap'
import VueAMap from 'vue-amap'
let amapManager = new VueAMap.AMapManager()
Vue.use(mavonEditor)

export default {
  data: function () {
    return {
      menu1: false,
      menu2: false,
      menu3: false,
      menu4: false,
      dialog: false,
      newsForm: {
        valid: false,
        form: {
          title: "",
          text: "",
          category: "",
          events: [],
          markdownText: "",
          relationships: []
        },
        loading: false,
        ok: false,
        error: false,
        categories: [],
        colors: []
      },
      addrBox: {
        search: null,
        items: [],
        staticNum: 0
      },
      eventForm: {
        form: {
          startDate: "",
          startTime: "",
          endDate: "",
          endTime: "",
          addr: "",
          description: "",
        },
        valid: false
      },

    }
  },
  asyncData (context) {
    return context.app.$axios
      .get(`/api/news/${context.route.params.id}/`)
      .then(
        res => {
          console.log(res.data)
          let colors = ['green', 'purple', 'indigo', 'cyan', 'teal', 'orange']
          let cnt = 0
          let rr = []
          for (let idx in res.data.relationships) {
            let item = res.data.relationships[idx]
            rr.push({
              text: `${item.relationship} : ${item.id}`,
              color: colors[(cnt++) % 6]
            })
          }
          return {
            newsForm: {
              form: res.data,
              valid: false,
              loading: false,
              ok: false,
              error: false,
              categories: [
                '天津', '北京', '广东', '新疆', '黑龙江', '吉林', '辽宁', '内蒙古', '河北', '山西', '陕西', '宁夏', '甘肃', '青海', '西藏', '四川', '云南', '广西', '贵州', '重庆', '河南', '山东', '湖北', '安徽', '江苏', '湖南', '江西', '浙江', '福建', '澳门', '香港', '台湾', '上海', '南海诸岛', '海南'
              ],
              colors: [
                'cyan', 'green', 'pink', 'amber', 'orange'
              ]
            },
            comboBox: {
              rawRelationships: rr,
              rules: [],
              editing: null,
              colors: ['green', 'purple', 'indigo', 'cyan', 'teal', 'orange'],
              colorCnt: cnt,
              search: null,
              items: [
                { header: '选取或新建联系，流行病传染关系请用 From : (id)的格式创建' },
                {
                  text: 'From : (id)',
                  color: 'blue',
                }
              ]
            },
          }
        },
        () => {
          context.error({ statusCode: 400, message: "error" })
          return {}
        }
      )
  },
  watch: {
    "addrBox.search": function (val, prev) {
      if (!val) return
      var autoOptions = {
        city: this.newsForm.category ? this.newsForm.category : "全国"
      }
      let self = this
      var autoComplete = new AMap.Autocomplete(autoOptions);
      autoComplete.search(val, function (status, result) {
        // 搜索成功时，result即是对应的匹配数据
        //console.log(result)
        self.addrBox.items = []
        for (let tip of result.tips) {
          self.addrBox.items.push(tip.name)
          //console.log(tip.name)
        }
      })
    },
    "comboBox.rawRelationships": function (val, prev) {
      if (val.length === prev.length) return

      this.comboBox.rawRelationships = val.map(v => {
        if (typeof v === 'string') {
          v = {
            text: v,
            color: this.comboBox.colors[this.comboBox.colorCnt % this.comboBox.colors.length],
          }

          this.comboBox.items.push(v)

          this.comboBox.colorCnt++
        }

        return v
      })
    },
  },
  mounted () {
    this.comboBox.rules.push(this.comboBoxRules)
  },
  methods: {
    addrBoxFilter (item, queryText, itemText) {
      if (item.header) return false

      const hasValue = val => val != null ? val : ''

      const text = hasValue(itemText)
      const query = hasValue(queryText)

      return text.toString()
        .toLowerCase()
        .indexOf(query.toString().toLowerCase()) > -1
    },
    addrBoxFilter (item, queryText, itemText) {
      const hasValue = val => val != null ? val : ''
      const text = hasValue(itemText)
      const query = hasValue(queryText)
      return text.toString()
        .toLowerCase()
        .indexOf(query.toString().toLowerCase()) > -1
    },
    comboBoxRules: function (labels) {
      for (let i in labels) {
        let label = labels[i].text

        let idx = label.indexOf(':')
        if (idx == -1)
          return "输入的格式应为 '联系的文字说明 : 另一患者id'"
        let tag = label.substring(0, idx).trim()
        let id = label.substring(idx + 1, label.length).trim()
        console.log(label, tag, id)
        id = parseInt(id)
        if (isNaN(id))
          return "输入的格式应为 '联系的文字说明 : 另一患者id'"
      }
      return true
    },
    changeText: function (markdown, html) {
      this.newsForm.form.text = html
    },
    clearEventForm: function () {
      this.eventForm.form.startDate = ""
      this.eventForm.form.startTime = ""
      this.eventForm.form.endDate = ""
      this.eventForm.form.endTime = ""
      this.eventForm.form.addr = ""
      this.eventForm.form.description = ""
      this.dialog = false
      this.eventForm.valid = true
    },
    deleteEvent (event) {
      this.newsForm.form.events = this.newsForm.form.events.filter(e => !(e.startDate == event.startDate && e.startTime == event.startTime && e.endDate == event.endDate && e.endTime == event.endTime && e.addr == event.addr
        && e.description == event.description))
    },
    validateTime: function () {
      let start = moment(this.eventForm.form.startDate + ' ' + this.eventForm.form.startTime).valueOf()
      let end = moment(this.eventForm.form.endDate + ' ' + this.eventForm.form.endTime).valueOf()
      if (start >= end)
        return false
      else
        return true
    },
    submitEventForm: function () {
      this.newsForm.form.events.push(JSON.parse(JSON.stringify(this.eventForm.form)))
      this.newsForm.form.events = this.newsForm.form.events.sort(function (a, b) {
        return moment(a.startDate + ' ' + a.startTime) - moment(b.startDate + ' ' + b.startTime)
      })
      this.clearEventForm()
      this.dialog = false
    },
    submit: function () {
      this.newsForm.loading = true
      this.newsForm.form.relationships = []
      for (let i in this.comboBox.rawRelationships) {
        let label = this.comboBox.rawRelationships[i].text
        let idx = label.indexOf(':')
        let tag = label.substring(0, idx).trim()
        let id = label.substring(idx + 1, label.length).trim()
        id = parseInt(id)
        this.newsForm.form.relationships.push({
          relationship: tag,
          id: id
        })
      }
      this.$axios
        .post(`/api/news/${this.$route.params.id}/`, this.newsForm.form)
        .then(
          res => {
            this.newsForm.error = false
            this.newsForm.ok = true
            setTimeout(() => {
              this.$router.push({ path: `/news/${this.$route.params.id}/` })
              console.log("shit!!")
            }, 500)
          },
          err => {
            this.newsForm.error = true
          },
          () => {

          }
        )
        .then(() => {
          this.newsForm.loading = false
        })
    },
    comboBoxEdit (index, item) {
      if (!this.comboBox.editing) {
        this.comboBox.editing = item
      } else {
        this.comboBox.editing = null
      }
    },
    comboBoxFilter (item, queryText, itemText) {
      if (item.header) return false

      const hasValue = val => val != null ? val : ''

      const text = hasValue(itemText)
      const query = hasValue(queryText)

      return text.toString()
        .toLowerCase()
        .indexOf(query.toString().toLowerCase()) > -1
    },
  }
}
</script>
