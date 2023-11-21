package wfil

// https://github.com/glifio/wfil

type Deposit struct {
	From   string
	Amount string
}

type Withdrawal struct {
	To     string
	Amount string
}
