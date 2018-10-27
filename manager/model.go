package manager

import "github.com/go-pg/pg"

// Policy ...
type Policy struct {
}

var upMigrations = []string{
	`CREATE TABLE IF NOT EXISTS ladon_policy (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		description STRING NOT NULL,
		effect
			STRING
			NOT NULL
			CHECK (effect = 'allow' OR effect = 'deny'),
		conditions bytea NOT NULL
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_subject (
		compiled STRING NOT NULL,
		template VARCHAR(1023) NOT NULL,
		policy VARCHAR(255) NOT NULL,
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_permission (
		compiled STRING NOT NULL,
		template VARCHAR(1023) NOT NULL,
		policy VARCHAR(255) NOT NULL,
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_resource (
		compiled STRING NOT NULL,
		template VARCHAR(1023) NOT NULL,
		policy VARCHAR(255) NOT NULL,
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_resource (
		compiled STRING NOT NULL,
		template VARCHAR(1023) NOT NULL,
		policy VARCHAR(255) NOT NULL,
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_resource (
		compiled STRING NOT NULL,
		template VARCHAR(1023) NOT NULL,
		policy VARCHAR(255) NOT NULL,
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_resource (
		compiled STRING NOT NULL,
		template VARCHAR(1023) NOT NULL,
		policy VARCHAR(255) NOT NULL,
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_resource (
		compiled STRING NOT NULL,
		template VARCHAR(1023) NOT NULL,
		policy VARCHAR(255) NOT NULL,
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_subject (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		has_regex BOOL NOT NULL,
		compiled VARCHAR(511) NOT NULL UNIQUE,
		template VARCHAR(511) NOT NULL UNIQUE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_action (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		has_regex BOOL NOT NULL,
		compiled VARCHAR(511) NOT NULL UNIQUE,
		template VARCHAR(511) NOT NULL UNIQUE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_resource (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		has_regex BOOL NOT NULL,
		compiled VARCHAR(511) NOT NULL UNIQUE,
		template VARCHAR(511) NOT NULL UNIQUE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_subject_rel (
		policy VARCHAR(255) NOT NULL,
		subject VARCHAR(64) NOT NULL,
		PRIMARY KEY (policy, subject),
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE,
		FOREIGN KEY (subject) REFERENCES ladon_subject (id) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_action_rel (
		policy VARCHAR(255) NOT NULL,
		action VARCHAR(64) NOT NULL,
		PRIMARY KEY (policy, action),
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE,
		FOREIGN KEY (action) REFERENCES ladon_action (id) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS ladon_policy_resource_rel (
		policy VARCHAR(255) NOT NULL,
		resource VARCHAR(64) NOT NULL,
		PRIMARY KEY (policy, resource),
		FOREIGN KEY (policy) REFERENCES ladon_policy (id) ON DELETE CASCADE,
		FOREIGN KEY (resource) REFERENCES ladon_resource (id) ON DELETE CASCADE
	);`,
	`CREATE INDEX ladon_subject_compiled_idx ON ladon_subject (
		compiled
	);`,
	`CREATE INDEX ladon_permission_compiled_idx ON ladon_action (
		compiled
	);`,
	`CREATE INDEX ladon_resource_compiled_idx ON ladon_resource (
		compiled
	);`,
}

// MigrateUp ...
func MigrateUp(db *pg.DB) error {
	// Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// Rollback On Error, call commit if success
	defer tx.Rollback()
	for _, query := range upMigrations {
		// Execute Query in Transaction
		_, err = tx.Exec(query)
		if err != nil {
			return err
		}
	}
	// Commit Transaction
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
