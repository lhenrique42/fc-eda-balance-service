package database

import (
	"balance-service/internal/entity"
	"database/sql"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{DB: db}
}

func (b *BalanceDB) Get(accountID string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	stmt, err := b.DB.Prepare("SELECT account_id, balance, updated_at FROM balances WHERE account_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(accountID)
	if err := row.Scan(&balance.AccountID, &balance.Amount, &balance.UpdatedAt); err != nil {
		return nil, err
	}
	return balance, nil
}

func (b *BalanceDB) Upsert(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare(`INSERT INTO balances (account_id, balance, updated_at)
	VALUES ($1, $2, $3)
	ON CONFLICT (account_id) DO UPDATE SET balance = $2, updated_at = $3`)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(balance.AccountID, balance.Amount, balance.UpdatedAt)
	return err
}
