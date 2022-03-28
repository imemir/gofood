package order

import (
	"github.com/gorilla/mux"
	"github.com/imemir/gofood/pkg/httpext"
	"github.com/imemir/gofood/pkg/validatorext"
	"net/http"
)

var HttpController controller

func init() {
	HttpController = controller{
		Service: Service,
	}
}

type controller struct {
	Service service
}

func (c controller) Register(router *mux.Router) {
	router.HandleFunc("/api/order", c.Create).Methods(http.MethodPost)
}

func (c controller) Create(w http.ResponseWriter, r *http.Request) {
	var (
		model orderModel
		err   error
	)
	err = httpext.BindModel(r, &model)
	if err != nil {
		httpext.SendError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	err = validatorext.Struct(model)
	if err != nil {
		httpext.SendError(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}
	err = c.Service.Create(model.OrderID, model.Price, model.Title)
	if err != nil {
		httpext.SendError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	httpext.SendModel(w, r, http.StatusOK, model)
}
