<template>
  <v-container>
    <v-row>
      <v-col
        cols="12"
        md="8"
      >
        <news-list />
      </v-col>
      <v-col
        cols="6"
        md="4"
      >
        <v-card
          :loading="user.loading"
          class="mx-auto my-4"
        >
          <v-hover v-slot:default="{ hover }">
            <v-img src="https://bkimg.cdn.bcebos.com/pic/377adab44aed2e736d5390c38001a18b86d6faa3?x-bce-process=image/resize,m_lfit,w_268,limit_1">
              <v-expand-transition>
                <div
                  v-if="hover"
                  class="d-flex transition-fast-in-fast-out black darken-2 v-card--reveal"
                  style="height: 100%;opacity:0.8;"
                >
                  <v-btn
                    color="primary"
                    style="position: absolute;top:50%;left:50%;transform:translate(-50%,-50%);"
                    @click="$router.push({path:'/profile/edit/'})"
                  >
                    Edit
                  </v-btn>
                </div>
              </v-expand-transition>
            </v-img>
          </v-hover>
          <v-card-title>{{user.name}}</v-card-title>
          <v-card-subtitle class="pb-0">{{user.username}}</v-card-subtitle>
          <v-card-text>
            <div class="my-1 subtitle-1">
              • {{user.organization}}
            </div>

            <div>{{user.description}}</div>
          </v-card-text>

          <v-divider class="mx-4"></v-divider>

          <v-card-text>
            {{user.email}}
          </v-card-text>

          <v-card-actions>
            <v-btn
              color="deep-purple lighten-2"
              text
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
  auth: "guest",
  components: {
    NewsList,
  },
  data: function () {
    return {
      user: {
        username: "fnoi2014xtx",
        name: "徐天行",
        birth: "1999-2-14",
        organization: "天津",
        description: "得到覅大哥法国i噶放工返工的撒哥",
        email: "slothfulxtx@google.com",
        email_md5: "fdafdfafdgwegrgwerg",
        loading: true,
        error: false,
        ok: false,
      },

    }
  },
  /*asyncData (context) {
    return context.app.$axios.get('/api/profile/').then(
      res => {
        let data = res.data
        data.user.loading = false
        data.user.ok = true
        return data
      },
      () => {
        context.error({ statusCode: 401, message: "You do not have the proper privilege level!" })
      }
    )

  },*/
}
</script>
