<script lang="ts" setup>

import InlineButton from "../ui/InlineButton.vue";
import {useTodoStore} from "../../services/store/todo.ts";
import {computed} from "vue";

export type TodoFilter = {
    filter: 'all' | 'completed' | 'uncompleted' | 'type';
    todoType?: string;
}

type TodoFiltersProps = {
    modelValue: TodoFilter;
};

const props = defineProps<TodoFiltersProps>();

const emit = defineEmits<{
    "update:modelValue": (value: TodoFilter) => void;
}>();

const store = useTodoStore();

const todos = computed(() => store.todos);
const todoTypes = computed(() => store.todoTypes);

const todoTypesWithTodos = computed(() => todoTypes.value.filter(todoType => todos.value.some(({type}) => type === todoType.id)))

function setFilter(filter: TodoFilter["filter"]) {
    emit("update:modelValue", {
        filter,
        todoType: undefined,
    });
}

function setFilterByTodoType(todoType: TodoFilter["todoType"]) {
    emit("update:modelValue", {
        filter: "type",
        todoType,
    });
}

function isType(todoType: TodoFilter["todoType"]) {
    return props.modelValue.filter === "type" && props.modelValue.todoType === todoType;
}

</script>

<template>
    <div
        class="sticky top-20 my-4 bg-gradient-to-b from-white via-white/80 to-transparent pt-2 overflow-x-auto z-10 select-none">

        <!-- Filter by all -->
        <InlineButton :class="{'active': modelValue.filter === 'all'}"
                      @click="setFilter('all')"
                      v-text="'Alles'"/>

        <!-- Filter by open -->
        <InlineButton :class="{'active': modelValue.filter === 'uncompleted'}"
                      @click="setFilter('uncompleted')"
                      v-text="'Offen'"/>

        <!-- Filter by type -->
        <InlineButton v-for="todoType in todoTypesWithTodos"
                      :key="todoType.id"
                      :class="{'text-white': isType(todoType.id)}"
                      :style="{backgroundColor: isType(todoType.id) ? todoType.color : ''}"
                      @click="setFilterByTodoType(todoType.id)">
            <span :style="{color:isType(todoType.id) ? 'white' : todoType.color}"
                  class="mr-1" v-text="' ■'"/>
            <span v-text="todoType.name"/>
        </InlineButton>
    </div>
</template>

<style scoped>

.active {
    @apply bg-blue-600 text-white;
}

.active:hover {
    @apply bg-blue-700;
}

.active:active {
    @apply bg-blue-800;
}
</style>
