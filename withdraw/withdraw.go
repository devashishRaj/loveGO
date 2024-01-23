package withdraw

import "errors"

func Withdraw(balance int, amount int) (int, error) {
	if amount > balance {
		return 0, errors.New("Not enough balance , deposity required")
	}
	return balance - amount, nil
}
