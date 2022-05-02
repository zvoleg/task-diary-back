package user

const (
	createScript string = `INSERT INTO administration.users
							(user_id, account_id, name, description, role_id, created_at, is_deleted)
							VALUES($1, $2, $3, $4, $5, $6, $7)
							RETURNING *`
	getScript     string = `SELECT * FROM administration.users u WHERE u.user_id = $1`
	getListScript string = `SELECT * FROM administration.users`
	updateScript  string = `UPDATE administration.users
							SET name = $2,
								description = $3,
								role_id = $4,
								updated_at = $5
							WHERE user_id = $1
							RETURNING *`
	deleteScript string = `UPDATE administration.users
						   SET is_deleted = true
						   	   updated_at = $2
						   WHERE user_id = $1`
)
