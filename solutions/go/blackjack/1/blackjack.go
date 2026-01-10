package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var c int
    
    switch card {
        case "ace":
        	c = 11
        case "two":
        	c = 2
        case "three":
        	c = 3
        case "four":
        	c = 4
        case "five":
        	c = 5
        case "six":
        	c = 6
        case "seven":
        	c = 7
        case "eight":
        	c = 8
        case "nine":
        	c = 9
        case "ten":
        	c = 10
        case "jack":
        	c = 10
        case "queen":
        	c = 10
        case "king":
        	c = 10
        case "other":
        	c = 0
    }
    
    return c
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    
	var action string

    c1 := ParseCard(card1)
    c2 := ParseCard(card2)
    d1 := ParseCard(dealerCard)
    
	switch {
        case (c1 + c2) == 21 && d1 >= 10:
        	action = "S" // Stand
        case (c1 + c2) == 21 && d1 < 10:
        	action = "W" // Win        
        case c1 == c2 && c1 == 11:
        	action = "P" // Split
        case (c1 + c2) >= 17 && (c1 + c2) <= 20:
        	action = "S" // Stand
        case (c1 + c2) >= 12 && (c1 + c2) <= 16 && d1 < 7:
        	action = "S" // Stand    
        case (c1 + c2) >= 12 && (c1 + c2) <= 16 && d1 >= 7:
        	action = "H" // Hit          
        case (c1 + c2) <= 11:
        	action = "H" // Hit           
    }
    return action
}
