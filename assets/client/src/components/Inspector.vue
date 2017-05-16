<template>
<div id='inspector' v-if='shown'>
  <div :style='style'>
    <div class='resize' @mousedown='onResizeStart'></div>
    <div class='content'>
      <img class='logo' src='/assets/favicon.png'>

      <div class='tabs'>
        <div class='tab' :class='{active:activeTab==0}' @click='activeTab=0'>
          <i class='fa fa-cubes'></i>
        </div>
        <div class='tab' :class='{active:activeTab==1}' @click='activeTab=1'>
          <i class='fa fa-columns'></i>
        </div>
      </div>

      <TreeTable v-show='activeTab==0' v-bind='compData'/>
      <TreeTable v-show='activeTab==1' v-bind='pageData'/>
      <TreeTable v-bind='propsData'/>
    </div>
  </div>
</div>
</template>

<script>
import {TreeTableData} from '../structs.js'

export default {
  data() {
    return {
      activeTab: 0,
      shown: true,
      style: {
        width: '500px'
      },
      selectedComponent: 'app.$children.menu',

      compData: {
        data: new TreeTableData(
          () => { 
            return window.vue.$children[0].component
          },
          (data) => { 
            var keys = Object.keys(data.children)
            return keys.map((k) => {
              return data.children[k]
            })
          },
          (data, index) => {
            switch(index) {
              case 0:
                if(data.parentDef.title) return data.parentDef.title
                  return data.parentDef.name
              case 1:
                return `<img src='http://localhost:8081/assets/favicon.png'> ` + data.type
            }
          },
          (data) => {
            return data.namespace.string
          }
        ),
        columns: [
          {
            title: 'Component', 
            width: 60
          },
          {
            title: 'Type', 
            width: 40,
            style: {
              'text-align': 'right'
            }
          }
        ]
      },

      propsData: {
        skipRoot: true,
        data: new TreeTableData(
          () => {
            return new Component(this.selectedComponent)
          },
          (data, root) => {
            if(root) {
              var keys = Object.keys(data.properties)

              return keys.map((k) => {
                return data.properties[k]
              })

            }
            return data.properties
          },
          (data, index) => {
            switch(index) {
              case 0:
                return data.namespace.last
              case 1:
                if(typeof data.value === 'object') 
                {
                  var fn = data.def.render
                  var val = data.value

                  if(data.def.type === 'array') {
                    if(!Array.isArray(val))
                      val = [val]
                  }

                  return new Function('val', fn)(val)
                }
                else {
                  return data.def.type
                }
            }
          },
          (data) => {
            return data.namespace.string
          }
        ),
        columns: [
          {
            title: 'Property', 
            width: 50
          },
          {
            title: 'Value', 
            width: 50
          },
        ]
      },

      pageData: {
        skipRoot: true,
        data: new TreeTableData(
          () => {
            return window.vue.$router.options.routes
          },
          (data, isRoot) => {
            if(isRoot)
              return window.vue.$router.options.routes
            
            if(data.children)
              return data.children
            return []
          },
          (data, index) => {
            return data.path
          },
          (data) => {
            return data.path
          }
        ),
        columns: [
          {
            title: 'Page',
            width: 100
          }
        ]
      }
    
    }
  },
  mounted() {
    this.setAppSize()
    window.addEventListener("keydown", this.onToggle, false);
  },
  methods: {
    onSelect(arg) {
      this.selectedComponent = arg
    },
    onResizeStart(e) {
      window.frames[0].window.addEventListener('mousemove', this.onResize, false);
      window.frames[0].window.addEventListener('mouseup', this.onResizeStop, false);
      window.addEventListener('mousemove', this.onResize, false);
      window.addEventListener('mouseup', this.onResizeStop, false);
      this.lastMouse = e.x
    },
    onResize(e) {
      var dif = this.lastMouse - e.x
      var w = parseInt(this.style.width)
      w += dif
      if(w < 200) {
        w = 200
      }
      else if(w > 500) {
        w = 500
      }
      this.style.width = '' + w + 'px'
      this.lastMouse = e.x

      e.stopPropagation()
    },
    onResizeStop(e) {
      window.frames[0].window.removeEventListener('mousemove', this.onResize, false);
      window.frames[0].window.removeEventListener('mouseup', this.onResizeStop, false);
      window.removeEventListener('mousemove', this.onResize, false);
      window.removeEventListener('mouseup', this.onResizeStop, false);
      
      this.setAppSize()
    },
    setAppSize() {
      if(this.shown) 
        document.getElementById('app').style.width = 'calc(100% - ' + this.style.width + ')'
      else
        document.getElementById('app').style.width = '100%'
    },
    onToggle(e) {
      if(e.key === 'e' && e.ctrlKey) {
        this.shown = !this.shown
        event.preventDefault()
        this.setAppSize()
      }
    }
  },
  components: {
    TreeTable: require('./TreeTable.vue')
  }
}
</script>

<style lang='scss'>
#inspector {
  position: fixed;
  right: 0;
  height: 100vh;
  background-color: #333;
  vertical-align: top;

  & .resize {
    height: 100vh;
    width: 2px;
    background-color: gray;
    cursor: ew-resize;
    position: absolute;
  }
  & .content {
    vertical-align: top;
    padding: 20px;
    text-align: center;
    height: 100vh;
    overflow: auto;

    & .logo {
      width: 30px;
      margin: 0;
    }

    & .tabs {
      height: 40px;
      width: 100%;

      & .tab {
        display: inline-block;
        padding: 10px 30px;
        color: #aaa;
        cursor: pointer;

        &:hover {
          background-color: #444;
          color: #fefefe;
        }

        &.active {
          color: #fefefe;
        }
      }
    }
  }
}
</style>