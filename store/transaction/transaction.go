package transaction

import (
	"github.com/jinzhu/gorm"
	"github.com/maichain/eth-indexer/model"
)

const (
	TableName = "transactions"
)

type Store interface {
	Insert(data *model.Transaction) error
	FindTransaction(hash []byte) (result *model.Transaction, err error)
	FindTransactionsByBlockHash(blockHash []byte) (result []*model.Transaction, err error)
}

type store struct {
	db *gorm.DB
}

func NewWithDB(db *gorm.DB) Store {
	return &store{
		db: db.Table(TableName),
	}
}

func (t *store) Insert(data *model.Transaction) error {
	return t.db.Create(data).Error
}

func (t *store) FindTransaction(hash []byte) (result *model.Transaction, err error) {
	result = &model.Transaction{}
	err = t.db.Where("BINARY hash = ?", hash).Limit(1).Find(result).Error
	return
}

func (t *store) FindTransactionsByBlockHash(blockHash []byte) (result []*model.Transaction, err error) {
	err = t.db.Where("BINARY block_hash = ?", blockHash).Find(&result).Error
	return
}
