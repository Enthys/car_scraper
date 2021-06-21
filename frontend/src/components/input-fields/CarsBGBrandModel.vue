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
    <div class="five-on-row">
      <div class="row-child" v-for="(model, modelId) in models" :key="modelId">
        <input type="checkbox"
               name="models[]"
               v-model="selectedModels"
               @change="selectModel"
               :value=modelId /> {{ model }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class CarsBGBrandModel extends Vue {
  private brand = '';

  private selectedModels: string[] = [];

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
    this.$emit('brandInput', (event.target as HTMLInputElement).value);
    await this.setBrandModels();
  }

  private selectModel(event: InputEvent): void {
    if (!event) {
      return;
    }

    const { checked, value } = event.target as EventTarget & { checked: boolean, value: string};
    if (checked) {
      this.selectedModels.push(value);
    } else {
      this.selectedModels.splice(this.selectedModels.indexOf(value), 1);
    }
    this.$emit('modelInput', [this.selectedModels]);
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

<style lang="scss" scoped>
.five-on-row {
  width: 100%;
  display: flex;
  justify-content: space-around;
  flex-direction: row;
  flex-wrap: wrap;

  .row-child {
    width: 19.4%;
    padding: 4px 0;
    margin: 3px;
    border: 1px solid #EEE;
    border-radius: 4px;
    transition: border 0.5s;

    &:hover {
      border: 1px solid #000;
    }
  }
}
</style>
