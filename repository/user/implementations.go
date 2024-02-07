package user_repository

import "context"

func (r *UserRepository) Save(ctx context.Context, input User) (output *User, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO test name FROM test WHERE id = $1", input.Id).Scan(&output.FullName)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, input GetUserByIdInput) (output *User, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.FullName)
	if err != nil {
		return nil, err
	}
	return output, nil
}
