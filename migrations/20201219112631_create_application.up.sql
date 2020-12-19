create table applications(
    id int NOT NULL AUTO_INCREMENT,
    slug varchar(200) not null,
    body json null,
    status varchar (100) not null default('in_process'),

    PRIMARY KEY (ID)
)

