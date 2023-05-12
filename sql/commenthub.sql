DROP TABLE IF EXISTS COMMENT;
CREATE TABLE COMMENT(
    gh_owner TEXT NOT NULL,
    gh_repo TEXT NOT NULL,
    gh_issuenr TEXT NOT NULL,
    comment_text TEXT NOT NULL
);

INSERT INTO COMMENT(gh_owner, gh_repo, gh_issuenr, comment_text) values ('crossplane', 'crossplane', '2', 'my private comment');
