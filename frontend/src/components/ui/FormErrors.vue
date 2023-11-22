<script lang="ts" setup>

import {computed} from "vue";

type FormErrorProps = {
    error: Error | string | string[] | null | undefined,
}

defineProps<FormErrorProps>();


const hasErrors = computed(() => false);

</script>

<template>
    <div v-if="hasErrors"></div>
</template>
