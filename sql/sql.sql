CREATE TABLE lista (
    id int,
    nome text,
    completa int,
    dt_created text   
)

CREATE TABLE item (
    id int,
    lista int,
    item text,
    valor real,
    qtd int,
    dt_insert text
)