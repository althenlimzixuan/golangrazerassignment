package main

import (
	"context"
	"fmt"
	"hello/ent"
	"hello/ent/comprimiseemail"
	"log"
)

func ReadAllEmail(ctx context.Context, client *ent.Client) []*ent.ComprimiseEmail {

	record, _ := client.ComprimiseEmail.Query().All(ctx)
	return record
}

func AddCompromisedEmail(ctx context.Context, client *ent.Client, email string) (*ent.ComprimiseEmail, error) {
	newEmail, err := client.ComprimiseEmail.
		Create().
		SetEmail(email).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed adding new compromised email: %w", err)
	}
	log.Println("new email was added: ", newEmail)
	return newEmail, nil
}

func ReadEmailById(ctx context.Context, client *ent.Client, i int) (*ent.ComprimiseEmail, error) {
	record, err := client.ComprimiseEmail.Get(ctx, i)

	if record != nil {
		return record, nil
	}

	return nil, fmt.Errorf("failed reading CompromisedEmail: %w", err)
}

func IsEmailExist(ctx context.Context, client *ent.Client, email string) bool {

	record, err := client.ComprimiseEmail.Query().Select(comprimiseemail.FieldEmail).Where(comprimiseemail.Email(email)).First(ctx)

	if err != nil && record != nil {
		fmt.Println(err)
	}

	return record != nil
}

func IsEmailCompromised(connectionstring string, email string) bool {

	db, err := ent.Open("mysql", "root:!Razer1234.@tcp(mysql:3306)/test")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
		return false
	}

	defer db.Close()

	fmt.Println("Connect Success!")

	ctx := context.Background()

	return IsEmailExist(ctx, db, email)
}
