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
            <v-toolbar-title>注册</v-toolbar-title>
            <v-spacer />
          </v-toolbar>
          <v-card-text>
            <v-form
              ref="form"
              v-model="registerForm.valid"
            >
              <v-text-field
                label="邮箱"
                name="email"
                prepend-icon="mdi-account"
                type="text"
                v-model="registerForm.form.email"
                :counter="50"
                :rules="registerForm.emailRules"
                required
              />

              <v-text-field
                label="用户名"
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
                label="密码"
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
                label="确认密码"
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
                label="邀请码"
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
              注册失败!
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
          email: "",
          confirmPassword: "",
          invitationCode: "",
        },
        usernameRules: [
          v => !!v || '用户名不能为空',
          v => (v && v.length <= 20) || '用户名不能超过20个字符',
          v => /[a-zA-Z0-9].{7,}/.test(v) || '用户名格式不合法',
        ],
        emailRules: [
          v => !!v || '邮箱不能为空',
          v => (v && v.length <= 50) || '邮箱不能超过50个字符',
          v => /.+@.+\..+/.test(v) || '邮箱格式不合法',
        ],
        passwordRules: [
          v => !!v || '密码不能为空',
          v => (v && v.length <= 20) || '密码不能超过20个字符',
          v => /[a-zA-Z0-9].{7,}/.test(v) || '密码格式不合法'
        ],
        confirmPasswordRules: [
          v => !!v || '请再输入一遍密码',
          v => (v && v.length <= 20) || '确认密码不能超过20个字符',
          v => /[a-zA-Z0-9].{7,}/.test(v) || '确认密码格式不合法'
        ],
        invitationCodeRules: [
          v => !!v || '邀请码不能为空',
          v => (v && v.length == 20) || '邀请码格式不合法'
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
      return value == this.registerForm.form.password ? true : "确认密码与密码不一致"
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
            }, 1000)
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