<script lang="ts" setup>

import {computed} from "vue";
import {useTodoStore} from "../../services/store/todo.ts";
import FormField from "../ui/FormField.vue";
import {TodoType} from "../../services/backend";

type TodoTypeFieldProps = {
    modelValue: TodoType | null;
}

const props = defineProps<TodoTypeFieldProps>();

const emit = defineEmits<{
    'update:modelValue': [value: TodoType | null],
}>();

const store = useTodoStore();

const selectableTodoTypes = computed(() => store.todoTypes.map((type) => ({
    value: type.id,
    label: type.name,
})));

const valueProxy = computed<string>({
    get: () => props.modelValue ? props.modelValue.id : '',
    set: (value: string) => {
        emit('update:modelValue', store.todoTypes.find(type => type.id === value) ?? store.todoTypes[0]);
    },
});

</script>

<template>
    <FormField v-model="valueProxy"
               :select-options="selectableTodoTypes"
               label="Aufgabentyp"
               name="todoType"
               required
               type="select"/>
</template>
