import {createPinia} from "pinia";
import {useAuthStore} from "./auth";


export const pinia = createPinia();


export const initStores = async () => {
    await useAuthStore().initialize();
};
