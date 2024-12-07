package local

import (
	"github.com/codingsince1985/checksum"
	"github.com/dgraph-io/badger/v4"
	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/afero"
)

func (v *VFS) WriteFile(path string, data []byte) error {
	file, err := v.OpenOrCreate(path)
	defer func(file afero.File) {
		err := file.Close()
		if err != nil {
			v.Error(err)
		}
	}(file)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	go func() {
		mimetype.DetectFile(path)
		checksum.MD5sum(path)
		err = v.DB.Update(func(txn *badger.Txn) error {
			e := badger.NewEntry([]byte("answer"), []byte("42"))
			return txn.SetEntry(e)
		})
	}()
	err = file.Sync()
	if err != nil {
		return err
	}
	return file.Close()
}

func (v *VFS) OpenOrCreate(path string) (afero.File, error) {
	_, err := v.Fs.Stat(path)
	if err != nil {
		created, err := v.Fs.Create(path)
		if err != nil {
			return nil, err
		}
		return created, nil
	}
	open, err := v.Fs.Open(path)
	if err != nil {
		return nil, err
	}
	return open, nil
}
