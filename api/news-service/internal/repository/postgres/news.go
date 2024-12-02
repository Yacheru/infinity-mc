package postgres

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"news-service/init/logger"
	"news-service/internal/entities"
	"news-service/pkg/constants"
)

type NewsPostgres struct {
	db *sqlx.DB
}

func NewNewsPostgres(db *sqlx.DB) *NewsPostgres {
	return &NewsPostgres{db: db}
}

func (p *NewsPostgres) AddNew(ctx *gin.Context, discordId int, embed *entities.Embed) (*int, error) {
	tx, err := p.db.BeginTx(ctx.Request.Context(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}

	var fid *int
	if embed.Footer != nil {
		footerQuery := `INSERT INTO footers (text, icon_url) VALUES ($1, $2) RETURNING id`
		err = p.db.GetContext(ctx.Request.Context(), &fid, footerQuery, embed.Footer.Text, embed.Footer.IconURL)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	var imgid *int
	if embed.Image != nil {
		imageQuery := `INSERT INTO resources (url) VALUES ($1) RETURNING id`
		err = p.db.GetContext(ctx.Request.Context(), &imgid, imageQuery, embed.Image.URL)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	var tid *int
	if embed.Thumbnail != nil {
		thumbnailQuery := `INSERT INTO resources (url) VALUES ($1) RETURNING id`
		err = p.db.GetContext(ctx.Request.Context(), &tid, thumbnailQuery, embed.Thumbnail.URL)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	var vid *int
	if embed.Video != nil {
		videoQuery := `INSERT INTO resources (url) VALUES ($1) RETURNING id`
		err = p.db.GetContext(ctx.Request.Context(), &vid, videoQuery, embed.Video.URL)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	var aid *int
	if embed.Author != nil {
		authorQuery := `INSERT INTO authors (name, url, icon_url) VALUES ($1, $2, $3) RETURNING id`
		err = p.db.GetContext(ctx.Request.Context(), &aid, authorQuery, embed.Author.Name, embed.Author.URL, embed.Author.IconURL)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	fids := make([]int, 0)
	if embed.Fields != nil {
		var fid int
		fieldQuery := `INSERT INTO fields (name, value, inline) VALUES ($1, $2, $3) RETURNING id`
		for _, field := range *embed.Fields {
			err = p.db.GetContext(ctx.Request.Context(), &fid, fieldQuery, field.Name, field.Value, field.Inline)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
			fids = append(fids, fid)
		}
	}

	if embed.CreatedAt == nil {
		now := time.Now().UTC()
		embed.CreatedAt = &now
		embed.UpdatedAt = &now
	}

	var nid int
	newsQuery := `INSERT INTO news (discord_id, title, description, url, color, footer, image, thumbnail, video, author, created_at, updated_at) 
				  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
	err = p.db.GetContext(ctx.Request.Context(), &nid, newsQuery, discordId, embed.Title, embed.Description, embed.URL, embed.Color, fid, imgid, tid, vid, aid, embed.CreatedAt, embed.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if len(fids) != 0 {
		newsFieldsQuery := `INSERT INTO news_fields (news_id, field_id) VALUES ($1, $2)`
		for _, fid := range fids {
			_, err = p.db.ExecContext(ctx.Request.Context(), newsFieldsQuery, nid, fid)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	return &nid, tx.Commit()
}

func (p *NewsPostgres) GetAllNews(ctx *gin.Context) ([]*entities.Embed, error) {
	var entityEmbeds []*entities.Embed

	query := `
		SELECT 
		    discord_id, title, description, news.url AS news_url, color,
		    footers.text AS footer_text, footers.icon_url AS footer_icon_url,
		    img.url AS image_url,
		    thumb.url AS thumbnail_url,
		    vid.url AS video_url,
		    authors.name AS author_name, authors.url AS author_url, authors.icon_url AS author_icon_url,
			CASE 
				WHEN COUNT(fields.id) = 0 THEN NULL
				ELSE json_agg(json_build_object('name', fields.name, 'value', fields.value, 'inline', fields.inline)) 
    		END AS fields,
			created_at, news.updated_at
		FROM news
		LEFT JOIN footers ON news.footer = footers.id
		LEFT JOIN resources img ON news.image = img.id
		LEFT JOIN resources thumb ON news.thumbnail = thumb.id
		LEFT JOIN resources vid ON news.video = vid.id
		LEFT JOIN authors ON news.author = authors.id
		LEFT JOIN news_fields ON news.id = news_fields.news_id
		LEFT JOIN fields ON news_fields.field_id = fields.id
		GROUP BY discord_id, title, description, news.url, created_at, color, footers.text, footers.icon_url, img.url, thumb.url, vid.url, authors.name, authors.url, authors.icon_url, created_at, updated_at
	`

	rows, err := p.db.QueryContext(ctx.Request.Context(), query)
	if err != nil {
		logger.Error(err.Error(), constants.LoggerPostgres)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			discordId   int64
			title       sql.NullString
			description sql.NullString
			URL         sql.NullString
			color       sql.NullInt64

			footerText    sql.NullString
			footerIconURL sql.NullString
			imageURL      sql.NullString
			thumbnailURL  sql.NullString
			videoURL      sql.NullString

			authorName    sql.NullString
			authorURL     sql.NullString
			authorIconURL sql.NullString

			fields []byte

			createdAt sql.NullTime
			updatedAt sql.NullTime
		)

		err = rows.Scan(
			&discordId, &title, &description, &URL, &color,
			&footerText, &footerIconURL,
			&imageURL, &thumbnailURL, &videoURL,
			&authorName, &authorURL, &authorIconURL,
			&fields, &createdAt, &updatedAt,
		)
		if err != nil {
			logger.Error(err.Error(), constants.LoggerPostgres)
			return nil, err
		}

		var fieldsArr = new([]entities.EmbedField)
		if fields != nil {
			err = json.Unmarshal(fields, &fieldsArr)
			if err != nil {
				logger.Error(err.Error(), constants.LoggerPostgres)
				return nil, err
			}
		}

		embed := &entities.Embed{
			DiscordId:   discordId,
			Title:       title.String,
			Description: description.String,
			URL:         URL.String,
			Color:       color.Int64,
			Footer: &entities.EmbedFooter{
				Text:    footerText.String,
				IconURL: footerIconURL.String,
			},
			Image: &entities.EmbedResource{
				URL: imageURL.String,
			},
			Thumbnail: &entities.EmbedResource{
				URL: thumbnailURL.String,
			},
			Video: &entities.EmbedResource{
				URL: videoURL.String,
			},
			Author: &entities.EmbedAuthor{
				Name:    authorName.String,
				URL:     authorURL.String,
				IconURL: authorIconURL.String,
			},
			Fields:    fieldsArr,
			CreatedAt: &createdAt.Time,
			UpdatedAt: &updatedAt.Time,
		}

		if rows.Err() != nil {
			logger.Error(rows.Err().Error(), constants.LoggerPostgres)
			return nil, err
		}

		entityEmbeds = append(entityEmbeds, embed)
	}

	return entityEmbeds, nil
}

func (p *NewsPostgres) GetNewsById(ctx *gin.Context, discordId int) (*entities.Embed, error) {
	query := `
		SELECT 
		    n.discord_id, n.title, n.description, n.url, n.created_at, n.color,
		    f.text AS footer_text, f.icon_url AS footer_icon_url, 
		    ri.url AS image_url,
		    rt.url AS thumbnail_url,
		    rv.url AS video_url,
		    a.name AS author_name, a.url AS author_url, a.icon_url AS author_icon_url
		FROM news n
			JOIN footers f ON f.id = n.footer
			JOIN resources ri ON ri.id = n.image
			JOIN resources rt ON rt.id = n.thumbnail
			JOIN resources rv ON rv.id = n.video
			JOIN authors a ON a.id = n.author
			JOIN news_fields nf ON nf.news_id = n.id
		WHERE discord_id = $1
	`

	row := p.db.QueryRowxContext(ctx.Request.Context(), query, discordId)
	if row.Err() != nil {
		logger.Error(row.Err().Error(), constants.LoggerPostgres)
		return nil, row.Err()
	}

	var embed = new(entities.Embed)
	var footer entities.EmbedFooter
	var image, thumbnail, video entities.EmbedResource
	var author entities.EmbedAuthor

	err := row.Scan(
		&embed.DiscordId, &embed.Title, &embed.Description, &embed.URL, &embed.CreatedAt, &embed.Color,
		&footer.Text, &footer.IconURL,
		&image.URL, &thumbnail.URL, &video.URL,
		&author.Name, &author.URL, &author.IconURL,
	)
	if err != nil {
		logger.Error(err.Error(), constants.LoggerPostgres)
		return nil, err
	}

	embed.Footer = &footer
	embed.Image = &image
	embed.Thumbnail = &thumbnail
	embed.Video = &video
	embed.Author = &author

	return embed, nil
}
