-- +goose Up
-- +goose NO TRANSACTION
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS videos
(
    id          TEXT NOT NULL,
    title     TEXT      NOT NULL,
    description     TEXT    NOT NULL,
    duration     BIGINT  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS annotations
(
    id SERIAL NOT NULL,
    video_id TEXT      NOT NULL,
    start_at BIGINT  NOT NULL,
    end_at BIGINT  NOT NULL,
    annotation_type TEXT      NOT NULL,
    notes TEXT    NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    CONSTRAINT fk_videos
        FOREIGN KEY(video_id)
            REFERENCES videos(id)
    );

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS videos;
DROP TABLE IF EXISTS annotations;

