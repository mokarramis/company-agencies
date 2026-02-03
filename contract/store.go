package contract

type StorageInterface interface {
	StoreIntoFile(data []byte)
	ReadFromFile() []string
}
