// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testNeighbors(t *testing.T) {
	t.Parallel()

	query := Neighbors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNeighborsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNeighborsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Neighbors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNeighborsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NeighborSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNeighborsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NeighborExists(ctx, tx, o.CrawlID, o.PeerID)
	if err != nil {
		t.Errorf("Unable to check if Neighbor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NeighborExists to return true, but got false.")
	}
}

func testNeighborsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	neighborFound, err := FindNeighbor(ctx, tx, o.CrawlID, o.PeerID)
	if err != nil {
		t.Error(err)
	}

	if neighborFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNeighborsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Neighbors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNeighborsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Neighbors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNeighborsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	neighborOne := &Neighbor{}
	neighborTwo := &Neighbor{}
	if err = randomize.Struct(seed, neighborOne, neighborDBTypes, false, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}
	if err = randomize.Struct(seed, neighborTwo, neighborDBTypes, false, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = neighborOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = neighborTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Neighbors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNeighborsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	neighborOne := &Neighbor{}
	neighborTwo := &Neighbor{}
	if err = randomize.Struct(seed, neighborOne, neighborDBTypes, false, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}
	if err = randomize.Struct(seed, neighborTwo, neighborDBTypes, false, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = neighborOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = neighborTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func neighborBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Neighbor) error {
	*o = Neighbor{}
	return nil
}

func neighborAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Neighbor) error {
	*o = Neighbor{}
	return nil
}

func neighborAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Neighbor) error {
	*o = Neighbor{}
	return nil
}

func neighborBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Neighbor) error {
	*o = Neighbor{}
	return nil
}

func neighborAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Neighbor) error {
	*o = Neighbor{}
	return nil
}

func neighborBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Neighbor) error {
	*o = Neighbor{}
	return nil
}

func neighborAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Neighbor) error {
	*o = Neighbor{}
	return nil
}

func neighborBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Neighbor) error {
	*o = Neighbor{}
	return nil
}

func neighborAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Neighbor) error {
	*o = Neighbor{}
	return nil
}

func testNeighborsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Neighbor{}
	o := &Neighbor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, neighborDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Neighbor object: %s", err)
	}

	AddNeighborHook(boil.BeforeInsertHook, neighborBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	neighborBeforeInsertHooks = []NeighborHook{}

	AddNeighborHook(boil.AfterInsertHook, neighborAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	neighborAfterInsertHooks = []NeighborHook{}

	AddNeighborHook(boil.AfterSelectHook, neighborAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	neighborAfterSelectHooks = []NeighborHook{}

	AddNeighborHook(boil.BeforeUpdateHook, neighborBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	neighborBeforeUpdateHooks = []NeighborHook{}

	AddNeighborHook(boil.AfterUpdateHook, neighborAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	neighborAfterUpdateHooks = []NeighborHook{}

	AddNeighborHook(boil.BeforeDeleteHook, neighborBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	neighborBeforeDeleteHooks = []NeighborHook{}

	AddNeighborHook(boil.AfterDeleteHook, neighborAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	neighborAfterDeleteHooks = []NeighborHook{}

	AddNeighborHook(boil.BeforeUpsertHook, neighborBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	neighborBeforeUpsertHooks = []NeighborHook{}

	AddNeighborHook(boil.AfterUpsertHook, neighborAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	neighborAfterUpsertHooks = []NeighborHook{}
}

func testNeighborsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNeighborsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(neighborColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNeighborToOneCrawlUsingCrawl(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Neighbor
	var foreign Crawl

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, neighborDBTypes, false, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, crawlDBTypes, false, crawlColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Crawl struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CrawlID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Crawl().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := NeighborSlice{&local}
	if err = local.L.LoadCrawl(ctx, tx, false, (*[]*Neighbor)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Crawl == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Crawl = nil
	if err = local.L.LoadCrawl(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Crawl == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testNeighborToOnePeerUsingPeer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Neighbor
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, neighborDBTypes, false, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PeerID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Peer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := NeighborSlice{&local}
	if err = local.L.LoadPeer(ctx, tx, false, (*[]*Neighbor)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Peer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Peer = nil
	if err = local.L.LoadPeer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Peer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testNeighborToOneSetOpCrawlUsingCrawl(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Neighbor
	var b, c Crawl

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, neighborDBTypes, false, strmangle.SetComplement(neighborPrimaryKeyColumns, neighborColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, crawlDBTypes, false, strmangle.SetComplement(crawlPrimaryKeyColumns, crawlColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, crawlDBTypes, false, strmangle.SetComplement(crawlPrimaryKeyColumns, crawlColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Crawl{&b, &c} {
		err = a.SetCrawl(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Crawl != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Neighbors[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CrawlID != x.ID {
			t.Error("foreign key was wrong value", a.CrawlID)
		}

		if exists, err := NeighborExists(ctx, tx, a.CrawlID, a.PeerID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testNeighborToOneSetOpPeerUsingPeer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Neighbor
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, neighborDBTypes, false, strmangle.SetComplement(neighborPrimaryKeyColumns, neighborColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Peer{&b, &c} {
		err = a.SetPeer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Peer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Neighbors[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PeerID != x.ID {
			t.Error("foreign key was wrong value", a.PeerID)
		}

		if exists, err := NeighborExists(ctx, tx, a.CrawlID, a.PeerID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testNeighborsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNeighborsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NeighborSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNeighborsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Neighbors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	neighborDBTypes = map[string]string{`CrawlID`: `integer`, `PeerID`: `integer`, `NeighborIds`: `ARRAYinteger`}
	_               = bytes.MinRead
)

func testNeighborsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(neighborPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(neighborAllColumns) == len(neighborPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNeighborsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(neighborAllColumns) == len(neighborPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Neighbor{}
	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, neighborDBTypes, true, neighborPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(neighborAllColumns, neighborPrimaryKeyColumns) {
		fields = neighborAllColumns
	} else {
		fields = strmangle.SetComplement(
			neighborAllColumns,
			neighborPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := NeighborSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNeighborsUpsert(t *testing.T) {
	t.Parallel()

	if len(neighborAllColumns) == len(neighborPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Neighbor{}
	if err = randomize.Struct(seed, &o, neighborDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Neighbor: %s", err)
	}

	count, err := Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, neighborDBTypes, false, neighborPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Neighbor struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Neighbor: %s", err)
	}

	count, err = Neighbors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
