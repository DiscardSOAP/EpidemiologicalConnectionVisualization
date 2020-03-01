<template>
  <v-card>
    <v-card-title>
      <v-text-field
        v-model="search"
        append-icon="mdi-account-search"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
    <v-data-table
      v-model="selected"
      :headers="headers"
      :items="desserts"
      :search="search"
      item-key="name"
      show-select
      class="elevation-1"
    >
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  data () {
    return {
      search: '',
      headers: [
        { text: 'Username', value: 'username' },
        { text: 'Last login', value: 'last' },
        { text: 'Organization', value: 'organization' },
        { text: 'Description', value: 'description' },
        { text: 'Distribution', value: 'distribution' },
      ],
      desserts: [
        {
          username: 'Frozen Yogurt',
          last: '2020-3-2',
          organization: 'Tianjin Gov.',
          description: 'A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people',
          distribution: '4%'
        },
        {
          username: 'Frozen Yogurt',
          last: '2020-3-2',
          organization: 'Tianjin Gov.',
          description: 'A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people',
          distribution: '4%'
        },
        {
          username: 'Frozen Yogurt',
          last: '2020-3-2',
          organization: 'Tianjin Gov.',
          description: 'A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people',
          distribution: '4%'
        },
        {
          username: 'Frozen Yogurt',
          last: '2020-3-2',
          organization: 'Tianjin Gov.',
          description: 'A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people',
          distribution: '4%'
        },
        {
          username: 'Frozen Yogurt',
          last: '2020-3-2',
          organization: 'Tianjin Gov.',
          description: 'A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people',
          distribution: '4%'
        },
        {
          username: 'Frozen Yogurt',
          last: '2020-3-2',
          organization: 'Tianjin Gov.',
          description: 'A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people A doctor from Tianjin Cancer Hospital, 42 years old, male, keen on saving other people',
          distribution: '4%'
        },
      ],
    }
  },
}
</script>