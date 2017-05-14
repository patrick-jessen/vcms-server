<template>
<div class='treetable'>
  <div v-if='!_level' class='header'>
    <div class='column' v-for='(h, i) in columnStyles' :style='h'>
      <div v-if='i < columnStyles.length - 1' class='column-resize' @mousedown='onResizeStart($event, i)'></div>
      {{columns[i].title}}
    </div>
  </div>
  <div class='body'>
    <TreeTable :data='root' :level='0' :columns='columns' @select='onSelect'/>
  </div>
</div>
</template>

<script>
export default {
  props: ['columns', 'test', 'level'],
  computed: {
    root() {
      //return window.vue.$children[0].component
      return this.test
    },
    columnStyles() {
      return this.columns.map((c) => {
        return 'width: ' + c.width + '%'
      })
    },
    _level() {
      if(typeof this._level === 'undefined')
        return 0
      return this._level
    },
  },
  methods: {
    
    onSelect(arg) {
      this.$emit('select', arg)
    },
    onResizeStart(e, index) {
      window.addEventListener('mousemove', this.onResize, false);
      window.addEventListener('mouseup', this.onResizeStop, false);
      this.lastMouse = e.x
      this.resizeIndex = index
    },
    onResize(e) {
      var dif = this.lastMouse - e.x
      var w = parseFloat(this.columns[this.resizeIndex].width)

      var absW = this.$el.clientWidth
      var f = 100 / absW
      w -= dif * f

      this.columns[this.resizeIndex].width = w
      this.lastMouse = e.x
    },
    onResizeStop(e) {
      window.removeEventListener('mousemove', this.onResize, false);
      window.removeEventListener('mouseup', this.onResizeStop, false);
    },
  },
  components: {
    TreeTableItem: require('./TreeTableItem.vue')
  }
}
</script>

<style lang='scss'>
.treetable {
  border-radius: 2px;
  background-color: white;
  text-align: left;
  overflow: auto;
  margin: 10px 0;

  height: 300px;

  &::-webkit-scrollbar {
    width: 0px;
  }

  & * {
    user-select: none;
  }
}

.header {
  border-bottom: 1px solid black;
  background-color: #42b883;
  color: #fcfcfc;
  font-weight: bold;
  overflow: hidden;
  height: 35px;
  cursor: default;

  & .column {
    display: inline-block;
    padding: 7px 10px;
    text-align: center;
    position: relative;
    overflow: hidden;

    & .column-resize {
      display: none;
      background-color: hsl(153, 47%, 40%);
      width: 2px;
      height: 20px;
      position: absolute;
      right: 0;
      top: 10%;
      height: 80%;
      cursor: ew-resize;
    }
  }

  &:hover .column-resize{
    display: block;
  }
}
</style>