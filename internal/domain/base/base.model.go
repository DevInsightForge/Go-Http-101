package base

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	CreatedBy string             `bson:"createdBy" json:"createdBy"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	UpdatedBy string             `bson:"updatedBy" json:"updatedBy"`
}

func (bm *BaseModel) SetAuditFieldsBeforeCreate(createdBy string) error {
	bm.CreatedAt = time.Now()
	bm.CreatedBy = createdBy
	return nil
}

func (bm *BaseModel) SetAuditFieldsBeforeUpdate(updatedBy string) error {
	bm.UpdatedAt = time.Now()
	bm.UpdatedBy = updatedBy
	return nil
}
