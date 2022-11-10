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
type PostgresRepositoryUser struct {
	/* Este atributo nos permite construir y liberar conexiones con la base de datos.
	   This attribute allows us to build and release connections to the database. */
	db *sql.DB
}

// Esta funcion se encargara de construir el objeto o abrir la conexion con base de datos.
func NewPostgresRepositoryUser(url string) (*PostgresRepositoryUser, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepositoryUser{db}, nil
}

// Esta funcion se encargara de cerrar la conexion con la base de datos.
func (repo *PostgresRepositoryUser) Close() {
	repo.db.Close()
}

func (repo *PostgresRepositoryUser) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, name, email, address) values ($1, $2, $3, $4)", user.ID, user.Name, user.Email, user.Address)
	return err
}

func (repo *PostgresRepositoryUser) ListUsers(ctx context.Context) ([]*models.User, error) {
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
