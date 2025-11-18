package main

import "fmt"

type Notifier interface {
	Notify(messages string) error
}

type EmailNotifier struct {
	SMTPHost string
	Port     int
}

func (e *EmailNotifier) Notify(msg string) error {
	fmt.Printf("发送邮件通知:%s\n", msg)
	return nil
}

type SmsNotifier struct {
	APIKey   string
	TmplCode string
}

func (s *SmsNotifier) Notify(msg string) error {
	fmt.Printf("发送邮件通知:%s\n", msg)
	return nil
}

type BroadCastNotifier struct {
	NotifierS []Notifier
}

func (b *BroadCastNotifier) Notify(msg string) error {
	for _, notifier := range b.NotifierS {
		err := notifier.Notify(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

type OrderService struct {
	notifier Notifier
}

func (o *OrderService) SetNotifier(n Notifier) {
	o.notifier = n
}

func (o *OrderService) CreateOrder(product string, quantity int) {
	fmt.Printf("创建订单：%s x %d\n", product, quantity)
	err := o.notifier.Notify("订单已创建")
	if err != nil {
		fmt.Println("通知失败", err)
	}

}

func notifierMain() {
	orderService := &OrderService{}

	emailNotifier := &EmailNotifier{
		SMTPHost: "smtp.example.com",
		Port:     587,
	}

	smsNotifier := &SmsNotifier{
		APIKey:   "your-api-key",
		TmplCode: "your-template-Code",
	}

	broadCastNotifier := &BroadCastNotifier{
		NotifierS: []Notifier{emailNotifier, smsNotifier},
	}

	orderService.SetNotifier(emailNotifier)
	orderService.CreateOrder("手机-email", 1)

	orderService.SetNotifier(smsNotifier)
	orderService.CreateOrder("手机-sms", 1)

	orderService.SetNotifier(broadCastNotifier)
	orderService.CreateOrder("手机-broad", 1)

}
