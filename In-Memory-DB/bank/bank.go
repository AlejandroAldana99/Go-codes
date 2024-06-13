package bankingsystem

import (
	"fmt"
	"sort"
	"sync"
)

type Account struct {
	AccountID  string
	Balance    int
	Outgoing   int
	CreatedAt  int
	BalanceLog map[int]int
}

type AccountSummary struct {
	AccountID string
	Outgoing  int
}

type ByOutgoing []AccountSummary

func (o ByOutgoing) Len() int      { return len(o) }
func (o ByOutgoing) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
func (o ByOutgoing) Less(i, j int) bool {
	if o[i].Outgoing == o[j].Outgoing {
		return o[i].AccountID < o[j].AccountID
	}
	return o[i].Outgoing > o[j].Outgoing
}

type ScheduledPayment struct {
	ID        string
	AccountID string
	Amount    int
	ExecuteAt int
}

type BankingSystemImpl struct {
	accounts          map[string]*Account
	scheduledPayments []ScheduledPayment
	mu                sync.Mutex
	paymentCounter    int
}

func NewBankingSystemImpl() *BankingSystemImpl {
	return &BankingSystemImpl{
		accounts:          make(map[string]*Account),
		scheduledPayments: []ScheduledPayment{},
	}
}

func (b *BankingSystemImpl) CreateAccount(timestamp int, accountId string) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Process by timestams
	b.processPayments(timestamp)

	if _, exists := b.accounts[accountId]; exists {
		return false
	}

	b.accounts[accountId] = &Account{
		AccountID:  accountId,
		Balance:    0,
		CreatedAt:  timestamp,
		Outgoing:   0,
		BalanceLog: make(map[int]int),
	}

	return true
}

func (b *BankingSystemImpl) Deposit(timestamp int, accountId string, amount int) *int {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Process by timestams
	b.processPayments(timestamp)

	if account, exists := b.accounts[accountId]; exists {
		account.Balance += amount
		account.BalanceLog[timestamp] = account.Balance

		return &account.Balance
	}

	return nil
}

func (b *BankingSystemImpl) Transfer(timestamp int, sourceAccountId string, targetAccountId string, amount int) *int {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Process by timestams
	b.processPayments(timestamp)

	sourceAccount, sourceExists := b.accounts[sourceAccountId]
	targetAccount, targetExists := b.accounts[targetAccountId]

	if !sourceExists || !targetExists {
		return nil
	}

	if sourceAccountId == targetAccountId {
		return nil
	}

	if sourceAccount.Balance < amount {
		return nil
	}

	sourceAccount.Balance -= amount
	sourceAccount.Outgoing += amount
	sourceAccount.BalanceLog[timestamp] = sourceAccount.Balance
	targetAccount.Balance += amount
	targetAccount.BalanceLog[timestamp] = targetAccount.Balance

	return &sourceAccount.Balance
}

func (b *BankingSystemImpl) TopSpenders(timestamp int, n int) []string {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Process by timestams
	b.processPayments(timestamp)

	var result []string
	var summaries []AccountSummary
	count := 0

	for _, account := range b.accounts {
		summaries = append(summaries, AccountSummary{AccountID: account.AccountID, Outgoing: account.Outgoing})
	}

	sort.Sort(ByOutgoing(summaries))

	for _, summary := range summaries {
		if count >= n {
			break
		}
		result = append(result, fmt.Sprintf("%s(%d)", summary.AccountID, summary.Outgoing))
		count++
	}

	return result
}

func (b *BankingSystemImpl) SchedulePayment(timestamp int, accountId string, amount int, delay int) *string {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Process by timestams
	b.processPayments(timestamp)

	if _, exists := b.accounts[accountId]; exists {
		b.paymentCounter++
		paymentID := fmt.Sprintf("payment%d", b.paymentCounter)
		scheduledTime := timestamp + delay

		b.scheduledPayments = append(b.scheduledPayments, ScheduledPayment{
			ID:        paymentID,
			AccountID: accountId,
			Amount:    amount,
			ExecuteAt: scheduledTime,
		})

		return &paymentID
	}
	return nil
}

func (b *BankingSystemImpl) CancelPayment(timestamp int, accountId string, paymentId string) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Process by timestams
	b.processPayments(timestamp)

	for i, payment := range b.scheduledPayments {
		if payment.ID == paymentId {
			if payment.AccountID != accountId {
				return false
			}
			b.scheduledPayments = append(b.scheduledPayments[:i], b.scheduledPayments[i+1:]...)
			return true
		}
	}

	return false
}

func (b *BankingSystemImpl) MergeAccounts(timestamp int, accountId1 string, accountId2 string) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Process by timestams
	b.processPayments(timestamp)

	if accountId1 == accountId2 {
		return false
	}

	account1, exists1 := b.accounts[accountId1]
	account2, exists2 := b.accounts[accountId2]

	if !exists1 || !exists2 {
		return false
	}

	account1.Balance += account2.Balance
	account1.Outgoing += account2.Outgoing
	account1.BalanceLog[timestamp] = account1.Balance

	for i, payment := range b.scheduledPayments {
		if payment.AccountID == accountId2 {
			b.scheduledPayments[i].AccountID = accountId1
		}
	}

	delete(b.accounts, accountId2)

	return true
}

func (b *BankingSystemImpl) GetBalance(timestamp int, accountId string, timeAt int) *int {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Process by timestams
	b.processPayments(timestamp)

	account, exists := b.accounts[accountId]
	if !exists || account.CreatedAt > timeAt {
		return nil
	}

	closestTime := -1
	for t := range account.BalanceLog {
		if t <= timeAt && (closestTime == -1 || t > closestTime) {
			closestTime = t
		}
	}

	if closestTime != -1 {
		balance := account.BalanceLog[closestTime]
		return &balance
	}

	return nil
}

// Helpers
func (b *BankingSystemImpl) processPayments(currentTimestamp int) {
	if len(b.scheduledPayments) < 1 {
		return
	}

	newQueue := []ScheduledPayment{}

	for _, payment := range b.scheduledPayments {
		if payment.ExecuteAt <= currentTimestamp {
			account, exists := b.accounts[payment.AccountID]
			if exists && account.Balance >= payment.Amount {
				account.Balance -= payment.Amount
				account.Outgoing += payment.Amount
				account.BalanceLog[currentTimestamp] = account.Balance
			}
		} else {
			newQueue = append(newQueue, payment)
		}
	}

	b.scheduledPayments = newQueue
}
