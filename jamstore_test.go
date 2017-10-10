package app

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine"
	"strings"
	"github.com/stretchr/testify/mock"
	"context"
)

type MockGetContext struct {
	mock.Mock
}

func (m *MockGetContext)getContext() context.Context {
	args := m.Called()
	return args.Get(0).(context.Context)
}

func buildMockGetContext(t *testing.T, ctx context.Context) *MockGetContext {
	mockGetContext := new(MockGetContext)
	mockGetContext.On("getContext").Return(ctx)
	return mockGetContext
}

func buildContext(t *testing.T, inst aetest.Instance) context.Context {
	req, err := inst.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := appengine.NewContext(req)
	return ctx
}
func buildInstance(t *testing.T) (aetest.Instance) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatal(err)
	}
	return inst
}

func TestStoreJam__calls_get_context(t *testing.T) {
	inst := buildInstance(t)
	defer inst.Close()
	ctx := buildContext(t, inst)
	mockGetContext := buildMockGetContext(t, ctx)
	StoreJam(mockGetContext, "test jam", true)
	mockGetContext.AssertCalled(t, "getContext")
}

func TestStoreJam__correctly_stores_jam_in_datastore(t *testing.T) {
	inst := buildInstance(t)
	defer inst.Close()
	ctx := buildContext(t, inst)
	mockGetContext := buildMockGetContext(t, ctx)

	jamQuery := "foo"
	jamState := true
	if err := StoreJam(mockGetContext, jamQuery, jamState); err != nil {
		t.Fatal(err)
	}

	query := datastore.NewQuery("Jam").Filter("JamText =", jamQuery)

	var jams []Jam
	_, err := query.GetAll(ctx, &jams)
	if err != nil {
		t.Fatal(err)
	}

	if len(jams) == 0 {
		t.Fatal("No Jams found")
	}
	if query := jams[0].JamText; query != jamQuery {
		t.Errorf("Expected %v, got %v", jamQuery, query)
	}
	if state := jams[0].State; state != jamState {
		t.Errorf("Expected %v, got %v", jamState, state)
	}
}

func TestStoreJam__correctly_stores_jam_in_datastore_with_lower_case_text(t *testing.T) {
	inst := buildInstance(t)
	defer inst.Close()
	ctx := buildContext(t, inst)
	mockGetContext := buildMockGetContext(t, ctx)

	jamText := "FOO"
	jamState := true
	if err := StoreJam(mockGetContext, jamText, jamState); err != nil {
		t.Fatal(err)
	}

	lowerText := strings.ToLower(jamText)
	query := datastore.NewQuery("Jam").Filter("JamText =", lowerText)

	var jams []Jam
	_, err := query.GetAll(ctx, &jams)
	if err != nil {
		t.Fatal(err)
	}

	if len(jams) == 0 {
		t.Fatal("No Jams found")
	}
}
