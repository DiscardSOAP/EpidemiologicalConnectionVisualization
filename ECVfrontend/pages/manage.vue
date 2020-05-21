<template>
  <v-card>
    <v-card-title>
      所有用户
    </v-card-title>
    <v-card-subtitle>
      <v-text-field
        v-model="searchKeyWord"
        append-icon="mdi-account-search"
        label="搜索用户"
        single-line
        hide-details
      >

      </v-text-field>
    </v-card-subtitle>
    <v-data-table
      v-model="selectedUsers"
      :headers="headers"
      :items="users"
      :search="searchKeyWord"
      item-key="username"
      show-select
      class="elevation-1"
      show-expand
      single-expand
      :expanded.sync="expanded"
    >
      <template v-slot:expanded-item="{ headers, item }">
        <td :colspan="headers.length">{{item.description}}</td>
      </template>
    </v-data-table>
    <v-card-actions>
      <v-btn
        text
        color='red'
        :disabled='selectedUsers.length == 0'
        @click.stop="deleteDialog = true"
      >Delete</v-btn>
      <v-dialog
        v-model="deleteDialog"
        max-width="60%"
        scrollable
      >
        <v-card>
          <v-card-title
            class="headline"
            style="color:red"
          >警告</v-card-title>

          <v-card-text>
            是否删除以下选中用户？
          </v-card-text>
          <v-card-content>
            <v-list three-line>
              <template v-for="(user, index) in selectedUsers">

                <v-list-item :key="index">
                  <v-list-item-avatar>
                    <v-img :src="'https://cdn.v2ex.com/gravatar/'
          +
          user.email_md5
          + '?s=1000&d=mm'"></v-img>
                  </v-list-item-avatar>

                  <v-list-item-content>
                    <v-list-item-title v-html="user.username"></v-list-item-title>
                    <v-list-item-subtitle v-html="user.email"></v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
              </template>
            </v-list>
          </v-card-content>
          <v-card-actions>
            <v-spacer></v-spacer>

            <v-btn
              color="primary darken-1"
              text
              @click="deleteDialog = false"
            >
              Disagree
            </v-btn>

            <v-btn
              color="red darken-1"
              text
              @click="deleteSelectedUser()"
            >
              Agree
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data () {
    return {
      searchKeyWord: '',
      deleteDialog: false,
      selectedUsers: [],
      headers: [
        { text: '用户名', value: 'username' },
        { text: '邮箱', value: 'email' },
        { text: '昵称', value: 'name' },
        { text: '所属机构', value: 'organization' },
        //{ text: 'Distribution', value: 'distribution' },
        { text: '简介', value: 'data-table-expand' },
      ],
      users: [],
    }
  },
  asyncData (context) {
    return context.app.$axios
      .get(`/api/root/manage/`)
      .then(
        res => {
          let users = res.data.users
          //console.log(users)
          return { users }
        },
        () => {
          context.error({ statusCode: 401, message: "No authority!" })
          return {}
        }
      )
  },
  methods: {
    deleteSelectedUser: function () {
      this.deleteDialog = false
      this.$axios.delete('/api/root/manage/', { users: this.selectedUsers }).then(
        () => {
          this.users = this.users.filter(user => {
            let flag = true
            for (var targetUser in this.selectedUsers) {
              if (targetUser.username == user.username) {
                flag = false
                break
              }
            }
            return flag
          })
        }
      )
    },
  }
}
</script>