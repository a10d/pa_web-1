<script lang="ts" setup>

import {Todo, TodoType, User, ValidationError} from "../../services/backend";
import Modal from "../ui/Modal.vue";
import {computed, reactive, ref, watch} from "vue";
import PopButton from "../ui/PopButton.vue";
import FormErrors from "../ui/FormErrors.vue";
import TodoAssigneesField from "./TodoAssigneesField.vue";
import FormField from "../ui/FormField.vue";
import {useErrorHandler} from "../../services/errorHandler";
import {useTodoStore} from "../../services/store/todo.ts";
import {useEventBus} from "../../services/eventBus";
import TodoTypeField from "./TodoTypeField.vue";
import TodoDueDateField from "./TodoDueDateField.vue";

type TodoEditModalProps = {
    todo: Todo | null;
};

const props = defineProps<TodoEditModalProps>();

const emit = defineEmits<{
    close: [],
    updateTodo: [todo: Todo],
}>();

const store = useTodoStore();
const eventBus = useEventBus();

const isModalVisible = computed(() => props.todo !== null);
const isSubmitting = ref(false);
const formError = ref<ValidationError | null>(null);

const todoForm = reactive<{
    title: string;
    description: string | null;
    dueDate: Date;
    completed: boolean;
    type: TodoType | null,
    assignees: User[],
}>({
    title: "",
    description: "",
    dueDate: new Date(),
    completed: false,
    type: null,
    assignees: [],
});

watch(() => props.todo, () => {
    if (props.todo === null) return;

    todoForm.title = props.todo.title;
    todoForm.description = props.todo.description;
    todoForm.dueDate = props.todo.dueDate;
    todoForm.completed = props.todo.completed;
    todoForm.type = props.todo.type;
    todoForm.assignees = props.todo.assignees;
})


function cancel() {
    if (isSubmitting.value) return;
    close();
}

function close() {
    emit('close');
    formError.value = null;
}

async function submitEditForm() {
    if (props.todo === null) return;
    isSubmitting.value = true;
    try {
        await store.updateTodo({
            id: props.todo.id,
            ...todoForm,
            type: todoForm.type as TodoType,
        });
        eventBus.emit('playSound', 'updateTodo');
        close();
    } catch (e: any) {
        formError.value = e;
        useErrorHandler().handle(e);
        eventBus.emit('playSound', 'formError');
    } finally {
        isSubmitting.value = false;
    }
}

async function destroy() {
    if (props.todo === null) return;
    try {
        await store.deleteTodo(props.todo);
        eventBus.emit('playSound', 'deleteTodo');
        close();
    } catch (e: any) {
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
    <Modal :closeable="!isSubmitting" :show="isModalVisible" @close="cancel">
        <form @submit.prevent="submitEditForm">
            <fieldset :disabled="isSubmitting" class="p-4">
                <!-- Errors -->
                <FormErrors :error="formError"/>

                <!-- Title -->
                <FormField
                    v-model="todoForm.title"
                    label="Titel"
                    name="title"
                    required
                />
                <FormErrors :error="formError" field="title"/>

                <FormField v-model="todoForm.completed" label="Erledigt" name="completed" type="checkbox"/>
                <FormErrors :error="formError" field="completed"/>

                <!-- Description -->
                <FormField
                    v-model="todoForm.description"
                    label="Genauere Beschreibung"
                    name="description"
                    required
                    type="textarea"
                />
                <FormErrors :error="formError" field="description"/>

                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div>
                        <!-- DueDate -->
                        <TodoDueDateField v-model="todoForm.dueDate"/>
                        <FormErrors :error="formError" field="dueDate"/>
                    </div>
                    <div>
                        <!-- TodoType -->
                        <TodoTypeField v-model="todoForm.type"/>
                        <FormErrors :error="formError" field="type"/>
                    </div>
                </div>

                <!-- Assignees -->
                <label class="block text-sm font-medium text-slate-800 px-2 mb-2">Zugewiesen an</label>
                <TodoAssigneesField v-model="todoForm.assignees"/>
                <FormErrors :error="formError" field="assignees"/>

            </fieldset>
            <div class="p-4 border-t bg-gray-50 flex justify-end gap-4 flex-wrap">
                <button :disabled="isSubmitting"
                        class="mr-auto text-red-800 font-medium underline"
                        type="button"
                        @click="destroy">
                    Aufgabe löschen
                </button>

                <PopButton :disabled="isSubmitting" color="gray" label="Abbrechen" type="button"
                           @click="cancel"/>
                <PopButton :disabled="isSubmitting" color="green" label="Sichern" type="submit"/>
            </div>
        </form>
    </Modal>
</template>
