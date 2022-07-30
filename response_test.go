package httputils

import "testing"

func TestSuccessResponseEmptyTransactionId(t *testing.T) {
	_, response := SuccessResponse(200, "OK", nil, nil)
	if response.Transaction != nil {
		t.Errorf("Expected empty transaction id, got %s", *response.Transaction)
	}
}
func TestSuccessResponseNonEmptyTransactionId(t *testing.T) {
	transactionId := "transactionId"
	_, response := SuccessResponse(200, "OK", nil, &transactionId)
	if response.Transaction == nil {
		t.Errorf("Expected non-empty transaction id, got %s", *response.Transaction)
	}
}

func TestErrorResponseEmptyTransactionId(t *testing.T) {
	_, response := ErrorResponse(500, "Internal Server Error", "", nil, nil)
	if response.Transaction != nil {
		t.Errorf("Expected empty transaction id, got %s", *response.Transaction)
	}
}

func TestErrorResponseNonEmptyTransactionId(t *testing.T) {
	transactionId := "transactionId"
	_, response := ErrorResponse(500, "Internal Server Error", "", nil, &transactionId)
	if response.Transaction == nil {
		t.Errorf("Expected non-empty transaction id, got %s", *response.Transaction)
	}
}
