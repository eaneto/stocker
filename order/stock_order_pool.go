package order

type Task interface {
	Run() error
}

// StockOrderError Structure to save the order request and the error
// that happend during processing.
type StockOrderError struct {
	Error   error
	Request StockOrderRequest
}

type StockOrderPool struct {
	Errors               chan StockOrderError
	OrderRequestsChannel chan StockOrderRequest
}

func NewStockOrder() StockOrderPool {
	return StockOrderPool{
		OrderRequestsChannel: make(chan StockOrderRequest),
		Errors:               make(chan StockOrderError),
	}
}

func (pool *StockOrderPool) Publish(order StockOrderRequest) {
	pool.OrderRequestsChannel <- order
}

func (pool *StockOrderPool) Poll() *StockOrderRequest {
	select {
	case request := <-pool.OrderRequestsChannel:
		return &request
	}
}
