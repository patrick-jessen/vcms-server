<template>
<div id='inspector' v-if='shown'>
  <div :style='style'>
    <div class='resize' @mousedown='onResizeStart'></div>
    <div class='content'>
      <img class='logo' src='/assets/favicon.png'>
    </div>
  </div>
</div>
</template>

<script>
export default {
  data() {
    return {
      shown: true,
      style: {
        width: '300px'
      }
    }
  },
  mounted() {
    this.setAppSize()
    window.addEventListener("keydown", this.onToggle, false);
  },
  methods: {
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
  }
}
</script>

<style lang='scss'>
#inspector {
  pointer-events: all;
  position: fixed;
  right: 0;
  height: 100vh;
  background-color: #141E30;
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

    & .logo {
      width: 30px;
      margin: 0;
    }
  }
}
</style>