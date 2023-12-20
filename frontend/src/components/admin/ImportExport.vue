<script lang="ts" setup>

import PopButton from "../ui/PopButton.vue";
import {useBackend} from "../../services/backend";
import {useTodoStore} from "../../services/store/todo.ts";
import {ref} from "vue";
import {useErrorHandler} from "../../services/errorHandler";

const store = useTodoStore();
const backend = useBackend();

const isWorking = ref(false);

async function openImport() {
    const inputFile = document.createElement('input');
    inputFile.type = 'file';
    inputFile.accept = '.json';
    inputFile.addEventListener('change', (event) => {
        const file = (event.target as HTMLInputElement).files?.[0];
        if (!file) return;

        const reader = new FileReader();
        reader.addEventListener('load', (event) => importData(event.target?.result ?? undefined));
        reader.readAsText(file);
    });
    inputFile.click();
}

async function importData(data: string | ArrayBuffer | undefined) {
    isWorking.value = true;
    try {
        if (!data) {
            throw new Error('Data stream not readable');
        }

        // If data is ArrayBuffer, convert to string
        if (data instanceof ArrayBuffer) {
            const decoder = new TextDecoder('utf-8');
            data = decoder.decode(data);
        }

        await backend.importData(data);
        await store.initialize();
    } catch (e: any) {
        useErrorHandler().handle(e);
    }
    isWorking.value = false;
}


async function downloadExport() {
    isWorking.value = true;
    const data = await backend.createExport();


    const blob = new Blob([data], {type: 'application/json'});
    const url = URL.createObjectURL(blob);

    const link = document.createElement('a');
    link.href = url;
    link.download = 'export.json';
    link.click();

    isWorking.value = false;
}

</script>

<template>

    <!-- Working overlay -->
    <transition enter-active-class="transition-all duration-1000"
                enter-from-class="backdrop-saturate-100"
                enter-to-class="backdrop-saturate-0">
        <div v-if="isWorking" class="z-50 fixed inset-0"/>
    </transition>

    <!-- Title -->
    <div class=" sticky top-0 bg-white/80 backdrop-blur-sm py-6 border-b border-white z-40">
        <h2 class="text-xl font-medium">Datenimport und -export</h2>
    </div>


    <!-- Import -->
    <h3 class="text-lg font-medium mb-2">Daten importieren</h3>
    <div class="md:flex items-center flex-wrap justify-between gap-2 mb-6">
        <p>
            Zuvor exportierte Daten können hier importiert werden.<br/>
            <span class="text-red-800 font-bold">Beim Importieren werden alle vorhandenen Daten gelöscht!</span>
        </p>
        <PopButton :disabled="isWorking" color="amber" @click="openImport">Datei wählen...</PopButton>
    </div>

    <!-- Export -->
    <h3 class="text-lg font-medium mb-2">Export herunterladen</h3>
    <div class="md:flex items-center flex-wrap justify-between gap-2">
        <p>
            Hier kannst du deine Daten exportieren.<br/>
            Du erhältst eine JSON-Datei mit allen deinen Daten.
        </p>
        <PopButton :disabled="isWorking" color="green" @click="downloadExport">Herunterladen</PopButton>
    </div>

</template>
