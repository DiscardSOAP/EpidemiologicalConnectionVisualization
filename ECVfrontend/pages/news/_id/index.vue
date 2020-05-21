<template>
  <v-layout
    column
    justify-center
    align-center
  >
    <v-flex
      xs12
      sm8
      md6
      style="width:100%;"
    >
      <v-card
        flat
        style="width:100%; margin:0 auto; text-align:center"
      >
        <v-card-text>

          <p class="display-3 text--primary text-center">
            {{article.title}}
          </p>

          <p class="text-sm-center font-weight subtitle-1">
            <v-icon> mdi-map-marker </v-icon> {{article.category}}&nbsp;&nbsp;
            {{article.name}} 发布于 {{formatTime(article.publishTime)}}
          </p>
          <v-chip
            class="ma-2"
            color="white"
            outlined
            label
          >
            <v-icon left>mdi-account-circle-outline</v-icon>
            ID : {{$route.params.id}}
          </v-chip>
        </v-card-text>
        <v-card-text
          class="text-justify text--primary font-weight-thin headline"
          style="width:85%; margin:0 auto;"
          v-html="article.text"
        >
        </v-card-text>

        <v-card-text>
          <v-btn
            x-large
            text
            color="lime"
            class="display-1"
            :to="`/news/${$route.params.id}/track/`"
          >
            <v-icon
              left
              x-large
            >mdi-google-maps</v-icon>患者运动轨迹
          </v-btn>
          <v-timeline v-if="article.events.length>0">
            <v-timeline-item
              v-for="(event, i) in article.events"
              :key="i"
              :color="colors[i%colors.length]"
              small
              right
            >
              <template v-slot:opposite>
                <span
                  :class="`headline font-weight-bold ${colors[i%colors.length]}--text`"
                  v-text="event.startDate+' '+event.startTime+' ~ '+event.endDate+' '+event.endTime"
                ></span>
              </template>
              <div class="py-2">
                <v-row>
                  <h2 :class="`headline ${colors[i%colors.length]}--text`">
                    {{event.addr}}
                  </h2>

                </v-row>
                <v-row>
                  {{event.description}}
                </v-row>
              </div>
            </v-timeline-item>
          </v-timeline>
        </v-card-text>
        <hr class="my-1">
        <v-card-actions>
          <div class="my-2">
            <v-btn
              large
              text
              color="blue"
              :to="`/news/${article.prevArticle.id}/`"
            >
              <v-icon>mdi-arrow-left</v-icon>
              {{article.prevArticle.title}}

            </v-btn>
          </div>
          <v-spacer />
          <div class="my-2">
            <v-btn
              large
              text
              v-if="$auth.loggedIn"
              color="purple"
              :to="`/news/${$route.params.id}/edit/`"
            >
              <v-icon>mdi-pencil</v-icon>
              Edit

            </v-btn>
          </div>
          <v-spacer />
          <div class="my-2">
            <v-btn
              large
              text
              color="blue"
              :to="`/news/${article.nextArticle.id}/`"
            >
              {{article.nextArticle.title}}
              <v-icon>mdi-arrow-right</v-icon>
            </v-btn>
          </div>
        </v-card-actions>
      </v-card>

    </v-flex>
  </v-layout>

</template>
<script>
import moment from "moment"
export default {
  auth: false,
  data: function () {
    return {
      article: {
        id: 0,
        title: '',
        name: '',
        username: '',
        category: '',
        publishTime: '',
        prevArticle: {
          title: '',
          id: 0,
        },
        nextArticle: {
          title: '',
          id: 0,
        },
        text: '',
        events: [],
      },
      colors: [
        'cyan', 'green', 'pink', 'amber', 'orange'
      ]
    }
  },
  asyncData (context) {
    return context.app.$axios
      .get(`/api/news/${context.route.params.id}/`)
      .then(
        res => {
          //console.log(res.data)
          return { article: res.data }
        },
        () => {
          context.error({ statusCode: 400, message: "error" })
          return {}
        }
      )
  },
  methods: {
    formatTime (timeStr) {
      return moment(timeStr).format('YYYY-MM-DD HH:mm:ss')
    }
  }
}
</script>
