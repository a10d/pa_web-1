<template>
    <div v-if="isCheckboxLike" class="mb-2">
        <label
            :for="name"
            class="inline-flex items-center gap-2 px-2 "
        >
            <input-field
                v-model:value="valueProxy"
                :disabled="disabled ?? false"
                :name="name"
                :required="required ?? false"
                :type="type ?? 'checkbox'"
            />
            <span class="text-sm font-medium text-slate-800 cursor-pointer" v-text="label ?? name"/>
        </label>
    </div>
    <div v-if="!isCheckboxLike" class="mb-4">
        <label
            :for="name"
            class="block text-sm font-medium text-slate-800 px-2 mb-2 "
            v-text="label ?? name"
        />
        <input-field
            v-model:value="valueProxy"
            :disabled="disabled"
            :max="max"
            :min="min"
            :multiple="multiple"
            :name="name"
            :placeholder="placeholder"
            :required="required"
            :select-options="selectOptions ?? []"
            :step="step"
            :type="type ?? 'text'"
        />
    </div>
</template>

<script lang="ts" setup>

import {computed} from "vue";
import InputField from "./InputField.vue";

type FormFieldProps = {
    modelValue: string | number | boolean | object | null | any[];
    type?: "text" | "textarea" | "select" | "file" | "checkbox" | "radio" | string;
    selectOptions?: SelectOption[];
    label?: string;
    name: string;
    required?: boolean;
    disabled?: boolean;
    multiple?: boolean;
    min?: number;
    max?: number;
    step?: number;
    placeholder?: string;
};

export interface SelectOption {
    value: string | number | boolean;
    label: string;
    insecureAsHtml?: boolean;
    disabled?: boolean;
}

const isCheckboxLike = computed(() => ["checkbox", "radio"].includes(props.type.toLowerCase()));

const props = withDefaults(
    defineProps<FormFieldProps>(),
    {
        modelValue: null,
        type: "text",
        required: false,
        multiple: false,
        disabled: false,
        placeholder: '',
    },
);

const emit = defineEmits<{
    "update:modelValue": [value: string | number | boolean | object | null | any[]],
}>();

const valueProxy = computed({
    get: () => props.modelValue,
    set: (v) => emit("update:modelValue", v),
});
</script>
