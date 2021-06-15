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
    <select v-model="model" @change="$emit('modelInput', $event.target.value)">
      <option value=""></option>
      <option v-for="(model, modelId) in models"
              :key="modelId"
              :value="modelId">{{ model }}
      </option>
    </select>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class CarsBGBrandModel extends Vue {
  private brand = '';

  private model = '';

  @Prop() private readonly type!: 'car' | 'bus' | 'bike';

  @Prop() private readonly brands!: { [id: number]: string };

  private models: { [id: number]: string } = {};

  private mounted(): void {
    const brand = Object.keys(this.brands).shift();
    if (brand) {
      this.brand = brand;
    }
    this.$emit('brandInput', this.brand);
  }

  private async setBrand(event: InputEvent) {
    const { value } = event.target as HTMLInputElement;
    this.$emit('brandInput', (event.target as HTMLInputElement).value);
    await this.setBrandModels();
    console.log(JSON.stringify(this.models));
  }

  private async setBrandModels(): Promise<void> {
    if (!this.brand) {
      this.models = {};
    }

    const resp = await this.$http.get(`/carsbg/${this.type}/brands/${this.brand}/models`, {
      responseType: 'text',
    });
    const modelsHtml = document.createElement('div');
    modelsHtml.innerHTML = resp.data;
    console.log(modelsHtml.querySelectorAll('label[for*="modelId_"]'));
    const reduceResult: { [id: string]: string } = {};

    this.models = Array.from(modelsHtml.querySelectorAll('label[for*="modelId_"]'))
      .reduce((result: { [id: string]: string }, label: Element) => {
        const resultHolder = result;
        const modelId = label.getAttribute('for')?.split('_').pop();

        if (modelId === undefined) {
          return resultHolder;
        }

        resultHolder[modelId] = (label as HTMLElement).innerText;

        return resultHolder;
      }, reduceResult);
  }

  private getModelLink(brandId: string): string {
    switch (this.type) {
      case 'car':
        return `https://www.cars.bg/carmodel.php?brandId=${brandId}`;
      case 'bus':
        return `https://www.cars.bg/carmodel.php?brandId=${brandId}`;
      case 'bike':
        return `https://www.cars.bg/carmodel.php?brandId=${brandId}`;
      default:
        throw new Error('Invalid Type!');
    }
  }
}
</script>

<style scoped>

</style>
