<template>
  <div>
    <v-container>
      <v-dialog
        max-width="60%"
        scrollable
        v-model="dialog"
        persistent
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
            <v-icon>mdi-filter-outline</v-icon>
          </v-btn>

        </template>
        <v-card>
          <v-card-title>关键词筛选</v-card-title>
          <v-divider></v-divider>
          <v-card-text style="height: 300px;">
            <v-treeview
              v-model="selection"
              :items="items"
              selection-type="leaf"
              selectable
              return-object
              open-all
            ></v-treeview>
          </v-card-text>
          <v-card-text>
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
                      v-model="startPublishDate"
                      class="mt-1"
                      label="开始时间"
                      prepend-icon="mdi-calendar"
                      dense
                      readonly
                      outlined
                      v-on="on"
                      required
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="startPublishDate"
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
                      v-model="endPublishDate"
                      class="mt-1"
                      label="结束时间"
                      prepend-icon="mdi-calendar"
                      dense
                      readonly
                      outlined
                      v-on="on"
                      required
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="endPublishDate"
                    scrollable
                    no-title
                    @input="menu2=false"
                    light
                  >
                  </v-date-picker>
                </v-menu>
              </v-col>
            </v-row>
            <v-row v-if="startPublishDate!=''||endPublishDate!=''">
              <div class="center">
                <v-btn
                  rounded
                  outlined
                  color="purple"
                  small
                  class="ma-1"
                  v-if="startPublishDate!=''"
                  @click.native="startPublishDate = ''"
                >
                  <v-icon left>mdi-minus</v-icon> 晚于 {{startPublishDate}}
                </v-btn>
                <v-btn
                  rounded
                  outlined
                  color="purple"
                  small
                  class="ma-1"
                  v-if="endPublishDate!=''"
                  @click.native="endPublishDate = ''"
                >
                  <v-icon left>mdi-minus</v-icon> 早于 {{endPublishDate}}
                </v-btn>
              </div>
            </v-row>
            <v-row v-if="selection.length>0">
              <div class="center">
                <v-btn
                  v-for="node in selection"
                  :key="node.id"
                  rounded
                  outlined
                  color="primary"
                  small
                  class="ma-1"
                  @click.native="remove_keyword(node)"
                >
                  <v-icon left>mdi-minus</v-icon> {{node.name}}
                </v-btn>
              </div>
            </v-row>
          </v-card-text>
          <v-card-actions>
            <v-spacer />
            <v-btn
              x-large
              text
              color='primary'
              @click="apply()"
            >Apply</v-btn>
            <v-btn
              x-large
              text
              color='red'
              @click="cancel()"
            >Cancel</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <div>
        <v-btn
          text
          value='left'
          color="purple"
          large
          v-if="startPublishDate!=''"
          @click.native="startPublishDate = ''"
        >
          <v-icon left>mdi-minus</v-icon>晚于 {{startPublishDate}}
        </v-btn>
        <v-btn
          text
          value='left'
          color="purple"
          large
          v-if="endPublishDate!=''"
          @click.native="endPublishDate = ''"
        >
          <v-icon left>mdi-minus</v-icon>早于 {{endPublishDate}}
        </v-btn>
      </div>
      <div>
        <v-btn
          value="left"
          v-for="node in selection"
          :key="node.id"
          text
          large
          @click.native="remove_keyword(node)"
        >
          <v-icon left>mdi-minus</v-icon>
          {{node.name}}
        </v-btn>
      </div>
      <news-list :news-list="newsList" />
      <div class="text-center">
        <v-pagination
          v-model="page"
          :length="totalPage"
          :total-visible="10"
        ></v-pagination>
      </div>
    </v-container>
  </div>
</template>
<script>

import NewsList from '~/components/NewsList.vue'

export default {
  auth: false,
  components: {
    NewsList,
  },
  data () {
    return {
      page: 1,
      totalPage: 1,
      dialog: false,
      menu1: false,
      menu2: false,
      newsList: [],
      selection: [],
      startPublishDate: "",
      endPublishDate: "",
      items: [
        {
          id: 1,
          name: '华中',
          children: [
            { id: 2, name: '河南' },
            { id: 3, name: '湖北' },
            { id: 4, name: '湖南' },
          ],
        },
        {
          id: 5,
          name: '华东',
          children: [
            { id: 6, name: '山东' },
            { id: 7, name: '江苏' },
            { id: 8, name: '安徽' },
            { id: 9, name: '上海' },
            { id: 10, name: '浙江' },
            { id: 11, name: '江西' },
            { id: 12, name: '福建' },
            { id: 13, name: '台湾' },
          ],
        },
        {
          id: 14,
          name: '华南',
          children: [
            { id: 15, name: '广东' },
            { id: 16, name: '广西' },
            { id: 17, name: '海南' },
            { id: 18, name: '香港' },
            { id: 19, name: '澳门' },
          ],
        },
        {
          id: 20,
          name: '华北',
          children: [
            { id: 21, name: '河北' },
            { id: 22, name: '山西' },
            { id: 23, name: '北京' },
            { id: 24, name: '天津' },
            { id: 25, name: '内蒙古' },
          ],
        },
        {
          id: 26,
          name: '西北',
          children: [
            { id: 27, name: '陕西' },
            { id: 28, name: '甘肃' },
            { id: 29, name: '青海' },
            { id: 30, name: '宁夏' },
            { id: 31, name: '新疆' },
          ],
        },
        {
          id: 32,
          name: '西南',
          children: [
            { id: 32, name: '重庆' },
            { id: 33, name: '四川' },
            { id: 34, name: '贵州' },
            { id: 35, name: '云南' },
            { id: 36, name: '西藏' },
          ],
        },
        {
          id: 37,
          name: '东北',
          children: [
            { id: 38, name: '黑龙江' },
            { id: 39, name: '吉林' },
            { id: 40, name: '辽宁' },
          ],
        },
        {
          id: 41,
          name: '南海',
          children: [
            { id: 42, name: '南海诸岛' },
          ],
        },
      ],
      searchCondition: {
        categories: [],
        username: "",
        startPublishDate: "",
        endPublishDate: "",
        pageNum: 1,
        newsPerPage: 10
      },
      clone: ""
    }
  },
  asyncData (context) {
    return context.app.$axios.post(`/api/newslib/`, { categories: [], username: "", startPublishDate: "", endPublishDate: "", pageNum: 1, newsPerPage: 10 }).then(
      res => {
        return {
          newsList: res.data.news,
          totalPage: res.data.totalPage > 1 ? res.data.totalPage : 1
        }
      },
      () => {
        context.error({ statusCode: 400, message: "error" })
        return {}
      }
    )
  },
  watch: {
    "page": function (newVal) {
      this.searchCondition.pageNum = newVal
      if (this.dialog) return
      this.$axios.post(`/api/newslib/`, this.searchCondition).then(
        res => {
          this.newsList = res.data.news
          this.totalPage = res.data.totalPage > 1 ? res.data.totalPage : 1
        },
        () => {
          context.error({ statusCode: 400, message: "error" })
          return
        }
      )
    },
    "dialog": function (newVal) {
      if (!newVal) return
      let tmpObj = {
        selection: this.selection,
        startPublishDate: this.startPublishDate,
        endPublishDate: this.endPublishDate,
        searchCondition: this.searchCondition,
        page: this.page
      }
      this.clone = JSON.stringify(tmpObj)
    },
    "selection": {
      handler (newVal, oldVal) {
        this.searchCondition.categories = []
        for (let idx in newVal) {
          //console.log(sel)
          this.searchCondition.categories.push(newVal[idx].name)
        }
        this.searchCondition.pageNum = 1
        if (this.dialog) return
        this.$axios.post(`/api/newslib/`, this.searchCondition).then(
          res => {
            this.newsList = res.data.news,
              this.totalPage = res.data.totalPage > 1 ? res.data.totalPage : 1
          },
          () => {
            context.error({ statusCode: 400, message: "error" })
            return
          }
        )
      },
      deep: true
    },
    "startPublishDate": function (newVal) {
      if (this.dialog) return
      this.searchCondition.startPublishDate = newVal
      this.$axios.post(`/api/newslib/`, this.searchCondition).then(
        res => {
          this.newsList = res.data.news,
            this.totalPage = res.data.totalPage > 1 ? res.data.totalPage : 1
        },
        () => {
          context.error({ statusCode: 400, message: "error" })
          return
        }
      )
    },
    "endPublishDate": function (newVal) {
      if (this.dialog) return
      this.searchCondition.endPublishDate = newVal
      this.$axios.post(`/api/newslib/`, this.searchCondition).then(
        res => {
          this.newsList = res.data.news,
            this.totalPage = res.data.totalPage > 1 ? res.data.totalPage : 1
        },
        () => {
          context.error({ statusCode: 400, message: "error" })
          return
        }
      )
    },
  },
  methods: {
    remove_keyword (nd) {
      this.selection = this.selection.filter(node => {
        return node.id != nd.id
      })
    },
    cancel () {
      let tmpObj = JSON.parse(this.clone)
      this.selection = tmpObj.selection
      this.page = tmpObj.page
      this.searchCondition = tmpObj.searchCondition
      this.startPublishDate = tmpObj.startPublishDate
      this.endPublishDate = tmpObj.endPublishDate
      this.dialog = false
    },
    apply () {
      //console.log("???")
      this.dialog = false
      this.searchCondition.categories = []
      for (let idx in this.selection) {
        //console.log(sel)
        this.searchCondition.categories.push(this.selection[idx].name)
      }
      //this.searchCondition.categories = this.selection
      this.searchCondition.startPublishDate = this.startPublishDate
      this.searchCondition.endPublishDate = this.endPublishDate
      if (this.page == 1) {
        this.$axios.post(`/api/newslib/`, this.searchCondition).then(
          res => {
            this.newsList = res.data.news
            this.totalPage = res.data.totalPage > 1 ? res.data.totalPage : 1
          },
          () => {
            context.error({ statusCode: 400, message: "error" })
            return
          }
        )
      } else {
        this.page = 1
      }
    }
  }
}
</script>