package client

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestCheckoutService_Create(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusCreated, stubs.CheckoutGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	checkout, response, err := client.Checkouts.Create(context.Background(), &CheckoutCreateParams{
		CustomPrice:     5000,
		EnabledVariants: []int{1},
		ButtonColor:     "#2DD272",
		CustomData:      map[string]string{"user_id": "123"},
		ExpiresAt:       time.Now().UTC(),
		StoreID:         "1",
		VariantID:       "1",
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.CheckoutGetResponse(), *response.Body)
	assert.Equal(t, "5e8b546c-c561-4a2c-a586-40c18bb2a195", checkout.Data.ID)

	// Teardown
	server.Close()
}

func TestCheckoutService_CreateWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Checkouts.Create(context.Background(), &CheckoutCreateParams{})

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestCheckoutService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.CheckoutGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	checkout, response, err := client.Checkouts.Get(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.CheckoutGetResponse(), *response.Body)
	assert.Equal(t, "5e8b546c-c561-4a2c-a586-40c18bb2a195", checkout.Data.ID)

	// Teardown
	server.Close()
}

func TestCheckoutService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Checkouts.Get(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestCheckoutService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.CheckoutListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	checkouts, response, err := client.Checkouts.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.CheckoutListResponse(), *response.Body)
	assert.Equal(t, 1, len(checkouts.Data))

	// Teardown
	server.Close()
}

func TestCheckoutService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Checkouts.List(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
