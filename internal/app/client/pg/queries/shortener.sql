{{define "PublishEvent"}}
SELECT events.publish($1, $2, $3, $4);
{{
end}}

{{define "CreateURL"}}
INSERT INTO shortener.urls (url, user_id) VALUES ($1, $2) ON CONFLICT (url) DO NOTHING RETURNING id, url;
{{end}}

{{define "GetURL"}}
SELECT url, is_deleted FROM shortener.urls WHERE id = $1;
{{end}}

{{define "GetURLID"}}
SELECT id FROM shortener.urls WHERE url = $1;
{{end}}

{{define "GetUserURLs"}}
SELECT id, url FROM shortener.urls WHERE user_id = $1;
{{end}}

{{define "DeleteURLs"}}
UPDATE shortener.urls
SET is_deleted = true
WHERE id IN ({{.IDs}})
    AND user_id = $1
{{end}}