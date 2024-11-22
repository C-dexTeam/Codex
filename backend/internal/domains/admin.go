package domains

import "context"

// IAdminService is the interface that provides the methods for the user service.
type IAdminService interface {
	GetAllUsers(ctx context.Context, id, username, email, page, limit string) ([]User, error)
}
