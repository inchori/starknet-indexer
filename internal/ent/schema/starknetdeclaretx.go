package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// StarknetDeclareTx holds the schema definition for the StarknetDeclareTx entity.
type StarknetDeclareTx struct {
	ent.Schema
}

// Fields of the StarknetDeclareTx.
func (StarknetDeclareTx) Fields() []ent.Field {
	return []ent.Field{
		field.Int("block_number"),
		field.String("declare_tx_hash"),
		field.String("class_hash"),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the StarknetDeclareTx.
func (StarknetDeclareTx) Edges() []ent.Edge {
	return nil
}
