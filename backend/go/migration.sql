create table if not exists users
(
    id        char(36)     not null primary key,
    firstName varchar(255) not null,
    lastName  varchar(255) not null,
    email     varchar(255) not null,
    createdAt datetime,
    updatedAt datetime
);

create trigger if not exists users_updateTrigger
    after update
    on users
begin
    update users
    set updatedAt = datetime('now')
    where id = new.id;
end;

create trigger if not exists users_createTrigger
    after insert
    on users
begin
    update users
    set createdAt = datetime('now'),
        updatedAt = datetime('now')
    where id = new.id;
end;

create trigger if not exists users_deleteTrigger
    after delete
    on users
begin
    delete
    from todoAssignees
    where userId = old.id;
end;

create table if not exists todoTypes
(
    id           char(36)     not null primary key,
    name         varchar(255) not null,
    description  text         not null,
    color        varchar(100) not null,
    reminderTime int          not null,
    createdAt    datetime,
    updatedAt    datetime
);

create trigger if not exists todoTypes_updateTrigger
    after update
    on todoTypes
begin
    update todoTypes
    set updatedAt = datetime('now')
    where id = new.id;
end;

create trigger if not exists todoTypes_createTrigger
    after insert
    on todoTypes
begin
    update todoTypes
    set createdAt = datetime('now'),
        updatedAt = datetime('now')
    where id = new.id;
end;

create table if not exists todos
(
    id          char(36)     not null primary key,
    todoTypeId  char(36)     not null,
    title       varchar(255) not null,
    description text         not null,
    dueDate     datetime     not null,
    completed   integer(1)   not null default 0,
    createdAt   datetime,
    updatedAt   datetime,
    foreign key (todoTypeId) references todoTypes (id)
);

create trigger if not exists todos_updateTrigger
    after update
    on todos
begin
    update todos
    set updatedAt = datetime('now')
    where id = new.id;
end;

create trigger if not exists todos_createTrigger
    after insert
    on todos
begin
    update todos
    set createdAt = datetime('now'),
        updatedAt = datetime('now')
    where id = new.id;
end;

create trigger if not exists todos_deleteTrigger
    after delete
    on todos
begin
    delete
    from todoAssignees
    where todoId = old.id;
end;

create table if not exists todoAssignees
(
    todoId char(36) not null,
    userId char(36) not null,
    primary key (todoId, userId),
    foreign key (todoId) references todos (id),
    foreign key (userId) references users (id)
);
