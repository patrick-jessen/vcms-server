<template>
<div v-if='data' class='item-wrapper'>
  <div class='item' :class='{selected}' @dblclick='toggleExpansion' @click='onClick' @mouseenter='hover(true)' @mouseleave='hover(false)'>
    <div class='column' v-for='(c, i) in columns' :style='style(i)'>
       <i v-if='i === 0' class='expand fa' 
        :class='expandIcon' 
        @click='toggleExpansion' 
        @mouseenter='hoverExpansion(true)' 
        @mouseleave='hoverExpansion(false)'>
      </i>
      <div class='render' v-html='data.render(i)'></div>
    </div>
  </div>
  <div v-if='expanded'>
    <div v-for='c in data.children'>
      <TreeTableItem :data='c' :level='level+1' :columns='columns' @select='onSelect'/>
    </div>
  </div>
</div>
</template>

<script>
export default {
  props: ['data', 'level', 'columns'],
  data() {
    return {
      expanded: false,
      expansionHovered: false,
      selected: false
    }
  },
  beforeCreate() {
    this.$options.components.TreeTableItem = require('./TreeTableItem.vue')
  },
  mounted() {
    window.vcmsvue.$on('select', (ns) => {
      if(this.data.namespace.string === ns)
        this.selected = true
      else
        this.selected = false
    })
  },
  computed: {
    expandIcon() {
      if(this.data.children.length === 0)
        return 'fa-square-o'

      var str = ''
      if(this.expanded)
        str = 'fa-minus-square'
      else
        str = 'fa-plus-square'

      if(!this.expansionHovered)
        str += '-o'

      return str
    },
  },
  methods: {
    style(col) {
      var pad = ''
      if(col === 0)
        pad = '' + (this.level * 15 + 10) + 'px'

      return {
        'padding-left': pad,
        'width': '' + this.columns[col].width + '%'
      }
    },
    toggleExpansion() {
      this.expanded = !this.expanded
    },
    hoverExpansion(val) {
      this.expansionHovered = val
    },
    onClick(e) {
      e.stopPropagation()
      this.$emit('select', this.data.data.namespace.string)
    },
    onSelect(val) {
      this.$emit('select', val)
    },
    hover(enter) {
      if(enter)
        window.vcmsvue.$emit('hover', this.data.data.namespace.string)
      else
        window.vcmsvue.$emit('hover', '')
    }
  }
}
</script>

<style lang='scss'>
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
      cursor: pointer;
    }
    &.selected {
      background-color: #f5f5f5;
    }
    & .column {
      display: inline-block;
      padding: 5px;
      height: 29px;
      vertical-align: middle;

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