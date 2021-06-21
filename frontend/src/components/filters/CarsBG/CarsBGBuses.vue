<template>
  <div class="search">
    <carsbg-brand-model
      :brands="brands"
      type="bus"
      v-on:brandInput="brand = arguments[0]"
      v-on:modelInput="models = arguments[0]"
    />
    <br>
    <currency-range
      :min-value="0"
      :max-value="1000"
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
import CarsBGBrandModel from '@/components/input-fields/CarsBGBrandModel.vue';
import CurrencyRange from '@/components/input-fields/CurrencyRange.vue';
import Range from '@/components/input-fields/Range.vue';
import { Component } from 'vue-property-decorator';
import FilterBase from '../FilterBase';
import brands from './CarsBGBuses.brands';

@Component({
  components: {
    'carsbg-brand-model': CarsBGBrandModel,
    CurrencyRange,
    Range,
    BrandModel,
  },
})
export default class CarsBGBuses extends FilterBase {
  protected type = 'CarsBGBus';

  private readonly brands = brands;

  private brand = '';

  private models: string[] = [];

  private yearStart = 1980;

  private yearEnd = 2015;

  private priceStart = 0;

  private priceEnd = 0;

  public static getTitle(): string {
    return 'CarsBG Buses';
  }

  protected createFilter(): void {
    this.$emit('createFilter', {
      type: this.type,
      data: {
        type: 'bus',
        brandId: this.brand,
        models: this.models,
        yearStart: String(this.yearStart),
        yearEnd: String(this.yearEnd),
        priceStart: String(this.priceStart),
        priceEnd: String(this.priceEnd),
      },
    });
  }
}
</script>

<style scoped>

</style>
