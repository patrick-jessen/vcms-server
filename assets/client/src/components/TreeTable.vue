<template>
<div>
  <!-- Base component -->
  <template v-if='typeof level === "undefined"'>
    <div class='treetable'>
      <div class='header'>
        <div class='column' v-for='(h, i) in columnStyles' :style='h'>
          <div v-if='i < columnStyles.length - 1' class='column-resize' @mousedown='onResizeStart($event, i)'></div>
          {{columns[i].title}}
        </div>
      </div>
      <div class='body'>
        <TreeTable :data='data' :level='0' :columns='columns' :skipRoot='skipRoot'/>
      </div>
    </div>
  </template>

  <!-- Tree items -->
  <template v-else>
    <div class='item-wrapper'>
      <div class='item' 
        :class='{selected}'
        @click='onClick'
        @dblclick='onToggleExpansion'
        @mouseenter='onHover(true)' 
        @mouseleave='onHover(false)'
        v-if='!skipRoot'>

        <div class='column' v-for='(c, i) in columns' :style='columnStyles[i]'>
          <div class='column-content'>
            <i v-if='i === 0' class='expand' 
              :class='expandIcon' 
              @click.stop='onToggleExpansion' 
              @mouseenter='onHoverExpansion(true)' 
              @mouseleave='onHoverExpansion(false)'>
            </i>
            <div class='render' v-html='data.render(i)'></div>
          </div>
        </div>
      </div>
      <template v-if='skipRoot'>
        <TreeTable v-for='c in data.children' :data='c' :level='level' :columns='columns'/>
      </template>
      <template v-else-if='expanded'>
        <TreeTable v-for='c in data.children' :data='c' :level='level+1' :columns='columns'/>
      </template>
    </div>
  </template>
</div>
</template>

<script>
export default {
  props: ['columns', 'data', 'level', 'skipRoot'],
  data() {
    return {
      expanded: false,
      selected: false,
      expansionHovered: false,
    }
  },
  beforeCreate() {
    this.$options.components.TreeTable = require('./TreeTable.vue')
  },
  computed: {
    expandIcon() {
      if(this.data.children.length === 0)
        return 'fa fa-square-o'

      var str = ''
      if(this.expanded)
        str = 'fa fa-minus-square'
      else
        str = 'fa fa-plus-square'

      if(!this.expansionHovered)
        str += '-o'

      return str
    },
    columnStyles() {
      var arr = []
      for(var i = 0; i < this.columns.length; i++) {
        var pad = ''
        if(i === 0)
          pad = '' + (this.level * 15 + 10) + 'px'

        arr.push({
          'padding-left': pad,
          'width': '' + this.columns[i].width + '%'
        })
      }

      return arr
    },
  },
  methods: {
    onToggleExpansion() {
      this.expanded = !this.expanded
    },
    onClick(e) {
      console.log('Click', this.data.identifier)
    },
    onHover(enter) {
      if(enter)
        console.log('Hover', this.data.identifier)
    },
    onHoverExpansion(enter) {
      this.expansionHovered = enter
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

      // Adjust neighbor
      if(this.columns.length >= this.resizeIndex + 1) {
        this.columns[this.resizeIndex+1].width += dif * f
      }
      else {
        this.columns[this.resizeIndex-1].width += dif * f
      }

      this.columns[this.resizeIndex].width = w
      this.lastMouse = e.x
    },
    onResizeStop(e) {
      window.removeEventListener('mousemove', this.onResize, false);
      window.removeEventListener('mouseup', this.onResizeStop, false);
    },
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

  height: 225px;

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

.item-wrapper {
  padding: 0;

  & .item {
    border-bottom: 1px solid lightgray;
    color: #444;
    cursor: default;
    height: 30px;
    overflow: hidden;

    &:hover {
      background-color: #fafafa;
    }
    & .expand {
      vertical-align: middle;
      margin-right: 5px;
      width: 12px;
      height: 14px;
      color: #42b883;

      &:not(.fa-square-o) {
        cursor: pointer;
      }
    }
    &.selected {
      background-color: #f5f5f5;
    }
    & .column {
      display: inline-block;
      padding: 5px;
      height: 29px;
      vertical-align: middle;
      overflow: hidden;

      & .column-content {
        width: 1000px;
        height: 100%;
      }

      & .render {
        display: inline-block;
        height: 100%;

        & * {
          vertical-align: middle;
          max-height: 100%;
          width: auto;
          max-width: 100%;
          height: auto;
        }
      }
    }
  }
}
</style>