import { Prop, Vue } from 'vue-property-decorator';

export default abstract class FilterBase extends Vue {
  @Prop() private readonly storage: unknown;

  protected abstract type: string;

  protected createFilter(): void {
    this.$emit('createFilter', {
      type: this.type,
      data: this.storage,
    });
  }

  public static getTitle(): string {
    return 'Undefined';
  }
}
