package trello

import (
	"context"
	"encoding/json"
	"fmt"
	trello "github.com/adlio/trello"
	"github.com/cloudevents/sdk-go"
	result "github.com/oms-services/trello/result"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Payload struct {
	EventId     string     `json:"eventID"`
	EventType   string     `json:"eventType"`
	ContentType string     `json:"contentType"`
	Data        TrelloArgs `json:"data"`
}

type Subscribe struct {
	Data      TrelloArgs `json:"data"`
	Endpoint  string     `json:"endpoint"`
	ID        string     `json:"id"`
	IsTesting bool       `json:"istesting"`
}

//TrelloArgs struct
type TrelloArgs struct {
	ListName    string `json:"listName,omitempty"`
	BoardName   string `json:"boardName,omitempty"`
	BoardID     string `json:"boardId,omitempty"`
	ListID      string `json:"listId,omitempty"`
	CardID      string `json:"cardId,omitempty"`
	CardName    string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Existing    bool   `json:"existing,omitempty"`
	Username    string `json:"username,omitempty"`
}

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

var Listener = make(map[string]Subscribe)
var rtmStarted bool
var isExistingPrinted bool
var newClient *trello.Client
var board *trello.Board
var finalList *trello.List
var cards []*trello.Card
var oldCard *trello.Card

//GetCards trello
func GetCards(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	var trelloCards []*trello.Card
	client := trello.NewClient(apiKey, token)
	board, err := client.GetBoard(param.BoardID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	lists, err := board.GetLists(trello.Defaults())

	for _, list := range lists {
		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			result.WriteErrorResponse(responseWriter, err)
			return
		}

		trelloCards = append(trelloCards, cards...)
	}
	bytes, _ := json.Marshal(trelloCards)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//GetBoard trello
func GetBoard(responseWriter http.ResponseWriter, request *http.Request) {
	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)
	board, err := client.GetBoard(param.BoardID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(board)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//GetLists trello
func GetLists(responseWriter http.ResponseWriter, request *http.Request) {
	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)
	board, err := client.GetBoard(param.BoardID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	lists, err := board.GetLists(trello.Defaults())

	bytes, _ := json.Marshal(lists)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//AddCard trello
func AddCard(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var card TrelloArgs
	decodeErr := decoder.Decode(&card)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)

	list, err := client.GetList(card.ListID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}
	list.AddCard(&trello.Card{Name: card.CardName, Desc: card.Description}, trello.Defaults())

	message := Message{"true", "Card added successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//MoveCard trello
func MoveCard(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)

	card, err := client.GetCard(param.CardID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	moveErr := card.MoveToList(param.ListID, trello.Defaults())
	if moveErr != nil {
		result.WriteErrorResponse(responseWriter, moveErr)
		return
	}

	message := Message{"true", "Card moved successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//CopyCard trello
func CopyCard(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)

	card, err := client.GetCard(param.CardID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	copy, copyErr := card.CopyToList(param.ListID, trello.Defaults())
	if copyErr != nil {
		result.WriteErrorResponse(responseWriter, copyErr)
		return
	}

	bytes, _ := json.Marshal(copy)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//CreateBoard trello
func CreateBoard(responseWriter http.ResponseWriter, request *http.Request) {
	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)

	board := trello.NewBoard(param.BoardName)

	err := client.CreateBoard(&board, trello.Defaults())
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	message := Message{"true", "Board created successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//DeleteBoard trello
func DeleteBoard(responseWriter http.ResponseWriter, request *http.Request) {
	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)

	board, err := client.GetBoard(param.BoardID, trello.Defaults())
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	deleteErr := board.Delete(trello.Defaults())
	if deleteErr != nil {
		bytes, _ := json.Marshal(err.Error())
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	message := Message{"true", "Board deleted successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//CreateList trello
func CreateList(responseWriter http.ResponseWriter, request *http.Request) {
	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var args trello.Arguments
	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)

	board, err := client.GetBoard(param.BoardID, trello.Defaults())
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	list, listErr := board.CreateList(param.ListName, args)
	if listErr != nil {
		bytes, _ := json.Marshal(listErr.Error())
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	bytes, _ := json.Marshal(list)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//GetAllBoards trello
func GetAllBoards(responseWriter http.ResponseWriter, request *http.Request) {
	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)

	member, err := client.GetMember(param.Username, trello.Defaults())
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	bytes, _ := json.Marshal(boards)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//SubscribeCard
func SubscribeCard(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	var sub Subscribe
	decoder := json.NewDecoder(request.Body)
	decodeError := decoder.Decode(&sub)
	if decodeError != nil {
		result.WriteErrorResponse(responseWriter, decodeError)
		return
	}

	newClient = trello.NewClient(apiKey, token)

	//BoardID
	var err error
	board, err = newClient.GetBoard(sub.Data.BoardID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}

	Listener[sub.Data.ListID] = sub
	if !rtmStarted {
		go RTSTrello()
		rtmStarted = true
	}

	bytes, _ := json.Marshal("Subscribed")
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//RTSTrello function
func RTSTrello() {
	isTest := false
	for {
		if len(Listener) > 0 {
			for ListID, Sub := range Listener {
				go getMessageUpdates(ListID, Sub, Sub.Data.Existing)
				isTest = Sub.IsTesting
			}
		} else {
			rtmStarted = false
			break
		}
		time.Sleep(5 * time.Second)
		if isTest {
			break
		}
	}
}

func getMessageUpdates(listID string, sub Subscribe, existing bool) {

	var finalCard *trello.Card
	var finalCards []*trello.Card
	if listID != "" {

		lists, _ := board.GetLists(trello.Defaults())

		for _, list := range lists {

			if list.ID == listID {
				finalList = list
			}
		}
		if finalList != nil {
			cards, _ = finalList.GetCards(trello.Defaults())
		} else {
			log.Fatalln("Error : Please provide valid list ID")
			return
		}

	} else {
		cards, _ = board.GetCards(trello.Defaults())
	}

	if isExistingPrinted == false {
		if existing {
			finalCards = cards
		} else {
			if finalCard == nil {
				finalCard = latestCard(cards)
			}
		}
	} else {
		finalCard = latestCard(cards)
	}

	contentType := "application/json"

	t, err := cloudevents.NewHTTPTransport(cloudevents.WithTarget(sub.Endpoint), cloudevents.WithStructuredEncoding())
	if err != nil {
		log.Printf("failed to create transport, %v", err)
		return
	}

	c, err := cloudevents.NewClient(t, cloudevents.WithTimeNow())
	if err != nil {
		log.Printf("failed to create client, %v", err)
		return
	}

	source, err := url.Parse(sub.Endpoint)
	event := cloudevents.Event{
		Context: cloudevents.EventContextV01{
			EventID:     sub.ID,
			EventType:   "card",
			Source:      cloudevents.URLRef{URL: *source},
			ContentType: &contentType,
		}.AsV01(),
		Data: "",
	}

	if finalCards != nil {
		event.Data = finalCards

	} else {
		event.Data = finalCard
	}

	if oldCard == nil && finalCard != nil {
		oldCard = finalCard
	}

	if existing && !isExistingPrinted {
		resp, evt, err := c.Send(context.Background(), event)
		if err != nil {
			log.Printf("failed to send: %v (%v)", err, evt)
		}
		fmt.Printf("Response1: \n%s\n", resp)
		finalCards = nil
		isExistingPrinted = true
	} else if oldCard != nil && finalCard.ID != oldCard.ID {
		resp, evt, err := c.Send(context.Background(), event)
		if err != nil {
			log.Printf("failed to send: %v (%v)", err, evt)
		}
		fmt.Printf("Response2: \n%s\n", resp)
		oldCard = finalCard
		finalCard = nil
	}
}

func latestCard(cards []*trello.Card) *trello.Card {
	if len(cards) == 0 {
		return nil
	}
	latest := cards[0]
	for _, card := range cards {
		if card.ID > latest.ID {
			latest = card
		}
	}
	return latest
}
