package downstream

const packageName string = "account.downstream"

type Downstream struct{}

func New() *Downstream {
	return &Downstream{}
}
