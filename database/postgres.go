package database

import (
	//El paquete contexto nos permite definir variables tipo contexto, que controlaran las solicitudes entrantes a un servidor.
	//The context package allows us to define context variables that will control incoming requests to a server.
	"context"
	/*El paquete database/sql entrega una interfaz a nuestra aplicacion para interactuar con bases de datos sql.
	The database/sql package provides an interface to our application to interact with sql databases.*/
	"database/sql"
	/*Es el controlador de base de datos para postgresql, que nos ayuda con las interacciones del paquete database/sql
	It is the database driver for postgresql, which helps us with the interactions of the database/sql package.*/
	_ "github.com/lib/pq"
	"southpandas.com/go/cqrs/models"
)

//In order to connect the postgres data base, use the PostgresRepository
type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() {
	repo.db.Close()
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, name, email, address) values ($1, $2, $3, $4)", user.ID, user.Email, user.Address, user.Name)
	return err
}

func (repo *PostgresRepository) ListUsers(ctx context.Context) ([]*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, email, address, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

/* func (repo *PostgresRepository) InsertUserClient(ctx context.Context, userClient *models.UserClient) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO usersClient (id, premium, user_id) values ($1, $2, $3)", userClient.ID, userClient.Premium, userClient.User_ID)
	return err
}

func (repo *PostgresRepository) ListUsersClient(ctx context.Context) ([]*models.UserClient, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, premium, user_id, created_at FROM usersClient")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usersclient := []*models.UserClient{}
	for rows.Next() {
		userClient := &models.UserClient{}
		if err := rows.Scan(&userClient.ID, &userClient.Premium, &userClient.User_ID, &userClient.CreatedAt); err != nil {
			return nil, err
		}
		usersclient = append(usersclient, userClient)
	}
	return usersclient, nil
}
*/
