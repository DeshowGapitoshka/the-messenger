package lists

import (
	"database/sql"
	"log"
	"os"

	"time"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Group struct {
	Id          int           `json:"id"`
	Owner_id    string        `json:"owner_id"`
	Name        string        `json:"name"`
	Info        string        `json:"info"`
	Users_id    pq.Int64Array `json:"users_id"`
	Admins_id   pq.Int64Array `json:"admins_id"`
	Enemies_id  pq.Int64Array `json:"enemies_id"`
	Created_at  *string       `json:"created_at"`
	LastMessage *string       `json:"last_message"`
	AvatarLink  *string       `json:"avatar_link"`
	UpdatedAt   *time.Time    `json:"updated_at"`
	UnreadCount *int          `json:"unread_count"`
}

func init() {
	var err error
	err = godotenv.Load()
	errorHandler(err)
	connToPG := os.Getenv("DB_CONNECT_TO_BASEDATA")
	if connToPG == "" {
		log.Fatal("connToPG is empty")
	}

	database, err = sql.Open("postgres", connToPG)
	errorHandler(err)
}

func AddNewGroup(owner_id int, name string, info string, users_id []int, admins_id []int) int {
	dbQuery := os.Getenv("DB_GROUP_CREATE")
	if dbQuery == "" {
		log.Println("DB_GROUP_CREATE is empty")
		return 500
	}
	_, err := database.Exec(dbQuery, owner_id, name, info, users_id, admins_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	return 200
}

func GetUserGroupsById(userId int) (int, []Group) {

	dbQuery := os.Getenv("DB_GET_USER_GROUPS")
	if dbQuery == "" {
		log.Println("DB_GET_USER_GROUPS is empty")
		return 500, nil
	}

	rows, err := database.Query(dbQuery, userId)
	if err != nil {
		log.Printf("Error getting groups: %v", err)
		return 500, nil
	}
	defer rows.Close()

	var groups []Group

	for rows.Next() {
		var group Group

		err := rows.Scan(
			&group.Id,
			&group.Owner_id,
			&group.Name,
			&group.Info,
			&group.Users_id,
			&group.Admins_id,
			&group.Enemies_id,
			&group.Created_at,
			&group.LastMessage,
			&group.AvatarLink,
			&group.UpdatedAt,
			&group.UnreadCount,
		)

		if err != nil {
			log.Printf("Error scanning chat: %v", err)
			continue
		}

		groups = append(groups, group)
	}

	return 200, groups
}
