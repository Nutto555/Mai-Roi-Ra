package controllers

import (
	"fmt"
	"log"
	"net/http"

	st "github.com/2110366-2566-2/Mai-Roi-Ra/backend/pkg/struct"
	"github.com/2110366-2566-2/Mai-Roi-Ra/backend/services"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	ServiceGateway services.ServiceGateway
}

func NewEventController(
	sg services.ServiceGateway,
) *EventController {
	return &EventController{
		ServiceGateway: sg,
	}
}

// CreateEvent creates a new event.
// @Summary Create new event
// @Description Create a new event with the provided details.
// @Tags events
// @Accept json
// @Produce json
// @Param request body structure.CreateEventRequest true "Create Event Request" example:{"organizer_id": "org123", "admin_id": "admin456", "location_id": "loc789", "start_date": "2024-01-15", "end_date": "2024-01-20", "status": "planned", "participant_fee": 50.0, "description": "Annual tech conference focusing on the future of technology.", "event_name": "TechFuture 2024", "deadline": "2023-12-31", "activities": "Keynotes, Workshops, Panels", "event_image": "http://example.com/image.jpg"}
// @Success 200 {object} structure.CreateEventResponse
// @Failure 400 {object} object "Bad Request"
// @Failure 500 {object} object "Internal Server Error"
// @Router /events [post]
func (c *EventController) CreateEvent(ctx *gin.Context, req *st.CreateEventRequest) {
	log.Println("[CTRL: CreateEvent] Input:", req)

	res, err := c.ServiceGateway.EventService.CreateEvent(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("[CTRL: CreateEvent] Output:", res)
	ctx.JSON(http.StatusOK, res)
}

// @Summary GetEventLists
// @Description Get list of events
// @Tags events
// @Accept json
// @Produce json
// @Param organizer_id query string false "For My events"
// @Param filter query string false "Status query. i.e. Approved"
// @Param sort query string false "Sort order. i.e. start_date ASC"
// @Param offset query int false "offset i.e. 0"
// @Param limit query int false "Items per page i.e. 10"
// @Success 200 {object} structure.GetEventListsResponse
// @Failure 400 {object} object "Bad Request"
// @Failure 500 {object} object "Internal Server Error"
// @Router /events [get]
func (c *EventController) GetEventLists(ctx *gin.Context, req *st.GetEventListsRequest) {
	log.Println("[CTRL: GetEventLists] Input:", req)
	res, err := c.ServiceGateway.EventService.GetEventLists(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("[CTRL: GetEventLists] Output:", res)
	ctx.JSON(http.StatusOK, res)
}

// @Summary GetEventDataById
// @Description Get a test message
// @Tags events
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Success 200 {object} structure.GetEventDataByIdResponse
// @Failure 400 {object} object "Bad Request"
// @Failure 500 {object} object "Internal Server Error"
// @Router /events/{event_id} [get]
func (c *EventController) GetEventDataById(ctx *gin.Context, req st.GetEventDataByIdRequest) {
	log.Println("[CTRL: GetEventDataById] Input:", req)
	res, err := c.ServiceGateway.EventService.GetEventDataById(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("[CTRL: GetEventDataById] Output:", res)
	ctx.JSON(http.StatusOK, res)
}

func UpdateEvent(c *gin.Context) {
	// Get the eventid parameter from the URL
	eventID := c.Param("eventid")

	// Parse the JSON request body into an Event struct
	var updatedEvent st.UpdateEventRequest
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return a successful response
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Updated event with ID: %s", eventID),
		"event":   updatedEvent,
	})
}

func DeleteEvent(c *gin.Context) {
	// Get the eventid parameter from the URL
	eventID := c.Param("eventid")

	// Return a successful response
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Deleted event with ID: %s", eventID),
	})
}