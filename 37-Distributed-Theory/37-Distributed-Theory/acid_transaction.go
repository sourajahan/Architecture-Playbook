func TransferFunds(ctx context.Context, db *sql.DB, from, to string, amount float64) error {
    tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
    if err != nil { return err }
    defer tx.Rollback()

    // Atomicity & Consistency: Ensuring both debits and credits succeed
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, from)
    if err != nil { return err }

    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, to)
    if err != nil { return err }

    return tx.Commit() // Durability: Persistence guaranteed here
}
