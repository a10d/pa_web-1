<script lang="ts" setup>

import {useTodoStore} from "../../services/store/todo.ts";
import {computed, nextTick, reactive, ref} from "vue";
import PopButton from "../ui/PopButton.vue";
import Modal from "../ui/Modal.vue";
import FormField from "../ui/FormField.vue";
import {User} from "../../services/backend";

const store = useTodoStore();

const users = computed(() => store.users);

const createForm = reactive({
  name: "",
})
const createModalOpen = ref(false);

function openCreateModal() {
  createModalOpen.value = true;
  createForm.name = "";

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
    await store.createUser(createForm);
    cancelCreateModal();
  } catch (e) {
    // TODO: Error handling
  } finally {
    isSubmitting.value = false;
  }
}

async function deleteUser(user: User) {
  try {
    await store.deleteUser(user);
  } catch (e) {
    // TODO: Error handling
  }
}


const editingUser = ref<User | null>(null);
const editingModalOpen = computed(() => !!editingUser.value);
const editForm = reactive({
  name: "",
});

function editUser(user: User) {
  editingUser.value = user;
  editForm.name = user.name;

  nextTick(() => {
    const input = document.querySelector("input[name=name]") as HTMLInputElement;
    input.focus();
  });
}

function cancelEditingUser() {
  editingUser.value = null;
}

async function submitEditForm() {
  isSubmitting.value = true;
  try {
    await store.updateUser({
      ...editingUser.value,
      ...editForm,
    } as User);
    cancelEditingUser();
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
    <h2 class="text-xl font-medium">Personen</h2>
    <PopButton color="gray" label="Erfassen" type="button" @click="openCreateModal"/>
  </div>


  <!-- List -->
  <div class="p-2 pb-0.5 rounded-lg bg-gray-100 shadow-inner">

    <!-- No Content -->
    <p v-if="users.length === 0" class="text-center text-sm text-gray-800 my-12">Es wurden noch keine Personen
      erfasst...</p>

    <!-- List Items -->
    <div v-for="user in users" :key="user.id" class="rounded bg-white p-2 mb-2 shadow flex gap-2 group items-center">
      <p class="mr-auto font-medium text-lg" v-text="user.name"/>

      <PopButton class="group-hover:opacity-100 opacity-0" color="gray" label="Bearbeiten" type="button"
                 @click="editUser(user)"/>
      <PopButton class="group-hover:opacity-100 opacity-0" color="red" label="Löschen" type="button"
                 @click="deleteUser(user)"/>
    </div>
  </div>

  <!-- Create Modal -->
  <Modal :closeable="!isSubmitting" :show="createModalOpen" max-width="2xl" @close="cancelCreateModal">
    <form @submit.prevent="submitCreateForm">
      <fieldset :disabled="isSubmitting" class="p-4">
        <h2 class="text-xl font-medium mb-4">Person erfassen</h2>

        <FormField v-model="createForm.name" label="Name" name="name" required/>
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
        <h2 class="text-xl font-medium mb-4">Person bearbeiten</h2>
        <FormField v-model="editForm.name" label="Name" name="name" required/>
      </fieldset>

      <div class="p-4 border-t bg-gray-50 flex justify-end gap-4 flex-wrap">
        <PopButton :disabled="isSubmitting" color="gray" label="Abbrechen" type="button" @click="cancelEditingUser"/>
        <PopButton :disabled="isSubmitting" color="green" label="Sichern" type="submit"/>
      </div>
    </form>
  </Modal>
</template>
