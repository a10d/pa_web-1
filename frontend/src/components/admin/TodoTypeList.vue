<script lang="ts" setup>

import {useTodoStore} from "../../services/store/todo.ts";
import {computed, nextTick, reactive, ref} from "vue";
import PopButton from "../ui/PopButton.vue";
import Modal from "../ui/Modal.vue";
import FormField from "../ui/FormField.vue";
import {TodoType, ValidationError} from "../../services/backend";
import FormErrors from "../ui/FormErrors.vue";
import {useErrorHandler} from "../../services/errorHandler";
import {useEventBus} from "../../services/eventBus";

const store = useTodoStore();
const eventBus = useEventBus();

const todoTypes = computed(() => store.todoTypes);

const createForm = reactive({
  name: "",
  description: "",
  color: "",
  reminderTime: 3,
})
const createModalOpen = ref(false);

function openCreateModal() {
  createModalOpen.value = true;
  formError.value = null;

  createForm.name = "";
  createForm.description = "";
  createForm.color = "#cccccc";
  createForm.reminderTime = 3;

  nextTick(() => {
    const input = document.querySelector("input[name=name]") as HTMLInputElement;
    input.focus();
  });
}

function cancelCreateModal() {
  createModalOpen.value = false;
}

const isSubmitting = ref(false);
const formError = ref<ValidationError | null>(null);

async function submitCreateForm() {
  isSubmitting.value = true;
  formError.value = null;
  try {

    createForm.reminderTime = parseInt(createForm.reminderTime.toString());

    await store.createTodoType(createForm);
    cancelCreateModal();
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

async function deleteTodoType(todoType: TodoType) {
  try {
    formError.value = null;
    await store.deleteTodoType(todoType);
  } catch (e) {
    useErrorHandler().handle(e)
  }
}


const editingTodoType = ref<TodoType | null>(null);
const editingModalOpen = computed(() => !!editingTodoType.value);
const editForm = reactive({
  name: "",
  description: "",
  color: "",
  reminderTime: 0,
});

function editTodoType(todoType: TodoType) {
  editingTodoType.value = todoType;
  formError.value = null;

  editForm.name = todoType.name;
  editForm.description = todoType.description;
  editForm.color = todoType.color;
  editForm.reminderTime = todoType.reminderTime;

  nextTick(() => {
    const input = document.querySelector("input[name=name]") as HTMLInputElement;
    input.focus();
  });
}

function cancelEditingTodoType() {
  editingTodoType.value = null;
}

async function submitEditForm() {
  isSubmitting.value = true;
  formError.value = null;
  try {

    editForm.reminderTime = parseInt(editForm.reminderTime.toString());

    await store.updateTodoType({
      ...editingTodoType.value,
      ...editForm,
    } as TodoType);
    cancelEditingTodoType();
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

  <div class="relative">
    <!-- Header -->
    <div
        class="flex items-center justify-between gap-2 sticky top-0 py-4 bg-white/80 backdrop-blur-sm border-b border-white z-40">
      <h2 class="text-xl font-medium">Aufgabentypen</h2>
      <PopButton color="gray" label="Erfassen" type="button" @click="openCreateModal"/>
    </div>


    <!-- List -->
    <div class="p-2 pb-0.5 rounded-lg bg-gray-100 shadow-inner">

      <!-- No Content -->
      <p v-if="todoTypes.length === 0" class="text-center text-sm text-gray-800 my-12">Es wurden noch keine
        Aufgabentypen
        erfasst...</p>

      <!-- List Items -->
      <div v-for="todoType in todoTypes" :key="todoType.id"
           class="rounded bg-white p-2 mb-2 shadow flex flex-wrap justify-end gap-2 group items-center">
        <div class="mr-auto w-full sm:w-auto">
          <p class="font-medium text-lg">
            <span :style="{color: todoType.color}" class="mr-1">■</span>
            <span v-text="todoType.name"/>
          </p>
          <p class="text-sm text-gray-600" v-text="todoType.description"/>
        </div>


        <PopButton class="group-hover:opacity-100 sm:opacity-0" color="gray" label="Bearbeiten" type="button"
                   @click="editTodoType(todoType)"/>
        <PopButton class="group-hover:opacity-100 sm:opacity-0" color="red" label="Löschen" type="button"
                   @click="deleteTodoType(todoType)"/>
      </div>
    </div>
  </div>

  <!-- Create Modal -->
  <Modal :closeable="!isSubmitting" :show="createModalOpen" max-width="2xl" @close="cancelCreateModal">
    <form @submit.prevent="submitCreateForm">
      <fieldset :disabled="isSubmitting" class="p-4">
        <h2 class="text-xl font-medium mb-4">Aufgabentyp erfassen</h2>

        <FormErrors :error="formError"/>

        <FormField v-model="createForm.name" label="Name" name="name" required/>
        <FormErrors :error="formError" field="name"/>
        <FormField v-model="createForm.description" label="Beschreibung" name="description" type="textarea"/>
        <FormErrors :error="formError" field="description"/>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <FormField v-model="createForm.color" label="Farbe" name="color" type="color"/>
            <FormErrors :error="formError" field="color"/>
          </div>
          <div>
            <FormField v-model="createForm.reminderTime" :max="30" :min="1" label="Erinnerung (Tage)"
                       name="reminderTime" type="number"/>
            <FormErrors :error="formError" field="reminderTime"/>
          </div>

        </div>
      </fieldset>

      <div class="p-4 border-t bg-gray-50 flex justify-between gap-4 flex-wrap">
        <PopButton :disabled="isSubmitting" color="gray" label="Abbrechen" type="button"
                   @click="cancelCreateModal"/>
        <PopButton :disabled="isSubmitting" color="green" label="Erstellen" type="submit"/>
      </div>
    </form>
  </Modal>

  <!-- Edit Modal -->
  <Modal :closeable="!isSubmitting" :show="editingModalOpen" max-width="2xl" @close="cancelCreateModal">
    <form @submit.prevent="submitEditForm">
      <fieldset :disabled="isSubmitting" class="p-4">
        <h2 class="text-xl font-medium mb-4">Aufgabentyp bearbeiten</h2>

        <FormErrors :error="formError"/>

        <FormField v-model="editForm.name" label="Name" name="name" required/>
        <FormErrors :error="formError" field="name"/>
        <FormField v-model="editForm.description" label="Beschreibung" name="description" type="textarea"/>
        <FormErrors :error="formError" field="description"/>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <FormField v-model="editForm.color" label="Farbe" name="color" type="color"/>
            <FormErrors :error="formError" field="color"/>
          </div>
          <div>
            <FormField v-model="editForm.reminderTime" :max="60" :min="1" label="Frist (Tage)"
                       name="reminderTime" type="number"/>
            <FormErrors :error="formError" field="reminderTime"/>
          </div>
        </div>
      </fieldset>

      <div class="p-4 border-t bg-gray-50 flex justify-between gap-4 flex-wrap">
        <PopButton :disabled="isSubmitting" color="gray" label="Abbrechen" type="button"
                   @click="cancelEditingTodoType"/>
        <PopButton :disabled="isSubmitting" color="green" label="Sichern" type="submit"/>
      </div>
    </form>
  </Modal>
</template>
