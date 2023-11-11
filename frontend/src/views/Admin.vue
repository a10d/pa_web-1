<script lang="ts" setup>

import AppLayout from "../components/layouts/AppLayout.vue";
import UserList from "../components/admin/UserList.vue";
import TodoTypeList from "../components/admin/TodoTypeList.vue";
</script>

<template>
  <AppLayout>
    <h1>Admin</h1>
    <UserList/>
    <TodoTypeList/>
  </AppLayout>
</template>
