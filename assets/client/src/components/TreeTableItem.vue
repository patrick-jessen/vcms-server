<template>
<div class='item-wrapper'>
  <div class='item' :class='{selected}' @dblclick='toggleExpansion' @click='select'>
    <div class='column' v-for='(c, i) in columns' :style='style(i)'>
       <i v-if='i === 0' class='expand fa' 
        :class='expandIcon' 
        @click='toggleExpansion' 
        @mouseenter='hoverExpansion(true)' 
        @mouseleave='hoverExpansion(false)'>
      </i>
      {{c.func(data)}}
    </div>
  </div>
  <div v-if='expanded'>
    <div v-for='ck in childKeys'>
      <TreeTableItem :data='childData(ck)' :level='level+1' :columns='columns'/>
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
      if(this.childKeys.length === 0)
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
    childKeys() {
      return Object.keys(this.data.children)
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
    childData(key) {
      return this.data.children[key]
    },
    toggleExpansion() {
      this.expanded = !this.expanded
    },
    hoverExpansion(val) {
      this.expansionHovered = val
    },
    select(e) {
      window.vcmsvue.$emit('select', this.data.namespace.string)
      e.stopPropagation()
      e.preventDefault()
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
    }
  }
}
</style>