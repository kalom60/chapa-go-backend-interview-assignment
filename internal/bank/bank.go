package bank

func New(bankRepo BankRepo) Handler {
	svc := NewService(bankRepo)
	return NewHandler(svc)
}
