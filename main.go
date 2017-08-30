package main;

import (
    "urawesome/ur"
    // "github.com/gorilla/schema"
    "fmt"
    "io"
    "math/rand"
    "net/http"
    "os"
    "path"
    "strconv"
    "sync"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    ruleset := ur.NewFinkelRuleset(&ur.UrBoard)
    game := ur.NewGame(&ur.UrBoard, &ruleset)
    var game_mutex sync.Mutex

    throwDicesHandler := func(w http.ResponseWriter, r *http.Request){
        fmt.Println("THROW DICES")
        game_mutex.Lock()

        if game.Phase == ur.ThrowDices {
            game.ThrowDices()
        } else {
            //TODO
        }
        fmt.Fprintf(w, ur.WebView(&game))

        game_mutex.Unlock()
    }

    chooseMoveHandler := func(w http.ResponseWriter, r *http.Request){
        fmt.Println("CHOOSE MOVE")
        game_mutex.Lock()

        if game.Phase == ur.ChooseMove {
            idx, err := strconv.Atoi(path.Base(r.URL.Path))
            if err != nil {
                panic(err) // TODO
            }
            game.ChooseMove(uint(idx))
        } else {
            //TODO
        }
        fmt.Fprintf(w, ur.WebView(&game))

        game_mutex.Unlock()
    }
    viewHandler := func(w http.ResponseWriter, r *http.Request){
        fmt.Println("VIEW")
        game_mutex.Lock()

        f, _ := os.Open("views/header.html")
        io.Copy(w, f)
        fmt.Fprintf(w, ur.WebView(&game))
        f, _ = os.Open("views/footer.html")
        io.Copy(w, f)

        game_mutex.Unlock()
    }

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/throw-dices", throwDicesHandler)
    http.HandleFunc("/choose-move/", chooseMoveHandler)
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8080", nil)
}

/*

*/
