//The good 'ol farmer
type Farmer struct {
}
//The gopher (goose) that the farmer purchased at the market
type Gopher struct {
}
func (g *Gopher) Eat(b *BagOfBeans, f *Farmer) {

    if b != nil && f == nil {
        log.Fatal("OOPS: Gopher just ate all the beans")
    }
}
//The bag of beans that the farmer purchased at the market
type BagOfBeans struct {
}
//The fox that the farmer purchased at the market
type Fox struct {
}
func (fox *Fox) Eat(g *Gopher, f *Farmer) {
    if g != nil && f == nil {
        log.Fatal("OOPS: Fox ate gopher just now.")
    }
}
//The boat which at any given time will only hold the farmer and ONE other item
type Boat struct {
    farmer *Farmer
    gopher *Gopher
    fox    *Fox
    beans  *BagOfBeans
}
//Defines a const of East or West
type Side int
const (
    East = 0
    West = 1
)
//The bank will represent the lands on either side of the river
type Bank struct {
    farmer *Farmer
    fox    *Fox
    gopher *Gopher
    beans  *BagOfBeans
    side   Side
}
//The river that the farmer must cross with all his belongings
var river = make(chan Boat, 1)
As you can see, the types representing each of our 
