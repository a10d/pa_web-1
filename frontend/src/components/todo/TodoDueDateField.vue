<script lang="ts" setup>

import FormField from "../ui/FormField.vue";
import {computed} from "vue";

type TodoDueDateFieldProps = {
    modelValue: string | Date;
}

const props = defineProps<TodoDueDateFieldProps>();

const emit = defineEmits<{
    'update:modelValue': [value: string],
}>();

const valueProxy = computed<string>({
    get: () => {
        const date = new Date(props.modelValue);
        date.setMinutes(date.getMinutes() - date.getTimezoneOffset());
        return date.toISOString().slice(0, 16);
    },
    set: (value: string) => {
        emit('update:modelValue', new Date(value).toISOString());
    },
});

</script>

<template>
    <FormField v-model="valueProxy"
               label="Fällig am"
               name="dueDate"
               required
               type="datetime-local"/>
</template>
