package credentials

import (
	"context"
	"testing"

	"oss-sdk-go/oss"
)

func TestStaticCredentialsProvider(t *testing.T) {
	s := StaticCredentialsProvider{
		Value: oss.Credentials{
			AccessKeyID:     "AKID",
			SecretAccessKey: "SECRET",
			SessionToken:    "",
		},
	}

	creds, err := s.Retrieve(context.Background())
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "AKID", creds.AccessKeyID; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "SECRET", creds.SecretAccessKey; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if l := creds.SessionToken; len(l) != 0 {
		t.Errorf("expect no token, got %v", l)
	}
}

func TestStaticCredentialsProviderIsExpired(t *testing.T) {
	s := StaticCredentialsProvider{
		Value: oss.Credentials{
			AccessKeyID:     "AKID",
			SecretAccessKey: "SECRET",
			SessionToken:    "",
		},
	}

	creds, err := s.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if creds.Expired() {
		t.Errorf("expect static credentials to never expire")
	}
}
