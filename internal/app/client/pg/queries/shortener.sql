{{define "PublishEvent"}}
SELECT events.publish($1, $2, $3, $4);
{{
end}}

{{define "CreateURL"}}
INSERT INTO shortener.urls (url) VALUES ($1) ON CONFLICT (url) DO NOTHING RETURNING id, url;
{{end}}

{{define "GetURL"}}
SELECT url FROM shortener.urls WHERE id = $1;
{{end}}

{{define "GetURLID"}}
SELECT id FROM shortener.urls WHERE url = $1;
{{end}}
