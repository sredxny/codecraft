package warehouse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cDStore CdStore

func init() {
	cDStore.cds = []CD{
		{
			title:  "my-title",
			stock:  10,
			artist: "artist1",
		},
		{
			title:  "my-title2",
			stock:  10,
			artist: "artist2",
		},
		{
			title:  "my-title3",
			stock:  0,
			artist: "artist3",
		},
	}
}

type PaymentProviderStub struct {
	accepted bool
}

func (paymentProviderStub PaymentProviderStub) accept() bool {
	return paymentProviderStub.accepted
}

func Test_SearchByTitle(t *testing.T) {

	t.Run("found title", func(t *testing.T) {
		title := "my-title"

		foundIndex := cDStore.SearchByTitle(title)
		assert.NotEqual(t, -1, foundIndex)
	})

	t.Run("not found title", func(t *testing.T) {
		title := "my-title4"

		foundIndex := cDStore.SearchByTitle(title)
		assert.Equal(t, -1, foundIndex)
	})
}

func Test_SearchByArtist(t *testing.T) {

	t.Run("artist found", func(t *testing.T) {
		artist := "artist1"
		foundIndex := cDStore.SearchByArtist(artist)
		assert.NotEqual(t, -1, foundIndex)
	})

	t.Run("artist not found", func(t *testing.T) {
		artist := "artist4"
		foundIndex := cDStore.SearchByArtist(artist)
		assert.Equal(t, -1, foundIndex)
	})
}

func Test_BuyCD(t *testing.T) {

	t.Run("payment successful", func(t *testing.T) {
		paymentProvider := PaymentProviderStub{accepted: true}
		cd := CD{stock: 10, paymentProvider: paymentProvider}
		cd.Buy(1)

		assert.Equal(t, 9,cd.stock)
	})

	t.Run("payment failure", func(t *testing.T) {
		paymentProvider := PaymentProviderStub{accepted: false}
		cd := CD{stock: 10, paymentProvider: paymentProvider}
		cd.Buy(1)

		assert.Equal(t, 10,cd.stock)
	})

	t.Run("buy more than stock", func(t *testing.T) {
		paymentProvider := PaymentProviderStub{accepted: true}
		cd := CD{stock: 5, paymentProvider: paymentProvider}
		buySuccessful := cd.Buy(10)

		assert.Equal(t,false,buySuccessful)
		assert.Equal(t, 5,cd.stock)
	})
}

func Test_Restock(t *testing.T){
	cd := CD{stock:0}
	cd.Restock(10)
	assert.Equal(t,10, cd.stock)
}

// search cd
// buy
// pay cd
// ask for review
