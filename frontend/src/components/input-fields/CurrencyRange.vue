<template>
  <div class="input">
    <div>
      Currency
      <select v-model="selectedCurrency" @change="$emit('currencyInput', $event.target.value)">
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

  private selectedCurrency = '';

  protected mounted(): void {
    this.min = this.minValue;
    if (this.maxValue !== undefined) {
      this.max = this.maxValue;
    }

    this.handleMinInput(this.min);
    this.handleMaxInput(this.max);

    [this.selectedCurrency] = this.currencyOptions;
    this.$emit('currencyInput', this.selectedCurrency);
  }
}
</script>

<style scoped>

</style>
