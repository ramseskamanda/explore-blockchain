package blockchain

import "github.com/dgraph-io/badger"

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (iterator *BlockChainIterator) Next() *Block {
	var block *Block

	err := iterator.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iterator.CurrentHash)
		HandleError(err)

		err = item.Value(func(val []byte) error {
			block = Deserialize(val)
			return nil
		})
		HandleError(err)
		return err
	})

	HandleError(err)

	iterator.CurrentHash = block.PrevHash

	return block
}
