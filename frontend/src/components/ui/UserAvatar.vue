<script lang="ts" setup>

import {computed} from "vue";
import {User} from "../../services/backend";

type UserAvatarProps = {
    user: User,
};

const props = defineProps<UserAvatarProps>();

const initials = computed(() => [props.user.firstName, props.user.lastName]
    .map((name) => name[0])
    .join(""));

const userColor = computed(() => {
    let hash = 0;
    props.user.id.split('').forEach(char => {
        hash = char.charCodeAt(0) + ((hash << 5) - hash)
    })
    let colour = '#'
    for (let i = 0; i < 3; i++) {
        const value = (hash >> (i * 8)) & 0xff
        colour += value.toString(16).padStart(2, '0')
    }
    return colour
});

</script>

<template>
    <div class="aspect-square rounded-full shadow-inner bg-white inline-flex relative select-none z-10">
        <div :style="{backgroundColor: userColor}" class="absolute inset-0 rounded-full opacity-60 z-0"/>
        <span class="text-black/30 m-auto z-20" v-text="initials"/>
    </div>
</template>
