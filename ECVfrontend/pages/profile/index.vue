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
          :loading="loading"
          class="mx-auto my-4"
        >
          <v-hover v-slot:default="{ hover }">
            <v-img src="https://cdn.vuetifyjs.com/images/cards/cooking.png">
              <v-expand-transition>
                <div
                  v-if="hover"
                  class="d-flex transition-fast-in-fast-out black darken-2 v-card--reveal"
                  style="height: 100%;opacity:0.8;"
                >
                  <v-btn
                    color="primary"
                    style="position: absolute;top:50%;left:50%;transform:translate(-50%,-50%);"
                    @click="reserve"
                  >
                    Reserve
                  </v-btn>
                </div>
              </v-expand-transition>
            </v-img>
          </v-hover>
          <v-card-title>Cafe Badilico</v-card-title>
          <v-card-subtitle class="pb-0">Number 10</v-card-subtitle>
          <v-card-text>
            <div class="my-1 subtitle-1">
              $ • Italian, Cafe
            </div>

            <div>Small plates, salads & sandwiches - an intimate setting with 12 indoor seats plus patio seating.</div>
          </v-card-text>

          <v-divider class="mx-4"></v-divider>

          <v-card-title>Tonight's availability</v-card-title>

          <v-card-text>
            <v-chip-group
              v-model="selection"
              active-class="deep-purple accent-4 white--text"
              column
            >
              <v-chip>5:30PM</v-chip>

              <v-chip>7:30PM</v-chip>

              <v-chip>8:00PM</v-chip>

              <v-chip>9:00PM</v-chip>
            </v-chip-group>
          </v-card-text>

          <v-card-actions>
            <v-btn
              color="deep-purple lighten-2"
              text
              @click="reserve"
            >
              Reserve
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
    }
  },
}
</script>
