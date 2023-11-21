<script lang="ts" setup>

import AppLayout from "../components/layouts/AppLayout.vue";
import UserList from "../components/admin/UserList.vue";
import TodoTypeList from "../components/admin/TodoTypeList.vue";
import ImportExport from "../components/admin/ImportExport.vue";
</script>

<template>
    <AppLayout>
        <h1 class="text-2xl font-medium tracking-wider">Einstellungen</h1>
        <UserList/>
        <TodoTypeList/>
        <ImportExport/>
    </AppLayout>
</template>
