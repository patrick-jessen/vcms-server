export class TreeTableData {
  constructor(dataFn, childrenFn, renderFn, identifierFn, root) {
    this._dataFn = dataFn
    this._childrenFn = childrenFn
    this._renderFn = renderFn
    this._identifierFn = identifierFn
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
        this._identifierFn,
        false)
    })
  }

  get identifier() {
    return this._identifierFn(this.data)
  }

  render(index) {
    return this._renderFn(this.data, index)
  }
}