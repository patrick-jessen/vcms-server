// import Vue from 'vue'
import App from './App.vue'

new Vue({
  el: '#vcms',
  name: 'VCMS',
  render: h => h(App)
})

// Prevent server from stopping (in case of refresh)
var client = new XMLHttpRequest();
client.open("GET", "/preventStop", true);
client.send();

// Inform server that browser has closed
window.addEventListener('unload', stopServer, false);
function stopServer() {
  var client = new XMLHttpRequest();
  client.open("GET", "/stop", false);
  client.send();
}

var script = document.createElement('script');
script.src = 'http://localhost:8080/dist/build.js';
document.head.appendChild(script); //or something of the likes