<script lang="ts" setup>

import {useEventBus} from "../../services/eventBus";

type PopButtonProps = {
    label?: string;
    type?: "submit" | "button";
    color?: "green" | "red" | "gray" | "amber";
    disabled?: boolean;
};

const props = withDefaults(defineProps<PopButtonProps>(), {
    label: "",
    type: "button",
    color: "gray",
    disabled: false,
});

const emit = defineEmits<{
    click: [];
}>();

function onButtonClick() {
    emit('click');

    useEventBus().emit('playSound', 'buttonClick', {color: props.color});
}
</script>

<template>
    <button :class="['boettong', props.color]" :disabled="disabled" :type="type" @click="onButtonClick">
        <span v-if="label" v-text="label"/>
        <slot v-else/>
    </button>
</template>

<style>

.boettong {
    @apply relative rounded-lg px-4 py-2 pb-3 font-medium text-sm transition-all duration-200 ease-in-out cursor-pointer;
    text-shadow: 0 1px 1px rgba(0, 0, 0, .2);
}

.boettong:disabled {
    @apply opacity-50 cursor-not-allowed;
}

.boettong:not(:disabled) {
    @apply shadow-sm;
}

.boettong:hover:not(:disabled) {
    @apply shadow-md;
}

.boettong:active:not(:disabled) {
    @apply scale-x-105 scale-y-95 translate-y-0.5 transform shadow-sm;
}

.boettong::after {
    @apply absolute inset-0 rounded-lg;
    content: '';
    box-shadow: inset 0 -.25rem 0 rgba(0, 0, 0, .2);
}

.boettong::before {
    @apply absolute inset-px rounded-lg transition-all;
    bottom: calc(.25rem + 1px);
    content: '';
    box-shadow: inset 0 0 0.125rem rgba(255, 255, 255, .2);
}

.boettong:hover:not(:disabled)::before {
    box-shadow: inset 0 0 0.125rem rgba(255, 255, 255, .4);
}

.boettong:active:not(:disabled)::after {
    @apply absolute inset-0 rounded-lg;
    content: '';
    box-shadow: inset 0 -.25rem 0 rgba(0, 0, 0, .2);
}


.boettong.green {
    @apply text-white bg-green-600 shadow-green-800;
}

.boettong.green:not(:disabled) {
    @apply hover:bg-green-500 hover:shadow-green-500 focus:bg-green-500 focus:shadow-green-500 active:bg-green-700;
}


.boettong.red {
    @apply text-white  bg-red-600  shadow-red-800 ;
}

.boettong.red:not(:disabled) {
    @apply hover:bg-red-500 hover:shadow-red-500  focus:bg-red-500 focus:shadow-red-500  active:bg-red-700;
}

.boettong.gray {
    @apply text-black bg-gray-100 shadow-gray-200;
}

.boettong.gray:not(:disabled) {
    @apply hover:bg-gray-200 hover:shadow-gray-300  focus:bg-gray-200 focus:shadow-gray-300  active:bg-gray-300;
}

.boettong.amber {
    @apply text-black bg-amber-400 shadow-amber-300;
}

.boettong.amber:not(:disabled) {
    @apply hover:bg-amber-300 hover:shadow-amber-300 focus:bg-amber-300 focus:shadow-amber-300 active:bg-amber-400;
}


</style>
