-- name: ListLoadouts :many
SELECT *
FROM public.loadouts
ORDER BY title;
-- name: ListGames :many
SELECT *
FROM public.games
ORDER BY name;
-- name: CreateLoadout :one
INSERT INTO public.loadouts (
        title,
        source,
        source_url,
        weapon_name,
        weapon_category,
        created_by,
        game_id,
        attachments
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;