package common

import (
	"fmt"
	"log"

	"github.com/keybase/go-keychain"
)

func ConfigVaultKeySave(key string) {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(ConfigIdentifier)
	item.SetAccount("Strawhouse Key v1")
	_ = keychain.DeleteItem(item)

	item = keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(ConfigIdentifier)
	item.SetLabel("Strawhouse")
	item.SetAccount("Strawhouse Key v1")
	item.SetData([]byte(key))
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	err := keychain.AddItem(item)
	if err != nil {
		log.Fatalf("Error adding item to Keychain: %v", err)
	} else {
		fmt.Println("Key stored in Keychain successfully.")
	}
}

func ConfigVaultKeyGet() string {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(ConfigIdentifier)
	query.SetAccount("Strawhouse Key v1")
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)

	results, err := keychain.QueryItem(query)
	if err != nil {
		log.Fatalf("Error querying Keychain: %v", err)
	}

	if len(results) == 0 {
		log.Fatalf("Key not found in Keychain.")
	}

	return string(results[0].Data)
}
