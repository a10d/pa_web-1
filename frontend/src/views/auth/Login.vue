<script lang="ts" setup>

import {computed, reactive, ref} from "vue";
import AppLayout from "../../components/layouts/AppLayout.vue";
import FormField from "../../components/ui/FormField.vue";
import PopButton from "../../components/ui/PopButton.vue";
import {useAuthStore} from "../../services/store/auth";

const authStore = useAuthStore();

const isSubmitting = ref(false);

const loginForm = reactive({
    username: "",
    password: "",
});

function setUsernameAndPassword(to: string) {
    loginForm.username = to;
    loginForm.password = to;
}

async function submitForm() {
    isSubmitting.value = true;

    try {
        await authStore.login(loginForm);
    } catch (e) {
        console.error(e);
    }

    isSubmitting.value = false;
}

const isFirstVisit = computed<boolean>(() => {
    const firstVisit = localStorage.getItem("firstVisit");
    if (firstVisit === null) {
        localStorage.setItem("firstVisit", "true");
    }
    return (firstVisit === null);
});

</script>

<template>
    <AppLayout>

        <div class="max-w-lg mx-auto border rounded-xl shadow-lg py-4 px-2 lg:py-6 lg:px-4">

            <div class="px-2">
                <h1 class="text-4xl tracking-wide font-medium">Anmelden</h1>
                <p v-if="isFirstVisit" class="my-2">
                    HAllo erstmal. Willkommen bei der echt voll nützlichen Todo-App, die echt voll nützlich ist!<br>
                    Um loszulegen, musst du dich erstmal anmelden.
                </p>
                <p v-if="!isFirstVisit" class="my-2">
                    Ja moin moin. Schön, dass du wieder da bist!
                </p>

                <p class="my-2">
                    Du kannst dich mit einem der folgenden Benutzer anmelden:
                    <button class="text-sm underline mr-1" @click="setUsernameAndPassword('admin')">admin</button>
                    <button class="text-sm underline mr-1" @click="setUsernameAndPassword('user1')">user1</button>
                    <button class="text-sm underline mr-1" @click="setUsernameAndPassword('user2')">user2</button>
                </p>
            </div>

            <form @submit.prevent="submitForm">
                <fieldset :disabled="isSubmitting">
                    <FormField v-model="loginForm.username" label="Benutzername" name="username" required/>
                    <FormField v-model="loginForm.password" label="Passwort" name="password" required type="password"/>

                    <div class="px-2 flex justify-end">
                        <PopButton color="green" label="Anmelden" type="submit"/>
                    </div>
                </fieldset>
            </form>

        </div>
    </AppLayout>
</template>
