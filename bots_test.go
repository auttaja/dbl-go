package dbl_go

import (
	"net/http"
	"os"
	"testing"
)

func TestDBL_GetBot(t *testing.T) {
	dbl := NewDBL(os.Getenv("TOKEN"), &http.Client{})
	_, err := dbl.GetBot("242730576195354624")
	if err != nil {
		t.Errorf("GetBot threw an error: %v", err)
	}
}

func TestDBL_CheckUserVote(t *testing.T) {
	dbl := NewDBL(os.Getenv("TOKEN"), &http.Client{})
	_, err := dbl.CheckUserVote("242730576195354624", "435066999987634177")
	if err != nil {
		t.Errorf("GetRecentVotes threw an error: %v", err)
	}
}

func TestDBL_GetBotStats(t *testing.T) {
	dbl := NewDBL(os.Getenv("TOKEN"), &http.Client{})
	_, err := dbl.GetBotStats("242730576195354624")
	if err != nil {
		t.Errorf("GetBotStats threw an error: %v", err)
	}
}

func TestDBL_PostBotStats(t *testing.T) {
	dbl := NewDBL(os.Getenv("TOKEN"), &http.Client{})
	err := dbl.PostBotStats("242730576195354624", &Stats{
		ServerCount: 69727,
		ShardCount:  70,
	})
	if err != nil {
		t.Errorf("PostBotStats threw an error: %v", err)
	}
}
