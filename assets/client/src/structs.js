export class TreeTableData {
  constructor(dataFn, childrenFn, renderFn) {
    this._dataFn = dataFn
    this._childrenFn = childrenFn
    this._renderFn = renderFn
  }

  get data() {
    return this._dataFn()
  }

  get children() {
    return this._childrenFn(this.data).map((c) => {
      return new TreeTableData(
        ()=>{return c},
        this._childrenFn,
        this._renderFn)
    })
  }

  render(index) {
    return this._renderFn(this.data, index)
  }
}