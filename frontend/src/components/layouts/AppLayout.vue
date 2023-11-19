<script lang="ts" setup>

</script>

<template>

    <div class="mx-auto my-4 p-2 max-w-2xl flex flex-col min-h-screen relative">

        <header class="static sm:absolute right-full h-20 w-20">
            <svg  style="fill-rule:evenodd;clip-rule:evenodd;stroke-linecap:round;stroke-linejoin:round;"
                  viewBox="0 0 500 500" class="w-16 h-16"
                 xml:space="preserve" xmlns="http://www.w3.org/2000/svg">
                <g transform="matrix(1.68746,0,0,1.68746,-182.182,-279.438)">
                    <g id="layers.svg" transform="matrix(7.64305,0,0,8.09962,234.195,334.213)">
                        <g transform="matrix(1,0,0,1,-12,-12)">
                            <path d="M2,17L12,22L22,17"
                                  style="fill:none;fill-rule:nonzero;stroke:black;stroke-width:2px;"/>
                            <path d="M2,12L12,17L22,12"
                                  style="fill:none;fill-rule:nonzero;stroke:black;stroke-width:2px;"/>
                        </g>
                    </g>
                    <g id="check.svg" transform="matrix(9.0051,0,0,8.71785,259.839,295.236)">
                        <g transform="matrix(1,0,0,1,-12,-12)">
                            <path d="M20,6L9,17L4,12"
                                  style="fill:none;fill-rule:nonzero;stroke:black;stroke-width:2px;"/>
                        </g>
                    </g>
                </g>
            </svg>

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
