<template>
  <v-container>
    <v-dialog
      max-width="60%"
      scrollable
    >
      <template v-slot:activator="{ on }">
        <v-btn
          dark
          fab
          bottom
          right
          fixed
          x-large
          color="primary"
          v-on="on"
        >
          <v-icon>mdi-filter-outline</v-icon>
        </v-btn>

      </template>
      <v-card>
        <v-card-title>Keywords</v-card-title>
        <v-divider></v-divider>
        <v-card-text style="height: 300px;">
          <v-treeview
            v-model="selection"
            :items="items"
            selection-type="leaf"
            selectable
            return-object
            open-all
          ></v-treeview>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-text>
          <div class="center">
            <v-btn
              v-for="node in selection"
              :key="node.id"
              rounded
              outlined
              color="primary"
              small
              class="ma-1"
              @click.native="remove_keyword(node)"
            >
              <v-icon left>mdi-minus</v-icon> {{node.name}}
            </v-btn>
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer>
          </v-spacer>
          <v-btn text>
            Confirm
          </v-btn>
          <v-btn text>
            Cancel
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <div>
      <v-btn
        value="left"
        v-for="node in selection"
        :key="node.id"
        text
        large
        @click.native="remove_keyword(node)"
      >
        <v-icon left>mdi-minus</v-icon>
        {{node.name}}
      </v-btn>
    </div>
    <v-row dense>

      <v-col
        cols="12"
        v-for="item in [0,1,2,3,4]"
        :key="item"
      >
        <v-card
          class="mx-auto"
          width="100%"
        >
          <v-card-title>
            <v-icon
              large
              left
            >
              mdi-bullhorn
            </v-icon>
            <span class="title font-weight">be•nev•o•lent</span>
          </v-card-title>

          <v-card-text class="headline font-weight-light text-justify">
            <p style="
                overflow : hidden;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-line-clamp: 3;
                -webkit-box-orient: vertical;">
              When I first started as a designer, I created designs that gave no consideration for the end user at all. Nor was I familiar with the intended business outcomes. Users? Who’s that? Lets just make this shit as cool as possible. We'll make this 3D cube spin and then explode out. It’ll be great. The companies that were actually thinking about users, were creating horrible looking work and just creating 100 page justification documents.
            </p>
          </v-card-text>
          <v-card-actions>
            <v-btn text>Jesus Brock</v-btn>

            <v-spacer></v-spacer>

            <v-btn
              color="primary"
              text
            >
              Posted on january 26, 2019
            </v-btn>

          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <div class="text-center">
      <v-pagination
        v-model="page"
        :length="15"
        :total-visible="10"
      ></v-pagination>
    </div>

  </v-container>
</template>
<script>
export default {
  data () {
    return {
      page: 1,
      selection: [],
      items: [
        {
          id: 1,
          name: 'Root',
          children: [
            { id: 2, name: 'Child #1' },
            { id: 3, name: 'Child #2' },
            {
              id: 4,
              name: 'Child #3',
              children: [
                { id: 5, name: 'Grandchild #1' },
                { id: 6, name: 'Grandchild #2' },
                { id: 7, name: 'Grandchild #1' },
                { id: 8, name: 'Grandchild #2' },
                { id: 9, name: 'Grandchild #1' },
                { id: 10, name: 'Grandchild #2' },
                { id: 11, name: 'Grandchild #1' },
                { id: 12, name: 'Grandchild #2' },
              ],
            },
          ],
        },
      ],
    }
  },
  methods: {
    remove_keyword (nd) {
      this.selection = this.selection.filter(node => {
        return node.id != nd.id
      })
    }
  }
}
</script>