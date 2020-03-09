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
      <v-card flat>
        <v-card-text>
          <p class="text-sm-center font-weight-light subtitle-1">Word of the Day</p>
          <p class="display-4 text--primary text-center">
            be•nev•o•lent
          </p>
          <p class="text-lg-center font-weight-light subtitle-1">
            Posted by Jesus Brock on january 26, 2019
          </p>
        </v-card-text>
        <v-card-text
          class="text-justify text--primary font-weight-thin headline"
          style="width:85%; margin:0 auto;"
        >
          <p>When I first started as a designer, I created designs that gave no consideration for the end user at all. Nor was I familiar with the intended business outcomes. Users? Who’s that? Lets just make this shit as cool as possible. We'll make this 3D cube spin and then explode out. It’ll be great. The companies that were actually thinking about users, were creating horrible looking work and just creating 100 page justification documents. I'm not sure they were really focused on servicing people. There weren't many companies that excelled at thinking holistically at designing for the web. It was an era of marketing sites and doing neat micro sites in flash. How designers and agencies could out do one another by having as much gratuitous animation as possible and try to win awards. Don’t get me wrong, it was a fun time and I learned a ton during this period, but over the years my focus has been less and less about that, and more about creating outcomes with good design.</p>
          <p>
            But what is good design you ask? In my opinion, it is work that meets and exceeds the needs and expectations of the end user, whilst delivering on business objectives and is beautiful and highly considered. There are many moving parts that go into creating these outcomes and to create good work. Unfortunately, there's too much to write about in just one post and it's not what I intend to write about today, but merely a segue to a small tool that we can add to our designers toolkit to help us get there. Something that can gauge the activities of users in a live environment and see how they are interacting with our work. And that is by using heat maps. Heat maps give you a page by page breakdown of where users are clicking and tapping. It provides us with a ton of data that we can extrapolate and make informed decisions about our design. Here are some of the benefits of using heat maps
          </p>
          <p>
            But what is good design you ask? In my opinion, it is work that meets and exceeds the needs and expectations of the end user, whilst delivering on business objectives and is beautiful and highly considered. There are many moving parts that go into creating these outcomes and to create good work. Unfortunately, there's too much to write about in just one post and it's not what I intend to write about today, but merely a segue to a small tool that we can add to our designers toolkit to help us get there. Something that can gauge the activities of users in a live environment and see how they are interacting with our work. And that is by using heat maps. Heat maps give you a page by page breakdown of where users are clicking and tapping. It provides us with a ton of data that we can extrapolate and make informed decisions about our design. Here are some of the benefits of using heat maps
          </p>
          <p>
            Sed non neque elit. Sed ut imperdiet nisi. Proin condimentum fermentum nunc. Etiam pharetra, erat sed fermentum feugiat, velit mauris egestas quam, ut aliquam massa nisl quis neque. Suspendisse in orci enim.
          </p>
          <p>Thank you for developing with Vuetify and I look forward to bringing more exciting features in the future.</p>

        </v-card-text>
        <hr class="my-1">
        <v-card-actions>
          <div class="my-2">
            <v-btn
              large
              text
            >
              <v-icon>mdi-arrow-left</v-icon>
              Where is this ship sailing to?

            </v-btn>
          </div>
          <v-spacer />
          <div class="my-2">
            <v-btn
              large
              text
            >
              Where is this ship sailing to?
              <v-icon>mdi-arrow-right</v-icon>
            </v-btn>
          </div>
        </v-card-actions>
      </v-card>
    </v-flex>
  </v-layout>

</template>
<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  components: {
    Logo,
    VuetifyLogo
  }
}
</script>
