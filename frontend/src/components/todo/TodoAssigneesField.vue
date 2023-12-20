<script lang="ts" setup>

import UserAvatar from "../ui/UserAvatar.vue";
import {useTodoStore} from "../../services/store/todo.ts";
import {computed} from "vue";

type TodoAssigneesFieldProps = {
    modelValue: string[];
}

const props = defineProps<TodoAssigneesFieldProps>();

const emit = defineEmits<{
    'update:modelValue': [value: string[]],
}>();

const valueProxy = computed({
    get: () => props.modelValue,
    set: (value: string[]) => emit('update:modelValue', value),
});

const users = computed(() => useTodoStore().users);

</script>

<template>
    <div
        :class="{'bg-gray-100': users.length === 0, 'bg-white': users.length > 0}"
        class="grid grid-cols-1 sm:grid-cols-2 gap-2 mb-4 p-2 border rounded-md shadow-sm border-gray-300 focus-within:ring">

        <div v-if="users.length === 0" class="text-center text-sm text-gray-500 my-6 sm:col-span-2">Es wurden noch keine
            Personen
            erfasst...
        </div>

        <label v-for="user in users" :key="user.id"
               :class="{'bg-blue-500 text-white': modelValue.includes(user.id)}"
               class="relative flex group focus-within:ring items-center rounded-full p-2 select-none">

            <input v-model="valueProxy" :value="user.id"
                   class="opacity-0 absolute inset-0 cursor-pointer" type="checkbox">

            <UserAvatar :user="user" class="w-8 h-8 mr-2"/>
            <span class="font-medium text-sm" v-text="`${user.firstName} ${user.lastName}`"/>

        </label>
    </div>
</template>
