import {createPinia} from "pinia";
import {useTodoStore} from "./todo.ts";


export const pinia = createPinia();


export const initStores = async () => {
    await useTodoStore().initialize();
};
