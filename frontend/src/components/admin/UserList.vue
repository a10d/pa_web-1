<script lang="ts" setup>

import {useTodoStore} from "../../services/store/todo.ts";
import {computed, nextTick, reactive, ref} from "vue";
import PopButton from "../ui/PopButton.vue";
import Modal from "../ui/Modal.vue";
import FormField from "../ui/FormField.vue";
import {User, ValidationError} from "../../services/backend";
import FormErrors from "../ui/FormErrors.vue";
import {useErrorHandler} from "../../services/errorHandler";
import UserAvatar from "../ui/UserAvatar.vue";
import {useEventBus} from "../../services/eventBus";

const store = useTodoStore();
const eventBus = useEventBus();

const users = computed(() => store.users);

const createForm = reactive({
    firstName: "",
    lastName: "",
    email: "",
})
const createModalOpen = ref(false);

function openCreateModal() {
    createModalOpen.value = true;
    createForm.firstName = "";
    createForm.lastName = "";
    createForm.email = "";
    formError.value = null;
    nextTick(() => {
        const input = document.querySelector("input[name=firstName]") as HTMLInputElement;
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
        await store.createUser(createForm);
        cancelCreateModal();
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

async function deleteUser(user: User) {
    try {
        await store.deleteUser(user);
    } catch (e: any) {
        useErrorHandler().handle(e);
    }
}


const editingUser = ref<User | null>(null);
const editingModalOpen = computed(() => !!editingUser.value);
const editForm = reactive({
    firstName: "",
    lastName: "",
    email: "",
});

function editUser(user: User) {
    editingUser.value = user;
    editForm.firstName = user.firstName;
    editForm.lastName = user.lastName;
    editForm.email = user.email;
    formError.value = null;
    nextTick(() => {
        const input = document.querySelector("input[name=firstName]") as HTMLInputElement;
        input.focus();
    });
}

function cancelEditingUser() {
    editingUser.value = null;
}

async function submitEditForm() {
    isSubmitting.value = true;
    formError.value = null;
    try {
        await store.updateUser({
            ...editingUser.value,
            ...editForm,
        } as User);
        cancelEditingUser();
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
    <div class="relative">
        <!-- Header -->
        <div
            class="flex items-center justify-between gap-2 sticky top-0 py-4 bg-white/80 backdrop-blur-sm border-b border-white z-40">
            <h2 class="text-xl font-medium">Personen</h2>
            <PopButton color="gray" label="Erfassen" type="button" @click="openCreateModal"/>
        </div>


        <!-- List -->
        <div class="p-2 pb-0.5 rounded-lg bg-gray-100 shadow-inner">

            <!-- No Content -->
            <p v-if="users.length === 0" class="text-center text-sm text-gray-800 my-12">Es wurden noch keine Personen
                erfasst...</p>

            <!-- List Items -->
            <div v-for="user in users" :key="user.id"
                 class="rounded bg-white p-2 mb-2 shadow flex flex-wrap justify-end gap-2 group items-center">


                <div class="mr-auto w-full sm:w-auto flex items-center">
                    <UserAvatar :user="user" class="h-8"/>
                    <div class="ml-3">
                        <p class="font-medium text-lg" v-text="`${user.firstName} ${user.lastName}`"/>
                        <p class="text-base text-gray-600" v-text="user.email"/>
                    </div>
                </div>


                <PopButton class="group-hover:opacity-100 sm:opacity-0" color="gray" label="Bearbeiten" type="button"
                           @click="editUser(user)"/>
                <PopButton class="group-hover:opacity-100 sm:opacity-0" color="red" label="Löschen" type="button"
                           @click="deleteUser(user)"/>
            </div>
        </div>
    </div>

    <!-- Create Modal -->
    <Modal :closeable="!isSubmitting" :show="createModalOpen" max-width="2xl" @close="cancelCreateModal">
        <form @submit.prevent="submitCreateForm">
            <fieldset :disabled="isSubmitting" class="p-4">
                <h2 class="text-xl font-medium mb-4">Person erfassen</h2>
                <FormErrors :error="formError"/>
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div>
                        <FormField v-model="createForm.firstName" label="Vorname" name="firstName" required/>
                        <FormErrors :error="formError" field="firstName"/>
                    </div>
                    <div>
                        <FormField v-model="createForm.lastName" label="Nachname" name="name" required/>
                        <FormErrors :error="formError" field="lastName"/>
                    </div>
                </div>
                <FormField v-model="createForm.email" label="E-Mail" name="email" required type="email"/>
                <FormErrors :error="formError" field="email"/>
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
                <h2 class="text-xl font-medium mb-4">Person bearbeiten</h2>
                <FormErrors :error="formError"/>
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div>
                        <FormField v-model="editForm.firstName" label="Vorname" name="firstName" required/>
                        <FormErrors :error="formError" field="firstName"/>
                    </div>
                    <div>
                        <FormField v-model="editForm.lastName" label="Nachname" name="name" required/>
                        <FormErrors :error="formError" field="lastName"/>
                    </div>
                </div>
                <FormField v-model="editForm.email" label="E-Mail" name="email" required type="email"/>
                <FormErrors :error="formError" field="email"/>
            </fieldset>

            <div class="p-4 border-t bg-gray-50 flex justify-between gap-4 flex-wrap">
                <PopButton :disabled="isSubmitting" color="gray" label="Abbrechen" type="button"
                           @click="cancelEditingUser"/>
                <PopButton :disabled="isSubmitting" color="green" label="Sichern" type="submit"/>
            </div>
        </form>
    </Modal>
</template>
