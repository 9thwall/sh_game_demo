package cards

import (
	"math/rand"
	"reflect"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

/*
TABLE CONSISTS OF MIN 2 -> MAX 10 SEATS
DEALER BUTTON SHOULD BE FIRST AUTOMATICALLY RANDOM SET AT BEGINNING
THERE WILL BE NO MORE NEED TO SET THE DEALER BUTTON AFTER FIRST INITIALIZATION
BUTTON ROTATES RIGHT
FIRST CARD TO BE DELT ALWAYS STARTS ON SMALL BLIND
*/


var Deck = []string{
"ah", "ad", "ac", "as",
"2h", "2d","2c","2s",
"3h","3d","3c","3s",
"4h","4d","4c","4s",
"5h","5d","5c","5s",
"6h","6d","6c","6s",
"7h","7d","7c","7s",
"8h","8d","8c","8s",
"9h","9d","9c","9s",
"10h","10d","10c","10s",
"jh","jd","jc","js",
"qh","qd","qc","qs",
"kh","kd","kc","ks",
}

func Shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}

func GetRandomDeck() []string {

	Deck := []string{
		"ah", "ad", "ac", "as",
		"2h", "2d","2c","2s",
		"3h","3d","3c","3s",
		"4h","4d","4c","4s",
		"5h","5d","5c","5s",
		"6h","6d","6c","6s",
		"7h","7d","7c","7s",
		"8h","8d","8c","8s",
		"9h","9d","9c","9s",
		"10h","10d","10c","10s",
		"jh","jd","jc","js",
		"qh","qd","qc","qs",
		"kh","kd","kc","ks",
	}
	Shuffle(Deck)
	return Deck
}

//DealPreFlop ONLY HAPPENS AT THE BEGINNING OF EACH HAND
func DealPreFlop(deck []string, allPlayerSeats []int, playerCount int, dealerSeat int){

	seatNumers := make([]int,playerCount)

	switch playerCount{
	case 9:
		//Sets up order slice depending on how many players there are because full table.
		for i := 0; i < playerCount; i++ {
			seatNumers[i] = i + 1
		}
	default:
		//Set up order slice depending on allPlayerSeats because table is not full and we need which seats are taken

	}

	//Depending on how many players and dealer button we can determine the start position of deal ONLY for DealPreFlop

	//fmt.Println("seatNumbers")
	//fmt.Println(seatNumers)

	//SET UP THE ARRAY TO CHOOSE WHERE THE CARDS GO TO EACH PLAYER. [6,7,8,9(dealer),1,2,3,4,5]

}




