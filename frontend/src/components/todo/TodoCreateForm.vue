<script lang="ts" setup>

import {computed, nextTick, reactive, ref} from "vue";
import FormField from "../ui/FormField.vue";
import PopButton from "../ui/PopButton.vue";
import {useTodoStore} from "../../services/store/todo.ts";

const store = useTodoStore();


const selectableTodoTypes = computed(() => store.todoTypes.map((type) => ({
    value: type.id,
    label: type.name,
})));

const selectableAssignees = computed(() => store.users.map((user) => ({
    value: user.id,
    label: user.name,
})));

const todoForm = reactive({
    title: "",
    description: "",
    dueDate: new Date(),
    type: 0,
    assignees: [],
});

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

    todoForm.title = "";
    todoForm.description = "";
    todoForm.dueDate = new Date();
    todoForm.type = store.todoTypes[0].id;
    todoForm.assignees = [];

    formVisible.value = true;
    await nextTick();
    const input = document.querySelector("input[name=title]") as HTMLInputElement;
    input.focus();
}

function cancel() {
    formVisible.value = false;
}

async function submitForm() {
    isSubmitting.value = true;
    try {
        await store.createTodo(todoForm);
        cancel();
    } catch (e) {
        // TODO: Handle error
    } finally {
        isSubmitting.value = false;
    }
}

</script>

<template>

    <!--  Action bar -->
    <input
        :placeholder="todoCTA"
        class="transition-all sticky top-6 bg-white focus:shadow-lg border rounded-lg w-full p-4 focus:ring outline-0 tracking-wide font-medium"
        @focus="openActionBar"
    />

    <div v-if="formVisible" class="fixed inset-0 bg-white/10 backdrop-blur-sm" @click="cancel"/>

    <!-- Form -->
    <transition enter-active-class="transition-all transform transform-gpu duration-250 ease-out origin-top"
                enter-from-class="scale-y-0 opacity-0"
                leave-active-class="transition-all transform transform-gpu duration-300 ease-out origin-top"
                leave-to-class="scale-y-0 opacity-0">
        <div v-if="formVisible" class="fixed top-4 left-1/2 -translate-x-1/2 w-full max-w-2xl p-2">
            <div class="mx-auto rounded-lg bg-white p-4 shadow-2xl border">

                <fieldset :disabled="isSubmitting">
                    <label
                        for="title"
                        class="inline-flex items-center gap-2 px-2 text-sm font-medium"
                    >
                        Titel
                    </label>

                    <input v-model="todoForm.title" :placeholder="todoCTA"
                           class="w-full outline-none py-1 mb-4 tracking-wide font-medium"
                           name="title" required type="text"/>

                    <FormField
                        v-model="todoForm.description"
                        label="Genauere Beschreibung"
                        name="description"
                        placeholder="Beschreibe die Aufgabe etwas genauer..."
                        type="textarea"
                    />

                    <div class="grid grid-cols-2 gap-4">

                        <div class="col-span-2 sm:col-span-1">
                            <FormField v-model="todoForm.dueDate" label="Fällig am" name="dueDate" type="date"/>
                        </div>

                        <div class="col-span-2 sm:col-span-1">
                            <FormField v-model="todoForm.type"
                                       :select-options="selectableTodoTypes"
                                       label="Aufgabentyp"
                                       name="todoType"
                                       type="select"/>
                        </div>

                    </div>



                    <FormField v-model="todoForm.assignees"
                               :select-options="selectableAssignees"
                               label="Zuständige Personen"
                               multiple
                               name="assignees"
                               type="select"/>

                    <div class="flex justify-between gap-4">
                        <PopButton
                            color="gray"
                            label="Abbrechen"
                            @click="cancel"
                        />
                        <PopButton
                            color="green"
                            label="Loslegen"
                            @click="submitForm"
                        />
                    </div>

                </fieldset>

            </div>
        </div>
    </transition>

</template>
