-- goose Up

CREATE TABLE events (
    id serial primary key,
    user_id INT,
    owner text,
    title text,
    description text
--     start_date date not null,
--     start_time time,
--     end_date date not null,
--     end_time time
);


INSERT INTO events (id, user_id, owner, title,  description)
VALUES
    ('1234', 12, 'this owner', 'this title 1', 'desc 1'),
    ('123', 12, 'this owner2', 'this title 2', 'desc 2'),
    ('1235', 12, 'this owner3', 'this title 3', 'desc 3'),
    ('123', 12, 'this owner4', 'this title 4', 'desc 4');


-- +goose Down
drop table events;