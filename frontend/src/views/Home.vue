<script lang="ts" setup>

import AppLayout from "../components/layouts/AppLayout.vue";
import TodoCreateForm from "../components/task/TodoCreateForm.vue";
import {useTodoStore} from "../services/store/todo.ts";
import {computed} from "vue";


const store = useTodoStore();

const todos = computed(() => store.todos);
</script>

<template>
    <AppLayout>
        <TodoCreateForm/>

      <pre>{{ todos }}</pre>

    </AppLayout>
</template>
