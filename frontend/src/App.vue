<script lang="ts" setup>

import {computed, onMounted} from "vue";
import {initStores} from "./services/store";
import {useAuthStore} from "./services/store/auth";
import Login from "./views/auth/Login.vue";

onMounted(() => initStores());
const isAuthenticated = computed(() => useAuthStore().isAuthenticated);

</script>

<template>
    <Login v-if="!isAuthenticated"/>
    <RouterView v-if="isAuthenticated"/>
</template>
