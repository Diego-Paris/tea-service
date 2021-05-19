package config

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
SetupDB Sets up connection to database.

Establishes a default configuration for the mgm library,
to then be used throughout the project. It then
makes a call to pingDB() which uses the default mongo
driver to connect a client to be used for a ping, if
no errors occurred then the ping was successful
and we can then use mgm library.
*/
func SetupDB() error {

	var err error

	// Client Options
	clientOptions := options.Client().ApplyURI(MongoCredentials)

	// Setup mgm default config
	err = mgm.SetDefaultConfig(nil, "Cluster0", clientOptions)
	if err != nil {
		return err
	}

	// Ping database
	err = pingDB(clientOptions)
	if err != nil {
		return err
	}

	return nil
}

/*
Creates a new mongo client to ping database.
If no error is returned, then the ping was successful.
*/
func pingDB(clientOptions *options.ClientOptions) error {

	var err error

	// Establish a client and context to ping and test db connection
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // avoids memory leaks

	err = client.Connect(ctx)
	defer client.Disconnect(ctx)

	if err != nil {
		return err
	}

	// Finally ping database and check if any
	// errors occurred establishing a connection
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return err
	}

	return nil
}
