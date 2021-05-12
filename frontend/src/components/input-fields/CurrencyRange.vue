<template>
  <div class="input">
    <div>
      Currency
      <select @change="$emit('currencyInput', $event.target.value)">
        <option value=""></option>
        <option v-for="(currency, index) in currencyOptions"
                :key="index"
                :value="currency">{{ currency }}</option>
      </select>
      Min
      <input
        type="number"
        :min="minValue"
        placeholder="0"
        v-model="min"
        @input="handleMinInput"
      >
      -
      Max
      <input
        type="number"
        min="0"
        placeholder="0"
        v-model="max"
        @blur="handleMaxInput"
      >

    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop } from 'vue-property-decorator';
import Range from './Range.vue';

@Component
export default class CurrencyRange extends Range {
  @Prop() private readonly currencyOptions!: string[];
}
</script>

<style scoped>

</style>
