// Code generated by github.com/euank/ifacepropagate

package case02

func (p *partialOverride) propogateInterfaces() If1 {
	_, ok0 := p.If1.(If2)
	switch {
	case ok0:
		return struct {
			If1
			If2
		}{p, p}
	case !ok0:
		return struct {
			If1
		}{p}
	default:
		panic("unreachable")
	}
}
func (p *partialOverride) Method4() {
	p.If1.(If2).Method4()
}
