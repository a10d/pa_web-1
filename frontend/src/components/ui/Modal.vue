<script lang="ts" setup>
import {computed, nextTick, onMounted, onUnmounted, watch} from "vue";

type ModalProps = {
    show: boolean;
    maxWidth?: "sm" | "md" | "lg" | "xl" | "2xl" | "4xl" | "6xl";
    closeable?: boolean;
}

const props = withDefaults(
    defineProps<ModalProps>(),
    {
        show: false,
        maxWidth: '2xl',
        closeable: true,
    }
);

const emit = defineEmits<{
    'close': [];
    'show': [];
}>();

watch(() => props.show, () => {
    if (props.show) {
        nextTick().then(() => emit("show"));
        document.body.style.overflow = "hidden";
    } else {
        document.body.style.overflow = "auto";
    }
});

const close = () => {
    if (props.closeable) {
        emit("close");
    }
};

const closeOnEscape = (e: KeyboardEvent) => {
    if (e.key === "Escape" && props.show) {
        close();
    }
};

onMounted(() => document.addEventListener("keydown", closeOnEscape));

onUnmounted(() => {
    document.removeEventListener("keydown", closeOnEscape);
    document.body.style.overflow = "auto";
});

const maxWidthClass = computed(() => {
    return {
        "sm": "sm:max-w-sm",
        "md": "sm:max-w-md",
        "lg": "sm:max-w-lg",
        "xl": "sm:max-w-xl",
        "2xl": "sm:max-w-2xl",
        "4xl": "sm:max-w-4xl",
        "6xl": "sm:max-w-6xl",
    }[props.maxWidth];
});
</script>


<template>
    <teleport to="body">
        <transition leave-active-class="duration-200">
            <div v-show="show" class="fixed flex inset-0 overflow-y-auto px-4 py-6 sm:px-0 z-50" scroll-region>

                <transition leave-active-class="transition duration-200">
                    <div v-show="show" class="fixed inset-0 transform" @click="close">
                        <transition
                            enter-active-class="ease-out duration-300"
                            enter-from-class="opacity-0"
                            enter-to-class="opacity-100 "
                            leave-active-class="ease-in duration-200"
                            leave-from-class="opacity-100"
                            leave-to-class="opacity-0"
                        >
                            <div v-show="show"
                                 class="absolute inset-0 bg-white/10 backdrop-blur-md backdrop-saturate-0"/>
                        </transition>
                    </div>
                </transition>


                <transition
                    enter-active-class="ease-out duration-300"
                    enter-from-class="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                    enter-to-class="opacity-100 translate-y-0 sm:scale-100"
                    leave-active-class="ease-in duration-200"
                    leave-from-class="opacity-100 translate-y-0 sm:scale-100"
                    leave-to-class="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                >
                    <div v-show="show"
                         :class="maxWidthClass"
                         class="border my-auto bg-white/80 backdrop-blur-xl rounded-lg overflow-hidden shadow-xl transform transition-all sm:w-full sm:mx-auto">
                        <slot/>
                    </div>
                </transition>
            </div>
        </transition>
    </teleport>
</template>
