package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UsedMixin{},
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(), //seq
		field.String("user_id"),     // 사용자 아이디
		field.String("user_name"),   // 반려묘/견 이름
		field.String("birthday"),    // 생일
		field.Float("weight"),       //몸무게
		field.String("type"),        //종류
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
