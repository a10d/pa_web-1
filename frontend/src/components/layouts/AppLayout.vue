<script lang="ts" setup>

</script>

<template>

    <div class="mx-auto my-4 p-2 max-w-2xl flex flex-col min-h-screen relative">

        <header class="static sm:absolute right-full h-20 w-20">
            <h1 class="text-xl text-center">Logo</h1>
        </header>

        <main class="flex-grow">
            <slot/>
        </main>

        <footer>
            <p class="text-gray-500 text-xs select-none">
                <span>&copy; 2023 Aaron Schmied</span>
                &bull;
                <span>Arbeit der ipso Bildung</span>
            </p>
        </footer>
    </div>
</template>
