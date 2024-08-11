package downstream

const packageName = "gateway.downstream"

type Downstream struct{}

func New() *Downstream {
	return &Downstream{}
}
