<script lang="ts" setup>

import {onMounted, ref} from "vue";
import {initStores} from "./services/store";
import SkeletonLayout from "./components/layouts/SkeletonLayout.vue";

const isInitialized = ref(false);
onMounted(() => initStores().then(() => isInitialized.value = true));

</script>

<template>
  <!-- Skeleton -->
  <SkeletonLayout v-if="!isInitialized"/>

  <!-- Main App Router View -->
  <RouterView v-if="isInitialized"/>
</template>
