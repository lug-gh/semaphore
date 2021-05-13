package sql

import "github.com/ansible-semaphore/semaphore/db"

func (d *SqlDb) GetEnvironment(projectID int, environmentID int) (db.Environment, error) {
	var environment db.Environment
	err := d.getObject(projectID, db.EnvironmentProps, environmentID, &environment)
	return environment, err
}

func (d *SqlDb) GetEnvironments(projectID int, params db.RetrieveQueryParams) ([]db.Environment, error) {
	var environment []db.Environment
	err := d.getObjects(projectID, db.EnvironmentProps, params, &environment)
	return environment, err
}

func (d *SqlDb) UpdateEnvironment(env db.Environment) error {
	_, err := d.sql.Exec(
		"update project__environment set name=?, json=? where id=?",
		env.Name,
		env.JSON,
		env.ID)
	return err
}

func (d *SqlDb) CreateEnvironment(env db.Environment) (newEnv db.Environment, err error) {
	res, err := d.sql.Exec(
		"insert into project__environment (project_id, name, json, password) values (?, ?, ?, ?)",
		env.ProjectID,
		env.Name,
		env.JSON,
		env.Password)

	if err != nil {
		return
	}

	insertID, err := res.LastInsertId()

	if err != nil {
		return
	}

	newEnv = env
	newEnv.ID = int(insertID)
	return
}

func (d *SqlDb) DeleteEnvironment(projectID int, environmentID int) error {
	return d.deleteObject(projectID, db.EnvironmentProps, environmentID)
}

func (d *SqlDb) DeleteEnvironmentSoft(projectID int, environmentID int) error {
	return d.deleteObjectSoft(projectID, db.EnvironmentProps, environmentID)
}
