<script lang="ts" setup>

import {useTodoStore} from "../../services/store/todo.ts";
import {computed, ref} from "vue";
import InlineButton from "../ui/InlineButton.vue";

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
    <div class="py-2">
        <InlineButton v-if="hasAnyCompletedTodos"
                      :disabled="isSubmitting"
                      @click="deleteCompletedTodos">
            Erledigte Aufgaben löschen
        </InlineButton>
    </div>

</template>
