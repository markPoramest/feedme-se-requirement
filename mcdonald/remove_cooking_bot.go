package mcdonald

func (m *McDonald) RemoveCookingBot() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.Bots) == 0 {
		return
	}

	removedBot := m.Bots[len(m.Bots)-1]
	m.Bots = m.Bots[:len(m.Bots)-1]

	logger := GetLogger()
	if removedBot.IsProcessing {
		logger.LogBotRemoved(removedBot.ID, "PROCESSING")
		removedBot.StopChannel <- true
		return
	}

	logger.LogBotRemoved(removedBot.ID, "IDLE")
}
