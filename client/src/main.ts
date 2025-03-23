import { createApp } from "vue";
import "./style.css";

import App from "./App.vue";
import ErrorPage from "./ErrorPage.vue";

import { checkServerConnection } from "./lib/service/app.service";

import router from "./router";

(async () => {
  const isConnected = await checkServerConnection();

  if (isConnected) {
    const app = createApp(App);
    app.use(router);
    app.mount("#app");
  } else {
    createApp(ErrorPage).mount("#app");
  }
})();
