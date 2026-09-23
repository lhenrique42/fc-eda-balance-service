# eda-balances-fc

Event-driven wallet system with two services sharing one Kafka broker:

- **wallet-core** — HTTP API for clients/accounts/transactions. Writes to
  MySQL. On each transaction, publishes a `BalanceUpdated` event to the
  Kafka `balances` topic.
- **balance-service** — consumes `BalanceUpdated` from Kafka and keeps a
  read-only balance per account in Postgres. Exposes one endpoint to query
  it.

```
wallet-core --(BalanceUpdated event)--> Kafka "balances" topic --> balance-service --> Postgres
    |                                                                      |
   MySQL                                                            GET /balances/{id}
```

## Running everything

From the repo root:

```bash
docker compose up -d --build
```

This starts both services plus MySQL, Postgres, Zookeeper, Kafka, and the
Kafka control-center UI (`http://localhost:9021`).

Both databases reseed with fixed fake data on every container start (see
`wallet-core/mysql/init/` and `balance-service/postgres/init/`). Data does
not persist across restarts by design.

## Services

| Service         | Port | Datastore       | Requests                          |
| --------------- | ---- | --------------- | --------------------------------- |
| wallet-core     | 8080 | MySQL (3306)    | `wallet-core/api/client.http`     |
| balance-service | 3003 | Postgres (5433) | `balance-service/api/client.http` |

## wallet-core endpoints

```
POST /clients
POST /accounts
POST /transactions
```

## balance-service endpoint

```
GET /balances/{accountId}
```

Returns `{ "account_id": "...", "balance": 0 }`, or `404` if the account
has no row yet.

Seed data ships with three fake accounts:

```
aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa  balance 1600.50
bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb  balance 270.25
cccccccc-cccc-cccc-cccc-cccccccccccc  balance 1180.75
```

## Testing the full event flow

This proves a transaction in `wallet-core` updates the balance in
`balance-service` via Kafka, not via a direct call.

1. Create a transaction in `wallet-core`:
   ```bash
   curl http://localhost:8080/transactions \
     -X POST -H "Content-Type: application/json" \
     -d '{"account_id_from":"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa","account_id_to":"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb","amount":100}'
   ```
2. Query both accounts in `balance-service` and confirm the balances moved:
   ```bash
   curl http://localhost:3003/balances/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa
   curl http://localhost:3003/balances/bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb
   ```

If the numbers don't move, check `docker compose logs balance-service` —
the consumer prints `error handling message: ...` on any failure (bad JSON,
DB error, etc).

## Direct DB access (debugging)

```bash
docker compose exec mysql mysql -uroot -proot wallet -e "SELECT * FROM accounts;"
docker compose exec balance-db psql -U postgres -d balance -c "SELECT * FROM balances;"
```
