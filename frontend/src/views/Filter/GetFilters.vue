<template>
  <h1>Get filters</h1>
</template>

<script lang="ts">
import FilterService, { FilterInterface } from '@/services/FilterService/FilterService';
import { Component, Vue } from 'vue-property-decorator';

@Component({
  inject: ['filterService'],
})
export default class GetFilters extends Vue {
  private readonly filterService!: FilterService

  private filters: FilterInterface[] = [];

  private async mounted(): Promise<void> {
    this.filters = await this.filterService.getFilters();
  }
}
</script>

<style scoped>

</style>
