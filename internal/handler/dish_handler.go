package handler

import (
	"GourmetOS/internal/domain"
	"GourmetOS/internal/handler/dto"
	"GourmetOS/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DishHandler struct {
	service *service.DishService
}

func NewDishHandler(service *service.DishService) *DishHandler {
	return &DishHandler{
		service: service,
	}
}

func (h *DishHandler) CreateDish(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	var req domain.Dish
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат данных"+err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidateError", "Неверный формат данных"+err.Error())
		return
	}
	created, err := h.service.CreateDish(r.Context(), conn, &req)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrCreateDish", err.Error())
		return
	}
	resp := dto.DishResponseFromDomain(*created)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) GetDish(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	dishID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "InvalidID", "Нету такого id"+err.Error())
	}
	dish, err := h.service.GetDish(r.Context(), conn, dishID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetDish", err.Error())
		return
	}
	resp := dto.DishResponseFromDomain(*dish)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) GetAllDishes(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	dishes, err := h.service.GetAllDishes(r.Context(), conn, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetAllDishes", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, dishes)
}

func (h *DishHandler) GetDishesByCategory(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	category := vars["category"]

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	dishes, err := h.service.GetDishesByCategory(r.Context(), conn, category, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetDishesByCategory", err.Error())
		return
	}
	resp := dto.NewDishListResponse(dishes, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) GetDishesByCuisine(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	cuisine := vars["cuisine"]
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	dishes, err := h.service.GetDishesByCuisine(r.Context(), conn, cuisine, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetDishesByCuisine", err.Error())
		return
	}
	resp := dto.NewDishListResponse(dishes, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) GetDishesByPriceRange(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	minR, _ := strconv.Atoi(r.URL.Query().Get("min"))
	maxR, _ := strconv.Atoi(r.URL.Query().Get("max"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	dishes, err := h.service.GetDishesByPriceRange(r.Context(), conn, minR, maxR, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetDishesByCuisine", err.Error())
		return
	}
	resp := dto.NewDishListResponse(dishes, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) GetAvailableDishes(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	dishes, err := h.service.GetAvailableDishes(r.Context(), conn, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetAvailableDishes", err.Error())
		return
	}
	resp := dto.NewDishListResponse(dishes, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) GetVegetarianDishes(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	dishes, err := h.service.GetVegetarianDishes(r.Context(), conn, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetVegetarianDishes", err.Error())
		return
	}
	resp := dto.NewDishListResponse(dishes, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) GetVeganDishes(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	dishes, err := h.service.GetVeganDishes(r.Context(), conn, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetVeganDishes", err.Error())
		return
	}
	resp := dto.NewDishListResponse(dishes, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) GetGlutenFreeDishes(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	dishes, err := h.service.GetGlutenFreeDishes(r.Context(), conn, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetGlutenFreeDishes", err.Error())
		return
	}
	resp := dto.NewDishListResponse(dishes, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) UpdateDish(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	dishId, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "ErrGetDishId", err.Error())
		return
	}

	var rep dto.UpdateDishRequest
	if err := json.NewDecoder(r.Body).Decode(&rep); err != nil {
		sendError(w, http.StatusBadRequest, "ErrUpdateDish", err.Error())
		return
	}
	if err := rep.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ErrUpdateDish", err.Error())
		return
	}
	updates := make(map[string]interface{})
	if rep.Category != "" || len(rep.Category) != 0 {
		updates["category"] = rep.Category
	}
	if rep.Name != "" || len(rep.Description) != 0 {
		updates["name"] = rep.Name
	}
	if rep.Price != 0 {
		updates["price"] = rep.Price
	}
	if rep.IsVegan != false {
		updates["is_vegan"] = rep.IsVegan
	}
	if rep.IsGlutenFree != false {
		updates["is_gluten_free"] = rep.IsGlutenFree
	}
	if rep.IsVegetarian != false {
		updates["is_vegetarian"] = rep.IsVegetarian
	}
	if rep.Description != "" || len(rep.Description) != 0 {
		updates["description"] = rep.Description
	}
	if rep.Calories != nil {
		updates["calories"] = rep.Calories
	}
	if rep.ImageURL != nil {
		updates["image_url"] = rep.ImageURL
	}
	if rep.CookingTime != 0 {
		updates["cooking_time"] = rep.CookingTime
	}
	if rep.Cuisine != "" || len(rep.Cuisine) != 0 {
		updates["cuisine"] = rep.Cuisine
	}
	dish, err := h.service.UpdateDishService(r.Context(), conn, dishId, updates)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrUpdateDish", err.Error())
		return
	}
	resp := dto.DishResponseFromDomain(*dish)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) DeleteDish(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	dishId, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "ErrGetDishId", err.Error())
		return
	}
	if err := h.service.DeleteDish(r.Context(), conn, dishId); err != nil {
		sendError(w, http.StatusInternalServerError, "ErrDeleteDish", err.Error())
		return
	}
	sendSuccess(w, http.StatusNoContent, nil)
}

func (h *DishHandler) UpdateAvailability(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	dishID, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, http.StatusBadRequest, "InvalidDishID", "Неверный ID блюда")
		return
	}

	var req dto.UpdateAvailabilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	err = h.service.UpdateAvailability(r.Context(), conn, dishID, req.IsAvailable)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "UpdateError", err.Error())
		return
	}

	sendSuccess(w, http.StatusOK, map[string]interface{}{"message": "Доступность блюда обновлена", "dish_id": dishID, "is_available": req.IsAvailable})
}

func (h *DishHandler) AddDecorator(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}

	vars := mux.Vars(r)
	dishId, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "ErrGetDishId", err.Error())
		return
	}

	var req struct {
		DecoratorType string `json:"decorator_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if req.DecoratorType == "" {
		sendError(w, http.StatusBadRequest, "ErrAddDecorator", "decorator_type не может быть пустым")
		return
	}

	decor, err := h.service.AddDecorator(r.Context(), conn, dishId, req.DecoratorType)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrAddDecorator", err.Error())
		return
	}
	resp := dto.DishResponseFromDomain(*decor)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) RemoveDecorator(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	dishId, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "ErrGetDishId", err.Error())
		return
	}
	decoratorType := vars["decoratorType"]
	decor, err := h.service.RemoveDecorator(r.Context(), conn, dishId, decoratorType)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrRemoveDecorator", err.Error())
		return
	}
	resp := dto.DishResponseFromDomain(*decor)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *DishHandler) GetMenuTree(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	menu, err := h.service.GetMenuTree(r.Context(), conn)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetMenuTree", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, menu)
}

func (h *DishHandler) GetComboMeal(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	comboType := vars["comboType"]
	combo, err := h.service.GetComboMeal(r.Context(), conn, comboType)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrGetComboMeal", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, combo)
}

func (h *DishHandler) SearchDishes(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	query := r.URL.Query().Get("query")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if query == "" {
		sendError(w, http.StatusBadRequest, "ErrSearchDishes", "Поисковый запрос не может быть пустым")
		return
	}

	search, total, err := h.service.SearchDishes(r.Context(), conn, query, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrSearchDishes", err.Error())
		return
	}
	resp := dto.NewDishListResponse(search, total, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}
