package web

import (
	"encoding/json"
	"net/http"

	"github.com/AndersonOdilo/otel/internal/entity"
	"github.com/AndersonOdilo/otel/internal/usecase"
	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type WebTempHandler struct {
	LocationRepository 		entity.LocationRepositoryInterface
	TempRepository 			entity.TempRepositoryInterface
	Tracer            		trace.Tracer
}

func NewWebTempHandler(locationRepository entity.LocationRepositoryInterface, tempRepository entity.TempRepositoryInterface, tracer  trace.Tracer) *WebTempHandler {
	return &WebTempHandler{
		LocationRepository: locationRepository,
		TempRepository: tempRepository,
		Tracer: tracer,
	};
}

func (h *WebTempHandler) Get(w http.ResponseWriter, r *http.Request) {

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	// ctx, span := h.Tracer.Start(ctx, "getLocationAndTemp")
	// defer span.End();

	inputDTO := usecase.GetTempInputDTO{
		Cep:chi.URLParam(r, "cep") ,
	};

	getTempUseCase := usecase.NewGetTempUseCase(h.LocationRepository, h.TempRepository, ctx, h.Tracer);

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
