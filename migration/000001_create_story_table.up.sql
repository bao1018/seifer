create table user_story(
    id serial primary key,
    title text not null Default '',
    body text not null Default '',
    title_vector int [],
    vector int
)