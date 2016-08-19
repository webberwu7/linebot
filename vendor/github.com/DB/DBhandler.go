/*
使用者是否在遊戲中
他手上有哪兩張牌


*/
package DB
import(
	"os"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
/*傳入使用者MID
回傳使用者是否正在遊戲*/
func UserGamming(MID string) bool{
	var GameID int
	GameID = 0
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	db.QueryRow("SELECT GameID FROM LineBot.GameAction WHERE MID = ? and Cancel = 0", MID ).Scan(&GameID)
	db.close()
	if GameID == 0{
		return false
	}else{
		return true
	}
}
func UserCards(MID string) (int ,int){
	var card1 int
	var card2 int
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	db.QueryRow("SELECT PlayerCards FROM LineBot.GameAction WHERE MID = ? and Cancel = 0", MID ).Scan(&GameID)
	db.close()
	//strings.Split(",a,b,c,d,e,", ",")

}