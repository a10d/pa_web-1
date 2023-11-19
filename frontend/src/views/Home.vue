<script lang="ts" setup>

import AppLayout from "../components/layouts/AppLayout.vue";
import TodoCreateForm from "../components/todo/TodoCreateForm.vue";
import {useTodoStore} from "../services/store/todo.ts";
import {computed} from "vue";
import ListTodo from "../components/todo/ListTodo.vue";


const store = useTodoStore();

const todos = computed(() => store.todos);
</script>

<template>
    <AppLayout>
        <TodoCreateForm/>

        <ListTodo v-for="todo in todos" :key="todo.id" :todo="todo"/>

    </AppLayout>
</template>
