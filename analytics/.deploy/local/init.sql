create schema if not exists inbox;

create table inbox."event"(
    id serial primary key,
    "key" text not null,
    topic text not null,
    payload jsonb not null,
    idempotency_key text not null,
    created_at timestamptz not null default now(),
    locked_until timestamptz,
    attempts int not null default 0,
    error text,
    processed_at timestamptz
);

create table public.user_course_progress(
    user_id bigint not null,
    course_id bigint not null,
    completion_percentage numeric not null default 0,
    last_updated timestamptz not null default now(),
    completed_lessons int not null default 0,
    completed_quizzes int not null default 0,
    completed_tasks int not null default 0,
    primary key (user_id, course_id)
);

CREATE TABLE public.course_rating (
    user_id bigint not null,
    course_id bigint not null,
    rating numeric not null default 0,
    position int not null,
    last_updated timestamptz not null default now(),
    primary key (user_id, course_id)
);

CREATE TABLE analytics.global_rating (
    user_id bigint primary key,
    rating numeric not null default 0,
    position int not null,
    last_updated timestamptz not null default now()
);