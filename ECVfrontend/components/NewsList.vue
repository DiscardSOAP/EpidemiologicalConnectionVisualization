<template>
  <div>
    <v-container>
      <v-row dense>

        <v-col
          cols="12"
          v-for="item in [0,1,2,3,4]"
          :key="item"
        >
          <v-card
            class="mx-auto"
            width="100%"
          >
            <v-card-title>
              <v-icon
                large
                left
              >
                mdi-bullhorn
              </v-icon>
              <span class="title font-weight">be•nev•o•lent</span>
            </v-card-title>

            <v-card-text class="headline font-weight-light text-justify">
              <p style="
                overflow : hidden;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-line-clamp: 3;
                -webkit-box-orient: vertical;">
                When I first started as a designer, I created designs that gave no consideration for the end user at all. Nor was I familiar with the intended business outcomes. Users? Who’s that? Lets just make this shit as cool as possible. We'll make this 3D cube spin and then explode out. It’ll be great. The companies that were actually thinking about users, were creating horrible looking work and just creating 100 page justification documents.
              </p>
            </v-card-text>
            <v-card-actions>
              <v-btn text>Jesus Brock</v-btn>

              <v-spacer></v-spacer>

              <v-btn
                color="primary"
                text
              >
                Posted on january 26, 2019
              </v-btn>

            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
      <div class="text-center">
        <v-pagination
          v-model="page"
          :length="15"
          :total-visible="10"
        ></v-pagination>
      </div>
    </v-container>
  </div>
</template>
<script>
export default {
  props: {

  },
  data: function () {
    return {
      page: 1
    }
  }

}
</script>
