package twitcher

import "time"

type Twitcher struct {
        C <-chan time.Time
        Work bool
}

type caller func()

func (tw *Twitcher) Do(fn caller) {
        tw.Work = true

        go func() {
                for {
                        <-tw.C

                        if(tw.Work) {
                                fn()
                        } else {
                                break
                        }
                }
        }()
}

func (tw *Twitcher) Stop() {
        tw.Work = false
}

func (tw *Twitcher) Update(t time.Duration) {
        tw.C = time.NewTicker(t).C
}

func NewTwitcher(t time.Duration) *Twitcher {
        tw := new(Twitcher)
        tw.C = time.NewTicker(t).C

        return tw
}
