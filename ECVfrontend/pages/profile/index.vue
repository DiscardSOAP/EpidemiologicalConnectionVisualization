<template>
  <v-container>
    <v-row>
      <v-col
        cols="12"
        md="8"
      >
        <news-list :news-list="newsList" />
        <div class="text-center">
          <v-pagination
            v-model="page"
            :length="totalPage"
            :total-visible="10"
          ></v-pagination>
        </div>
      </v-col>
      <v-col
        cols="6"
        md="4"
      >
        <v-card class="mx-auto my-4">
          <v-hover v-slot:default="{ hover }">
            <v-img :src="'https://cdn.v2ex.com/gravatar/'
          +
          $auth.$state.user.email_md5
          + '?s=1000&d=mm'">
              <v-expand-transition>
                <div
                  v-if="hover"
                  class="d-flex transition-fast-in-fast-out black darken-2 v-card--reveal"
                  style="height: 100%;opacity:0.8;"
                >
                  <a
                    is='v-btn'
                    color="primary"
                    style="position: absolute;top:50%;left:50%;transform:translate(-50%,-50%);"
                    href="https://en.gravatar.com/"
                    target="_blank"
                  >
                    Edit Image
                  </a>
                </div>
              </v-expand-transition>
            </v-img>
          </v-hover>
          <v-card-title>{{$auth.$state.user.name}}</v-card-title>
          <v-card-subtitle class="pb-0">
            <div>
              <v-icon>mdi-account</v-icon>{{$auth.$state.user.username}}
            </div>
            <div>
              <v-icon>mdi-account-box-outline</v-icon>{{$auth.$state.user.organization}}
            </div>
            <div>
              <v-icon>mdi-email</v-icon>{{$auth.$state.user.email}}
            </div>
          </v-card-subtitle>
          <v-card-text>

            <div>{{$auth.$state.user.description}}</div>
          </v-card-text>

          <v-divider class="mx-4"></v-divider>
          <v-card-actions>
            <v-btn
              color="deep-purple lighten-2"
              text
              block
              @click="$router.push({path:'/profile/edit/'})"
            >
              Edit
            </v-btn>
          </v-card-actions>
        </v-card>

      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import NewsList from '~/components/NewsList.vue'


export default {
  components: {
    NewsList,
  },
  asyncData (context) {
    return context.app.$axios.post(`/api/newslib/`, { categories: [], username: context.$auth.$state.user.username, startPublishDate: "", endPublishDate: "", newsPerPage: 10, pageNum: 1 }).then(
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
      this.$axios.post(`/api/newslib/`, { categories: [], username: this.$auth.$state.user.username, startPublishDate: "", endPublishDate: "", newsPerPage: 10, pageNum: newVal }).then(
        res => {
          this.newsList = res.data.news
          this.totalPage = res.data.totalPage > 1 ? res.data.totalPage : 1
        },
        () => {
          context.error({ statusCode: 400, message: "error" })
          return
        }
      )
    }
  },
  data: function () {
    return {
      newsList: [],
      page: 1,
      totalPage: 1,
    }
  },
}
</script>
