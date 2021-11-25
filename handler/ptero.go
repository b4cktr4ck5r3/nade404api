package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/b4cktr4ck5r3/nade404api/config"
	"github.com/b4cktr4ck5r3/nade404api/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func GetPteroServerByIpAndPort(c *fiber.Ctx) error {
	var ip string
	var port int
	if c.Query("ip") != "" {
		ip = c.Query("ip")
	} else {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": "No IP query find",
		})
	}

	if c.Query("port") != "" {
		tempPort, convertError := strconv.Atoi(c.Query("port"))
		if convertError != nil {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": "Port Query cannot be cast in int value",
			})
		}
		port = tempPort
	} else {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": "No Port query find",
		})
	}

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

			server, err := findServerInList(t, ip, port)
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

	return c.Status(500).JSON(&fiber.Map{
		"succes":  false,
		"message": "No server found",
	})
}

func findServerInList(t *model.PteroServerList, ip string, port int) (model.PteroServerListDatum, error) {
	for i := 0; i < len(t.Data); i++ {
		data := t.Data[i]
		if len(data.Attributes.Relationships.Allocations.Data) > 0 {
			currentIp := data.Attributes.Relationships.Allocations.Data[0].Attributes.IP
			currentPort := data.Attributes.Relationships.Allocations.Data[0].Attributes.Port
			if currentIp == ip && currentPort == port {
				return t.Data[i], nil
			}
		}
	}

	return model.PteroServerListDatum{}, errors.New("Server not found")
}
