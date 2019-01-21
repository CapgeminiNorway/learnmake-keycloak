import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'

Vue.config.productionTip = false

import VueKeyCloak from '@dsb-norge/vue-keycloak-js';
Vue.use(VueKeyCloak, {
  config: {
    authRealm: "KeycloakExampleApp",
    authUrl: "http://localhost:8081/auth",
    authClientId: "login-app",
    logoutRedirectUri: ""
  },
  onReady: (keycloak) => {
    // eslint-disable-next-line
    console.log(`Keycloak response: ${keycloak}`)
    new Vue({
      render: h => h(App),
    }).$mount('#app')
  }
});


/*
new Vue({
  render: h => h(App),
}).$mount('#app')
*/