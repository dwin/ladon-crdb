package manager

import (
	"github.com/go-pg/pg"
	"github.com/ory/ladon"
)

// CRDBManager ...
type CRDBManager struct {
	db *pg.DB
}

// NewCRDBManager ...
func NewCRDBManager(db *pg.DB) *CRDBManager {
	c := CRDBManager{
		db: db,
	}

	return &c
}

// Create ...
func (c *CRDBManager) Create(policy ladon.Policy) error {

	return nil
}

// Update ...
func (c *CRDBManager) Update(policy ladon.Policy) error {

	return nil
}

// Get ..
func (c *CRDBManager) Get(id string) (ladon.Policy, error) {

	return nil, nil
}

// Delete ...
func (c *CRDBManager) Delete(id string) error {

	return nil
}

// FindRequestCandidates ...
func (c *CRDBManager) FindRequestCandidates(r *ladon.Request) (ladon.Policies, error) {
	return nil, nil
}

// FindPoliciesForResource ...
func (c *CRDBManager) FindPoliciesForResource(resource string) (ladon.Policies, error) {
	return nil, nil
}

// FindPoliciesForSubject ...
func (c *CRDBManager) FindPoliciesForSubject(subject string) (ladon.Policies, error) {
	return nil, nil
}
