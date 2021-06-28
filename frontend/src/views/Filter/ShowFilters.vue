<template>
  <div class="filter-container">
    <div class="filter" v-for="(filter, index) in filters" :key="index">
      {{ filter.type }}<br>
      Brand: {{ filter.brand }} |
      <span v-if="filter.model">Model: {{
          Array.isArray(filter.model) ? filter.model.join(", ") : filter.model
        }} <br></span>
      <span v-if="filter.yearFrom !== undefined || filter.yearTo !== undefined">
        Year From {{ filter.yearFrom }} | Year To {{ filter.yearTo }} <br>
      </span>
      <span v-if="filter.priceFrom !== undefined || filter.priceTo !== undefined">
        Price From {{ filter.priceFrom }} | Price To {{ filter.priceTo }} <br>
      </span>

      <button class="btn btn-danger" @click="deleteFilter(filter.id)">Delete</button>
    </div>
  </div>
</template>

<script lang="ts">
import FilterService, { FilterSearchInterface, FilterType } from '@/services/FilterService/FilterService';
import { Component, Vue } from 'vue-property-decorator';

@Component({
  inject: ['filterService'],
})
export default class ShowFilters extends Vue {
  private readonly filterService!: FilterService;

  private mobileBGFilterTypes = [
    FilterType.MOBILEBG_CARS,
    FilterType.MOBILEBG_BUSES,
    FilterType.MOBILEBG_BIKES,
  ];

  private carsBGFilterTypes = [
    FilterType.CARSBG_CARS,
    FilterType.CARSBG_BUSES,
    FilterType.CARSBG_BIKES,
  ];

  private filters: FilterSearchInterface[] = [];

  private async mounted(): Promise<void> {
    await this.setupFilters();
  }

  private async setupFilters() {
    this.filters = (await this.filterService.getFilters())
      .reduce((result: FilterSearchInterface[], filter) => {
        if (this.mobileBGFilterTypes.indexOf(filter.type) !== -1) {
          result.push(this.decodeMobileBGSearch(filter.id, filter.type, filter.search));
        }
        if (this.carsBGFilterTypes.indexOf(filter.type) !== -1) {
          result.push(this.decodeCarsBGSearch(filter.id, filter.type, filter.search));
        }
        return result;
      }, []);
  }

  private async deleteFilter(filterId: number): Promise<void> {
    await this.filterService.deleteFilter(filterId);
    await this.setupFilters();
  }

  private decodeMobileBGSearch(
    id: number,
    type: FilterType,
    search: string,
  ): FilterSearchInterface {
    const searchData = JSON.parse(search);
    return {
      id,
      type,
      brand: searchData.f5,
      model: searchData.f6,
      priceTo: searchData.f7 ? `${searchData.f7} ${searchData.f9}` : undefined,
      priceFrom: searchData.f8 ? `${searchData.f8} ${searchData.f9}` : undefined,
      yearTo: searchData.f10 ? `${searchData.f10}` : undefined,
      yearFrom: searchData.f11 ? `${searchData.f11}` : undefined,
    };
  }

  private decodeCarsBGSearch(id: number, type: FilterType, search: string): FilterSearchInterface {
    const searchData = JSON.parse(search);
    return {
      id,
      type,
      brand: searchData.brandId,
      model: searchData.models ?? searchData.model_moto,
      priceTo: searchData.priceTo ? `${searchData.priceTo} .лв` : undefined,
      priceFrom: searchData.priceFrom ? `${searchData.priceFrom} .лв` : undefined,
      yearTo: searchData.yearTo ? `${searchData.yearTo}` : undefined,
      yearFrom: searchData.yearFrom ? `${searchData.yearFrom}` : undefined,
    };
  }
}
</script>

<style scoped lang="scss">
.filter-container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-between;

  .filter {
    width: 350px;
    margin: 10px;
    border: 1px solid black;
    border-radius: 15px;
  }
}
</style>
