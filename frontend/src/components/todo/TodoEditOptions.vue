<script lang="ts" setup>

import {useTodoStore} from "../../services/store/todo.ts";
import {computed, ref} from "vue";

const store = useTodoStore();

const isSubmitting = ref(false);

const hasAnyCompletedTodos = computed(() => store.todos.some((todo) => todo.completed));
const deleteCompletedTodos = async () => {
    isSubmitting.value = true;
    await Promise.all(store.todos.filter(({completed}) => completed).map(todo => store.deleteTodo(todo)));
    isSubmitting.value = false;
};
</script>

<template>
    <div class="sticky bottom-0 bg-white/80 backdrop-blur-lg py-2 overflow-x-auto">
        <button v-if="hasAnyCompletedTodos"
                :disabled="isSubmitting"
                @click="deleteCompletedTodos">
            Erledigte Aufgaben löschen
        </button>


    </div>

</template>

<style scoped>

button {
    @apply rounded-full font-medium border bg-white px-4 py-1 text-sm transition-all mr-4;
}

button:hover {
    @apply bg-gray-50;
}

button:active {
    @apply bg-gray-100;
}

button:disabled {
    @apply opacity-50 cursor-not-allowed;
}

</style>
