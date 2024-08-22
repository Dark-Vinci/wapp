package downstream

const packageName string = "account.downstream"

type Downstream struct{}

type DownStreamOperations interface{}

func New() DownStreamOperations {
	return DownStreamOperations(&Downstream{})
}
