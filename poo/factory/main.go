package main

import "fmt"

// Factory para el SMS

type INotificationFactory interface {
	SendNotification()  // Enviar notificacion
	GetSender() ISender // Obtener remitente
}

type ISender interface {
	GetSenderMethod() string  // Obtener método del remitente
	GetSenderChannel() string // Obtener canal del remitente
}

type SMSNotification struct {
}

// estamos creando el metodo para SendNotification
func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Factory para el Email

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	//suponemos que estamos implemetnado un servicio de amazon para enviar el correo.
	return "SES"
}

// Creamos la funcion para que nos retorne las clases concretas que quermos utilizar para enviar las notificaciones
// resivimos un tipo de notificacion como parametro (para saber cual es el que se tiene que enviar )
// Vamos a retornar es un INotificationFactory o  un error
func getNotificationFactory(notificationType string) (INotificationFactory, error) {

	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("Notification Type")
}

// Creao la funcion que me captura cual es la notificacion que se va enviar
// la salida del getNotificationFactory puede ser tanto INotificationFactory como error
func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {

	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
