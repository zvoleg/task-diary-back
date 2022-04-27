package role

const (
	getScript     string = `SELECT * FROM reference.roles r WHERE r.role_id = $1`
	getListScript string = `SELECT * FROM reference.roles`
)
