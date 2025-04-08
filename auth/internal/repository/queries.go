package repository

const (
	queryInsertUser = `
	insert into public."user" (login, email, hash_password, registration_date)
	values ($1, $2, $3, $4)
	returning id;
	`

	queryGetUserByLoginOrEmail = `
	select
		id, login, email, hash_password, registration_date
	from 
		public."user"
	where login = $1 or email = $1;
	`

	queryCheckUserExists = `
	select exists(
    	select 1 
    	from public."user" 
    	where login = $1 or email = $2
	) as user_exists;
	`

	queryUpdateUser = `
	update 
		public."user"
	set
		login = $2,
		email = $3,
		hash_password = $4
	where
		id = $1;
	`
)
