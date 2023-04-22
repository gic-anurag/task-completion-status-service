package taskCompletionStatus

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskCompletionStatus struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	WorkflowName       string             `bson:"workflowName" json:"workflowName"`
	CompletedEventName string             `bson:"completedEventName" json:"completedEventName"`
	TransactioinId     string             `bson:"transactioinId" json:"transactioinId"`
	EventId            string             `bson:"eventId" json:"eventId"`
	Status             string             `bson:"status" json:"status"`
	CreationDateTime   time.Time          `bson:"creationDateTime" json:"creationDateTime"`
}

type filterRequest struct {
	Id                 string    `bson:"_id,omitempty" json:"id,omitempty"`
	WorkflowName       string    `bson:"workflowName" json:"workflowName"`
	CompletedEventName string    `bson:"completedEventName" json:"completedEventName"`
	TransactioinId     string    `bson:"transactioinId" json:"transactioinId"`
	EventId            string    `bson:"eventId" json:"eventId"`
	Status             string    `bson:"status" json:"status"`
	CreationDateTime   time.Time `bson:"creationDateTime" json:"creationDateTime"`
}

type GetTaskCompletionStatus struct {
	Id string `bson:"id,omitempty"`
}
