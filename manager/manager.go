package manager

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/blake2b"

	"github.com/go-pg/pg"
	jsoniter "github.com/json-iterator/go"
	"github.com/ory/ladon"
	"github.com/ory/ladon/compiler"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

// Create a new pollicy to Manager.
func (c *CRDBManager) Create(policy ladon.Policy) error {
	conditions := []byte("{}")
	if policy.GetConditions() != nil {
		cs := policy.GetConditions()
		conditionsJSON, err := json.Marshal(&cs)
		if err != nil {
			return err
		}
		conditions = conditionsJSON
	}
	// Begin Transaction
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	// Rollback on error before commit
	defer tx.Rollback()
	// Insert Policy
	_, err = tx.Exec("INSERT INTO ladon_policy (id, description, effect, conditions) SELECT ?, ?, ?, ? WHERE NOT EXISTS (SELECT 1 FROM ladon_policy WHERE id = ?)", policy.GetID(), policy.GetDescription(), policy.GetEffect(), conditions, policy.GetID())
	if err != nil {
		return err
	}
	// Process Relations
	type relation struct {
		p []string
		t string
	}
	var relations = []relation{{p: policy.GetActions(), t: "action"}, {p: policy.GetResources(), t: "resource"}, {p: policy.GetSubjects(), t: "subject"}}
	for _, v := range relations {
		for _, template := range v.p {
			// Hash Template
			h := blake2b.Sum256([]byte(template))
			id := fmt.Sprintf("%x", h)
			// Compile Regex
			compiled, err := compiler.CompileRegex(template, policy.GetStartDelimiter(), policy.GetEndDelimiter())
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
			// Insert Relations
			query := fmt.Sprintf("INSERT INTO ladon_%s (id, template, compiled, has_regex) SELECT ?, ?, ?, ? WHERE NOT EXISTS (SELECT 1 FROM ladon_%[1]s WHERE id = ?)", v.t)
			_, err = tx.Exec(query, id, template, compiled.String(), strings.Index(template, string(policy.GetStartDelimiter())) > -1, id)
			if err != nil {
				return err
			}
			query = fmt.Sprintf("INSERT INTO ladon_policy_%s_rel (policy, %[1]s) SELECT ?, ? WHERE NOT EXISTS (SELECT 1 FROM ladon_policy_%[1]s_rel WHERE policy = ? AND %[1]s = ?)", v.t)
			if _, err := tx.Exec(query, policy.GetID(), id, policy.GetID(), id); err != nil {

				return err
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// Update updates an existing policy.
func (c *CRDBManager) Update(policy ladon.Policy) error {

	return nil
}

// Get retrieves a policy.
func (c *CRDBManager) Get(id string) (ladon.Policy, error) {

	return nil, nil
}

// GetAll returns all policies.
func (c *CRDBManager) GetAll(limit, offset int64) (ladon.Policies, error) {

	return nil, nil
}

// Delete removes a policy.
func (c *CRDBManager) Delete(id string) error {

	return nil
}

// FindRequestCandidates returns candidates that could match the request object. It either returns
// a set that exactly matches the request, or a superset of it. If an error occurs, it returns nil and
// the error.
func (c *CRDBManager) FindRequestCandidates(r *ladon.Request) (ladon.Policies, error) {
	return nil, nil
}

// FindPoliciesForResource returns policies that could match the resource. It either returns
// a set of policies that apply to the resource, or a superset of it.
// If an error occurs, it returns nil and the error.
func (c *CRDBManager) FindPoliciesForResource(resource string) (ladon.Policies, error) {
	return nil, nil
}

// FindPoliciesForSubject returns policies that could match the subject. It either returns
// a set of policies that applies to the subject, or a superset of it.
// If an error occurs, it returns nil and the error.
func (c *CRDBManager) FindPoliciesForSubject(subject string) (ladon.Policies, error) {
	return nil, nil
}
