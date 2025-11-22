package sql_storage

import (
    "database/sql"
    "time"

    "github.com/go-next-pizza/internal/app/storage"
)

type SMSCodeRepository struct {
    storage *SQLStorage
}

func (r *SMSCodeRepository) SaveCode(phone string, codeHash string, expiresAt int64) error {
    _, err := r.storage.db.Exec(
        "INSERT INTO sms_codes (phone, code, expires_at) VALUES ($1, $2, to_timestamp($3))",
        phone, codeHash, expiresAt,
    )
    return err
}

func (r *SMSCodeRepository) GetLatestCodeHash(phone string) (string, int64, error) {
    var codeHash string
    var expiresAt time.Time
    err := r.storage.db.QueryRow(
        "SELECT code, expires_at FROM sms_codes WHERE phone = $1 ORDER BY id DESC LIMIT 1",
        phone,
    ).Scan(&codeHash, &expiresAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", 0, storage.ErrRecordNotFound
        }
        return "", 0, err
    }
    return codeHash, expiresAt.Unix(), nil
}

func (r *SMSCodeRepository) DeleteCodes(phone string) error {
    _, err := r.storage.db.Exec("DELETE FROM sms_codes WHERE phone = $1", phone)
    return err
}


