package main

import (
	cu "github.com/Davincible/chromedp-undetected"
	"github.com/chromedp/chromedp"
	"time"
)

func main() {
	thread := 3
	for i := 0; i < thread; i++ {
		go func() {
			for {
				RunThread()
			}
		}()
	}
	forever := make(chan bool)
	<-forever
}

func RunThread() error {
	ctx, cancel, err := cu.New(cu.NewConfig(
	// Remove this if you want to see a browser window.
	//cu.WithHeadless(),
	))
	if err != nil {
		return err
	}
	defer cancel()

	if err := chromedp.Run(ctx,
		// Check if we pass anti-bot measures.
		chromedp.Navigate("https://nowsecure.nl"),
	); err != nil {
		return err
	}
	time.Sleep(1000 * time.Second)
	return nil
}
