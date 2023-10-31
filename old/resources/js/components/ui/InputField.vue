<template>
  <select
      v-if="type === 'select'"
      :id="name"
      :class="[inputClasses, 'pb-2 pt-3 pr-3']"
      :disabled="disabled"
      :name="name"
      :required="required"
      :value="value"
      @change="onInput"
  >
    <option
        v-for="{value, label, disabled} in selectOptions"
        :disabled="disabled"
        :value="value"
        v-text="label"
    />
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
}

type InputFieldProps = {
  name: string;
  type: "text" | "number" | "textarea" | "select" | "file" | string;
  required: boolean;
  disabled: boolean;
  selectOptions?: SelectOption[];
  value: string | number | boolean | object | any[] | null;
};

withDefaults(
    defineProps<InputFieldProps>(),
    {
      type: "text",
      required: false,
      disabled: false,
      value: "",
    },
);

const emit = defineEmits(["update:value"]);

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
