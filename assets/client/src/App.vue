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
    <Inspector v-if='validApp'/>
    <iframe v-show='!showBase' id='app' :src='iframeUrl'></iframe>
  </div>
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
      appUrl: '127.0.0.1:8080',
      showBase: true,
      validApp: false,
      iframeUrl: ''
    }
  },
  created() {
    this.$root.$on('hover', (arg) => {
      window.vue.$set(window.vue, 'hoveredNamespace', arg)
    })
    this.$root.$on('select', (arg) => {
      
    })
  },
  mounted() {

    this.connect()
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
      this.connect()     
    },
    onRecent(url) {
      this.appUrl = url
      this.connect()
    },

    connect() {
      var protocol, host, port
      var urlFrags = this.appUrl.split(':')
      switch(urlFrags.length) {
        case 1:
          protocol = 'http'
          host = urlFrags[0]
          port = '80'
          break
        case 2:
          // Check if port
          port = '' + parseInt(urlFrags[1])
          if(port.length > 0) {
            protocol = 'http'
            host = urlFrags[0]
          }
          else {
            protocol = urlFrags[0]
            host = urlFrags[1].replace('//', '')
            port = '80'
          }
          break
        case 3:
          protocol = urlFrags[0]
          host = urlFrags[1].replace('//', '')
          port = urlFrags[2]
      }
      
      var client = new XMLHttpRequest();
      client.open("GET", "/api/connect/" + protocol + '/' + host + '/' + port, true);
      client.onload = () => {
        if (client.readyState === 4) {
          if (client.status === 200) {
            
            this.iframeUrl = 'http://127.0.0.1:1337'
            document.getElementById('app').onload = () => {
              window.vue = window.frames[0].window.vue
              this.validApp = true

              window.Component = window.frames[0].window.Component
            }

            return
          }
        }
        this.validApp = false
      }
      client.send();
    }
  },
  components: {
    Inspector: require('./components/Inspector.vue'),
  }
}
</script>

<style lang="scss">
body {
  margin: 0;
}

#vcms {
  height: 100vh;
  overflow: hidden;
  font-family: 'Roboto', sans-serif;

  & .base{
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

#app {
  height: 100vh;
  border: none;
}

.slide-enter-active, .slide-leave-active {
  transition: all 0.0s ease-in;
  
}
.slide-enter, .slide-leave-to {
  transform: translateY(-100%);
}
</style>
