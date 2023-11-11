<script lang="ts" setup>

import {useTodoStore} from "../../services/store/todo.ts";
import {computed, nextTick, reactive, ref} from "vue";
import PopButton from "../ui/PopButton.vue";
import Modal from "../ui/Modal.vue";
import FormField from "../ui/FormField.vue";
import {TodoType} from "../../services/backend";

const store = useTodoStore();

const todoTypes = computed(() => store.todoTypes);


const createForm = reactive({
  name: "",
  description: "",
  color: "",
  reminderTime: 0,
})
const createModalOpen = ref(false);

function openCreateModal() {
  createModalOpen.value = true;

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

async function submitCreateForm() {
  isSubmitting.value = true;

  try {
    await store.createTodoType(createForm);
    cancelCreateModal();
  } catch (e) {
    // TODO: Error handling
  } finally {
    isSubmitting.value = false;
  }
}

async function deleteTodoType(todoType: TodoType) {
  try {
    await store.deleteTodoType(todoType);
  } catch (e) {
    // TODO: Error handling
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
  try {
    await store.updateTodoType({
      ...editingTodoType.value,
      ...editForm,
    } as TodoType);
    cancelEditingTodoType();
  } catch (e) {
    // TODO: Error handling
  } finally {
    isSubmitting.value = false;
  }
}

</script>

<template>

  <!-- Header -->
  <div class="flex items-center justify-between gap-2 sticky top-0 py-4 bg-white/80 backdrop-blur-lg">
    <h2 class="text-xl font-medium">Aufgabentypen</h2>
    <PopButton color="gray" label="Erfassen" type="button" @click="openCreateModal"/>
  </div>


  <!-- List -->
  <div class="p-2 pb-0.5 rounded-lg bg-gray-100 shadow-inner">

    <!-- No Content -->
    <p v-if="todoTypes.length === 0" class="text-center text-sm text-gray-800 my-12">Es wurden noch keine Aufgabentypen
      erfasst...</p>

    <!-- List Items -->
    <div v-for="todoType in todoTypes" :key="todoType.id"
         class="rounded bg-white p-2 mb-2 shadow flex gap-2 group items-center">
      <p class="mr-auto font-medium text-lg" v-text="todoType.name"/>

      <PopButton class="group-hover:opacity-100 opacity-0" color="gray" label="Bearbeiten" type="button"
                 @click="editTodoType(todoType)"/>
      <PopButton class="group-hover:opacity-100 opacity-0" color="red" label="Löschen" type="button"
                 @click="deleteTodoType(todoType)"/>
    </div>
  </div>

  <!-- Create Modal -->
  <Modal :closeable="!isSubmitting" :show="createModalOpen" max-width="2xl" @close="cancelCreateModal">
    <form @submit.prevent="submitCreateForm">
      <fieldset :disabled="isSubmitting" class="p-4">
        <h2 class="text-xl font-medium mb-4">Aufgabentyp erfassen</h2>

        <FormField v-model="createForm.name" label="Name" name="name" required/>
        <FormField v-model="createForm.description" label="Beschreibung" name="description" type="textarea"/>
        <FormField v-model="createForm.color" label="Farbe" name="color" type="color"/>
        <FormField v-model="createForm.reminderTime" :max="30" :min="1" label="Erinnerung (Tage)"
                   name="reminderTime" type="number"/>
      </fieldset>

      <div class="p-4 border-t bg-gray-50 flex justify-end gap-4 flex-wrap">
        <PopButton :disabled="isSubmitting" color="gray" label="Abbrechen" type="button" @click="cancelCreateModal"/>
        <PopButton :disabled="isSubmitting" color="green" label="Erstellen" type="submit"/>
      </div>
    </form>
  </Modal>

  <!-- Edit Modal -->
  <Modal :closeable="!isSubmitting" :show="editingModalOpen" max-width="2xl" @close="cancelCreateModal">
    <form @submit.prevent="submitEditForm">
      <fieldset :disabled="isSubmitting" class="p-4">
        <h2 class="text-xl font-medium mb-4">Aufgabentyp bearbeiten</h2>

        <FormField v-model="editForm.name" label="Name" name="name" required/>
        <FormField v-model="editForm.description" label="Beschreibung" name="description" type="textarea"/>
        <FormField v-model="editForm.color" label="Farbe" name="color" type="color"/>
        <FormField v-model="editForm.reminderTime" :max="30" :min="1" label="Erinnerung (Tage)" name="reminderTime"
                   type="number"/>
      </fieldset>

      <div class="p-4 border-t bg-gray-50 flex justify-end gap-4 flex-wrap">
        <PopButton :disabled="isSubmitting" color="gray" label="Abbrechen" type="button"
                   @click="cancelEditingTodoType"/>
        <PopButton :disabled="isSubmitting" color="green" label="Sichern" type="submit"/>
      </div>
    </form>
  </Modal>
</template>
