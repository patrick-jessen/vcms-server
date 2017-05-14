<template>
  <div id='vcms'>
    <transition name='slide'>
      <div v-if='showBase' class='base'>
        <img class='logo' src='/assets/favicon.png'>
        <h1>VCMS</h1>
        <h5>Enter URL of website</h5>
        <input id='url' :value='appUrl' @input='onInput'>
        <button @click='onNext' :disabled='!validApp'>Next</button>
        <br><br>
        <h5>Recent:</h5>
        <ul class='recent'>
          <li v-for='a in recent()' @click='onRecent(a)'>
            {{a}}
          </li>
        </ul>
      </div>
    </transition>
    <Inspector v-if='!showBase'/>
  </div>
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
      appUrl: 'localhost:8080',
      showBase: true,
      validApp: false,
    }
  },
  created() {
    this.$root.$on('hover', (arg) => {
      window.vue.$set(window.vue, 'hoveredNamespace', arg)
    })
    this.$root.$on('select', (arg) => {
      
    })

    this.loadScript()
  },
  methods: {
    recent() {
      var recent = window.localStorage.getItem('recentApps')
      if(!recent)
        return []
      return recent.split(',')
    },
    onNext() {
      this.showBase = false
      var recent = window.localStorage.getItem('recentApps')
      if(!recent)
        recent = []
      else
        recent = recent.split(',')

      if(recent.indexOf(this.appUrl) === -1)
        recent.push(this.appUrl);

      window.localStorage.setItem('recentApps', recent)
      document.body.style.overflow = 'auto';
    },
    onInput(e) {
      this.appUrl = e.target.value
      this.loadScript()     
    },
    onRecent(url) {
      this.appUrl = url
      this.loadScript()
    },

    loadScript() {
      var prevScript = document.getElementById('appscript')
      if(prevScript) {
        document.head.removeChild(prevScript)
      }

      var url = 'http://' + this.appUrl + '/dist/build.js';
      var script = document.createElement('script');
      script.id = 'appscript';
      script.src = url;
      script.onload = () => {
        this.validApp = true;
        this.showBase = false;
      }
      document.head.appendChild(script);
    }
  },
  components: {
    Inspector: require('./components/Inspector.vue'),
  }
}
</script>

<style lang="scss">
#vcms {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  pointer-events: none;

  & .base{
    pointer-events: all;
    width: 100%;
    height: 100%;
    background: #141E30;  /* fallback for old browsers */
    background: -webkit-linear-gradient(to top, #243B55, #141E30);  /* Chrome 10-25, Safari 5.1-6 */
    background: linear-gradient(to top, #243B55, #141E30); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
    text-align: center;
    z-index: 9999;
  }
  & .logo {
    margin: 30px auto;
    width: 150px;
  }
  & h1, h5, p {
    color: white;
  }

  & button[disabled] {
    color: red;
  }

  & .recent {
    max-width: 250px;
    margin: auto;
    padding: 0;

    & li {
      list-style: none;
      padding: 10px;
      cursor: pointer;
      border: 1px solid black;
      color: white;

      &:hover {
        background-color: #243B55;
      }
    }
  }
}

.slide-enter-active, .slide-leave-active {
  transition: all 0.0s ease-in;
  
}
.slide-enter, .slide-leave-to {
  transform: translateY(-100%);
}
</style>
