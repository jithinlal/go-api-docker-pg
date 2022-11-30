package store

import (
	"context"
	"fmt"
	"github.com/jithinlal-gelato/go_api/objects"
	"math/rand"
	"time"
)

type IEventStore interface {
	Get(ctx context.Context, in *objects.GetRequest) (*objects.Event, error)
	List(ctx context.Context, in *objects.ListRequest) ([]*objects.Event, error)
	Create(ctx context.Context, in *objects.CreateRequest) error
	UpdateDetails(ctx context.Context, in *objects.UpdateDetailsRequest) error
	Cancel(ctx context.Context, in *objects.CancelRequest) error
	Reschedule(ctx context.Context, in *objects.RescheduleRequest) error
	Delete(ctx context.Context, in *objects.DeleteRequest) error
}

func init() {
	rand.Seed(time.Now().UTC().Unix())
}

func GenerateUniqueID() string {
	word := []byte("0987654321")
	rand.Shuffle(len(word), func(i, j int) {
		word[i], word[j] = word[j], word[i]
	})
	now := time.Now().UTC()
	return fmt.Sprintf("%010v-%010v-%s", now.Unix(), now.Nanosecond(), string(word))
}
