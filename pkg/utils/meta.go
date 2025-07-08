package utils

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Pagination struct {
	Page     int
	PageSize int
}

type PaginatedResponseBanks[T any] struct {
	Banks []T   `json:"banks"`
	Meta  Meta  `json:"meta"`
	Links Links `json:"links"`
}

type PaginatedResponseTransactions[T any] struct {
	Transactions []T   `json:"transactions"`
	Meta         Meta  `json:"meta"`
	Links        Links `json:"links"`
}

type PaginatedResponseTransfers[T any] struct {
	Transfers []T   `json:"transfers"`
	Meta      Meta  `json:"meta"`
	Links     Links `json:"links"`
}

type Meta struct {
	ItemsPerPage int         `json:"itemsPerPage"`
	TotalItems   int64       `json:"totalItems"`
	CurrentPage  int         `json:"currentPage"`
	TotalPages   int         `json:"totalPages"`
	SortBy       [][2]string `json:"sortBy"`
}

type Links struct {
	Current string `json:"current"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
	First   string `json:"first,omitempty"`
	Last    string `json:"last,omitempty"`
}

func BuildLinks(r *http.Request, page, total, limit int, sortBy [][2]string) Links {
	baseURL := fmt.Sprintf("https://%s%s", r.Host, r.URL.Path)
	query := r.URL.Query()

	links := Links{
		Current: buildURL(baseURL, query, page, limit, sortBy),
	}

	if page > 1 {
		query.Set("page", "1")
		links.First = buildURL(baseURL, query, 1, limit, sortBy)

		query.Set("page", strconv.Itoa(page-1))
		links.Prev = buildURL(baseURL, query, page-1, limit, sortBy)
	}

	if page*limit < total {
		query.Set("page", strconv.Itoa(page+1))
		links.Next = buildURL(baseURL, query, page+1, limit, sortBy)

		totalPages := int(math.Ceil(float64(total) / float64(limit)))
		query.Set("page", strconv.Itoa(totalPages))
		links.Last = buildURL(baseURL, query, totalPages, limit, sortBy)
	}

	return links
}

func buildURL(base string, query url.Values, page, limit int, sortBy [][2]string) string {
	query.Set("page", strconv.Itoa(page))
	query.Set("limit", strconv.Itoa(limit))

	for k := range query {
		if k == "sortBy" {
			delete(query, k)
		}
	}

	for _, sort := range sortBy {
		query.Add("sortBy", fmt.Sprintf("%s:%s", sort[0], sort[1]))
	}

	return fmt.Sprintf("%s?%s", base, query.Encode())
}

func ParseSortParams(sortBy []string) [][2]string {
	var sortParams [][2]string
	for _, param := range sortBy {
		parts := strings.Split(param, ":")
		if len(parts) == 2 {
			sortParams = append(sortParams, [2]string{parts[0], parts[1]})
		}
	}
	return sortParams
}
