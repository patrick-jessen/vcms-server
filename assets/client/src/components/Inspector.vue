<template>
<div id='inspector' v-if='shown'>
  <div :style='style'>
    <div class='resize' @mousedown='onResizeStart'></div>
    <div class='content'>
      <img class='logo' src='/assets/favicon.png'>
      <TreeTable v-bind='pageData' @select='onSelect'/>

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
      shown: true,
      style: {
        width: '300px'
      },
      selectedComponent: 'app',

      pageData: {

        test: new TreeTableData(
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
        test: new TreeTableData(
          () => {
            return new Component(this.selectedComponent)
          },
          (data) => {
            console.log("here", data.properties)
            var keys = Object.keys(data.properties)
            return keys.map((k) => {
              return data.properties[k]
            })
          },
          (data, index) => {
            return data.namespace.last
          }
        ),
        columns: [
          {
            title: 'Property', 
            width: 100
          },
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
      console.log('select', arg)
      this.selectedComponent = arg
    },
    onResizeStart(e) {
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
    },
    onResizeStop(e) {
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
  pointer-events: all;
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
    width: 100%;
    vertical-align: top;
    padding: 20px;
    text-align: center;
    height: 100vh;
    overflow: auto;

    & .logo {
      width: 30px;
      margin: 0;
    }
  }
}
</style>