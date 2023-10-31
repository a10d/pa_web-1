<script setup lang="ts">

import { useSound } from "@vueuse/sound";
import { computed } from "vue";
import popSound from "../../../assets/pop-or-bloop-7008.mp3";

type PopButtonProps = {
  label?: string;
  type?: "submit" | "button";
  color?: "green" | "red" | "gray" | "amber";
};

const props = withDefaults(defineProps<PopButtonProps>(), {
  label: "",
  type: "button",
  color: "gray",
});

const { play } = useSound(popSound, {
  sprite: {
    green: [150, 200],
    red: [520, 200],
    gray: [950, 200],
    amber: [1480, 200],
  },
});

const onMouseUp = () => {
  play({
    id: props.color,
    volume: 0.7,
    interrupt: true,
    sprite: true,
  });
};

const buttonClassNames = computed(() => {

  const classNames = [
    "rounded-lg px-4 py-2 pb-3",
    "font-medium text-sm",
    "transition-all",
    "active:translate-y-0.5 active:scale-y-95 active:scale-x-105",
    "duration-200 ease-in-out",
    "hover:shadow focus:shadow",
  ];

  if ( props.color === "green" ) {
    classNames.push(...[
      "text-white",
      "bg-green-600",
      "shadow-green-800",
      "hover:bg-green-500 hover:shadow-green-500",
      "focus:bg-green-500 focus:shadow-green-500",
      "active:bg-green-700",
    ]);
  }

  if ( props.color === "red" ) {
    classNames.push(...[
      "text-white",
      "bg-red-600",
      "shadow-red-800",
      "hover:bg-red-500 hover:shadow-red-500",
      "focus:bg-red-500 focus:shadow-red-500",
      "active:bg-red-700",
    ]);
  }

  if ( props.color === "gray" ) {
    classNames.push(...[
      "text-black",
      "bg-gray-100",
      "shadow-gray-200",
      "hover:bg-gray-200 hover:shadow-gray-300",
      "focus:bg-gray-200 focus:shadow-gray-300",
      "active:bg-gray-300",
    ]);
  }

  if ( props.color === "amber" ) {
    classNames.push(...[
      "text-black",
      "bg-amber-400",
      "shadow-amber-300",
      "hover:bg-amber-300 hover:shadow-amber-300",
      "focus:bg-amber-300 focus:shadow-amber-300",
      "active:bg-amber-400",
    ]);
  }

  return classNames;
});

</script>

<template>
  <button :type="type" :class="buttonClassNames" @mouseup="onMouseUp">
    <span v-if="label" v-text="label"/>
    <slot v-else/>
  </button>
</template>

<style scoped>

button {
  @apply relative;
  text-shadow: 0 1px 1px rgba(0, 0, 0, .2);
}

button::after {
  @apply absolute inset-0 rounded-lg;
  content: '';
  box-shadow: inset 0 -.25rem 0 rgba(0, 0, 0, .2);
}

button::before {
  @apply absolute inset-px rounded-lg transition-all;
  bottom: calc(.25rem + 1px);
  content: '';
  box-shadow: inset 0 0 0.125rem rgba(255, 255, 255, .2);
}

button:hover::before {
  box-shadow: inset 0 0 0.125rem rgba(255, 255, 255, .4);
}

button:active::after {
  @apply absolute inset-0 rounded-lg;
  content: '';
  box-shadow: inset 0 -.25rem 0 rgba(0, 0, 0, .2);
}

</style>
