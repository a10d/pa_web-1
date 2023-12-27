<script lang="ts" setup>

import {computed} from "vue";

type Quote = {
    quote: string,
    author: string,
}

const quotes: Quote[] = [
    {
        quote: 'Das Geheimnis des Erfolgs ist anzufangen.',
        author: 'Mark Twain',
    },
    {
        quote: 'Der Weg ist das Ziel.',
        author: 'Konfuzius',
    },
    {
        quote: 'Wer nichts waget, der darf nichts hoffen.',
        author: 'Friedrich von Schiller',
    },
    {
        quote: 'Räum dein Zimmer auf.',
        author: 'Mama',
    },
];

const quote = computed<Quote>(() => quotes[Math.floor(Math.random() * quotes.length)]);

</script>

<template>
    <div class="my-24 select-none">
        <p class="text-xl italic text-gray-500">
            „{{ quote.quote }}“
        </p>
        <p class="text-gray-500 text-sm font-medium ml-3">
            — {{ quote.author }}
        </p>
    </div>
</template>
