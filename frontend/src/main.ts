import {createApp} from "vue";
import App from "./App.vue";
import {router} from "./services/router";
import {pinia} from "./services/store";
import "./style.css"

const app = createApp(App);

app.use(pinia);
app.use(router);


app.mount("#app");
