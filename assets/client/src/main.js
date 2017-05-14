// import Vue from 'vue'
import App from './App.vue'
import './structs.js'

window.vcmsvue = new Vue({
  name: 'VCMS',
  el: '#vcms',
  render: h => h(App),
  data: {
    selected: null
  }
})

// Prevent server from stopping (in case of refresh)
var client = new XMLHttpRequest();
client.open("GET", "/api/preventStop", true);
client.send();

// Inform server that browser has closed
window.addEventListener('unload', stopServer, false);
function stopServer() {
  var client = new XMLHttpRequest();
  client.open("GET", "/api/stop", false);
  client.send();
}
