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
type PostgresRepositoryWorkerOfClient struct {
	/* Este atributo nos permite construir y liberar conexiones con la base de datos.
	   This attribute allows us to build and release connections to the database. */
	db *sql.DB
}

// Esta funcion se encargara de construir el objeto o abrir la conexion con base de datos.
func NewPostgresRepositoryWorkerOfClient(url string) (*PostgresRepositoryWorkerOfClient, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepositoryWorkerOfClient{db}, nil
}

// Esta funcion se encargara de cerrar la conexion con la base de datos.
func (repo *PostgresRepositoryWorkerOfClient) Close() {
	repo.db.Close()
}

func (repo *PostgresRepositoryWorkerOfClient) InsertWorkerOfClient(ctx context.Context, workerOfClient *models.WorkerOfClient) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO workersofclient (id, description, userclient_id) values ($1, $2, $3)", workerOfClient.ID, workerOfClient.Description, workerOfClient.UserClient_ID)
	return err
}

func (repo *PostgresRepositoryWorkerOfClient) ListWorkersOfClient(ctx context.Context) ([]*models.WorkerOfClient, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, description, userclient_id, created_at FROM workersofclient")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	workers := []*models.WorkerOfClient{}
	for rows.Next() {
		worker := &models.WorkerOfClient{}
		if err := rows.Scan(&worker.ID, &worker.Description, &worker.UserClient_ID, &worker.CreatedAt); err != nil {
			return nil, err
		}
		workers = append(workers, worker)
	}
	return workers, nil
}
