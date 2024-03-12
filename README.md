# DOCUMENTATION FOR TODO REST API

## Technologies

<ul>
<li>Go</li>
<li>Postgres SQL</li>
<li>migrate(SQL Migration)</li>
</ul>


### Database Migration Set Up

1. Installing Migrate CLI

> go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

2. Create Sample up and down migration files

>  migrate create -seq -ext=.sql -dir=./migrations create_movies_table

3. Execute Migration Up files

> migrate -path=./migrations -database=$TODO_DB_DSN up

#Links

[The-Ultimate-Markdown-Cheat-Sheet](https://github.com/lifeparticle/Markdown-Cheatsheet)