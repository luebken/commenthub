# CommentHub

## Idea
A service which allows to comment on a public issue which is only visible to a certain set of users.

## Milestones
- [ ] JSON Rest service which allows to post & get a comment per owner/repo/issuenr
- [ ] Some web UI
- [ ] Authenticate with Github OAuth
- [ ] Allow users from the same org read shared comments
- [ ] Chrome-plugin which includes CommentHub into GitHub

## Run

Local:
```sh
export DB_PATH_FILE=<some path to some file>
sqlite3 $DB_PATH_FILE < sql/commenthub.sql

go run cmd/main/main.go

curl localhost:8080/comments/crossplane/crossplane/issues/2

curl localhost:8080/comments/ \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"owner": "crossplane", "repo: "crossplane", "issuenr" :"2", "comment_text": "A second very private comment"}'
```