package handlers

import (
	"errors"
	"net/http"
	"urlshortener/internal/models"
	"urlshortener/internal/service"

	"github.com/labstack/echo/v4"
)

type CreateShortUrlRequest struct {
	URL string `json:"url"`
}

type CreateShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

// CreateShortUrl godoc
// @Summary      Создание короткой ссылки
// @Description  Создание короткой ссылки
// @Tags         shortUrl
// @Accept       json
// @Produce      json
// @Param        rq   body       CreateShortUrlRequest true  "Входящие данные"
// @Success	     200  {object}  CreateShortUrlResponse
// @Failure      400  {object}  echo.HTTPError  "Некорректный URL"
// @Router       /shortener [post]
func (h *Handler) CreateShortUrl(c echo.Context) error {
	var req CreateShortUrlRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if !models.ValidateURL(req.URL) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid URL")
	}

	shortURL, err := h.srv.Create(c.Request().Context(), req.URL)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, CreateShortUrlResponse{ShortUrl: shortURL})
}

// GetURLs godoc
// @Summary      Получение списка коротких ссылок
// @Description  Возвращает список всех созданных коротких ссылок
// @Tags         shortUrl
// @Accept       json
// @Produce      json
// @Success      200  {array}   string  "Список коротких ссылок"
// @Failure      500  {object}  echo.HTTPError  "Ошибка сервера"
// @Router       /shortener [get]
func (h *Handler) GetURLs(c echo.Context) error {
	shortURLs, err := h.srv.GetURLs(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, shortURLs)
}

// GetLongURL godoc
// @Summary      Получение оригинальной ссылки
// @Description  Перенаправляет на оригинальную ссылку, связанную с короткой ссылкой
// @Tags         shortUrl
// @Accept       json
// @Produce      json
// @Param        link  path      string  true  "Короткая ссылка"
// @Success      302   {string}  string  "Перенаправление на оригинальную ссылку"
// @Failure      404   {object}  echo.HTTPError  "Короткая ссылка не найдена"
// @Failure      500   {object}  echo.HTTPError  "Ошибка сервера"
// @Router       /{link} [get]
func (h *Handler) GetLongURL(c echo.Context) error {
	url := c.Param("link")
	longURL, err := h.srv.GetLongURL(c.Request().Context(), url)
	if err != nil {
		if errors.Is(err, service.ErrLinkNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Link not found")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusFound, longURL)
}

// DeleteURL godoc
// @Summary      Удаление короткой ссылки
// @Description  Удаляет короткую ссылку из базы данных
// @Tags         shortUrl
// @Accept       json
// @Produce      json
// @Param        link  path      string  true  "Короткая ссылка"
// @Success      200   {object}  nil    "Короткая ссылка успешно удалена"
// @Failure      404   {object}  echo.HTTPError  "Короткая ссылка не найдена"
// @Failure      500   {object}  echo.HTTPError  "Ошибка сервера"
// @Router       /{link} [delete]
func (h *Handler) DeleteURL(c echo.Context) error {
	url := c.Param("link")

	err := h.srv.DeleteURL(c.Request().Context(), url)
	if err != nil {
		if errors.Is(err, service.ErrLinkNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Link not found")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

// GetStats godoc
// @Summary      Получение статистики по ссылке
// @Description  Возвращает статистику по количеству переходов по короткой ссылке
// @Tags         shortUrl
// @Accept       json
// @Produce      json
// @Param        link  path      string  true  "Короткая ссылка"
// @Success      200   {object}  urlshortener_internal_models.Stat  "Статистика по ссылке"
// @Failure      404   {object}  echo.HTTPError  "Статистика не найдена"
// @Failure      500   {object}  echo.HTTPError  "Ошибка сервера"
// @Router       /stats/{link} [get]
func (h *Handler) GetStats(c echo.Context) error {
	url := c.Param("link")
	stat, err := h.srv.GetStats(c.Request().Context(), url)
	if err != nil {
		if errors.Is(err, service.ErrStatNotFound) || errors.Is(err, service.ErrLinkNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Stat not found")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stat)
}
