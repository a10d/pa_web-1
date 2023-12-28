<script lang="ts" setup>

import {computed} from "vue";
import {ValidationError} from "../../services/backend";

type FormErrorProps = {
    error: ValidationError | null,
    field?: string | null,
}

const props = withDefaults(defineProps<FormErrorProps>(), {
    field: null,
});

const hasErrors = computed(() => props.error !== null);
const isSingleField = computed(() => props.field !== null);

</script>

<template>
    <!-- Single field error -->
    <div v-if="hasErrors && isSingleField && props.error && props.field && props.error.has(props.field)"
         class="text-red-800 px-2 rounded-lg -mt-3 mb-4">
        <p v-for="message in props.error.get(props.field)" :key="message" class="text-sm" v-text="message"/>
    </div>

    <!-- All fields error -->
    <div v-if="hasErrors && !isSingleField"
         class="bg-red-100 border border-red-300 text-red-900 p-2 rounded-lg mb-4 shadow">
        <ul v-if="error" class="mt-1 text-sm list-disc pl-4">
            <template v-for="messages in error.all()" :key="messages.join()">
                <li v-for="message in messages" :key="message" v-text="message"/>
            </template>
        </ul>
    </div>
</template>
