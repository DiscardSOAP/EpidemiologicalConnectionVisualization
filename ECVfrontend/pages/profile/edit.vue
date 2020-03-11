<template>
  <v-card class="elevation-12">
    <v-toolbar
      color="primary"
      dark
      flat
    >
      <v-toolbar-title>Profile</v-toolbar-title>
      <v-spacer />
    </v-toolbar>
    <v-card-text>
      <v-form
        ref="form"
        v-model="profileForm.valid"
      >
        <v-container>
          <v-row>
            <v-col>
              <v-text-field
                label="Name"
                name="name"
                prepend-icon="mdi-account"
                type="text"
                v-model="profileForm.form.name"
                :counter="10"
                :rules="profileForm.nameRules"
              />
            </v-col>
            <v-col>
              <v-text-field
                label="Organization"
                name="organization"
                prepend-icon="mdi-account-multiple"
                type="text"
                v-model="profileForm.form.organization"
                :counter="20"
                :rules="profileForm.organizationRules"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-text-field
                label="New Password"
                name="newPassword"
                prepend-icon="mdi-lock-open"
                type="password"
                v-model="profileForm.form.newPassword"
                :counter="20"
                :rules="profileForm.newPasswordRules"
              />
            </v-col>
            <v-col>
              <v-text-field
                label="Confirm New Password"
                name="confirmNewPassword"
                prepend-icon="mdi-lock-open-outline"
                type="password"
                v-model="profileForm.form.confirmNewPassword"
                :counter="20"
                :rules="profileForm.confirmNewPasswordRules"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-textarea
                outlined
                name="description"
                label="Description"
                v-model="profileForm.form.description"
              ></v-textarea>
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-text-field
                id="password"
                label="Password"
                name="password"
                prepend-icon="mdi-lock"
                type="password"
                v-model="profileForm.form.password"
                :counter="20"
                :rules="profileForm.passwordRules"
                required
              />
            </v-col>
          </v-row>
        </v-container>
      </v-form>
      <v-alert
        dense
        outlined
        type="error"
        v-if="profileForm.error"
      >
        Change profile failed!
      </v-alert>
    </v-card-text>
    <v-card-actions>
      <v-btn
        text
        :color="profileForm.error?'error':'primary'"
        x-large
        block
        @click="submit"
        :disabled="!profileForm.valid||profileForm.error||profileForm.loading"
        :loading="profileForm.loading"
      >
        Submit
      </v-btn>
      <!--v-spacer />
            <v-btn color="primary">Login</v-btn-->
    </v-card-actions>
  </v-card>
</template>
<script>
export default {
  auth: "guest",
  data: function () {
    return {
      profileForm: {
        valid: true,
        form: {
          name: "",
          organization: "",
          description: "",
          password: "",
          newPassword: "",
          confirmNewPassword: "",
        },
        nameRules: [
          v => (v.length <= 10) || 'Name must be less than 10 characters',
        ],
        organizationRules: [
          v => (v.length <= 20) || 'Organization must be less than 20 characters',
        ],
        passwordRules: [
          v => !!v || 'Password is required.',
          v => (v && v.length <= 20) || 'Password must be less than 20 characters',
          v => /[a-zA-Z0-9].{7,}/.test(v) || 'Password must be valid'
        ],
        newPasswordRules: [
          v => (v.length <= 20) || 'New Password must be less than 20 characters',
          /*v => (/[a-zA-Z0-9].{7,}/.test(v)) || 'New Password must be valid'*/
        ],
        confirmNewPasswordRules: [
          v => (v.length <= 20) || 'Confirm New Password must be less than 20 characters',
          /*v => (!!v || v.length == 0 || /[a-zA-Z0-9].{7,}/.test(v)) || 'Confirm New Password must be valid'*/
        ],
        loading: false,
        ok: false,
        error: false
      },
    }
  },
  watch: {
    "profileForm.form.password": function () {
      this.profileForm.error = false
      this.profileForm.loading = false
      this.profileForm.error = false
    }
  },
  mounted () {
    this.profileForm.confirmNewPasswordRules.push(this.sameAsPassword)
    this.profileForm.newPasswordRules.push(this.nullOrValid)
    this.profileForm.confirmNewPasswordRules.push(this.nullOrValid)
  },
  methods: {
    sameAsPassword: function (value) {
      return value == this.profileForm.form.newPassword ? true : "Confirm New Password must be the same as New Password."
    },
    nullOrValid: function (value) {
      return value == "" ? true : (/[a-zA-Z0-9].{7,}/.test(value) ? true : "New Password must be valid")
    },
    submit: function () {
      this.profileForm.loading = true

      this.$axios.post('/api/profile/', this.profileForm.form)
        .then(() => {
          this.profileForm.error = false
          this.profileForm.ok = true
        })
        .catch(err => {
          this.profileForm.error = true
        })
        .then(() => {
          this.profileForm.loading = false
        })
    }
  }
}
</script>