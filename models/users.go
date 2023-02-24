package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `bson:"_id, omitempty" json="id"`
	Name string           `bson:"name" json="name, omitempty"`
	Lastname string       `bson:"lastname" json="lastname, omitempty"`
	Birthdate time.Time   `bson:"birthdate" json="birthdate, omitempty"`
	Email string          `bson:"email" json="email"`
	Password string       `bson:"password" json="password, omitempty"`
	Avatar string         `bson:"avatar" json="avatar, omitempty"`
	Banner string         `bson:"password" json="password, omitempty"`
	Biography string      `bson:"biography" json="biography, omitempty"`
	Location string       `bson:"location" json="location, omitempty"`
	WebSite string        `bson:"web_site" json="web_site, omitempty`
}