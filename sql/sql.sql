CREATE TABLE lista (
    id integer PRIMARY KEY AUTOINCREMENT UNIQUE,
    nome text,
    completa int,
    dt_created text   
);

CREATE TABLE item (
    id integer PRIMARY KEY AUTOINCREMENT UNIQUE,
    lista int,
    item text,
    valor real,
    qtd int,
    dt_insert text,
    status int
);