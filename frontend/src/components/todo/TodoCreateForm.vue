<script lang="ts" setup>

import {computed, nextTick, reactive, ref, watch} from "vue";
import FormField from "../ui/FormField.vue";
import PopButton from "../ui/PopButton.vue";
import {useTodoStore} from "../../services/store/todo.ts";
import {useEventBus} from "../../services/eventBus";
import {useErrorHandler} from "../../services/errorHandler";
import FormErrors from "../ui/FormErrors.vue";
import TodoAssigneesField from "./TodoAssigneesField.vue";
import {TodoType, User, ValidationError} from "../../services/backend";
import TodoTypeField from "./TodoTypeField.vue";
import TodoDueDateField from "./TodoDueDateField.vue";

const eventBus = useEventBus();
const store = useTodoStore();

const selectableTodoTypes = computed(() => store.todoTypes.map((type) => ({
    value: type,
    label: type.name,
})));

const formDisabled = computed(() => isSubmitting.value || selectableTodoTypes.value.length === 0);

const todoForm = reactive<{
    title: string,
    description: string,
    dueDate: Date,
    type: TodoType | null,
    assignees: User[],
}>({
    title: "",
    description: "",
    dueDate: new Date(),
    type: null,
    assignees: [],
});

const formError = ref<any>(null);
const isSubmitting = ref(false);
const formVisible = ref(false);

const todoCTAs = [
    "Ohne Fleiss kein Preis! Wo fängst du an?",
    "Was steht noch an?",
    "Was ist noch zu tun?",
    "Was ist noch zu erledigen?",
    "Was ist der Plan?",
    "Was gibts zu tun?",
];
const todoCTA = computed(() => todoCTAs[Math.floor(Math.random() * todoCTAs.length)]);

async function openActionBar() {

    formError.value = null;

    const tomorrow = new Date();
    tomorrow.setDate(tomorrow.getDate() + 1);

    todoForm.title = "";
    todoForm.description = "";
    todoForm.dueDate = tomorrow;
    todoForm.type = store.todoTypes[0];
    todoForm.assignees = [];

    formVisible.value = true;
    await nextTick();
    const input = document.querySelector("input[name=title]") as HTMLInputElement;
    input.focus();
}

watch(() => formVisible, () => {
    if (formVisible.value) {
        document.body.style.overflow = "hidden";
        eventBus.emit('playSound', 'openTodoForm')
    } else {
        document.body.style.overflow = "auto";
        eventBus.emit('playSound', 'closeTodoForm')
    }
});

function cancel() {
    formVisible.value = false;

}

async function submitForm() {
    isSubmitting.value = true;
    try {
        await store.createTodo(todoForm);
        eventBus.emit('playSound', 'submitTodoForm')
        cancel();
    } catch (e) {
        eventBus.emit('playSound', 'formError');
        if (e instanceof ValidationError) {
            formError.value = e;
        } else {
            useErrorHandler().handle(e);
        }
    } finally {
        isSubmitting.value = false;
    }
}

</script>

<template>

    <!--  Action bar -->
    <input
        :placeholder="todoCTA"
        class="transition-all sticky top-6 bg-white focus:shadow-lg border rounded-lg w-full p-4 focus:ring outline-0 tracking-wide font-medium select-none z-20"
        @focus="openActionBar"
    />

    <div v-if="formVisible" class="fixed inset-0 bg-white/10 backdrop-blur-sm z-40" @click="cancel"/>

    <!-- Form -->
    <transition enter-active-class="transition-all transform transform-gpu duration-250 ease-out origin-top"
                enter-from-class="scale-y-0 opacity-0"
                leave-active-class="transition-all transform transform-gpu duration-300 ease-out origin-top"
                leave-to-class="scale-y-0 opacity-0">
        <div v-if="formVisible"
             class="fixed top-4 left-1/2 -translate-x-1/2 w-full max-w-2xl p-2 z-50 ">
            <div class="mx-auto rounded-lg bg-white p-4 shadow-2xl border max-h-[80vh] overflow-y-auto">

                <fieldset :disabled="formDisabled">
                    <!-- Errors -->
                    <FormErrors :error="formError"/>

                    <!-- Title -->
                    <FormField
                        v-model="todoForm.title"
                        :placeholder="todoCTA"
                        label="Titel"
                        name="title"
                        required
                    />

                    <!-- Description -->
                    <FormField
                        v-model="todoForm.description"
                        label="Genauere Beschreibung"
                        name="description"
                        placeholder="Beschreibe die Aufgabe etwas genauer..."
                        required
                        type="textarea"
                    />

                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <!-- DueDate -->
                        <TodoDueDateField v-model="todoForm.dueDate"/>

                        <!-- TodoType -->
                        <TodoTypeField v-model="todoForm.type"/>
                    </div>

                    <!-- Assignees -->
                    <label class="block text-sm font-medium text-slate-800 px-2 mb-2">Zugewiesen an</label>
                    <TodoAssigneesField v-model="todoForm.assignees"/>

                </fieldset>

                <!-- Buttons -->
                <div class="flex justify-between gap-4">
                    <PopButton
                        :disabled="isSubmitting"
                        color="gray"
                        label="Abbrechen"
                        @click="cancel"
                    />
                    <PopButton
                        :disabled="formDisabled"
                        color="green"
                        label="Loslegen"
                        @click="submitForm"
                    />
                </div>
            </div>
        </div>
    </transition>

</template>
