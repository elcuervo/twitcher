package twitcher

import (
	"testing"
	"time"
	"git.tideland.biz/gots/asserts"
)

func TestTwitcherExecution(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)
        twitcher := NewTwitcher(time.Second * 1)
        items := []int{}

        twitcher.Do(func() {
                items = append(items, 1)
        })

        time.Sleep(time.Second * 2)
        twitcher.Stop()
        assert.Equal(2, len(items))
}

func TestTwitcherUpdate(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)
        twitcher := NewTwitcher(time.Second * 1)
        items := []int{}

        twitcher.Do(func() {
                items = append(items, 1)
        })

        twitcher.Update(time.Millisecond * 500)

        time.Sleep(time.Second * 2)
        twitcher.Stop()
        assert.Equal(4, len(items))
}
