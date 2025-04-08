create table public."user" (
    id bigserial primary key,
    "login" text not null,
    email text not null,
    hash_password text not null,
    registration_date timestamp not null
);

CREATE UNIQUE INDEX user_login_idx ON public."user" (login);
CREATE UNIQUE INDEX user_email_idx ON public."user" (email);