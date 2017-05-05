
new Vue({
  template: `
  <div>
    <div v-if='showBase' class='base'>
      <img class='logo' src='/assets/favicon.png'>
      <h1>VCMS</h1>
      <p>Enter URL of website</p>
      <input id='url' :value='appUrl' @input='onInput'>
      <button @click='onNext()' :disabled='!urlIsOk'>Next</button>
    </div>
    <div v-show='!showBase' class='base'>
      	
    </div>
  </div>
  `,

  ///<iframe class='iframe' :src='appUrl' @load='appLoaded' :key='appUrl'></iframe>
  data() {
    return {
      showBase: true,
      appUrl: 'http://localhost:8080',
      urlIsOk: true
    }
  },
  el: '#client-app',
  methods: {
    onNext() {
      this.showBase = false;

      // var head = document.getElementsByTagName('head')[0],
      // script = document.createElement('script');
      // script.src = 'http://localhost:8080/dist/main.js';
      // head.appendChild(script);

      var script = document.createElement('script');
      script.src = 'http://localhost:8080/dist/build.js';
      document.head.appendChild(script); //or something of the likes
    },
    onInput(e) {
      console.log("here")
      this.appUrl = e.target.value;
    },
    appLoaded(e) {
      if(e.target)
        this.urlIsOk = true;
      else
        this.urlIsOk = false;
    },
  }
})

// Prevent server from stopping (in case of refresh)
var client = new XMLHttpRequest();
client.open("GET", "/preventStop", true);
client.send();

// Inform server that browser has closed
window.addEventListener('unload', stopServer, false);
function stopServer() {

  setTimeout(function() {
    console.log("timeout");
  }, 1000)

  var client = new XMLHttpRequest();
  client.open("GET", "/stop", false);
  client.send();
}

function onNext() {
  //window.location.replace(document.getElementById('url').value);
  document.getElementById('iframe').src = document.getElementById('url').value;
  document.getElementById('iframe').style.display = 'block';
}