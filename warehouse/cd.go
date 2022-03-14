package warehouse

type CD struct {
	stock           int
	paymentProvider PaymentProvider
	title           string
	artist string
}

func (cd *CD) Buy(quantity int) bool {

	if cd.stock - quantity < 0 {
		return false
	}

	if cd.paymentProvider.accept(){
		cd.stock -= quantity
		return true
	}
	return false
}

func (cd *CD) Restock(newItems int) {
	cd.stock += newItems
}

type CdStore struct{
	cds []CD
}

func (cdStore *CdStore) SearchByTitle(title string) int {

	for index := range cdStore.cds{
		if cdStore.cds[index].title == title {
			return index
		}
	}
	return -1
}

func (cdStore *CdStore) SearchByArtist(artist string) int {

	for index := range cdStore.cds {
		if cdStore.cds[index].artist == artist {
			return index
		}
	}
	return -1
}




type PaymentProvider interface {
	accept() bool
}
