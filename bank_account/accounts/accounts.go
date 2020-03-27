package accounts

import "errors"

// Account struct
type account struct { //얘는 퍼블릿이냐..
	owner   string
	balance int
}

var errNomoney = errors.New("Can't withdraw. you are poor")

// NewAccount creates account
func NewAccount(owner string) *account { //주의
	account := account{owner: owner, balance: 0}
	return &account
	//address를 리턴한다.
	//return new Object(); 를 하는건디 이건 자바에서 생성자가 아니면 안되기 때문에..
	//아님 Builder 객체에서 넘겨줘야 하기때문에.. 어쨋건 복사가 한번 일어나는디 얜 생성한 메모리주소
	//자체를 리턴해버림
}

//이건 method 다 메소드는 (a account) 같은 reciver가 있다.
func (a *account) Deposit(amount int) {
	a.balance += amount
}

func (a account) Balance() int {
	return a.balance
}

func (a *account) WithDraw(amount int) error {

	if a.balance < amount {
		return errNomoney
	} else {
		a.balance -= amount
		return nil
	}
}
