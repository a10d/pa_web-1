<script lang="ts" setup>

import {computed} from "vue";
import {ValidationError} from "../../services/backend";

type FormErrorProps = {
    error: ValidationError | null,
}

const props = defineProps<FormErrorProps>();

const hasErrors = computed(() => props.error !== null);

</script>

<template>
    <div v-if="hasErrors" class="bg-red-200 border border-red-300 text-red-900 p-4 rounded-lg mb-4 shadow">

        <p>Whups, da stimmt etwas nicht ganz.</p>

        <ul class="mt-1 text-sm list-disc pl-3">
            <template v-for="(messages, field) in error.all()" :key="field">
                <li v-for="message in messages" :key="message" v-text="message"/>
            </template>
        </ul>
    </div>
</template>
