<template>
  <h1>{{ text }}</h1>
</template>

<script>
export default {
  asyncData(context) {
    return context.$axios.get('/api/helloworld/').then((res) => {
      const text = res.data.msg
      return { text }
    })
  },
  data() {
    return {
      text: ''
    }
  }
}
</script>
