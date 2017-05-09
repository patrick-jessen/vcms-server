<template>
<div class='treetable'>
  <div class='header'>
    <div class='column' v-for='(h, i) in columnStyles' :style='h'>{{columns[i].title}}</div>
  </div>
  <div class='body'>
    <TreeTableItem :data='root' :level='0' :columns='columns'/>
  </div>
</div>
</template>

<script>
export default {
  computed: {
    root() {
      return window.vue.$children[0].component
    },
    columnStyles() {
      return this.columns.map((c) => {
        return 'width: ' + c.width + '%'
      })
    },
    columns() {
      return [
        {
          title: 'Component', 
          width: 60, 
          func: (i) => {
            if(i.parentDef.title) return i.parentDef.title
            return i.parentDef.name
          }
        },
        {
          title: 'Type', 
          width: 40,
          style: {
            'text-align': 'right'
          },
          func: (i) => {
           return i.type
          }
        }
      ]
    }
  },
  components: {
    TreeTableItem: require('./TreeTableItem.vue')
  }
}
</script>

<style lang='scss'>
.treetable {
  border-radius: 4px;
  background-color: white;
  text-align: left;
  overflow: auto;
  margin: 10px 0;

  height: 300px;
}

.header {
  border-bottom: 1px solid black;
  background-color: #42b883;
  color: #fcfcfc;
  font-weight: bold;

  & .column {
    display: inline-block;
    padding: 7px 10px;
    text-align: center;
  }
}
</style>