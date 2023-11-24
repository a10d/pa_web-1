<script lang="ts" setup>

import {User} from "../../services/backend";
import UserAvatar from "./UserAvatar.vue";
import {computed} from "vue";

type UserAvatarListProps = {
    users: User[],
};

const maxVisibleUsers = 4;

const props = defineProps<UserAvatarListProps>();

const count = computed(() => props.users.length);

const visibleUsers = computed<User[]>(() => props.users.slice(1, maxVisibleUsers));

</script>

<template>

    <div class="inline-flex text-xs font-bold items-center">
        <UserAvatar v-if="count > 0" :user="props.users[0]" class="h-6 z-20 border border-white shadow-sm"/>
        <UserAvatar v-for="(user, i) in visibleUsers" :key="user.id" :style="{zIndex: count - i}"
                    :user="user"
                    class="h-6 -ml-1 border border-white border-r-3 shadow-sm"/>

        <p v-if="count > maxVisibleUsers" class="ml-1 text-xs">
            +{{ count - maxVisibleUsers }}
        </p>
    </div>

</template>
