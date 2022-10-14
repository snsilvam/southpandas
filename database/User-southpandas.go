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
type PostgresRepositoryUserSouthPandas struct {
	/* Este atributo nos permite construir y liberar conexiones con la base de datos.
	   This attribute allows us to build and release connections to the database. */
	db *sql.DB
}

// Esta funcion se encargara de construir el objeto o abrir la conexion con base de datos.
func NewPostgresRepositoryUserSouthPandas(url string) (*PostgresRepositoryUserSouthPandas, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepositoryUserSouthPandas{db}, nil
}

// Esta funcion se encargara de cerrar la conexion con la base de datos.
func (repo *PostgresRepositoryUserSouthPandas) Close() {
	repo.db.Close()
}

func (repo *PostgresRepositoryUserSouthPandas) InsertUserSouthPandas(ctx context.Context, userSouthPandas *models.UserSouthPandas) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO usersouthpandas (id, type, user_id) values ($1, $2, $3)", userSouthPandas.ID, userSouthPandas.Type, userSouthPandas.User_ID)
	return err
}

func (repo *PostgresRepositoryUserSouthPandas) ListUsersSouthPandas(ctx context.Context) ([]*models.UserSouthPandas, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, type, user_id, created_at FROM usersouthpandas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usersSouthPandas := []*models.UserSouthPandas{}
	for rows.Next() {
		userSouthPandas := &models.UserSouthPandas{}
		if err := rows.Scan(&userSouthPandas.ID, &userSouthPandas.Type, &userSouthPandas.User_ID, &userSouthPandas.CreatedAt); err != nil {
			return nil, err
		}
		usersSouthPandas = append(usersSouthPandas, userSouthPandas)
	}
	return usersSouthPandas, nil
}
