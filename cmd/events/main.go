package main

import (
	"database/sql"
	"net/http"

	"github.com/paulo1402/imersao18-golang/internal/events/infra/repository"
	"github.com/paulo1402/imersao18-golang/internal/events/infra/service"
	"github.com/paulo1402/imersao18-golang/internal/events/usecase"

	httpHandler "github.com/paulo1402/imersao18-golang/internal/events/infra/http"
)

func main() {
		db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao18")

		if err != nil {
			panic(err)
		}

		defer db.Close()

		eventRepo, err := repository.NewMysqlEventRepository(db)

		if err != nil {
			panic(err)
		}

		parnerBaseURLs := map[int]string{
			'1': "http://localhost:8080",
			'2': "http://localhost:8081",
		}

		partnerFactory := service.NewPartnerFactory(parnerBaseURLs)

		listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
		getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
		listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
		buyTicketUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

		eventsHandler := httpHandler.NewEventsHandler(
			listEventsUseCase,
			listSpotsUseCase,
			getEventUseCase,
			buyTicketUseCase,
			nil,
			nil,
		)

		r := http.NewServeMux()
		r.HandleFunc("/events", eventsHandler.ListEvents)
		r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
		r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
		r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

		http.ListenAndServe(":8080", r)
}
