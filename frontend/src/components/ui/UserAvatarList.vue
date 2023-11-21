<script lang="ts" setup>

import {User} from "../../services/backend";
import UserAvatar from "./UserAvatar.vue";

type UserAvatarListProps = {
    users: User[],
};

defineProps<UserAvatarListProps>();

</script>

<template>

    <UserAvatar v-for="user in users" :key="user.id" :user="user" class="-mr-1.5"/>
</template>
