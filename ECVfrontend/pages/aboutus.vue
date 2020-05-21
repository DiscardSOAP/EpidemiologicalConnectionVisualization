<template>
  <v-layout
    column
    justify-center
    align-center
  >
    <v-flex
      xs12
      sm8
      md6
    >
      <div class="text-center">
        <logo />
        <vuetify-logo />
      </div>
      <v-card>
        <v-card-title class="headline">
          欢迎使用由DiscardSOAP小组开发的流行病学联系可视化的网站
        </v-card-title>
        <v-card-text>
          <p>ECV:流行病学联系可视化是一个开源的项目，集新闻发布，疫情数据可视化，流行病学联系可视化和患者轨迹追踪于一体的，尝试解决疫情信息不直观，发布平台杂乱等现象，给民众提供一个统一的新闻发布平台。</p>
          <p>
            如果您有任何的问题，请联系我们 <a
              href="https://github.com/fnoi2014xtx"
              target="_blank"
              title="contribute"
            >
              fnoi2014xtx
            </a>, <a
              href="https://github.com/jkxing"
              target="_blank"
              title="contribute"
            >
              jkxing
            </a>
          </p>
          <p>
            如果您发现任何bug，欢迎在Github上提出 <a
              href="https://github.com/DiscardSOAP/EpidemiologicalConnectionVisualization/issues"
              target="_blank"
              title="contribute"
            >
              issue board
            </a>
          </p>
          <p>本网站前端使用Nuxt.js + Vuetify搭建，在此感谢两个项目对前端做出的贡献</p>
          <hr class="my-3">
          <a
            href="https://nuxtjs.org/"
            target="_blank"
          >
            Nuxt Documentation
          </a>
          <br>
          <a
            href="https://vuetifyjs.com/en/"
            target="_blank"
          >
            Vuetify Documentation
          </a>
        </v-card-text>
        <!--v-card-actions>
          <v-spacer />
          <v-btn
            color="primary"
            nuxt
            to="/inspire"
          >
            Continue
          </v-btn>
        </v-card-actions-->
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  auth: false,
  components: {
    Logo,
    VuetifyLogo
  }
}
</script>
