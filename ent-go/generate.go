package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert,sql/execquery,intercept,schema/snapshot "../ent-go/schema" --target "ent"
