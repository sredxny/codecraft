package warehouse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cDStoreTemplate CdStore

func init() {
	cDStoreTemplate.cds = []CD{
		{
			title:  "my-title",
			stock:  10,
			artist: "artist1",
			price: 5,
		},
		{
			title:  "my-title",
			stock:  10,
			artist: "artist2",
			price: 2,
		},
		{
			title:  "my-title3",
			stock:  0,
			artist: "artist1",
			price: 99,
		},
	}
}

type PaymentProviderStub struct {
	accepted bool
}

func (paymentProviderStub PaymentProviderStub) accept() bool {
	return paymentProviderStub.accepted
}

type CompetitorDataStub struct {
	cds []CD
}

func (competitorDataStub CompetitorDataStub) Search(title, artist string) *CD {
	for index := range competitorDataStub.cds {
		if competitorDataStub.cds[index].title == title && competitorDataStub.cds[index].artist == artist {
			return &competitorDataStub.cds[index]
		}
	}
	return nil
}

type Top100Stub struct {
	cds []CD
}

func (top100Stub Top100Stub) Search(title, artist string) bool {
	for index := range  top100Stub.cds {
		cd := top100Stub.cds[index]
		if cd.title == title && cd.artist == artist {
			return true
		}
	}
	return false
}

func Test_SearchByTitle(t *testing.T) {
	cDStore := cDStoreTemplate
	cDStore.competitorData = CompetitorDataStub{}
	cDStore.top100Service = Top100Stub{}
	t.Run("found title with many results", func(t *testing.T) {
		foundIndexes := cDStore.SearchByTitle("my-title")
		assert.Len(t, foundIndexes, 2)
	})

	t.Run("found title that belongs to top 100", func(t *testing.T) {

		title := "my-title"

		competitorCDs := cDStore.cds
		for index := range competitorCDs {
			competitorCDs[index].price = 2
		}

		cDStore.top100Service = Top100Stub{cds:[]CD{cDStore.cds[0]}}

		competitorDataStub := CompetitorDataStub{cds:competitorCDs}
		cDStore.competitorData = competitorDataStub
		foundIndexes := cDStore.SearchByTitle(title)

		for index := range foundIndexes{
			cd := cDStore.cds[index]
			newPrice := competitorCDs[index].price
			assert.Equal(t, newPrice, cd.price)
		}
	})

	t.Run("not found title", func(t *testing.T) {
		title := "my-title4"

		foundIndexes := cDStore.SearchByTitle(title)
		assert.Len(t, foundIndexes, 0)
	})
}

func Test_SearchByArtist(t *testing.T) {
	cDStore := cDStoreTemplate
	t.Run("artist found", func(t *testing.T) {
		artist := "artist1"
		foundIndexes := cDStore.SearchByArtist(artist)
		assert.Len(t, foundIndexes, 2)
	})

	t.Run("artist not found", func(t *testing.T) {
		artist := "artist4"
		foundIndexes := cDStore.SearchByArtist(artist)
		assert.Len(t, foundIndexes, 0)
	})
}

type ChartNotificationMock struct {
	artist, title string
	quantity      int
}

func (chartNotificationMock *ChartNotificationMock) Notify(artist string, title string, quantity int) {
	chartNotificationMock.artist = artist
	chartNotificationMock.title = title
	chartNotificationMock.quantity = quantity
}

func Test_BuyCD(t *testing.T) {

	t.Run("payment successful", func(t *testing.T) {
		chartNotificationMock := ChartNotificationMock{}
		paymentProvider := PaymentProviderStub{accepted: true}

		buyQuantity := 1
		cd := CD{stock: 10,
			title:             "my-title",
			artist:            "my-artist",
			paymentProvider:   paymentProvider,
			chartNotification: &chartNotificationMock,
		}
		buySuccessful := cd.Buy(buyQuantity)

		assert.Equal(t, true, buySuccessful)
		assert.Equal(t, 9, cd.stock)
		assert.Equal(t, cd.artist, chartNotificationMock.artist)
		assert.Equal(t, cd.title, chartNotificationMock.title)
		assert.Equal(t, buyQuantity, chartNotificationMock.quantity)
	})

	t.Run("payment failure", func(t *testing.T) {
		paymentProvider := PaymentProviderStub{accepted: false}
		cd := CD{stock: 10, paymentProvider: paymentProvider}
		buySuccessful := cd.Buy(1)

		assert.Equal(t, false, buySuccessful)
		assert.Equal(t, 10, cd.stock)
	})

	t.Run("buy more than stock", func(t *testing.T) {
		initialStock := 5
		paymentProvider := PaymentProviderStub{accepted: true}
		cd := CD{stock: initialStock, paymentProvider: paymentProvider}
		buySuccessful := cd.Buy(10)

		assert.Equal(t, false, buySuccessful)
		assert.Equal(t, initialStock, cd.stock)
	})
}

func Test_Restock(t *testing.T) {
	cd := CD{stock: 0}
	cd.Restock(10)
	assert.Equal(t, 10, cd.stock)
}

func Test_AddNewCD(t *testing.T) {
	cDStore := cDStoreTemplate

	originalAvailableCDs := len(cDStore.cds)
	newAvailableCDs := originalAvailableCDs + 1

	cDStore.AddCd("cd-title", "cd-artist", 10)
	assert.Equal(t, newAvailableCDs, len(cDStore.cds))
}

func Test_Reviews(t *testing.T) {

	t.Run("add review with comment and valid rating", func(t *testing.T) {
		cd := CD{}
		cd.AddReview(10, "some review")
		assert.Len(t, cd.Reviews, 1)
	})

	t.Run("add review with negative rating", func(t *testing.T) {
		cd := CD{}
		cd.AddReview(-1, "some review")
		assert.Len(t, cd.Reviews, 0)
	})

	t.Run("add review with rating above 10", func(t *testing.T) {
		cd := CD{}
		cd.AddReview(100, "some review")
		assert.Len(t, cd.Reviews, 0)
	})
}
