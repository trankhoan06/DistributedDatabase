package worker

//
//import (
//	"context"
//	"log"
//	"main.go/config"
//)
//
//type CountJob struct {
//	cfg config.Configuration
//}
//
//var CountQueue chan CountJob
//
//func InitCountWorker(
//	ctx context.Context,
//	queueSize int,
//) {
//	CountQueue = make(chan CountJob, queueSize)
//	log.Printf("🚀 Initializing Count worker with queue size: %d", queueSize)
//	go func() {
//		for {
//			select {
//			case <-ctx.Done():
//				return
//
//			case job := <-CountQueue:
//				err := sender.SendMailResend(
//					job.Subject,
//					job.Content,
//					job.To,
//					nil,
//					nil,
//					job.File,
//					job.Buf,
//				)
//				if err != nil {
//					log.Println("❌ send email error:", err)
//				} else {
//					log.Println("send mail success")
//				}
//			}
//
//		}
//	}()
//}
