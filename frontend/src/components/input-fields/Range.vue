<template>
  <div class="input">
    <div>
      {{ label }}:&nbsp;
      Min
      <input
      type="number"
      min="0"
      placeholder="0"
      v-model="min"
      @input="handleMinInput"
    >
      -
      Max
      <input
      type="number"
      min="0"
      placeholder="0"
      v-model="max"
      @blur="handleMaxInput"
    >

    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class Range extends Vue {
  @Prop({ type: String }) protected readonly label!: string;

  @Prop({ type: Number }) protected readonly minValue!: number;

  protected min = 0;

  @Prop({ type: Number }) protected readonly maxValue?: number;

  protected max = 0;

  protected mounted(): void {
    this.min = this.minValue;
    if (this.maxValue !== undefined) {
      this.max = this.maxValue;
    }
    this.handleMinInput({ target: { value: this.min } });
    this.handleMaxInput({ target: { value: this.max } });
  }

  protected handleMinInput(event: InputEvent): void {
    this.min = Number((event.target as HTMLInputElement).value);
    if (this.maxValue !== undefined && this.min > this.maxValue) {
      this.min = this.maxValue;
    }

    if (this.min > this.max) {
      this.max = this.min;
    }

    if (event.target) {
      this.$emit('minInput', this.min);
    }
  }

  protected handleMaxInput(event: InputEvent): void {
    this.max = Number((event.target as HTMLInputElement).value);

    if (this.max < this.minValue) {
      this.max = this.minValue;
    }

    if (this.max < this.min) {
      this.max = this.min;
    }

    if (event.target) {
      this.$emit('maxInput', this.max);
    }
  }
}
</script>

<style scoped>

</style>
