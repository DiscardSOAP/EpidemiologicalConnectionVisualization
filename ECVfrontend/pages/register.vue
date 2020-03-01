<template>
  <v-container
    class="fill-height"
    fluid
  >
    <v-row
      align="center"
      justify="center"
    >
      <v-col
        cols="12"
        sm="8"
        md="5"
      >
        <v-card class="elevation-12">
          <v-toolbar
            color="primary"
            dark
            flat
          >
            <v-toolbar-title>Register form</v-toolbar-title>
            <v-spacer />
          </v-toolbar>
          <v-card-text>
            <v-form>
              <v-text-field
                label="Username"
                name="username"
                prepend-icon="mdi-account"
                type="text"
              />

              <v-text-field
                id="password"
                label="Password"
                name="password"
                prepend-icon="mdi-lock"
                type="password"
              />
              <v-text-field
                id="confirm-password"
                label="Confirm password"
                name="confirm-password"
                prepend-icon="mdi-lock-outline"
                type="password"
              />
              <v-text-field
                id="invitation-code"
                label="Invitation code"
                name="invitation-code"
                prepend-icon="mdi-security"
                type="text"
              />
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn
              text
              color="primary"
              x-large
              block
            >
              Register
            </v-btn>
            <!--v-spacer />
            <v-btn color="primary">Login</v-btn-->
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  props: {
    source: String,
  },
}
</script>