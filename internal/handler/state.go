package handler

type State int

const (
	InitState State = iota + 1

	WaitDepositName
	WaitDepositRate
	WaitDepositPeriod
	WaitDepositBalance
)
