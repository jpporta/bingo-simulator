package main

// Own -1 vertical win, 0 no win, 1 horizontal win
type Player struct {
	Card BingoCard
	Own  int
}

func CreatePlayer() Player {
	return Player{
		Card: CreateNewCard(),
		Own:  0,
	}
}

func (p *Player) CheckNewBall(ball int) bool {
	p.Card.DrawBall(ball)
	own := p.Card.HasWonHorizontal()
	if own {
		p.Own = 1
		return true
	}
	own = p.Card.HasWonVertical()
	if own {
		p.Own = -1
		return true
	}
	return false
}
