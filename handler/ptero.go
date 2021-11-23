package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/b4cktr4ck5r3/nade404api/config"
	"github.com/b4cktr4ck5r3/nade404api/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func GetPteroServerList(c *fiber.Ctx) error {
	serverNotFound := true
	client := http.Client{}
	url := "https://p.ezstrat.com/api/client/"

	for serverNotFound {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}

		req.Header = http.Header{
			"Accept":        []string{"application/json"},
			"Authorization": []string{"Bearer " + config.Config("PTERO_TOKEN")},
		}

		res, err := client.Do(req)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}

		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			t := new(model.PteroServerList)
			if err := json.NewDecoder(res.Body).Decode(t); err != nil {
				return c.Status(500).JSON(&fiber.Map{
					"succes":  false,
					"message": err,
				})
			}

			server, err := FindServerInList(t)
			if err != nil {
				if t.Meta.Pagination.CurrentPage < t.Meta.Pagination.TotalPages {
					url = t.Meta.Pagination.Links.Next
				} else {
					serverNotFound = false
				}

			} else {
				serverNotFound = false
				return c.Status(500).JSON(&fiber.Map{
					"succes":      true,
					"message":     "Server found",
					"serverInfos": server,
				})
			}
		} else {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": "Request didn't work",
			})
		}

	}

	return c.Status(200).JSON("null")
}

func FindServerInList(t *model.PteroServerList) (model.PteroServerListDatum, error) {
	fmt.Println("CALL FINDSERVER")
	for i := 0; i < len(t.Data); i++ {
		data := t.Data[i]
		if len(data.Attributes.Relationships.Allocations.Data) > 0 {
			ip := data.Attributes.Relationships.Allocations.Data[0].Attributes.IP
			port := data.Attributes.Relationships.Allocations.Data[0].Attributes.Port
			if ip == "51.158.82.97" && port == 27015 {
				return t.Data[i], nil
			}
		}
	}

	return model.PteroServerListDatum{}, errors.New("Server not found")
}

func GetPteroServerDetails(c *fiber.Ctx) error {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://p.ezstrat.com/api/client/servers/"+c.Params("server_id"), nil)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer odexNvfil7D21kXHc3UD9xRa5xFQOJ2PgVU74IwsZ0uV6OJK"},
	}

	res, err := client.Do(req)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		t := new(model.PteroServerListDatum)
		if err := json.NewDecoder(res.Body).Decode(t); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}
		fmt.Println(t.Attributes.Relationships.Allocations)
	}
	fmt.Println(res)
	return c.Status(200).JSON("null")
}
