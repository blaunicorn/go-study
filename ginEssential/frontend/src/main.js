import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

Vue.config.productionTip = false;

// import bootstrap
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";

// install bootstrapVue
Vue.use(BootstrapVue);
// optioanlly install the bootstrapVue icon components plugin
Vue.use(IconsPlugin);

// import scss style. if  bootstrap> 4.5.3, an error will be reported
import "./assets/scss/index.scss";

// import Vuelidate
import Vuelidate from "vuelidate";
Vue.use(Vuelidate);

import axios from "axios";
Vue.prototype.axios = axios;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
