{{define "PublishEvent"}}
SELECT events.publish($1, $2, $3, $4);
{{
end}}
