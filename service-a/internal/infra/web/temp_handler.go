package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AndersonOdilo/otel/service-a/internal/entity"
	"github.com/AndersonOdilo/otel/service-a/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type WebTempHandler struct {
	TempRepository 			entity.TempRepositoryInterface
	Tracer            		trace.Tracer
}

func NewWebTempHandler(tempRepository entity.TempRepositoryInterface, tracer  trace.Tracer) *WebTempHandler {
	return &WebTempHandler{
		TempRepository: tempRepository,
		Tracer: tracer,
	};
}

func (h *WebTempHandler) Get(w http.ResponseWriter, r *http.Request) {

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	resp, err := io.ReadAll(r.Body);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var inputDTO usecase.GetTempInputDTO;
	err = json.Unmarshal(resp, &inputDTO)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	getTempUseCase := usecase.NewGetTempUseCase(h.TempRepository, ctx, h.Tracer);

	output, err := getTempUseCase.Execute(inputDTO);

	if err != nil {

		if (err.Error() == "can not find zipcode"){

			http.Error(w, err.Error(), http.StatusNotFound)

		} else if (err.Error() == "invalid zipcode"){

			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		
		}else {

			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
