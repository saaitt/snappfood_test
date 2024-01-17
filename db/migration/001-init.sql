CREATE TABLE orders
(
    id            SERIAL      not null,
    delivery_time bigint      not null,
    created_at    timestamptz not null,
    agent_id      bigint      NULL,
    vendor_id     bigint      not null,
    PRIMARY KEY (id)
);

CREATE TABLE order_delay_reports
(
    id         SERIAL      not null,
    created_at timestamptz not null,
    order_id   bigint      not null,
    processed  bool        not null default false,
    PRIMARY KEY (id)
);

CREATE TABLE trips
(
    id         SERIAL      not null,
    created_at timestamptz not null,
    order_id   bigint      not null,
    processed  bool        not null default false,
    PRIMARY KEY (id)
)


