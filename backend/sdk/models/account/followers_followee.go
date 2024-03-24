package account

import "github.com/google/uuid"

type FollowerFollowee struct {
	ID          uuid.UUID
	EntityID    uuid.UUID
	EntityType  string //ENUM of user, page, company
	FollowingID uuid.UUID
	FollowerID  uuid.UUID
}
