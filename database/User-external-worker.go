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
type PostgresRepositoryUserExternalWorker struct {
	/* Este atributo nos permite construir y liberar conexiones con la base de datos.
	   This attribute allows us to build and release connections to the database. */
	db *sql.DB
}

// Esta funcion se encargara de construir el objeto o abrir la conexion con base de datos.
func NewPostgresRepositoryUserExternalWorker(url string) (*PostgresRepositoryUserExternalWorker, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepositoryUserExternalWorker{db}, nil
}

// Esta funcion se encargara de cerrar la conexion con la base de datos.
func (repo *PostgresRepositoryUserExternalWorker) Close() {
	repo.db.Close()
}

func (repo *PostgresRepositoryUserExternalWorker) InsertUserExternalWorker(ctx context.Context, userExternalWorker *models.UserExternalWorker) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO userexternalworker (id, contracttype, workexperience, workremote, willingnesstravel, currentsalary, expectedsalary, possibilityofrotation, profilelinkedln, workarea, descriptionworkarea, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", userExternalWorker.ID, userExternalWorker.ContractType, userExternalWorker.WorkExperience, userExternalWorker.WorkRemote, userExternalWorker.Willingnesstravel, userExternalWorker.CurrentSalary, userExternalWorker.ExpectedSalary, userExternalWorker.PossibilityOfRotation, userExternalWorker.Profilelinkedln, userExternalWorker.Workarea, userExternalWorker.DescriptionWorkArea, userExternalWorker.User_id)
	return err
}

func (repo *PostgresRepositoryUserExternalWorker) ListUsersExternalWorkers(ctx context.Context) ([]*models.UserExternalWorker, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, contracttype, workexperience, workremote, willingnesstravel, currentsalary, expectedsalary, possibilityofrotation, profilelinkedln, workarea, descriptionworkarea, user_id, created_at FROM userexternalworker")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usersExternalWorkers := []*models.UserExternalWorker{}
	for rows.Next() {
		userExternalWorker := &models.UserExternalWorker{}
		if err := rows.Scan(userExternalWorker.ID, userExternalWorker.ContractType, userExternalWorker.WorkExperience, userExternalWorker.WorkRemote, userExternalWorker.Willingnesstravel, userExternalWorker.CurrentSalary, userExternalWorker.ExpectedSalary, userExternalWorker.PossibilityOfRotation, userExternalWorker.Profilelinkedln, userExternalWorker.Workarea, userExternalWorker.DescriptionWorkArea, userExternalWorker.User_id, &userExternalWorker.CreatedAt); err != nil {
			return nil, err
		}
		usersExternalWorkers = append(usersExternalWorkers, userExternalWorker)
	}
	return usersExternalWorkers, nil
}
