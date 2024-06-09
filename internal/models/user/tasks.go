package scheduler

// import (
// 	"context"
// 	"time"
// )

// func CheckSubs(ctx context.Context, interval time.Duration) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			subs := getSubs()
// 			for _, s := range subs {
// 				if s.Birthday == today {
// 					sendReminderToEmail(s.Email)
// 				}
// 			}
// 		}
// 		time.Sleep(interval)
// 	}
// }

// func main1() {
// 	// создаём контекст с функцией завершения
// 	ctx, cancel := context.WithCancel(context.Background())
// 	// запускаем нашу горутину
// 	go task(ctx, time.Minute)
// 	// делаем паузу, чтобы дать горутине поработать
// 	time.Sleep(10 * time.Minute)
// 	// завершаем контекст, чтобы завершить горутину
// 	cancel()
// }
