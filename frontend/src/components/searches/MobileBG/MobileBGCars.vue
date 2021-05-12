<template>
  <div class="search">
    <brand-model
      :brands-and-models="brandModels"
      v-on:brandInput="brand = arguments[0]"
      v-on:modelInput="model = arguments[0]"
    />

    <currency-range
      :min-value="0"
      :currency-options="currencyOptions"
      v-on:currencyInput="currency = arguments[0]"
      v-on:minInput="priceStart = arguments[0]"
      v-on:maxInput="priceEnd = arguments[0]"
    />

    <range
      label="Year"
      :min-value="1980"
      :max-value="new Date().getFullYear()"
      v-on:minInput="yearStart = arguments[0]"
      v-on:maxInput="yearEnd = arguments[0]"
    />

    <button @click="createFilter">Create</button>
  </div>
</template>

<script lang="ts">
import BrandModel from '@/components/input-fields/BrandModel.vue';
import CurrencyRange from '@/components/input-fields/CurrencyRange.vue';
import Range from '@/components/input-fields/Range.vue';
import FilterBase from '@/views/Filter/FilterBase';
import { Component } from 'vue-property-decorator';
import { brandModels, currencyOptions } from './MobileBGCars.brands';

@Component({
  components: { CurrencyRange, Range, BrandModel },
})
export default class MobileBGCars extends FilterBase {
  protected type = 'MobileBGCars';

  private readonly brandModels = brandModels;

  private readonly currencyOptions = currencyOptions;

  private brand = '';

  private model = '';

  private yearStart = 1980;

  private yearEnd = 2015;

  private currency = 'USD';

  private priceStart = 0;

  private priceEnd = 0;

  public static getTitle(): string {
    return 'MobileBG Cars';
  }

  private protected createFilter(): void {
    this.$emit('createFilter', {
      type: this.type,
      data: {
        topmenu: '1',
        act: '3',
        f5: this.brand,
        f6: this.model,
        f10: this.yearStart,
        f11: this.yearEnd,
        f7: this.priceStart,
        f8: this.priceEnd,
        f9: this.currency,
      },
    });
  }
}
</script>

<style scoped>

</style>
