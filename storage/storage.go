package storage

type Storage interface {
	Open(options map[string]string) error
	Get(key string, val interface{}) error
	Put(key string, val interface{}) error
	Close() error
}

func NewStorage(dbtype string) (Storage, error) {
	if dbtype == "dynamo" {
		db := &DynamoDbStorage{}
		return db, nil
	} else {
		return nil, nil
	}
}
