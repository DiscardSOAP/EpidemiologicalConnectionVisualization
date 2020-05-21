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
          <v-img :src="'https://cdn.v2ex.com/gravatar/'
          +
          user.email_md5
          + '?s=1000&d=mm'">
          </v-img>
          <v-card-title>{{user.name}}</v-card-title>
          <v-card-subtitle class="pb-0">
            <div>
              <v-icon>mdi-account-box-outline</v-icon>
              {{user.organization}}
            </div>
            <div>
              <v-icon>mdi-email</v-icon>
              {{user.email}}
            </div>
          </v-card-subtitle>
          <v-card-text>

            <div>{{user.description}}</div>
          </v-card-text>
        </v-card>

      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import NewsList from '~/components/NewsList.vue'


export default {
  auth: false,
  components: {
    NewsList,
  },
  asyncData (context) {
    return context.app.$axios.get(`/api/profile/${context.route.params.id}/`).then(
      res => {
        let user = res.data.user

        return context.app.$axios.post(`/api/newslib/`, { categories: [], username: user.username, startPublishDate: "", endPublishDate: "", newsPerPage: 10, pageNum: 1 }).then(
          res => {
            return {
              user: user,
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
      () => {
        context.error({ statusCode: 400, message: "error" })
        return {}
      }
    )
  },
  watch: {
    "page": function (newVal) {
      this.$axios.post(`/api/newslib/`, { categories: [], username: this.user.username, startPublishDate: "", endPublishDate: "", newsPerPage: 10, pageNum: newVal }).then(
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
      user: {
        username: '',
        name: '',
        birth: '',
        organization: '',
        description: '',
        email: '',
        email_md5: '',
      },
      page: 1,
      totalPage: 1,
      newsList: []
    }
  },
}
</script>
