<template>
  <div class="search">
    Brand:
    <select v-model="brand" @change="setBrand">
      <option value=""></option>
      <option v-for="(brand, brandId) in brands"
              :key="brandId"
              :value="brandId">{{ brand }}
      </option>
    </select>

    Model:
    <input v-model="model" @change="$emit('modelInput', $event.target.value)">
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class CarsBGBikeBrandModel extends Vue {
  private brand = '';

  private model = '';

  @Prop() private readonly type!: 'car' | 'bus' | 'bike';

  @Prop() private readonly brands!: { [id: number]: string };

  private mounted(): void {
    const brand = Object.keys(this.brands).shift();
    if (brand) {
      this.brand = brand;
    }
    this.$emit('brandInput', this.brand);
  }

  private async setBrand(event: InputEvent) {
    this.$emit('brandInput', (event.target as HTMLInputElement).value);
  }
}
</script>

<style scoped>

</style>
