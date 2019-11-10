package route

import (
    "github.com/gorilla/mux"
    trello "github.com/oms-services/trello/trello"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "GetCards",
        "POST",
        "/getcards",
        trello.GetCards,
    },
    Route{
        "GetBoard",
        "POST",
        "/getboard",
        trello.GetBoard,
    },
    Route{
        "GetLists",
        "POST",
        "/getlists",
        trello.GetLists,
    },
    Route{
        "AddCard",
        "POST",
        "/addcard",
        trello.AddCard,
    },
    Route{
        "MoveCard",
        "POST",
        "/movecard",
        trello.MoveCard,
    },
    Route{
        "SubscribeCard",
        "POST",
        "/subscribe",
        trello.SubscribeCard,
    },
    Route{
        "CopyCard",
        "POST",
        "/copycard",
        trello.CopyCard,
    },
    Route{
        "CreateBoard",
        "POST",
        "/createboard",
        trello.CreateBoard,
    },
    Route{
        "DeleteBoard",
        "POST",
        "/deleteboard",
        trello.DeleteBoard,
    },
    Route{
        "CreateList",
        "POST",
        "/createlist",
        trello.CreateList,
    },
    Route{
        "GetAllBoards",
        "POST",
        "/getallboards",
        trello.GetAllBoards,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
