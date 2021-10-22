package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var mutex sync.Mutex

	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				counter = counter + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("data counter", counter)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (bank *BankAccount) AddBalance(balance int) {
	bank.RWMutex.Lock()
	bank.Balance = bank.Balance + balance
	bank.RWMutex.Unlock()
}

func (bank *BankAccount) GetBalance() int {
	bank.RWMutex.RLock()
	balance := bank.Balance
	bank.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	bank := BankAccount{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				bank.AddBalance(1)
				fmt.Println(bank.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("data counter", bank.GetBalance())
}

type UserBalance struct {
	sync.RWMutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.RWMutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.RWMutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(+amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	fmt.Println("Unlock user 1", user1.Name)
	user2.Unlock()
	fmt.Println("Unlock user 2", user2.Name)
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "oman",
		Balance: 1000,
	}

	user2 := UserBalance{
		Name:    "pradipta",
		Balance: 1000,
	}

	go Transfer(&user1, &user2, 100)
	go Transfer(&user2, &user1, 200)

	time.Sleep(10 * time.Second)

	fmt.Println("user", user1.Name, "balance", user1.Balance)
	fmt.Println("user", user2.Name, "balance", user2.Balance)
}
