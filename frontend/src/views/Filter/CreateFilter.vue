<template>
  <div>
    <template>
      <div>
        <div class="searchNavigation">
          <button v-on:click="index != 0 ? index-- : index = searches.length-1">◄</button>
          <h1>{{ currentSearch.getTitle() }}</h1>
          <button v-on:click="index > searches.length ? index = 0 : index++">►</button>
        </div>
        <keep-alive>
          <component :is="currentSearch" v-on:createFilter="createFilter"></component>
        </keep-alive>
        <div class="message">
          <div v-if="message !== ''"
               class="success"
               @click="clearMessage"
          >
            {{ message }}
          </div>
          <div v-if="error !== ''"
               class="error"
               @click="clearError"
          >
            {{ error }}
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script lang="ts">
import CarsBGBikes from '@/components/filters/CarsBG/CarsBGBikes.vue';
import CarsBGBuses from '@/components/filters/CarsBG/CarsBGBuses.vue';
import CarsBGCars from '@/components/filters/CarsBG/CarsBGCars.vue';
import MobileBGBikes from '@/components/filters/MobileBG/MobileBGBikes.vue';
import MobileBGCars from '@/components/filters/MobileBG/MobileBGCars.vue';
import MobileBGBuses from '@/components/filters/MobileBG/MobileBGBuses.vue';
import FilterService from '@/services/FilterService/FilterService';
import { VueConstructor } from 'vue';
import { Component, Vue } from 'vue-property-decorator';

@Component({
  inject: ['filterService'],
})
export default class CreateFilter extends Vue {
  private readonly filterService!: FilterService;

  private index = 0;

  private searches: VueConstructor[] = [
    MobileBGCars, MobileBGBuses, MobileBGBikes, CarsBGCars, CarsBGBuses, CarsBGBikes,
  ];

  private message = '';

  private error = '';

  get currentSearch(): VueConstructor {
    return this.searches[this.index % this.searches.length];
  }

  private async createFilter(filterData: { type: string, data: unknown }): Promise<void> {
    try {
      await this.filterService.createFilter(filterData.type, filterData.data);
      this.error = '';
      this.message = 'Filter Created!';
    } catch (err) {
      this.error = `Failed to create filter! Error: ${err.message}`;
      this.message = '';
    }
  }

  private clearMessage(): void {
    this.message = '';
  }

  private clearError(): void {
    this.error = '';
  }
}
</script>

<style lang="scss" scoped>
#wrapper {
  display: flex;
  flex-direction: column;
  width: 70vw;
  min-width: 600px;
  margin: 0 auto;
}

.searchNavigation {
  display: flex;
  justify-content: space-between;
  flex-direction: row;

  button {
    border: none;
    background: rgb(38, 176, 183);
    color: white;
    font-family: 'Courier New', Courier, monospace;
    font-weight: bold;
    font-size: 30px;
  }

}

.message {
  text-align: center;
  text-transform: uppercase;

  .success {
    display: block;
    margin: 10px auto;
    width: 400px;
    border: 1px solid rgb(150, 255, 150);
    border-radius: 10px;
    background-color: rgb(100, 255, 100);
    color: white;
  }

  .error {
    display: block;
    margin: 10px auto;
    width: 400px;
    border: 1px solid rgb(255, 55, 55);
    border-radius: 10px;
    background-color: rgb(255, 70, 70);
    color: black;
  }

}
</style>
