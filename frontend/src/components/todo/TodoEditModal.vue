<script lang="ts" setup>

import {Todo} from "../../services/backend";
import Modal from "../ui/Modal.vue";
import {computed, reactive, ref, watch} from "vue";
import PopButton from "../ui/PopButton.vue";
import FormErrors from "../ui/FormErrors.vue";
import TodoAssigneesField from "./TodoAssigneesField.vue";
import FormField from "../ui/FormField.vue";
import {useErrorHandler} from "../../services/errorHandler";
import {useTodoStore} from "../../services/store/todo.ts";
import {useEventBus} from "../../services/eventBus";

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
const formError = ref<Error | null>(null);

const todoForm = reactive<{
    type: string;
    title: string;
    description: string | null;
    dueDate: Date;
    assignees: string[];
    completed: boolean;
}>({
    title: "",
    description: "",
    dueDate: new Date(),
    completed: false,
    type: "",
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

const selectableTodoTypes = computed(() => store.todoTypes.map((type) => ({
    value: type.id,
    label: type.name,
})));

function cancel() {
    if (isSubmitting.value) return;
    close();
}

function close() {
    emit('close');
}

async function submitEditForm() {
    if (props.todo === null) return;
    isSubmitting.value = true;
    try {
        await store.updateTodo({
            ...props.todo,
            ...todoForm,
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
        formError.value = e;
        useErrorHandler().handle(e);
        eventBus.emit('playSound', 'formError');
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

                <FormField v-model="todoForm.completed" label="Erledigt" name="completed" type="checkbox"/>

                <!-- Description -->
                <FormField
                    v-model="todoForm.description"
                    label="Genauere Beschreibung"
                    name="description"
                    required
                    type="textarea"
                />

                <div class="grid grid-cols-2 gap-4">
                    <!-- DueDate -->
                    <div class="col-span-2 sm:col-span-1">
                        <FormField v-model="todoForm.dueDate"
                                   label="Fällig am"
                                   name="dueDate"
                                   required
                                   type="date"/>
                    </div>

                    <!-- TodoType -->
                    <div class="col-span-2 sm:col-span-1">
                        <FormField v-model="todoForm.type"
                                   :select-options="selectableTodoTypes"
                                   label="Aufgabentyp"
                                   name="todoType"
                                   required
                                   type="select"/>
                    </div>

                </div>


                <!-- Assignees -->
                <label class="block text-sm font-medium text-slate-800 px-2 mb-2">Zugewiesen an</label>
                <TodoAssigneesField v-model="todoForm.assignees"/>

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
