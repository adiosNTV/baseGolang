/* name: GetAllLabelsPaginated :many */
SELECT name, image
FROM labels
LIMIT ?
OFFSET ?;

/* name: GetLabelTotalCounts :one */
SELECT COUNT(*)
FROM labels;