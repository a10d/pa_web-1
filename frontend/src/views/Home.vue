<script lang="ts" setup>

import AppLayout from "../components/layouts/AppLayout.vue";
import TodoCreateForm from "../components/todo/TodoCreateForm.vue";
import {useTodoStore} from "../services/store/todo.ts";
import {computed} from "vue";
import ListTodo from "../components/todo/ListTodo.vue";
import TodoEditOptions from "../components/todo/TodoEditOptions.vue";

const store = useTodoStore();

const todos = computed(() => store.todos);
</script>

<template>
    <AppLayout>
        <TodoCreateForm/>

        <div v-if="todos.length === 0" class="my-24 select-none">
            <p class="text-xl italic text-gray-500">
                „Das Geheimnis des Erfolgs ist anzufangen.“
            </p>
            <p class="text-gray-500 text-sm font-medium ml-3">
                — Mark Twain oder so
            </p>
        </div>

        <ListTodo v-for="todo in todos" :key="todo.id" :todo="todo"/>

        <TodoEditOptions/>
    </AppLayout>
</template>
