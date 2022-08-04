INSERT INTO public.users
(created_at, updated_at, deleted_at, age, city, country, email, first_name, last_name, username, pronouns, active,
 password, role)
VALUES ('2022-08-03 16:38:41.992000 +00:00', '2022-08-03 16:38:44.034000 +00:00', null, 25, 'Pato Branco',
        'Brasil', 'gabsmanfroi@gmail.com', 'Gabriel', 'Manfroi', 'gabemanfroi', 'He/Him', true, 'm1nfroy900',
        'Fullstack Developer');

INSERT INTO public.user_websites (created_at, updated_at, deleted_at, url, type, user_id)
VALUES ('2022-08-03 16:42:59.184000 +00:00', '2022-08-03 16:43:01.985000 +00:00', null,
        'https://linkedin.com/in/gabemanfroi', 'linkedin', 1),
       ('2022-08-03 16:42:59.184000 +00:00', '2022-08-03 16:43:01.985000 +00:00', null,
        'https://github.com/gabemanfroi', 'github', 1),
       ('2022-08-03 16:42:59.184000 +00:00', '2022-08-03 16:43:01.985000 +00:00', null,
        'https://twitter.com/gabemanfroi', 'twitter', 1);

INSERT INTO public.carousel_items (created_at, updated_at, deleted_at, position, content, user_id)
VALUES ('2022-08-03 16:48:03.350000 +00:00', '2022-08-03 16:48:05.787000 +00:00', null, 0, 'mobile apps', 1);

INSERT INTO public.skill_categories(created_at, updated_at, deleted_at, name, user_id)
VALUES ('2022-08-03 16:49:06.215000 +00:00', '2022-08-03 16:49:07.712000 +00:00', null, 'Languages', null),
       ('2022-08-03 16:49:06.215000 +00:00', '2022-08-03 16:49:07.712000 +00:00', null,
        'Programming Languages', null),
       ('2022-08-03 16:49:06.215000 +00:00', '2022-08-03 16:49:07.712000 +00:00', null, 'Other', null);

INSERT INTO public.skills (created_at, updated_at, deleted_at, name, proficiency, user_id, show, skill_category_id)
VALUES ('2022-08-03 16:52:47.665000 +00:00', '2022-08-03 16:52:48.949000 +00:00', null, 'Python', 90, 1, true,
        2),
       ('2022-08-03 16:52:47.665000 +00:00', '2022-08-03 16:52:48.949000 +00:00', null, 'Javascript', 90, 1,
        true, 2),
       ('2022-08-03 16:52:47.665000 +00:00', '2022-08-03 16:52:48.949000 +00:00', null, 'React', 90, 1, true,
        2),
       ('2022-08-03 16:52:47.665000 +00:00', '2022-08-03 16:52:48.949000 +00:00', null, 'Typescript', 90, 1,
        true, 2),
       ('2022-08-03 16:52:47.665000 +00:00', '2022-08-03 16:52:48.949000 +00:00', null, 'Go', 80, 1, true, 2),
       ('2022-08-03 16:52:47.665000 +00:00', '2022-08-03 16:52:48.949000 +00:00', null, 'Node', 85, 1, true,
        2);

INSERT INTO public.project_categories (created_at, updated_at, deleted_at, user_id, name)
VALUES ('2022-08-03 16:57:12.236000 +00:00', '2022-08-03 16:57:12.988000 +00:00', null, 1, 'Web'),
       ('2022-08-03 16:57:12.236000 +00:00', '2022-08-03 16:57:12.988000 +00:00', null, 1, 'Mobile'),
       ('2022-08-03 16:57:12.236000 +00:00', '2022-08-03 16:57:12.988000 +00:00', null, 1, 'API');

INSERT INTO public.projects (created_at, updated_at, deleted_at, user_id, project_category_id, title, url,
                             image_url)
VALUES ('2022-08-03 16:59:59.596000 +00:00', '2022-08-03 17:00:00.609000 +00:00', null, 1, 1,
        'Instituional Website', 'https://google.com', null);

INSERT INTO public.work_experiences (created_at, updated_at, deleted_at, user_id, role, company_name, start_date,
                                     description, end_date)
VALUES ('2022-08-03 17:23:21.022000 +00:00', '2022-08-03 17:23:22.668000 +00:00', null, 1,
        'FullStack Web Developer', 'TBDC AgroSoftware', '2022-08-03 17:23:44.597000 +00:00', null,
        '2022-08-03 17:23:48.422000 +00:00');