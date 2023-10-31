<script lang="ts" setup>

import {computed, nextTick, reactive, ref} from "vue";
import FormField from "../ui/FormField.vue";
import PopButton from "../ui/PopButton.vue";


const taskForm = reactive({
    title: "",
    description: "",
});

const formVisible = ref(false);

const taskCTAs = [
    "Ohne Fleiss kein Preis! Wo fängst du an?",
    "Was steht noch an?",
    "Was ist noch zu tun?",
    "Was ist noch zu erledigen?",
];
const taskCTA = computed(() => taskCTAs[Math.floor(Math.random() * taskCTAs.length)]);

async function openActionBar() {
    formVisible.value = true;
    await nextTick();
    const input = document.querySelector("input[name=title]") as HTMLInputElement;
    input.focus();
}

function cancel() {
    formVisible.value = false;
}

async function submitForm() {
    alert("Submit");
}

</script>

<template>

    <!--  Action bar -->
    <input
        :placeholder="taskCTA"
        class="transition-all sticky top-6 bg-white focus:shadow-lg border rounded-lg w-full p-4 focus:ring outline-0" @focus="openActionBar"
    />

    <div v-if="formVisible" class="fixed inset-0 bg-white/80 backdrop-blur-lg" @click="cancel"/>

    <!-- Form -->
    <transition enter-active-class="transition-all transform transform-gpu duration-250 ease-out origin-top"
                enter-from-class="scale-y-0 opacity-0"
                leave-active-class="transition-all transform transform-gpu duration-300 ease-out origin-top"
                leave-to-class="scale-y-0 opacity-0">
        <div v-if="formVisible" class="fixed top-4 left-1/2 -translate-x-1/2 w-full max-w-2xl p-2">
            <div class="mx-auto rounded-lg bg-white p-4 shadow-2xl border">

                <fieldset>
                    <FormField
                        v-model="taskForm.title"
                        :placeholder="taskCTA"
                        label="Titel"
                        name="title"
                        required
                    />

                    <FormField
                        v-model="taskForm.description"
                        label="Beschreibung"
                        name="description"
                        placeholder="Beschreibe die Aufgabe etwas genauer..."
                        type="textarea"
                    />

                    <div class="flex justify-between gap-4">
                        <PopButton
                            color="gray"
                            label="Abbrechen"
                            @click="cancel"
                        />
                        <PopButton
                            color="green"
                            label="Sichern"
                            @click="submitForm"
                        />
                    </div>

                </fieldset>

            </div>
        </div>
    </transition>

</template>
