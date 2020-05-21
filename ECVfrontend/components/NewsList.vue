<template>
  <div>
    <v-container>
      <v-row dense>

        <v-col
          cols="12"
          v-for="(news,key) in newsList"
          :key="key"
        >
          <v-card
            class="mx-auto"
            width="100%"
          >
            <v-card-title>

              <v-btn
                text
                large
                color='primary'
                :to="`/news/${news.id}/`"
                class="title font-weight"
              >
                <v-icon
                  large
                  left
                >
                  mdi-bullhorn
                </v-icon>{{news.title}}
              </v-btn>
              <v-spacer />
              <v-icon>
                mdi-map-marker
              </v-icon>
              <span class="font-weight-light">
                {{news.category}}
              </span>
            </v-card-title>

            <v-card-text class="headline font-weight-light text-justify">
              <p style="
                overflow : hidden;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-line-clamp: 3;
                -webkit-box-orient: vertical;">
                {{news.markdownText}}
              </p>
            </v-card-text>
            <v-card-actions>
              <v-btn
                text
                :to="`/profile/${news.username}/`"
              >{{news.name}}</v-btn>

              <v-spacer></v-spacer>

              <v-btn
                color="primary"
                text
                disabled
              >
                发布日期 {{formatTime(news.publishTime)}}
              </v-btn>

            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>
<script>
import moment from "moment"
export default {
  props: {
    newsList: {
      type: Array,
      default: []
    }
  },
  data: function () {
    return {
    }
  },
  methods: {
    formatTime (timeStr) {
      return moment(timeStr).format('YYYY-MM-DD')
    }
  }

}
</script>
