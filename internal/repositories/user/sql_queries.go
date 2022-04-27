package user

const (
	createScript string = `INSERT INTO administration.users
							(user_id, account_id, name, description, role_id, created_at, is_deleted)
							VALUES($1, $2, $3, $4, $5, $6, $7)
							RETURNING *`
	getScript     string = `SELECT * FROM administration.users u WHERE u.user_id = $1`
	getListScript string = `SELECT * FROM administration.users`
	updateScript  string = `UPDATE administration.users
							SET account_id = $2,
								name = $3,
								description = $4,
								role_id = $5,
								created_at = $6,
								is_deleted = $7
							WHERE user_id = $1
							RETURNING *`
	deleteScript string = `UPDATE administration.users
						   SET is_deleted = true
						   WHERE user_id = $1`
)
