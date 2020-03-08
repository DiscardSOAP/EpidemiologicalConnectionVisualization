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
            <v-form
              ref="form"
              v-model="registerForm.valid"
            >
              <v-text-field
                label="Username"
                name="username"
                prepend-icon="mdi-account"
                type="text"
                v-model="registerForm.form.username"
                :counter="20"
                :rules="registerForm.usernameRules"
                required
              />

              <v-text-field
                id="password"
                label="Password"
                name="password"
                prepend-icon="mdi-lock"
                type="password"
                v-model="registerForm.form.password"
                :counter="20"
                :rules="registerForm.passwordRules"
                required
              />
              <v-text-field
                id="confirm-password"
                label="Confirm password"
                name="confirm-password"
                prepend-icon="mdi-lock-outline"
                type="password"
                v-model="registerForm.form.confirmPassword"
                :counter="20"
                :rules="registerForm.confirmPasswordRules"
                required
              />
              <v-text-field
                id="invitation-code"
                label="Invitation code"
                name="invitation-code"
                prepend-icon="mdi-security"
                type="text"
                v-model="registerForm.form.invitationCode"
                :counter="20"
                :rules="registerForm.invitationCodeRules"
                required
              />
            </v-form>
            <v-alert
              dense
              outlined
              type="error"
              v-if="registerForm.error"
            >
              Register Failed!
            </v-alert>
          </v-card-text>
          <v-card-actions>
            <v-btn
              text
              :color="registerForm.error?'error':'primary'"
              x-large
              block
              @click="register"
              :disabled="!registerForm.valid||registerForm.error||registerForm.loading"
              :loading="registerForm.loading"
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
  auth: "guest",
  data: function () {
    return {
      registerForm: {
        valid: true,
        form: {
          username: "",
          password: "",
          confirmPassword: "",
          invitationCode: "",
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
        confirmPasswordRules: [
          v => !!v || 'Confirm Password is required.',
          v => (v && v.length <= 20) || 'Confirm Password must be less than 20 characters',
          v => /[a-zA-Z0-9].{7,}/.test(v) || 'Confirm Password must be valid'
        ],
        invitationCodeRules: [
          v => !!v || 'Invitation Code is required.',
          v => (v && v.length == 20) || 'Invitation Code must be valid.'
        ],
        loading: false,
        ok: false,
        erorr: false
      }
    }
  },
  watch: {
    "registerForm.form.username": function () {
      this.registerForm.error = false
      this.registerForm.loading = false
      this.registerForm.error = false
    },
    "registerForm.form.password": function () {
      this.registerForm.error = false
      this.registerForm.loading = false
      this.registerForm.error = false
    },
    "registerForm.form.invitationCode": function () {
      this.registerForm.error = false
      this.registerForm.loading = false
      this.registerForm.error = false
    }
  },
  mounted () {
    this.registerForm.confirmPasswordRules.push(this.sameAsPassword)
  },
  methods: {
    sameAsPassword: function (value) {
      return value == this.registerForm.form.password ? true : "Confirm Password must be the same as Password."
    },
    register: function () {
      this.registerForm.loading = true
      this.$axios
        .post("/api/register/", this.registerForm.form)
        .then(
          () => {
            this.registerForm.error = false
            this.registerForm.ok = true
            setTimeout(() => {
              this.$router.push("/login/")
            }, 500)
          },
          err => {
            this.registerForm.error = true
          }
        )
        .then(() => {
          this.registerForm.loading = false
        })
    }
  },

}
</script>