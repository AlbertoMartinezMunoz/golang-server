package catalogue

import "errors"

type CatalogueStore interface {
	List() (map[string]Album, error)
	Add(id string, album Album) error
	Get(id string) (Album, error)
}

var (
	NotFoundErr   = errors.New("not found")
	IsRepeatedErr = errors.New("album name is in the catalogue")
)

type MemStore struct {
	catalogue map[string]Album
}

func NewMemStore() *MemStore {
	catalogue := map[string]Album{
		"blue-train":                       {Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		"jeru":                             {Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		"sarah-vaughan-and-clifford-brown": {Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	return &MemStore{
		catalogue,
	}
}

func (m MemStore) Add(id string, album Album) error {
	if _, ok := m.catalogue[id]; ok == true {
		return IsRepeatedErr
	}
	m.catalogue[id] = album
	return nil
}

func (m MemStore) Get(id string) (Album, error) {

	if val, ok := m.catalogue[id]; ok {
		return val, nil
	}

	return Album{}, NotFoundErr
}

func (m MemStore) List() (map[string]Album, error) {
	return m.catalogue, nil
}

func (m MemStore) Update(id string, album Album) error {

	if _, ok := m.catalogue[id]; ok {
		m.catalogue[id] = album
		return nil
	}

	return NotFoundErr
}

func (m MemStore) Remove(id string) error {
	delete(m.catalogue, id)
	return nil
}
