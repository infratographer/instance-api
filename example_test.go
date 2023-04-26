package main

import (
	"context"
	"fmt"
	"log"

	"go.infratographer.com/instance-api/internal/ent"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func Example_Todo() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	provider, err := client.InstanceProvider.Create().
		SetName("Test Provider").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a provider: %v", err)
	}

	instance, err := client.Instance.Create().
		SetName("Test Instance").
		SetInstanceProvider(provider).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a provider: %v", err)
	}

	p, _ := instance.QueryInstanceProvider().First(ctx)
	if provider.Name != p.Name {
		fmt.Println("something is fishy")
	}

	fmt.Printf("%+v\n", provider)
	fmt.Printf("%+v\n", instance)
	fmt.Printf("%+v\n", instance.Edges.InstanceProvider)
	fmt.Printf("%+v\n", p)
	// Output:
	// InstanceProvider(id=1)

}
