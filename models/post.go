package models

import (
 // "time"
 // "todo-api/models/db"
)

type Post struct {
 Title string "json:'title' bson:'title'"
 // SlugUrl string "json:”slug_url" bson:"slug_url"'
 // Content string "json:"content" bson:"content"'
 // PublishedAt time.Time "json:”published_at” bson:”published_at” ID bson.ObjectId `json:"id" bson:"_id"'
 // CreatedAt time.Time "json:"created_at,omitempty" bson:"created_at"'
 // UpdatedAt time.Time "json:"updated_at,omitempty" bson:"updated_at"'
}
