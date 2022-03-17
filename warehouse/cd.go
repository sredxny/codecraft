package warehouse

// CD struct to store
type CD struct {
	stock             int
	paymentProvider   PaymentProvider
	title             string
	artist            string
	Reviews           []Review
	chartNotification ChartNotification
	price             float32
}

// Review ...
type Review struct {
	rating  int
	comment string
}

// IsValid check if review is valid
func (r Review) IsValid() bool {
	return r.rating >= 1 && r.rating <= 10
}

// Buy buys a cd
func (cd *CD) Buy(quantity int) bool {

	if cd.stock-quantity < 0 {
		return false
	}

	if cd.paymentProvider.accept() {
		cd.stock -= quantity
		cd.chartNotification.Notify(cd.artist, cd.title, quantity)

		return true
	}
	return false
}

// Restock add items to stock
func (cd *CD) Restock(newItems int) {
	cd.stock += newItems
}

// AddReview add a review
func (cd *CD) AddReview(rating int, comment string) {

	review := Review{
		rating:  rating,
		comment: comment,
	}

	if review.IsValid() {
		cd.Reviews = append(cd.Reviews, review)
	}
}

type CdStore struct {
	cds []CD
	competitorData CompetitorData
	top100Service Top100
}

func (cdStore *CdStore) SearchByTitle(title string) []int {
	foundIndexes := []int{}
	for index := range cdStore.cds {
		currentCD := cdStore.cds[index]
		if currentCD.title == title {

			if cdStore.top100Service.Search(title, currentCD.artist){
				foundInCompetitor := cdStore.competitorData.Search(title,currentCD.artist)
				if foundInCompetitor != nil {
					cdStore.cds[index].price = foundInCompetitor.price -1
				}
			}

			foundIndexes = append(foundIndexes, index)
		}
	}
	return foundIndexes
}

func (cdStore *CdStore) SearchByArtist(artist string) []int {
	foundIndexes := []int{}
	for index := range cdStore.cds {
		if cdStore.cds[index].artist == artist {
			foundIndexes = append(foundIndexes, index)
		}
	}
	return foundIndexes
}

func (cdStore *CdStore) AddCd(cdTitle string, cdArtist string, cdStock int) {
	cd := CD{
		stock:           cdStock,
		title:           cdTitle,
		artist:          cdArtist,
	}

	cdStore.cds = append(cdStore.cds, cd)
}

type PaymentProvider interface {
	accept() bool
}

type ChartNotification interface {
	Notify(artist string, title string, quantity int)
}

type CompetitorData interface {
	Search(title, artist string) *CD
}

type Top100 interface {
	Search(title, artist string) bool
}


