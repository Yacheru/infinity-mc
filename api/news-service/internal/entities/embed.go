package entities

import "time"

type Embed struct {
	DiscordId   int64          `json:"discord_id" db:"discord_id"`
	Title       string         `json:"title,omitempty" db:"title"`
	Description string         `json:"description,omitempty" db:"description"`
	URL         string         `json:"url,omitempty" db:"url"`
	Color       int64          `json:"color,omitempty" db:"color"`
	Footer      *EmbedFooter   `json:"footer,omitempty" db:"footer"`
	Image       *EmbedResource `json:"image,omitempty" db:"image"`
	Thumbnail   *EmbedResource `json:"thumbnail,omitempty" db:"thumbnail"`
	Video       *EmbedResource `json:"video,omitempty" db:"video"`
	Author      *EmbedAuthor   `json:"author,omitempty" db:"author"`
	Fields      *[]EmbedField  `json:"fields,omitempty" db:"fields" binding:"max=25"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty" db:"updated_at"`
	CreatedAt   *time.Time     `json:"timestamp,omitempty" db:"created_at"` // ISO 8601 || RFC 3339
}

type EmbedResource struct {
	Id  int    `json:"id,omitempty" db:"id"`
	URL string `json:"url,omitempty" db:"url"`
}

type EmbedAuthor struct {
	Id      int    `json:"id,omitempty" db:"id"`
	Name    string `json:"name,omitempty" db:"name"`
	URL     string `json:"url,omitempty" db:"url"`
	IconURL string `json:"icon_url,omitempty" db:"icon_url"`
}

type EmbedFooter struct {
	Id      int    `json:"id,omitempty" db:"id"`
	Text    string `json:"text,omitempty" db:"text"`
	IconURL string `json:"icon_url,omitempty" db:"icon_url"`
}

type EmbedField struct {
	Id     int    `json:"id,omitempty" db:"id"`
	Name   string `json:"name" db:"name"`
	Value  string `json:"value" db:"value"`
	Inline *bool  `json:"inline,omitempty" db:"inline"`
}
