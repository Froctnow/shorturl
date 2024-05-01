{{define "PublishEvent"}}
SELECT events.publish($1, $2, $3, $4);
{{
end}}

{{define "CreateURL"}}
INSERT INTO shortener.urls (url) VALUES ($1) RETURNING id;
{{end}}

{{define "GetURL"}}
SELECT url FROM shortener.urls WHERE id = $1;
{{end}}
