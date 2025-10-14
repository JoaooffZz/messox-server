package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	crudCreate "crud/create"
	crudDel "crud/delete"
	crudRead "crud/read"
	crudUp "crud/update"
	test "test/db" // ajuste o import conforme seu projeto
)

// ──────────────────────────────
// Utils
// ──────────────────────────────
func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n👉 ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// ──────────────────────────────
// Main Menu
// ──────────────────────────────
func main() {
	time.Sleep(5 * time.Second)

	config := test.DatabaseConfig{}
	newConfig := config.New()
	db, err := test.GetConn(newConfig)
	if err != nil {
		fmt.Printf("❌ error connecting to the database: %v\n", err)
		return
	}
	defer db.Close()

	mock := test.MockDB{}

	for {
		clearScreen()
		fmt.Println("═══════════════════════════")
		fmt.Println("        🔷 TESTS MENU      ")
		fmt.Println("═══════════════════════════")
		fmt.Println("1. TestUpdateDB")
		fmt.Println("2. TestReadDB")
		fmt.Println("3. TestCreateDB")
		fmt.Println("4. TestDeleteDB")
		fmt.Println("0. Exit")
		fmt.Println("═══════════════════════════")

		choice := prompt()

		switch choice {
		case "1":
			runTestUpdateDB(db, mock)
		case "2":
			runTestReadDB(db, mock)
		case "3":
			runTestCreateDB(db, mock)
		case "4":
			runTestDeleteDB(db, mock)
		case "0":
			fmt.Println("👋 bye!")
			return
		default:
			fmt.Println("❌ invalid option")
			time.Sleep(1 * time.Second)
		}
	}
}

// ──────────────────────────────
// TestUpdateDB Panel
// ──────────────────────────────
func runTestUpdateDB(db *sql.DB, mock test.MockDB) {
	updateDB := test.TestUpdateDB{DB: db, Users: mock.GetNewUsers()}
	update := crudUp.Update{DB: db}

	for {
		clearScreen()
		fmt.Println("🔷 TestUpdateDB")
		fmt.Println("1. FlowRun")
		fmt.Println("2. UnitUpUserProfile")
		fmt.Println("3. UnitUpUserBio")
		fmt.Println("4. UnitUpHistoryChat")
		fmt.Println("0. Go Back")

		choice := prompt()
		var err error

		switch choice {
		case "1":
			err = updateDB.FlowRun()
		case "2":
			err = updateDB.UnitUpUserProfile(update)
		case "3":
			err = updateDB.UnitUpUserBio(update)
		case "4":
			err = updateDB.UnitUpHistoryChat(mock, update)
		case "0":
			return
		default:
			fmt.Println("❌ invalid option")
			time.Sleep(1 * time.Second)
			continue
		}

		printResult(err)
	}
}

// ──────────────────────────────
// TestReadDB Panel
// ──────────────────────────────
func runTestReadDB(db *sql.DB, mock test.MockDB) {
	readDB := test.TestReadDB{DB: db, Users: mock.GetNewUsers()}
	read := crudRead.Read{DB: db}

	for {
		clearScreen()
		fmt.Println("🔷 TestReadDB")
		fmt.Println("1. FlowRun")
		fmt.Println("2. UnitGetTotalUsers")
		fmt.Println("3. UnitGetUsers")
		fmt.Println("4. UnitGetInboxMessages")
		fmt.Println("5. UnitGetSentRequests")
		fmt.Println("6. UnitGetReceivedRequests")
		fmt.Println("0. Go Back")

		choice := prompt()
		var err error

		switch choice {
		case "1":
			err = readDB.FlowRun()
		case "2":
			err = readDB.UnitGetTotalUsers(read)
		case "3":
			err = readDB.UnitGetUsers(read)
		case "4":
			err = readDB.UnitGetInboxMessages(read)
		case "5":
			err = readDB.UnitGetSentRequests(read)
		case "6":
			err = readDB.UnitGetReceivedRequests(read)
		case "0":
			return
		default:
			fmt.Println("❌ invalid option")
			time.Sleep(1 * time.Second)
			continue
		}

		printResult(err)
	}
}

// ──────────────────────────────
// TestCreateDB Panel
// ──────────────────────────────
func runTestCreateDB(db *sql.DB, mock test.MockDB) {
	createDB := test.TestCreateDB{DB: db, Users: mock.GetNewUsers()}
	create := crudCreate.Create{DB: db}

	for {
		clearScreen()
		fmt.Println("🔷 TestCreateDB")
		fmt.Println("1. FlowCreate")
		fmt.Println("2. UnitCreateUser")
		fmt.Println("3. UnitCreateContacts")
		fmt.Println("4. UnitCreateHistoryChats")
		fmt.Println("5. UnitCreateInboxMessages")
		fmt.Println("6. UnitCreateInboxRequests")
		fmt.Println("0. Go Back")

		choice := prompt()
		var err error

		switch choice {
		case "1":
			err = createDB.FlowCreate()
		case "2":
			err = createDB.UnitCreateUser(create)
		case "3":
			err = createDB.UnitCreateContacts(create)
		case "4":
			err = createDB.UnitCreateHistoryChats(mock, create)
		case "5":
			err = createDB.UnitCreateInboxMessages(mock, create)
		case "6":
			err = createDB.UnitCreateInboxRequests(create)
		case "0":
			return
		default:
			fmt.Println("❌ invalid option")
			time.Sleep(1 * time.Second)
			continue
		}

		printResult(err)
	}
}

// ──────────────────────────────
// TestDeleteDB Panel
// ──────────────────────────────
func runTestDeleteDB(db *sql.DB, mock test.MockDB) {
	deleteDB := test.TestDeleteDB{DB: db, Users: mock.GetNewUsers()}
	del := crudDel.Delete{DB: db}

	for {
		clearScreen()
		fmt.Println("🔷 TestDeleteDB")
		fmt.Println("1. FlowRun")
		fmt.Println("2. UnitDeleteContacts")
		fmt.Println("3. UnitDeleteInboxMessages")
		fmt.Println("4. UnitDeleteInboxRequests")
		fmt.Println("0. Go Back")

		choice := prompt()
		var err error

		switch choice {
		case "1":
			err = deleteDB.FlowRun()
		case "2":
			err = deleteDB.UnitDeleteContacts(del)
		case "3":
			err = deleteDB.UnitDeleteInboxMessages(del)
		case "4":
			err = deleteDB.UnitDeleteInboxRequests(del)
		case "0":
			return
		default:
			fmt.Println("❌ invalid option")
			time.Sleep(1 * time.Second)
			continue
		}

		printResult(err)
	}
}

// ──────────────────────────────
// Helper for success/error
// ──────────────────────────────
func printResult(err error) {
	if err != nil {
		fmt.Printf("❌ error: %v\n", err)
	} else {
		fmt.Println("✅ executed successfully")
	}
	fmt.Println("\nPress ENTER to continue...")
	fmt.Scanln()
}


