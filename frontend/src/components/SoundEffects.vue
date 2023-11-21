<script lang="ts" setup>

import {useEventBus} from "../services/eventBus";
import {useSound} from "@vueuse/sound";
import popSound from '../assets/sounds.mp3'
import {onMounted} from "vue";
import {useStorage} from "@vueuse/core";

const soundMap = {
    'todoCompleted': 'scratch',
    'todoUncompleted': 'erase'
}

const {play} = useSound(popSound, {
    sprite: {
        'ding': [0, 350],
        'pock': [400, 300],
        'thud': [750, 250],
        'scratch': [1000, 600],
        'erase': [1600, 500],
        'klack': [2150, 300],
        'whop': [2500, 250],
        'tang': [2800, 300],
    },
});


// preload audio context by playing a sound
onMounted(() => play({id: 'ding', volume: 0.01,}));

const eventBus = useEventBus();

const soundsEnabled = useStorage('soundsEnabled', true);
const toggleSounds = () => {
    soundsEnabled.value = !soundsEnabled.value;
    eventBus.emit('playSound', 'enableSounds')
};

eventBus.on('playSound', (soundEvent: string) => {
    if (!soundsEnabled.value) return;

    if (!soundMap.hasOwnProperty(soundEvent)) {
        console.info('No mapped sound for event', soundEvent)
        return;
    }

    play({
        id: soundMap[soundEvent],
        volume: 0.4,
        interrupt: true,
        sprite: true,
    });
})

</script>

<template>
    <button
        :title="soundsEnabled ? 'Klicken um Soundeffekte zu deaktivieren' : 'Ziemlich coole Soundeffekte aktivieren'"
        class="fixed bottom-4 right-4 aspect-square w-6 text-black/50 rounded-full transition-all hover:text-black/75 focus:ring focus:scale-110 hover:scale-110 active:scale-90"
        @click="toggleSounds">
        <svg v-if="soundsEnabled" class="feather feather-volume-2" fill="none" height="24" stroke="currentColor"
             stroke-linecap="round" stroke-linejoin="round"
             stroke-width="2" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg">
            <polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon>
            <path d="M19.07 4.93a10 10 0 0 1 0 14.14M15.54 8.46a5 5 0 0 1 0 7.07"></path>
        </svg>
        <svg v-if="!soundsEnabled" class="feather feather-volume-x" fill="none" height="24" stroke="currentColor"
             stroke-linecap="round" stroke-linejoin="round"
             stroke-width="2" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg">
            <polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon>
            <line x1="23" x2="17" y1="9" y2="15"></line>
            <line x1="17" x2="23" y1="9" y2="15"></line>
        </svg>
    </button>


</template>
