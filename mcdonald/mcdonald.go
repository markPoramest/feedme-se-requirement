package mcdonald

import (
	"sync"
)

type McDonald struct {
	PendingOrder  []*Order
	CompleteOrder []*Order
	Bots          []*Bot

	nextOrderID int
	nextBotID   int

	mu             sync.Mutex
	orderProcessed chan *Order
}

func NewMcDonald() *McDonald {
	return &McDonald{
		PendingOrder:   make([]*Order, 0),
		CompleteOrder:  make([]*Order, 0),
		Bots:           make([]*Bot, 0),
		nextOrderID:    1,
		nextBotID:      1,
		orderProcessed: make(chan *Order, 100),
	}
}

func (m *McDonald) GetPendingOrders() []*Order {
	m.mu.Lock()
	defer m.mu.Unlock()

	result := make([]*Order, len(m.PendingOrder))
	copy(result, m.PendingOrder)
	return result
}

func (m *McDonald) GetCompleteOrders() []*Order {
	m.mu.Lock()
	defer m.mu.Unlock()

	result := make([]*Order, len(m.CompleteOrder))
	copy(result, m.CompleteOrder)
	return result
}

func (m *McDonald) GetBots() []*Bot {
	m.mu.Lock()
	defer m.mu.Unlock()

	result := make([]*Bot, len(m.Bots))
	copy(result, m.Bots)
	return result
}

func (m *McDonald) GetFinalStatus() {
	m.mu.Lock()
	defer m.mu.Unlock()

	completed := len(m.CompleteOrder)
	pending := len(m.PendingOrder)
	activeBots := len(m.Bots)
	vipCount := 0
	normalCount := 0

	for _, order := range m.CompleteOrder {
		if order.Customer == VIPCustomer {
			vipCount++
		} else {
			normalCount++
		}
	}

	for _, order := range m.PendingOrder {
		if order.Customer == VIPCustomer {
			vipCount++
		} else {
			normalCount++
		}
	}

	logger := GetLogger()
	logger.LogFinalStatus(completed, vipCount, normalCount, activeBots, pending)
}
