<script lang="ts" setup>

</script>

<template>
    <button class="inline-button">
        <slot/>
    </button>
</template>

<style scoped>

.inline-button {
    @apply rounded-full font-medium border bg-white px-3 py-1 text-sm transition-all mr-2 mb-2;
}

.inline-button:hover {
    @apply bg-gray-50;
}

.inline-button:active {
    @apply bg-gray-100;
}

.inline-button:disabled {
    @apply opacity-50 cursor-not-allowed;
}
</style>
