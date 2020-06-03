<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :clipped="true"
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>

      </v-list>

      <template v-slot:append>
        <blockquote class="blockquote"><span>&copy; {{ new Date().getFullYear() }}</span> </blockquote>

      </template>
    </v-navigation-drawer>
    <v-app-bar
      :clipped-left="true"
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-btn
        text
        v-if="!$auth.loggedIn"
        to="/login/"
      >
        登录
      </v-btn>
      <v-btn
        text
        v-if="!$auth.loggedIn"
        to="/register/"
      >

        注册
      </v-btn>
      <v-avatar v-if="$auth.loggedIn">
        <img :src="'https://cdn.v2ex.com/gravatar/'
          +
          $auth.$state.user.email_md5
          + '?s=1000&d=mm'">

      </v-avatar>
      <v-menu
        offset-y
        v-if="$auth.loggedIn"
      >
        <template v-slot:activator="{ on }">
          <v-btn
            color="primary"
            text
            v-on="on"
          >
            {{$auth.$state.user.username}}
          </v-btn>
          <!-- <v-avatar>
            <img
              src="https://cdn.vuetifyjs.com/images/john.jpg"
              alt="John"
              v-on="on"
            >

          </v-avatar>
          <container>
          {{$auth.$state.user.username}} -->
        </template>

        <v-list>
          <v-list-item-group>
            <v-list-item @click="$auth.logout()">
              <v-list-item-icon>
                <v-icon>mdi-logout</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>登出</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-list-item @click="$router.push({ path: '/profile/' })">
              <v-list-item-icon>
                <v-icon>mdi-account-circle</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>个人主页</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
          <!-- <v-list-item >
            <v-list-item-title>LOGOUT</v-list-item-title>
          </v-list-item>
          <v-list-item @click="$router.push({ path: '/profile/' })">
            <v-list-item-title>PROFILE</v-list-item-title>
          </v-list-item> -->
        </v-list>
      </v-menu>

    </v-app-bar>
    <v-content>
      <v-container style="height:100%">
        <nuxt />
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  transition: { name: "page", mode: "out-in" },
  middleware: ["auth"],
  mounted () {
    if (this.$auth.loggedIn) {
      this.items.push({
        icon: 'mdi-key',
        title: '管理用户',
        to: '/manage/'
      })
      this.items.push({
        icon: 'mdi-border-color',
        title: '发布新闻',
        to: '/news/edit/'
      })
    }
  },
  watch: {
    "$auth.loggedIn": function () {
      /*console.log("shit!")*/
      if (this.$auth.loggedIn) {
        this.items.push({
          icon: 'mdi-key',
          title: '管理用户',
          to: '/manage/'
        })
        this.items.push({
          icon: 'mdi-border-color',
          title: '发布新闻',
          to: '/news/edit/'
        })
      } else {
        this.items.pop()
        this.items.pop()
      }
    }
  },
  data () {
    return {
      drawer: false,
      userMenu: false,
      items: [
        {
          icon: 'mdi-apps',
          title: '首页',
          to: '/'
        },
        {
          icon: 'mdi-newspaper',
          title: '新闻',
          to: '/news/'
        },
        {
          icon: 'mdi-google-maps',
          title: '中国疫情',
          to: '/topology/'
        },
        {
          icon: 'mdi-information-outline',
          title: '关于我们',
          to: '/aboutus/'
        }
      ],
      title: 'ECV : 流行病学联系可视化'
    }
  }
}
</script>
