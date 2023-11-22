<script lang="ts" setup>

import SoundEffects from "../SoundEffects.vue";
</script>

<template>
    <div class="mx-auto p-2 max-w-2xl flex flex-col h-screen relative">
        <header
            class="static md:absolute right-full h-20 md:w-20 text-gray-400 flex md:flex-col gap-4 justify-center items-center my-8">
            <RouterLink active-class="text-black" to="/">
                <svg class="w-16 h-16"
                     style="fill-rule:evenodd;clip-rule:evenodd;stroke-linecap:round;stroke-linejoin:round;"
                     viewBox="0 0 500 500"
                     xml:space="preserve" xmlns="http://www.w3.org/2000/svg">
                    <g transform="matrix(1.68746,0,0,1.68746,-182.182,-279.438)">
                        <g id="layers.svg" transform="matrix(7.64305,0,0,8.09962,234.195,334.213)">
                            <g transform="matrix(1,0,0,1,-12,-12)">
                                <path d="M2,17L12,22L22,17"
                                      style="fill:none;fill-rule:nonzero;stroke:currentColor;stroke-width:2px;"/>
                                <path d="M2,12L12,17L22,12"
                                      style="fill:none;fill-rule:nonzero;stroke:currentColor;stroke-width:2px;"/>
                            </g>
                        </g>
                        <g id="check.svg" transform="matrix(9.0051,0,0,8.71785,259.839,295.236)">
                            <g transform="matrix(1,0,0,1,-12,-12)">
                                <path d="M20,6L9,17L4,12"
                                      style="fill:none;fill-rule:nonzero;stroke:currentColor;stroke-width:2px;"/>
                            </g>
                        </g>
                    </g>
                </svg>
            </RouterLink>
            <RouterLink active-class="text-black" to="/admin">
                <svg class="w-8 h-8" fill="none" height="24" stroke="currentColor" stroke-linecap="round"
                     stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" width="24"
                     xmlns="http://www.w3.org/2000/svg">
                    <circle cx="12" cy="12" r="3"></circle>
                    <path
                        d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
                </svg>
            </RouterLink>
        </header>

        <main class="flex-grow">
            <slot/>
        </main>

        <footer>
            <p class="text-gray-500 text-xs select-none my-8">
                <span>&copy; 2023 Aaron Schmied</span>
                &bull;
                <span>Arbeit der ipso Bildung</span>
            </p>
            <SoundEffects/>
        </footer>
    </div>
</template>
