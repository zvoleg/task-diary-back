package task

const (
	getScript    string = `SELECT * FROM storage.tasks t WHERE t.task_id = $1`
	createScript string = `INSERT INTO storage.task 
						  (task_id, board_id, title, description, type, status, 
						   author_id, created_at, updated_at, tags, is_deleted)
						   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
						   RETURNING *`
	updateScript string = `UPDATE storage.task 
						   SET title = $2,
						   	   description = $3,
						   	   type = $4,
							   status = $5,
							   updated_at = $6,
							   tags = $7
						   WHERE task_id = $1
						   RETURNING *`
	deleteScript string = `UPDATE storage.task
						   SET is_deleted = true
						   	   updated_at = $2
						   WHERE task_id = $1`
)
