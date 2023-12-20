<template>
    <select
        v-if="type === 'select'"
        :id="name"
        :class="[inputClasses, 'pb-2 pt-3 pr-3']"
        :disabled="disabled"
        :multiple="multiple"
        :name="name"
        :required="required"
        :value="value"
        @change="onInput"
    >
        <template v-for="{value, label, disabled, insecureAsHtml} in selectOptions" :key="value">
            <option v-if="insecureAsHtml === true" :disabled="disabled" :value="value" v-html="label"/>
            <option v-else :disabled="disabled" :value="value" v-text="label"/>
        </template>
    </select>
    <input
        v-else-if="type==='number'"
        :id="name"
        :class="inputClasses"
        :disabled="disabled"
        :name="name"
        :required="required"
        :type="type"
        :value="value"
        @input="onInput"
    />
    <textarea
        v-else-if="type==='textarea'"
        :id="name"
        :class="inputClasses"
        :disabled="disabled"
        :name="name"
        :required="required"
        :value="value as string | undefined"
        @input="onInput"
    />
    <input
        v-else-if="type==='file'"
        :id="name"
        :class="inputClasses"
        :disabled="disabled"
        :name="name"
        :required="required"
        type="file"
        @change="onSelectFile"
    />
    <input
        v-else-if="type==='checkbox'"
        :id="name"
        :checked="!!value"
        :disabled="disabled"
        :required="required"
        :type="type"
        @input="onCheckboxInput"
    />
    <div v-else-if="type === 'color'" :class="inputClasses" class="relative h-[42px]">
        <div :style="{backgroundColor: `${value}`}"
             class="absolute inset-1 rounded shadow-inner shadow border border-black/30"/>
        <input
            :disabled="disabled"
            :required="required"
            :value="value"
            class="absolute inset-0 w-full h-full opacity-0 cursor-pointer disabled:cursor-not-allowed"
            type="color"
            @input="onInput"
        />
    </div>

    <input
        v-else
        :id="name"
        :class="inputClasses"
        :disabled="disabled"
        :name="name"
        :required="required"
        :type="type"
        :value="value"
        @input="onInput"
    />
</template>

<script lang="ts" setup>

export interface SelectOption {
    value: string | number | boolean;
    label: string;
    disabled?: boolean;
    insecureAsHtml?: boolean;
}

type InputFieldProps = {
    name: string;
    type: "text" | "number" | "textarea" | "select" | "file" | string;
    required?: boolean;
    disabled?: boolean;
    multiple?: boolean;
    selectOptions?: SelectOption[];
    value: string | number | boolean | object | any[] | null;
};

withDefaults(
    defineProps<InputFieldProps>(),
    {
        type: "text",
        required: false,
        disabled: false,
        multiple: false,
        value: "",
    },
);

const emit = defineEmits<{
    'update:value': [value: string | number | boolean | object | any[] | null],
}>();

const onInput = (e: any) => emit("update:value", e.target.value);

const onCheckboxInput = (e: any) => emit("update:value", e.target.checked);

const onSelectFile = (e: any) => {
    emit("update:value", e.target.files[0]);
};

const inputClasses = [
    "w-full px-2 py-2",
    "border border-gray-300 focus:border-gray-500",
    "rounded-lg shadow-sm",
    "bg-white  text-black/80",
    "disabled:bg-gray-100 disabled:text-gray-500",
    "transition-all ease-in-out duration-200",
    "focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50",
];

</script>
