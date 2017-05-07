<template>
  <div id='vcms'>
   <Test/>
  </div>
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  components: {
    Test: require('./components/Test.vue')
  }
}
</script>

<style lang="scss">
</style>
