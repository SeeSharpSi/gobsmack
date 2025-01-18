package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	// "seesharpsi/gobsmack/assets"
	"seesharpsi/gobsmack/gamestate"
	"seesharpsi/gobsmack/templ"
)

type Games struct {
	// Takes a gamekey
	Progs map[string]*gamestate.Game

	// Takes a useragent
	Players map[string]*gamestate.Game
}

var games = Games{}

func main() {
	games.Progs = make(map[string]*gamestate.Game)
	games.Players = make(map[string]*gamestate.Game)

	port := flag.Int("port", 9779, "port the server runs on")
	address := flag.String("address", "http://localhost", "address the server runs on")
	flag.Parse()

	// ip parsing
	base_ip := *address
	ip := base_ip + ":" + strconv.Itoa(*port)
	root_ip, err := url.Parse(ip)
	if err != nil {
		log.Panic(err)
	}

	mux := http.NewServeMux()
	add_routes(mux)

	server := http.Server{
		Addr:    root_ip.Host,
		Handler: mux,
	}

	// start server
	log.Printf("running server on %s\n", root_ip.Host)
	err = server.ListenAndServe()
	defer server.Close()
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func add_routes(mux *http.ServeMux) {
	mux.HandleFunc("/", GetIndex)
	mux.HandleFunc("/static/{file}", ServeStatic)
	mux.HandleFunc("/test", GetTest)
	mux.HandleFunc("/game", SpawnGame)
	mux.HandleFunc("/loop", LoopGames)
	mux.HandleFunc("/action/{type}/{with}", Action)
	mux.HandleFunc("/map", Map)
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")
	log.Printf("got /static/%s request\n", file)
	http.ServeFile(w, r, "./static/"+file)
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("got / request\n")
	component := templ.Index()
	component.Render(context.Background(), w)
}

func GetTest(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /test request\n")
	component := templ.Test()
	component.Render(context.Background(), w)
}

func SpawnGame(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /game request\n")
	component := templ.Test()
	component.Render(context.Background(), w)
	gamekey := keygen(5)
	for _, ok := games.Progs[gamekey]; ok; _, ok = games.Progs[gamekey] {
		fmt.Printf("\nkey %s already exists, creating new key...\n", gamekey)
		gamekey = keygen(5)
	}
	fmt.Println("created new key ", gamekey)
	g := gamestate.Game{}
	g.Ship.NewShip()
	g.GameKey = gamekey
	games.Progs[gamekey] = &g
	games.Players[r.RemoteAddr] = &g
}

func LoopGames(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /loop request\n")
	component := templ.Test()
	component.Render(context.Background(), w)
	for _, v := range games.Progs {
		fmt.Printf("\n%+v\n", v)
		v.StartGame()
	}
	fmt.Printf("\n%+v\n", games)
}

func Action(w http.ResponseWriter, r *http.Request) {
	atype := r.PathValue("type")
	donewith := r.PathValue("with")
	intdonewith, err := strconv.Atoi(donewith)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	player := r.RemoteAddr
	rgame := games.Players[player]
	rgame.QueueAction(player, atype, intdonewith)
}

func Map(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /map request\n")
	player := r.RemoteAddr
	rgame := games.Players[player]
	component := templ.Minimap(rgame.Ship)
	component.Render(context.Background(), w)
}

func keygen(n int) string {
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func printContextInternals(ctx interface{}, inner bool) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				printContextInternals(reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
}
