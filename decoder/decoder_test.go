package decoder_test

import (
	"testing"

	"github.com/kondrushin/decoder/decoder"
	"github.com/stretchr/testify/assert"
)

func Test_OneChar(t *testing.T) {
	decodeNumber := decoder.Decode("")
	assert.Equal(t, 0, decodeNumber)
}

func Test_TwoChartsNoGrouping(t *testing.T) {
	decodeNumber := decoder.Decode("34")
	assert.Equal(t, 1, decodeNumber)
}

func Test_TwoChartsWithGrouping(t *testing.T) {
	decodeNumber := decoder.Decode("21")
	assert.Equal(t, 2, decodeNumber)
}

func Test_ThreeChartsNoGrouping(t *testing.T) {
	decodeNumber := decoder.Decode("467")
	assert.Equal(t, 1, decodeNumber)
}

func Test_ThreeChartsWithSingleGrouping(t *testing.T) {
	decodeNumber := decoder.Decode("183")
	assert.Equal(t, 2, decodeNumber)
}
func Test_ThreeChartsWithDoubleGrouping(t *testing.T) {
	decodeNumber := decoder.Decode("123")
	assert.Equal(t, 3, decodeNumber)
}

func Test_EndsWithZero_GroupingIsPossible(t *testing.T) {
	decodeNumber := decoder.Decode("120")
	assert.Equal(t, 1, decodeNumber)
}

func Test_EndsWithZero_GroupingNotPossible(t *testing.T) {
	decodeNumber := decoder.Decode("180")
	assert.Equal(t, 0, decodeNumber)
}

func Test_LongStringWithMultipleCombinations(t *testing.T) {
	decodeNumber := decoder.Decode("1253489220367151428347278482364238562017634722374868764827364238735628376587568476583745678345673485")
	assert.Equal(t, 576, decodeNumber)
}

func Test_EmptyString(t *testing.T) {
	decodeNumber := decoder.Decode("")
	assert.Equal(t, 0, decodeNumber)
}

func Test_StartsWithZero(t *testing.T) {
	decodeNumber := decoder.Decode("09")
	assert.Equal(t, 0, decodeNumber)
}
