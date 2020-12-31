create schema if not exists stocker;

create table if not exists customer(
    id bigint generated always as identity primary key,
    code uuid constraint uk_customer_code not null,
    name varchar not null,
	created_at timestamp not null,
	updated_at timestamp not null
);

create table if not exists stock(
    id bigint generated always as identity primary key,
    ticker varchar(20) constraint uk_stock_ticker unique not null,
	price numeric(22, 2) not null,
	created_at timestamp not null,
	updated_at timestamp not null
);

create index if not exists idx_stock_by_ticker on stock(ticker);

create table if not exists stock_history(
    id bigint generated always as identity primary key,
    stock_id bigint not null,
	price numeric(22, 2) not null,
	created_at timestamp not null,

    constraint fk_stock_history_stock_id
		foreign key(stock_id)
		references stock(id)
);

create index if not exists idx_stock_history_by_stock_id
    on stock_history(stock_id);

create table if not exists stock_order(
    id bigint generated always as identity primary key,
    code uuid constraint uk_stock_order_code unique not null,
    stock_id bigint not null,
    customer_id bigint not null,
	amount int not null,
    status varchar(20) not null,
	created_at timestamp not null,
	updated_at timestamp not null,

    constraint fk_stock_order_stock_id
		foreign key(stock_id)
		references stock(id),
    constraint fk_stock_order_customer_id
		foreign key(customer_id)
		references customer(id)
);

create index if not exists idx_order_by_code on stock_order(code);
