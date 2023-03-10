package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

type UsedMixin struct {
	mixin.Schema
}

func (UsedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_used").Default(true),
	}
}
