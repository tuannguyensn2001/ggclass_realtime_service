package models

import "time"

type Assignment struct {
	AssignmentId int        `bson:"assignmentId,omitempty"`
	Action       string     `bson:"action,omitempty"`
	CreatedAt    *time.Time `bson:"createdAt,omitempty"`
	UpdatedAt    *time.Time `bson:"updatedAt,omitempty"`
	Id           string     `bson:"_id,omitempty"`
}

func (Assignment) CollectionName() string {
	return "assignments"
}
