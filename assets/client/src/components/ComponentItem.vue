<template>
<div class='wrapper'>
  <div class='item' :style='style' :class='{selected}' @dblclick='toggleExpansion' @click='select'>
    <i class='fa' 
      :class='expandIcon' 
      @click='toggleExpansion' 
      @mouseenter='hoverExpansion(true)' 
      @mouseleave='hoverExpansion(false)'>
    </i>
    {{title}}
  </div>
  <div v-if='expanded'>
    <div v-for='ck in childKeys'>
      <ComponentItem :data='childData(ck)' :def='childDef(ck)' :level='level+1'/>
    </div>
  </div>
</div>
</template>

<script>
export default {
  props: ['data', 'def', 'level'],
  data() {
    return {
      expanded: false,
      expansionHovered: false,
      selected: false
    }
  },
  beforeCreate() {
    this.$options.components.ComponentItem = require('./ComponentItem.vue')
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
    title() {
      if(this.def.title)
        return this.def.title
      return this.def.name
    },
    style() {
      return {
        'padding-left': '' + (this.level * 25 + 10) + 'px'
      }
    },
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
    childData(key) {
      return this.data.children[key]
    },
    childDef(key) {
      return this.data.def.children.find((c) => {
        if(c.name === key)
          return true
        return false
      });
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
#vcms {
  & .component-hierarchy {
    & .wrapper {
      padding: 0;

      & .item {
        padding: 5px;
        &:hover {
          background-color: #eee;
          cursor: pointer;
        }
        & i {
          vertical-align: middle;
          margin-right: 5px;
          width: 12px;
          height: 14px;
        }

        &.selected {
          background-color: lightgray;
        }
      }
    }
    
  }
}
</style>