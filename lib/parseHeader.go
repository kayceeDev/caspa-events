package lib

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kayceeDev/caspa-events/pkg/config"


)

func ParseHeader(c *gin.Context) (*float64, error) {
	env := config.EnvVars()

	token := GetHeaderToken(c)
	decoded, err := DecodeJwtToken(*token, env.JWT_SECRET)

	if err != nil {
		return nil, err
	}

	uid, ok := decoded["userID"].(float64)
	if !ok {
		return nil, errors.New("error parsing decoded value")
	}
	return &uid, nil
}

func ParseCookie(c *gin.Context) (*float64, error) {
	env := config.EnvVars()

	if cookie, err := c.Cookie("refresh_token"); err == nil {
		decoded, err := DecodeJwtToken(cookie, env.JWT_SECRET)
		if err != nil {
			return nil, err
		}
		uid, ok := decoded["userID"].(float64)
		if !ok {
			return nil, errors.New("error parsing decoded value")
		}
		return &uid, nil
	} else {
		return nil, errors.New("cookie not found")
	}
}

// Update ParsePaginationCategoryQuery to handle cursors more explicitly
func ParsePaginationCategoryQuery(c *gin.Context, categories []string) (map[string]string, int) {
	cursors := make(map[string]string)

	// Get the cursor from the query
	cursorStr := c.Query("cursor")
	if cursorStr != "" {
		// Split the cursor string by the delimiter
		cursorParts := strings.Split(cursorStr, "_")

		// Map each part to the corresponding category
		for i, category := range categories {
			if i < len(cursorParts) {
				cursors[category] = cursorParts[i]
			}
		}
	}

	// Parse and validate limit
	limit := 10 // default limit
	if limitStr := c.Query("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	return cursors, limit
}
func ParsePaginationQuery(c *gin.Context) (*[]string, int) {
	cursor, ok := c.Request.URL.Query()["cursor"]
	if !ok {
		cursor = nil
	}
	limit, ok := c.Request.URL.Query()["limit"]
	if !ok {
		limit = append(limit, "10")
	}

	l, err := strconv.Atoi(limit[0])
	if err != nil {
		return &cursor, 10
	}
	return &cursor, l
}

func ParseUUIDParam(c *gin.Context) (*uuid.UUID, error) {
	uuid_param := c.Params.ByName("uuid")
	parsed_uuid, err := uuid.Parse(uuid_param)

	if err != nil {
		return nil, err
	}

	return &parsed_uuid, nil
}

func ParseIDParam(c *gin.Context) (*int, error) {
	id_param := c.Params.ByName("id")

	if strings.TrimSpace(id_param) == "" {
		return nil, errors.New("ID value missing")
	}
	parsed_id, err := strconv.Atoi(id_param)
	if err != nil {
		return nil, err
	}
	return &parsed_id, nil
}

func ParseNamedUUIDParam(name string, c *gin.Context) (*uuid.UUID, error) {
	if name == "" {
		return nil, fmt.Errorf("A valid UUID is required")
	}
	uuidString := c.Params.ByName(name)

	parsedUUID, err := uuid.Parse(uuidString)

	if err != nil {
		return nil, err
	}

	return &parsedUUID, nil
}
