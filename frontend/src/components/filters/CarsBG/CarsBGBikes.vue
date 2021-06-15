<template>
  <div class="search">
    <carsbg-bike-brand-model
      :brands="brands"
      type="bike"
      v-on:brandInput="brand = arguments[0]"
      v-on:modelInput="model = arguments[0]"
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
import CarsBGBikeBrandModel from '@/components/input-fields/CarsBGBikeBrandModel.vue';
import CurrencyRange from '@/components/input-fields/CurrencyRange.vue';
import Range from '@/components/input-fields/Range.vue';
import { Component } from 'vue-property-decorator';
import FilterBase from '../FilterBase';
import brands from './CarsBGBikes.brands';

@Component({
  components: {
    'carsbg-bike-brand-model': CarsBGBikeBrandModel,
    CurrencyRange,
    Range,
    BrandModel,
  },
})
export default class CarsBGBikes extends FilterBase {
  protected type = 'CarsBGBike';

  private readonly brands = brands;

  private brand = '';

  private model = '';

  private yearStart = 1980;

  private yearEnd = 2015;

  private priceStart = 0;

  private priceEnd = 0;

  public static getTitle(): string {
    return 'CarsBG Bikes';
  }

  protected createFilter(): void {
    this.$emit('createFilter', {
      type: this.type,
      data: {
        type: 'bike',
        brandId: this.brand,
        model: this.model,
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
