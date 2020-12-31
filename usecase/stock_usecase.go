package usecase

type StockUseCase interface {
	RegisterStock(ticker string, price uint) error
	SearchByTicker(ticker string) error
}
