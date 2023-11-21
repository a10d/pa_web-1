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

</script>

<template>
    <div class="aspect-square rounded-full shadow-inner">
        <span class="text-black/30" v-text="initials"/>
    </div>
</template>
