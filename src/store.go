package bot

import (
	"context"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func NewStore(ctx context.Context) (*sqlstore.Container, error) {
	dbLog := waLog.Stdout("Database", "WARN", true)
	return sqlstore.New(ctx, "sqlite3", "file:whatsapp.db?_foreign_keys=on", dbLog)
}
