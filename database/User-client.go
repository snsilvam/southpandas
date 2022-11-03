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

/* ----- Repository Design Pattern/ Patron de diseño repositorio ----- */
type PostgresRepositoryUserClient struct {
	/* Este atributo nos permite construir y liberar conexiones con la base de datos.
	   This attribute allows us to build and release connections to the database. */
	db *sql.DB
}

// Esta funcion se encargara de construir el objeto o abrir la conexion con base de datos.
func NewPostgresRepositoryUserClient(url string) (*PostgresRepositoryUserClient, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepositoryUserClient{db}, nil
}

// Esta funcion se encargara de cerrar la conexion con la base de datos.
func (repo *PostgresRepositoryUserClient) Close() {
	repo.db.Close()
}

func (repo *PostgresRepositoryUserClient) InsertUserClient(ctx context.Context, userClient *models.UserClient) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO userClient (id, premium, user_id) values ($1, $2, $3)", userClient.ID, userClient.Premium, userClient.User_ID)
	return err
}

func (repo *PostgresRepositoryUserClient) ListUsersClients(ctx context.Context) ([]*models.UserClient, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, premium, user_id, created_at FROM userClient")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usersClients := []*models.UserClient{}
	for rows.Next() {
		userClient := &models.UserClient{}
		if err := rows.Scan(&userClient.ID, &userClient.Premium, &userClient.User_ID, &userClient.CreatedAt); err != nil {
			return nil, err
		}
		usersClients = append(usersClients, userClient)
	}
	return usersClients, nil
}
