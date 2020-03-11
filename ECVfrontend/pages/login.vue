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
            <v-toolbar-title>Login form</v-toolbar-title>
            <v-spacer />
            <v-tooltip right>
              <template v-slot:activator="{ on }">
                <v-btn
                  icon
                  large
                  href="/register/"
                  v-on="on"
                >
                  <v-icon>mdi-account-plus</v-icon>
                </v-btn>
              </template>
              <span>Register</span>
            </v-tooltip>
          </v-toolbar>
          <v-card-text>
            <v-form
              ref="form"
              v-model="loginForm.valid"
            >
              <v-text-field
                label="Username"
                name="username"
                prepend-icon="mdi-account"
                type="text"
                v-model="loginForm.form.username"
                :counter="20"
                :rules="loginForm.usernameRules"
                required
              />

              <v-text-field
                id="password"
                label="Password"
                name="password"
                prepend-icon="mdi-lock"
                type="password"
                v-model="loginForm.form.password"
                :counter="20"
                :rules="loginForm.passwordRules"
                required
              />
            </v-form>
            <v-alert
              dense
              outlined
              type="error"
              v-if="loginForm.error"
            >
              Username or Password mismatch!
            </v-alert>
          </v-card-text>
          <v-card-actions>
            <v-btn
              text
              :color="loginForm.error?'error':'primary'"
              x-large
              block
              @click="login"
              :disabled="!loginForm.valid||loginForm.error||loginForm.loading"
              :loading="loginForm.loading"
            >
              Login
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
  auth: "guest",
  data: function () {
    return {
      loginForm: {
        valid: true,
        form: {
          username: "",
          password: ""
        },
        usernameRules: [
          v => !!v || 'Username is required.',
          v => (v && v.length <= 20) || 'Username must be less than 20 characters',
          v => /[a-zA-Z0-9].{7,}/.test(v) || 'Username must be valid',
        ],
        passwordRules: [
          v => !!v || 'Password is required.',
          v => (v && v.length <= 20) || 'Password must be less than 20 characters',
          v => /[a-zA-Z0-9].{7,}/.test(v) || 'Password must be valid'
        ],
        loading: false,
        ok: false,
        error: false
      },
    }
  },
  watch: {
    "loginForm.form.username": function () {
      this.loginForm.error = false
      this.loginForm.loading = false
      this.loginForm.error = false
    },
    "loginForm.form.password": function () {
      this.loginForm.error = false
      this.loginForm.loading = false
      this.loginForm.error = false
    }
  },
  methods: {
    login: function () {
      this.loginForm.loading = true

      this.$auth
        .loginWith("local", {
          data: this.loginForm.form
        })
        .then(() => {
          this.loginForm.error = false
          this.loginForm.ok = true
          this.$router.push({ path: '/profile/' })
        })
        .catch(err => {
          this.loginForm.error = true
        })
        .then(() => {
          this.loginForm.loading = false
        })
    }
  }
}
</script>