package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("event_id").Nillable(),                                   //イベントID                                  //イベント名
		field.Int("amount"),                                                //支払金額
		field.Int("paid_by"),                                               //支払った（お金を返した)ユーザーID
		field.Int("paid_to"),                                               //支払ってもらった（お金を返してもらった)ユーザーID
		field.Time("paid_at").Default(time.Now).Immutable(),                //支払日
		field.Time("created_at").Default(time.Now).Immutable(),             //作成時刻
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), //更新時刻
		field.Time("deleted_at").Nillable().Optional(),                     //削除フラグ 削除時刻
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{}
}
