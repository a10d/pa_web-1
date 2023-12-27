<script lang="ts" setup>

import AppLayout from "../components/layouts/AppLayout.vue";
import TodoCreateForm from "../components/todo/TodoCreateForm.vue";
import {useTodoStore} from "../services/store/todo.ts";
import {computed, ref} from "vue";
import ListTodo from "../components/todo/ListTodo.vue";
import TodoFilters, {TodoFilter} from "../components/todo/TodoFilters.vue";
import {useEventBus} from "../services/eventBus";
import TodoEditOptions from "../components/todo/TodoEditOptions.vue";
import NoContentFiller from "../components/todo/NoContentFiller.vue";
import {Todo} from "../services/backend";
import TodoEditModal from "../components/todo/TodoEditModal.vue";

const eventBus = useEventBus();
const store = useTodoStore();

const todos = computed(() => store.todos
    .filter(todo => {
        switch (filter.value.filter) {
            case 'all':
                return true;
            case 'completed':
                return todo.completed;
            case 'uncompleted':
                return !todo.completed;
            case 'type':
                return todo.type.id === filter.value.todoType;
        }
    })
    .sort((a, b) => new Date(a.dueDate).getTime() - new Date(b.dueDate).getTime()));

const filter = ref<TodoFilter>({
    filter: 'uncompleted',
});

// Reset filter when creating a new task with a different type
eventBus.on('createTodo', (todo) => {
    if (filter.value.todoType !== todo.type.id) {
        filter.value = {filter: 'uncompleted'}
    }
});

const editingTodo = ref<null | Todo>(null);

function editTodo(todo: Todo) {
    editingTodo.value = todo;
}

function closeEditTodo() {
    editingTodo.value = null;
}

</script>

<template>

    <AppLayout>

        <section class="relative pt-6">

            <TodoCreateForm/>

            <TodoFilters v-model="filter"/>

            <transition-group enter-active-class="transition-all duration-250 ease-out transform-gpu origin-top"
                              enter-from-class="opacity-0 scale-y-0 h-0"
                              enter-to-class="opacity-100 h-[110px]"
                              leave-active-class="transition-all duration-250 ease-in-out transform-gpu origin-top"
                              leave-from-class="opacity-100 h-[110px]"
                              leave-to-class="opacity-0 scale-y-0  h-0">

                <ListTodo
                    v-for="todo in todos"
                    :key="todo.id"
                    :todo="todo"
                    @editTodo="editTodo"/>

            </transition-group>

            <NoContentFiller v-if="todos.length === 0"/>

            <TodoEditOptions/>

        </section>

        <TodoEditModal :todo="editingTodo"
                       @close="closeEditTodo"/>

    </AppLayout>

</template>
