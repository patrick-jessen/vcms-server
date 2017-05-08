<template>
<div>
  <h1>Hierarchy</h1>
  <div class='component-hierarchy'>
    <ComponentItem :data='app' :def='def' :level='0'/>
  </div>
</div>
</template>

<script>
export default {
  computed: {
    app() {
      return window.vue.$children[0].component // 'app' component
    },
    def() {
      return {
        name:'app', 
        type:'app',
        title:'App',
        descr:'App description'
      }
    }
  },
  components: {
    ComponentItem: require('./ComponentItem.vue')
  }
}
</script>

<style lang='scss'>
#vcms {
  & .component-hierarchy {
    background-color: white;
    padding: 10px;
    border: 1px solid gray;
    text-align: left;
    height: 200px;
    overflow: auto;
  }
}
</style>