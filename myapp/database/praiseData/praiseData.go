package praiseData

import (
	"log"

	"myapp/database/userData"

	"github.com/jmoiron/sqlx"
)

type Praise struct {
	Id             int           `db:"id"`
	TargetUser     userData.User `db:"target_user"`
	RegisteredUser userData.User `db:"registered_user"`
	Content        string        `db:"content"`
	HasApproved    bool          `db:"has_approved"`
}

func AddPraise(content string, userId int) {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO gogo.praises (content, target_user_id, registered_user_id, has_approved) VALUES ($1, $2, 1, $3)", content, userId, false)
	tx.Commit()
}

func GetAllPraises(getApprovedOnly bool) []Praise {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	praises := []Praise{}
	if getApprovedOnly == true {
		db.Select(&praises, "SELECT p.id AS id, tu.id AS \"target_user.id\", tu.user_name AS \"target_user.user_name\", tu.invalid_flg AS \"target_user.invalid_flg\", ru.id AS \"registered_user.id\", ru.user_name AS \"registered_user.user_name\", ru.invalid_flg AS \"registered_user.invalid_flg\", p.content AS content, p.has_approved AS has_approved FROM gogo.praises p LEFT OUTER JOIN gogo.users tu ON p.target_user_id=tu.id LEFT OUTER JOIN gogo.users ru ON p.registered_user_id = ru.id WHERE p.has_approved = false ORDER BY p.id DESC")
	} else {
		db.Select(&praises, "SELECT p.id AS id, tu.id AS \"target_user.id\", tu.user_name AS \"target_user.user_name\", tu.invalid_flg AS \"target_user.invalid_flg\", ru.id AS \"registered_user.id\", ru.user_name AS \"registered_user.user_name\", ru.invalid_flg AS \"registered_user.invalid_flg\", p.content AS content, p.has_approved AS has_approved FROM gogo.praises p LEFT OUTER JOIN gogo.users tu ON p.target_user_id=tu.id LEFT OUTER JOIN gogo.users ru ON p.registered_user_id = ru.id ORDER BY p.id DESC")
	}

	log.Println(praises)
	return praises
}

func GetPraiseRandomly(target_user_id int) Praise {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	praise := Praise{}
	err = db.Get(&praise, "SELECT p.id AS id, tu.id AS \"target_user.id\", tu.user_name AS \"target_user.user_name\", tu.invalid_flg AS \"target_user.invalid_flg\", ru.id AS \"registered_user.id\", ru.user_name AS \"registered_user.user_name\", ru.invalid_flg AS \"registered_user.invalid_flg\", p.content AS content, p.has_approved AS has_approved FROM gogo.praises p LEFT OUTER JOIN gogo.users tu ON p.target_user_id=tu.id LEFT OUTER JOIN gogo.users ru ON p.registered_user_id = ru.id WHERE p.has_approved=true AND p.target_user_id=$1 ORDER BY random() LIMIT 1", target_user_id)
	return praise
}

func ApprovePraises(id int) {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec("UPDATE gogo.praises SET has_approved=true WHERE id=$1", id)

	tx.Commit()
}
