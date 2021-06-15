<template>
  <div class="input">
    <div>
    Price Range:
      <div v-if="currencyOptions">
        Currency
        <select v-model="selectedCurrency" @change="$emit('currencyInput', $event.target.value)">
          <option v-for="(currency, index) in currencyOptions"
                  :key="index"
                  :value="currency">{{ currency }}</option>
        </select>
      </div>
      Min
      <input
        type="number"
        min="0"
        v-model="min"
        @input="handleMinInput"
      >
      -
      Max
      <input
        type="number"
        min="0"
        v-model="max"
        @blur="handleMaxInput"
      >

    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class CurrencyRange extends Vue {
  @Prop({ required: false }) private readonly currencyOptions!: string[];

  private min = ''

  private max = ''

  private selectedCurrency = '';

  protected mounted(): void {
    this.handleMinInput(this.min);
    this.handleMaxInput(this.max);

    if (this.currencyOptions) {
      [this.selectedCurrency] = this.currencyOptions;
    }
    this.$emit('currencyInput', this.selectedCurrency);
  }

  protected handleMinInput(event: InputEvent | string): void {
    if (typeof event === 'string') {
      this.min = event;
    } else {
      this.min = (event.target as HTMLInputElement).value;
    }

    if (Number(this.min) > Number(this.max)) {
      this.max = this.min;
    }

    this.$emit('minInput', this.min);
  }

  protected handleMaxInput(event: InputEvent | string): void {
    if (typeof event === 'string') {
      this.max = event;
    } else {
      this.max = (event.target as HTMLInputElement).value;
    }

    if (Number(this.max) < Number(this.min)) {
      this.max = this.min;
    }

    this.$emit('maxInput', this.max);
  }
}
</script>

<style scoped>

</style>
