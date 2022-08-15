// package channel

// import (
// 	"context"
// 	"time"

// 	"k8s.io/klog/v2"
// )

// type WalkieTalkie struct {
// 	Content string
// }

// var voices = make(chan WalkieTalkie, 10)

// func (m *WalkieTalkie) Broadcast() {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
// 	defer cancel()

// 	klog.Infof("talkie: %s", m.Content)
// }

// func walkieTalkie(q <-chan WalkieTalkie) {
// 	for {
// 		select {
// 		case talkie := <-q:
// 			talkie.Broadcast()
// 		}
// 	}
// }

// func message(talkie string) {
// 	manager := &WalkieTalkie{
// 		Content: talkie,
// 	}
// 	voices <- *manager
// }

// func main() {
// 	go walkieTalkie(voices)

// 	message("10-4")

// 	// time.Sleep(5 * time.Second)
// }
