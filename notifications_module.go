package main

import (
	"beacon/models"
	"beacon/repositories"
	"beacon/services"
)

var notificationService *services.NotificationService

func initNotificationModule() {
	notificationService = services.NewNotificationService(repositories.NewNotificationRepository())
	notificationService.Initialize()
}

func dispatchAuditNotification(userID int, username, action, resource, status, message string) {
	if notificationService == nil {
		return
	}
	notificationService.DispatchAuditEvent(models.AuditNotificationEvent{
		UserID: userID, Username: username, Action: action,
		Resource: resource, Status: status, Message: message,
	})
}
