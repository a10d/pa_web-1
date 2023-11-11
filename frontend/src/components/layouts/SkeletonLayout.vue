<template>
    <div class="fixed inset-0">
        <div class="mx-auto max-w-2xl">
            <!-- Input Box -->
            <div class="animate-pulse bg-gray-200 rounded h-16 my-8"/>

            <!-- Content -->
            <div v-for="_ in 10" class="my-8">
                <div :style="{width: `${Math.random() * 40 + 30}%`}"
                     class="animate-pulse bg-gray-200 rounded h-6 mb-1"/>
                <div v-for="_ in 3" :style="{width: `${Math.random() * 60 + 10}%`}"
                     class="animate-pulse bg-gray-200 rounded h-4 mb-1"/>
            </div>
        </div>
    </div>
</template>
