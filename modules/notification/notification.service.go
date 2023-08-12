package notification

type Notification interface {
	Send(message string)
}

type EmailNotification struct {
}

func (EmailNotification) Send(message string) {

}

type SmsNotification struct {
}

func (SmsNotification) Send(message string) {
}

type SlackNotification struct {
}

func (SlackNotification) Send(message string) {
}

type NotificationDecorator struct {
	core         *NotificationDecorator
	notification Notification
}

func (nd NotificationDecorator) Send(message string) {
	nd.core.notification.Send(message)

	if nd.core != nil {
		nd.core.Send(message)
	}
}

func (nd NotificationDecorator) Decorate(notification Notification) NotificationDecorator {
	return NotificationDecorator{
		core:         &nd,
		notification: notification,
	}
}

func NewNotificationDecorator(notification Notification) NotificationDecorator {
	return NotificationDecorator{notification: notification}
}

type Service struct {
	notification Notification
}

func (s Service) SendNotification(message string) {
	s.notification.Send(message)
}

func send() {
	notification := NewNotificationDecorator(
		EmailNotification{}).Decorate(SmsNotification{}).Decorate(SlackNotification{})

	s := Service{
		notification: notification,
	}

	s.SendNotification("Hello")
}
