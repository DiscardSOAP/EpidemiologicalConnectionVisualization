<template>
  <v-app dark>
    <v-layout>
      <v-flex class="text-center">

        <v-alert
          prominent
          type="error"
          outlined
        >
          <v-row align="center">
            <v-col
              class="grow display-1 font-weight-bold"
              v-if="error.statusCode===404"
            >
              {{pageNotFound}}
            </v-col>
            <v-col
              class="grow display-1 font-weight-bold"
              v-else
            >
              {{ otherError }}
            </v-col>
            <v-col class="shrink">
              <v-btn>
                Return
                <v-icon
                  right
                  dark
                >mdi-reply</v-icon>
              </v-btn>
            </v-col>
          </v-row>
        </v-alert>
      </v-flex>
    </v-layout>
    <!--h1 v-if="error.statusCode === 404">
      {{ pageNotFound }}
    </h1>
    <h1 v-else>
      {{ otherError }}
    </h1-->

  </v-app>
</template>

<script>
export default {
  layout: 'empty',
  props: {
    error: {
      type: Object,
      default: null
    }
  },
  data () {
    return {
      pageNotFound: '404 Not Found',
      otherError: 'An error occurred'
    }
  },
  head () {
    const title =
      this.error.statusCode === 404 ? this.pageNotFound : this.otherError
    return {
      title
    }
  }
}
</script>

<style scoped>
h1 {
  font-size: 20px;
}
</style>
