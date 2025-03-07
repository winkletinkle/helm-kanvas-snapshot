package errors

import (
	"github.com/layer5io/meshkit/errors"
)

var (
	ErrInvalidChartURICode          = "kanvas-snapshot-1000"
	ErrCreatingMesheryDesignCode    = "kanvas-snapshot-1001"
	ErrGeneratingSnapshotCode       = "kanvas-snapshot-1002"
	ErrHTTPPostRequestCode          = "kanvas-snapshot-1003"
	ErrDecodingAPICode              = "kanvas-snapshot-1004"
	ErrRequiredFieldNotProvidedCode = "kanvas-snapshot-1005"
)

func ErrInvalidChartURI(err error) error {
	return errors.New(ErrInvalidChartURICode, errors.Alert,
		[]string{"Invalid or missing Helm chart URI."},
		[]string{err.Error()},
		[]string{"The provided URI for the Helm chart is either missing or invalid."},
		[]string{"Ensure the Helm chart URI is correctly provided."},
	)
}

func ErrCreatingMesheryDesign(err error) error {
	return errors.New(ErrCreatingMesheryDesignCode, errors.Alert,
		[]string{"Failed to create Meshery design."},
		[]string{err.Error()},
		[]string{"Meshery Design creation failed due to an error."},
		[]string{"Check Meshery API connection and ensure the payload is correct."},
	)
}

func ErrGeneratingSnapshot(err error) error {
	return errors.New(ErrGeneratingSnapshotCode, errors.Alert,
		[]string{"Failed to generate snapshot."},
		[]string{err.Error()},
		[]string{"Snapshot generation failed due to an error."},
		[]string{"Check Meshery Cloud API connection and payload."},
	)
}

func ErrHTTPPostRequest(err error) error {
	return errors.New(ErrHTTPPostRequestCode, errors.Alert,
		[]string{"Failed to perform HTTP POST request."},
		[]string{err.Error()},
		[]string{"HTTP POST request failed during interaction with Meshery API."},
		[]string{"Check Meshery API endpoint and ensure valid request payload."},
	)
}

func ErrDecodingAPI(err error) error {
	return errors.New(ErrDecodingAPICode, errors.Alert,
		[]string{"Failed to decode API response."},
		[]string{err.Error()},
		[]string{"API response could not be decoded into the expected format."},
		[]string{"Ensure the Meshery API response format is correct."},
	)
}
