<template>
  <v-container>
    <v-row>
      <v-col
        cols="12"
        md="8"
      >
      </v-col>
      <v-col
        cols="6"
        md="4"
      >
        <v-card class="mx-auto">
          <v-card-text>
            <div>组织</div>
            <p class="display-1 text--primary">
              昵称
            </p>
            <p>用户名</p>
            <div class="text--primary">
              自述
            </div>
          </v-card-text>
          <v-card-actions>
            <v-btn
              text
              color="deep-purple accent-4"
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
export default {
  auth: "guest",
  data: function () {
    return {
    }
  },
}
</script>
