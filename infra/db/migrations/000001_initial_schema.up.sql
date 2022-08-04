create table users
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    age        bigint,
    city       text,
    country    text,
    email      text,
    first_name text not null,
    last_name  text not null,
    username   text not null,
    pronouns   text,
    active     boolean default true,
    password   text not null,
    role       text
);

alter table users
    owner to postgres;

create index idx_users_deleted_at
    on users (deleted_at);

create table skill_categories
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text not null,
    user_id    bigint
        constraint fk_skill_categories_user
            references users
);

alter table skill_categories
    owner to postgres;

create index idx_skill_categories_deleted_at
    on skill_categories (deleted_at);

create table skills
(
    id                bigserial
        primary key,
    created_at        timestamp with time zone,
    updated_at        timestamp with time zone,
    deleted_at        timestamp with time zone,
    name              text     not null,
    proficiency       smallint not null,
    user_id           bigint   not null
        constraint fk_skills_user
            references users
        constraint fk_users_skills
            references users,
    show              boolean,
    skill_category_id bigint   not null
        constraint fk_skills_skill_category
            references skill_categories
);

alter table skills
    owner to postgres;

create index idx_skills_deleted_at
    on skills (deleted_at);

create table carousel_items
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    position   bigint not null,
    content    text   not null,
    user_id    bigint not null
        constraint fk_carousel_items_user
            references users
);

alter table carousel_items
    owner to postgres;

create index idx_carousel_items_deleted_at
    on carousel_items (deleted_at);

create table user_websites
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    url        text   not null,
    type       text   not null,
    user_id    bigint not null
        constraint fk_user_websites_user
            references users
);

alter table user_websites
    owner to postgres;

create index idx_user_websites_deleted_at
    on user_websites (deleted_at);

create table articles
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id    bigint not null
        constraint fk_articles_user
            references users,
    title      text   not null,
    content    text   not null
);

alter table articles
    owner to postgres;

create index idx_articles_deleted_at
    on articles (deleted_at);

create table project_categories
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id    bigint not null
        constraint fk_project_categories_user
            references users,
    name       text   not null
);

alter table project_categories
    owner to postgres;

create index idx_project_categories_deleted_at
    on project_categories (deleted_at);

create table projects
(
    id                  bigserial
        primary key,
    created_at          timestamp with time zone,
    updated_at          timestamp with time zone,
    deleted_at          timestamp with time zone,
    user_id             bigint not null
        constraint fk_projects_user
            references users,
    project_category_id bigint not null
        constraint fk_projects_project_category
            references project_categories,
    title               text   not null,
    url                 text   not null,
    image_url           text
);

alter table projects
    owner to postgres;

create index idx_projects_deleted_at
    on projects (deleted_at);

create table work_experiences
(
    id           bigserial
        primary key,
    created_at   timestamp with time zone,
    updated_at   timestamp with time zone,
    deleted_at   timestamp with time zone,
    user_id      bigint                   not null
        constraint fk_work_experiences_user
            references users,
    role         text                     not null,
    company_name text                     not null,
    start_date   timestamp with time zone not null,
    description  text,
    end_date     timestamp with time zone
);

alter table work_experiences
    owner to postgres;

create index idx_work_experiences_deleted_at
    on work_experiences (deleted_at);
