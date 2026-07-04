package sapliy

// NotificationsService handles webhook management and alerting.
type NotificationsService struct {
	c *Client
}

// No explicit public endpoints for notifications yet in openapi, but reserving for future.
// Workers in the backend handle processing. Webhook Management would go here.
