<template>
  <div class="search">
    <brand-model
      :brands-and-models="brandModels"
      v-on:brandInput="brand = arguments[0]"
      v-on:modelInput="model = arguments[0]"
    />
    <br>
    <currency-range
      :min-value="0"
      :max-value="1000"
      :currency-options="currencyOptions"
      v-on:currencyInput="currency = arguments[0]"
      v-on:minInput="priceStart = arguments[0]"
      v-on:maxInput="priceEnd = arguments[0]"
    />

    <br>

    <range
      label="Year"
      :min-value="1980"
      :max-value="new Date().getFullYear()"
      v-on:minInput="yearStart = arguments[0]"
      v-on:maxInput="yearEnd = arguments[0]"
    />

    <br>
    <button @click="createFilter">Create</button>
  </div>
</template>

<script lang="ts">
import BrandModel from '@/components/input-fields/BrandModel.vue';
import CurrencyRange from '@/components/input-fields/CurrencyRange.vue';
import Range from '@/components/input-fields/Range.vue';
import FilterBase from '@/components/filters/FilterBase';
import { Component } from 'vue-property-decorator';
import brandModels from './MobileBGBikes.brands';
import currencyOptions from './currencyOptions';

@Component({
  components: { CurrencyRange, Range, BrandModel },
})
export default class MobileBGBikes extends FilterBase {
  protected type = 'MobileBGBike';

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
    return 'MobileBG Bikes';
  }

  protected createFilter(): void {
    this.$emit('createFilter', {
      type: this.type,
      data: {
        topmenu: '1',
        act: '3',
        f5: this.brand,
        f6: this.model,
        f10: String(this.yearStart),
        f11: String(this.yearEnd),
        f7: String(this.priceStart),
        f8: String(this.priceEnd),
        f9: this.currency,
      },
    });
  }
}
</script>

<style scoped>

</style>
