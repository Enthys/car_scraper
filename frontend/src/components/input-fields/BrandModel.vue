<template>
  <div class="search">
    Brand:
    <select v-model="brand" @change="$emit('brandInput', $event.target.value)">
      <option value=""></option>
      <option v-for="(brandOption, index) in Object.keys(brandsAndModels)"
              :key="index"
              :value="brandOption">{{ brandOption }}</option>
    </select>

    Model:
    <select v-model="model" @change="$emit('modelInput', $event.target.value)">
      <option value=""></option>
      <option v-for="(modelOption, index) in brandModels"
              :key="index"
              :value="modelOption">{{ modelOption }}</option>
    </select>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class BrandModel extends Vue {
  private brand = '';

  private model = '';

  @Prop() private readonly brandsAndModels!: {[brandIndex: string]: string[]};

  private mounted(): void {
    const brand = Object.keys(this.brandsAndModels).shift();
    if (brand) {
      this.brand = brand;
    }
    this.$emit('brandInput', this.brand);
  }

  get brandModels(): string[] {
    if (this.brand) {
      return this.brandsAndModels[this.brand];
    }

    return [];
  }
}
</script>

<style scoped>

</style>
