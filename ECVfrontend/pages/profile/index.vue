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
        <v-card class="mx-auto my-4">
          <v-hover v-slot:default="{ hover }">
            <v-img src="https://www.baidu.com/img/bd_logo1.png">
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
          <v-card-subtitle class="pb-0">{{$auth.$state.user.username}}</v-card-subtitle>
          <v-card-text>
            <div class="my-1 subtitle-1">
              • {{$auth.$state.user.organization}}
            </div>

            <div>{{$auth.$state.user.description}}</div>
          </v-card-text>

          <v-divider class="mx-4"></v-divider>

          <v-card-text>
            {{$auth.$state.user.email}}
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
  components: {
    NewsList,
  },
  data: function () {
    return {
    }
  },
}
</script>
