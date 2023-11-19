<script lang="ts" setup>

import {User} from "../../services/backend";
import UserAvatar from "./UserAvatar.vue";

type UserAvatarListProps = {
    users: User[],
};

const props = defineProps<UserAvatarListProps>();

</script>

<template>

    <UserAvatar v-for="user in users" :key="user.id" class="-mr-1.5" user="user"/>
</template>
