package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Device struct {
	ID             primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string              `json:"name, omitempty" bson:"name, omitempty"`
	Serial         string              `json:"serial" bson:"serial"`
	Configurations []map[string]string `json:"configurations, omitempty" bson:"configurations"`
}
