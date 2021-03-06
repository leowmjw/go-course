package store

import (
	"sync"

	puppy "github.com/anz-bank/go-course/08_project/runnerdave/pkg/puppy"
)

// SyncStore implementation of the Storer interface
type SyncStore struct {
	sync.Mutex
	sync.Map
}

// NewSyncStore creates a new in-memory store with map intialised
func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy saves new puppy if not in store, if it is already returns error
func (s *SyncStore) CreatePuppy(p *puppy.Puppy) error {
	if err := puppy.ValidateValue(p.Value); err != nil {
		return err
	}
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(p.ID); ok {
		return puppy.Errorf(puppy.ErrUnknown, "puppy with id %d already exists", p.ID)
	}
	s.Store(p.ID, *p)
	return nil
}

// ReadPuppy gets a puppy from the store given an ID
func (s *SyncStore) ReadPuppy(id uint16) (puppy.Puppy, error) {
	if puppyData, ok := s.Load(id); ok {
		puppy, _ := puppyData.(puppy.Puppy)
		return puppy, nil
	}
	return puppy.Puppy{}, puppy.Errorf(puppy.ErrIDNotFound, "puppy with ID:%d not found", id)
}

// UpdatePuppy puts new puppy data to the store, error if id does not exist
func (s *SyncStore) UpdatePuppy(id uint16, p *puppy.Puppy) error {
	if err := puppy.ValidateValue(p.Value); err != nil {
		return err
	}
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return puppy.Errorf(puppy.ErrIDNotFound, "puppy with ID:%d not found", id)
	}
	s.Store(id, *p)
	return nil
}

// DeletePuppy deletes a puppy from the store
func (s *SyncStore) DeletePuppy(id uint16) error {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return puppy.Errorf(puppy.ErrIDNotFound, "puppy with ID:%d not found", id)
	}
	s.Delete(id)
	return nil
}
