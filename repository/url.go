package repository

import (
	"context"
	"log"
	"url-shortener-db-migrate/pkg/entity"
)

/*
note:

	r = receiver
*/

func (r *Postgres) InsertURL(ctx context.Context, url entity.URL) (id int, err error) {
	query := `INSERT INTO
				urls_short_and_target (url_short, url_target)
			VALUES
				($1, $2)
			RETURNING id`

	err = r.DB.QueryRowContext(ctx, query, url.URLShort, url.URLTarget).Scan(&id)
	if err != nil {
		log.Printf("Failed to insert URL: %v", err)
		return id, err
	}

	log.Println("URL inserted successfully")
	return id, nil
}

func (r *Postgres) UpdateURLShort(ctx context.Context, url entity.URL) error {
	query := `UPDATE
				urls_short_and_target
			SET
				url_short = $1
			WHERE
				id = $2`

	_, err := r.DB.Exec(query, url.URLShort, url.ID)
	if err != nil {
		log.Printf("Failed to update URL Short: %v", err)
		return err
	}

	return nil
}

func (r *Postgres) UpdateURLTarget(ctx context.Context, url entity.URL) error {
	query := `UPDATE
				urls_short_and_target
			SET
				url_target = $1
			WHERE
				id = $2`

	_, err := r.DB.Exec(query, url.URLTarget, url.ID)
	if err != nil {
		log.Printf("Failed to update URL Target: %v", err)
		return err
	}

	return nil
}

func (r *Postgres) GetURLTargetByURLShort(ctx context.Context, urlInput entity.URL) (entity.URL, error) {
	var urlOuput entity.URL

	query := `SELECT
				id,
				url_short,
				url_target
			FROM 
				urls_short_and_target
			WHERE
				url_short = $1`

	err := r.DB.QueryRowContext(ctx, query, urlInput.URLShort).Scan(
		&urlOuput.ID,
		&urlOuput.URLShort,
		&urlOuput.URLTarget,
	)
	if err != nil {
		log.Printf("Failed to get url target by url short: %v", err)
		return urlOuput, err
	}

	return urlOuput, nil
}

func (r *Postgres) GetAllURL(ctx context.Context) ([]entity.URL, error) {
	var urls []entity.URL

	query := `SELECT
				id,
				url_short,
				url_target
			FROM
				urls_short_and_target`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Failed to get all urls: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url entity.URL
		err := rows.Scan(
			&url.ID,
			&url.URLShort,
			&url.URLTarget,
		)
		if err != nil {
			log.Printf("Failed to scan url: %v", err)
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}
