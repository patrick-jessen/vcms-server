export class TreeTableData {
  constructor(dataFn, childrenFn, renderFn, root) {
    this._dataFn = dataFn
    this._childrenFn = childrenFn
    this._renderFn = renderFn
    if(typeof root === 'undefined')
      root = true
    this.root = root
  }

  get data() {
    return this._dataFn()
  }

  get children() {
    return this._childrenFn(this.data, this.root).map((c) => {
      return new TreeTableData(
        ()=>{return c},
        this._childrenFn,
        this._renderFn,
        false)
    })
  }

  render(index) {
    return this._renderFn(this.data, index)
  }
}