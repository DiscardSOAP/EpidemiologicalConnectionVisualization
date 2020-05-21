<template>
  <v-card class="elevation-12">
    <v-toolbar
      color="primary"
      dark
      flat
    >
      <v-toolbar-title>个人信息</v-toolbar-title>
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
                label="昵称"
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
                label="所属机构"
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
                label="新密码"
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
                label="确认新密码"
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
                label="简介"
                v-model="profileForm.form.description"
              ></v-textarea>
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-text-field
                id="password"
                label="密码"
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
        修改个人信息失败，密码错误！
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
  data: function () {
    return {
      profileForm: {
        valid: true,
        form: {
          name: this.$auth.$state.user.name,
          organization: this.$auth.$state.user.organization,
          description: this.$auth.$state.user.description,
          password: "",
          newPassword: "",
          confirmNewPassword: "",
        },
        nameRules: [
          v => (v.length <= 10) || '昵称不能超过10个字符',
        ],
        organizationRules: [
          v => (v.length <= 20) || '所属机构名不能超过20个字符',
        ],
        passwordRules: [
          v => !!v || '修改个人信息应当输入旧密码确认',
          v => (v && v.length <= 20) || '密码不能超过20个字符',
          v => /[a-zA-Z0-9].{7,}/.test(v) || '密码格式不合法'
        ],
        newPasswordRules: [
          v => (v.length <= 20) || '新密码不能超过20个字符',
          /*v => (/[a-zA-Z0-9].{7,}/.test(v)) || 'New Password must be valid'*/
        ],
        confirmNewPasswordRules: [
          v => (v.length <= 20) || '确认新密码不能超过20个字符',
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
      return value == this.profileForm.form.newPassword ? true : "确认新密码与新密码不一致"
    },
    nullOrValid: function (value) {
      return value == "" ? true : (/[a-zA-Z0-9].{7,}/.test(value) ? true : "新密码格式不合法")
    },
    submit: function () {
      this.profileForm.loading = true

      this.$axios.post('/api/profile/', this.profileForm.form)
        .then(() => {
          this.profileForm.error = false
          this.profileForm.ok = true
          this.$auth.fetchUser()
          setTimeout(() => {
            this.$router.push({ path: '/profile/' })
          }, 1000);
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