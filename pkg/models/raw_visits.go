// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// RawVisit is an object representing the database table.
type RawVisit struct {
	ID              int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	CrawlID         null.Int          `boil:"crawl_id" json:"crawl_id,omitempty" toml:"crawl_id" yaml:"crawl_id,omitempty"`
	VisitStartedAt  time.Time         `boil:"visit_started_at" json:"visit_started_at" toml:"visit_started_at" yaml:"visit_started_at"`
	VisitEndedAt    time.Time         `boil:"visit_ended_at" json:"visit_ended_at" toml:"visit_ended_at" yaml:"visit_ended_at"`
	DialDuration    null.String       `boil:"dial_duration" json:"dial_duration,omitempty" toml:"dial_duration" yaml:"dial_duration,omitempty"`
	ConnectDuration null.String       `boil:"connect_duration" json:"connect_duration,omitempty" toml:"connect_duration" yaml:"connect_duration,omitempty"`
	CrawlDuration   null.String       `boil:"crawl_duration" json:"crawl_duration,omitempty" toml:"crawl_duration" yaml:"crawl_duration,omitempty"`
	Type            string            `boil:"type" json:"type" toml:"type" yaml:"type"`
	AgentVersion    null.String       `boil:"agent_version" json:"agent_version,omitempty" toml:"agent_version" yaml:"agent_version,omitempty"`
	PeerMultiHash   string            `boil:"peer_multi_hash" json:"peer_multi_hash" toml:"peer_multi_hash" yaml:"peer_multi_hash"`
	Protocols       types.StringArray `boil:"protocols" json:"protocols,omitempty" toml:"protocols" yaml:"protocols,omitempty"`
	MultiAddresses  types.StringArray `boil:"multi_addresses" json:"multi_addresses,omitempty" toml:"multi_addresses" yaml:"multi_addresses,omitempty"`
	Error           null.String       `boil:"error" json:"error,omitempty" toml:"error" yaml:"error,omitempty"`
	ErrorMessage    null.String       `boil:"error_message" json:"error_message,omitempty" toml:"error_message" yaml:"error_message,omitempty"`
	CreatedAt       time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	AgentVersionID  null.Int          `boil:"agent_version_id" json:"agent_version_id,omitempty" toml:"agent_version_id" yaml:"agent_version_id,omitempty"`
	ProtocolIds     types.Int64Array  `boil:"protocol_ids" json:"protocol_ids,omitempty" toml:"protocol_ids" yaml:"protocol_ids,omitempty"`

	R *rawVisitR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L rawVisitL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RawVisitColumns = struct {
	ID              string
	CrawlID         string
	VisitStartedAt  string
	VisitEndedAt    string
	DialDuration    string
	ConnectDuration string
	CrawlDuration   string
	Type            string
	AgentVersion    string
	PeerMultiHash   string
	Protocols       string
	MultiAddresses  string
	Error           string
	ErrorMessage    string
	CreatedAt       string
	AgentVersionID  string
	ProtocolIds     string
}{
	ID:              "id",
	CrawlID:         "crawl_id",
	VisitStartedAt:  "visit_started_at",
	VisitEndedAt:    "visit_ended_at",
	DialDuration:    "dial_duration",
	ConnectDuration: "connect_duration",
	CrawlDuration:   "crawl_duration",
	Type:            "type",
	AgentVersion:    "agent_version",
	PeerMultiHash:   "peer_multi_hash",
	Protocols:       "protocols",
	MultiAddresses:  "multi_addresses",
	Error:           "error",
	ErrorMessage:    "error_message",
	CreatedAt:       "created_at",
	AgentVersionID:  "agent_version_id",
	ProtocolIds:     "protocol_ids",
}

var RawVisitTableColumns = struct {
	ID              string
	CrawlID         string
	VisitStartedAt  string
	VisitEndedAt    string
	DialDuration    string
	ConnectDuration string
	CrawlDuration   string
	Type            string
	AgentVersion    string
	PeerMultiHash   string
	Protocols       string
	MultiAddresses  string
	Error           string
	ErrorMessage    string
	CreatedAt       string
	AgentVersionID  string
	ProtocolIds     string
}{
	ID:              "raw_visits.id",
	CrawlID:         "raw_visits.crawl_id",
	VisitStartedAt:  "raw_visits.visit_started_at",
	VisitEndedAt:    "raw_visits.visit_ended_at",
	DialDuration:    "raw_visits.dial_duration",
	ConnectDuration: "raw_visits.connect_duration",
	CrawlDuration:   "raw_visits.crawl_duration",
	Type:            "raw_visits.type",
	AgentVersion:    "raw_visits.agent_version",
	PeerMultiHash:   "raw_visits.peer_multi_hash",
	Protocols:       "raw_visits.protocols",
	MultiAddresses:  "raw_visits.multi_addresses",
	Error:           "raw_visits.error",
	ErrorMessage:    "raw_visits.error_message",
	CreatedAt:       "raw_visits.created_at",
	AgentVersionID:  "raw_visits.agent_version_id",
	ProtocolIds:     "raw_visits.protocol_ids",
}

// Generated where

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_StringArray) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_StringArray) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var RawVisitWhere = struct {
	ID              whereHelperint
	CrawlID         whereHelpernull_Int
	VisitStartedAt  whereHelpertime_Time
	VisitEndedAt    whereHelpertime_Time
	DialDuration    whereHelpernull_String
	ConnectDuration whereHelpernull_String
	CrawlDuration   whereHelpernull_String
	Type            whereHelperstring
	AgentVersion    whereHelpernull_String
	PeerMultiHash   whereHelperstring
	Protocols       whereHelpertypes_StringArray
	MultiAddresses  whereHelpertypes_StringArray
	Error           whereHelpernull_String
	ErrorMessage    whereHelpernull_String
	CreatedAt       whereHelpertime_Time
	AgentVersionID  whereHelpernull_Int
	ProtocolIds     whereHelpertypes_Int64Array
}{
	ID:              whereHelperint{field: "\"raw_visits\".\"id\""},
	CrawlID:         whereHelpernull_Int{field: "\"raw_visits\".\"crawl_id\""},
	VisitStartedAt:  whereHelpertime_Time{field: "\"raw_visits\".\"visit_started_at\""},
	VisitEndedAt:    whereHelpertime_Time{field: "\"raw_visits\".\"visit_ended_at\""},
	DialDuration:    whereHelpernull_String{field: "\"raw_visits\".\"dial_duration\""},
	ConnectDuration: whereHelpernull_String{field: "\"raw_visits\".\"connect_duration\""},
	CrawlDuration:   whereHelpernull_String{field: "\"raw_visits\".\"crawl_duration\""},
	Type:            whereHelperstring{field: "\"raw_visits\".\"type\""},
	AgentVersion:    whereHelpernull_String{field: "\"raw_visits\".\"agent_version\""},
	PeerMultiHash:   whereHelperstring{field: "\"raw_visits\".\"peer_multi_hash\""},
	Protocols:       whereHelpertypes_StringArray{field: "\"raw_visits\".\"protocols\""},
	MultiAddresses:  whereHelpertypes_StringArray{field: "\"raw_visits\".\"multi_addresses\""},
	Error:           whereHelpernull_String{field: "\"raw_visits\".\"error\""},
	ErrorMessage:    whereHelpernull_String{field: "\"raw_visits\".\"error_message\""},
	CreatedAt:       whereHelpertime_Time{field: "\"raw_visits\".\"created_at\""},
	AgentVersionID:  whereHelpernull_Int{field: "\"raw_visits\".\"agent_version_id\""},
	ProtocolIds:     whereHelpertypes_Int64Array{field: "\"raw_visits\".\"protocol_ids\""},
}

// RawVisitRels is where relationship names are stored.
var RawVisitRels = struct {
}{}

// rawVisitR is where relationships are stored.
type rawVisitR struct {
}

// NewStruct creates a new relationship struct
func (*rawVisitR) NewStruct() *rawVisitR {
	return &rawVisitR{}
}

// rawVisitL is where Load methods for each relationship are stored.
type rawVisitL struct{}

var (
	rawVisitAllColumns            = []string{"id", "crawl_id", "visit_started_at", "visit_ended_at", "dial_duration", "connect_duration", "crawl_duration", "type", "agent_version", "peer_multi_hash", "protocols", "multi_addresses", "error", "error_message", "created_at", "agent_version_id", "protocol_ids"}
	rawVisitColumnsWithoutDefault = []string{"crawl_id", "visit_started_at", "visit_ended_at", "dial_duration", "connect_duration", "crawl_duration", "type", "agent_version", "peer_multi_hash", "protocols", "multi_addresses", "error", "error_message", "created_at", "agent_version_id", "protocol_ids"}
	rawVisitColumnsWithDefault    = []string{"id"}
	rawVisitPrimaryKeyColumns     = []string{"id"}
)

type (
	// RawVisitSlice is an alias for a slice of pointers to RawVisit.
	// This should almost always be used instead of []RawVisit.
	RawVisitSlice []*RawVisit
	// RawVisitHook is the signature for custom RawVisit hook methods
	RawVisitHook func(context.Context, boil.ContextExecutor, *RawVisit) error

	rawVisitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	rawVisitType                 = reflect.TypeOf(&RawVisit{})
	rawVisitMapping              = queries.MakeStructMapping(rawVisitType)
	rawVisitPrimaryKeyMapping, _ = queries.BindMapping(rawVisitType, rawVisitMapping, rawVisitPrimaryKeyColumns)
	rawVisitInsertCacheMut       sync.RWMutex
	rawVisitInsertCache          = make(map[string]insertCache)
	rawVisitUpdateCacheMut       sync.RWMutex
	rawVisitUpdateCache          = make(map[string]updateCache)
	rawVisitUpsertCacheMut       sync.RWMutex
	rawVisitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var rawVisitBeforeInsertHooks []RawVisitHook
var rawVisitBeforeUpdateHooks []RawVisitHook
var rawVisitBeforeDeleteHooks []RawVisitHook
var rawVisitBeforeUpsertHooks []RawVisitHook

var rawVisitAfterInsertHooks []RawVisitHook
var rawVisitAfterSelectHooks []RawVisitHook
var rawVisitAfterUpdateHooks []RawVisitHook
var rawVisitAfterDeleteHooks []RawVisitHook
var rawVisitAfterUpsertHooks []RawVisitHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RawVisit) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawVisitBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RawVisit) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawVisitBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RawVisit) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawVisitBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RawVisit) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawVisitBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RawVisit) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawVisitAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RawVisit) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawVisitAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RawVisit) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawVisitAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RawVisit) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawVisitAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RawVisit) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rawVisitAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRawVisitHook registers your hook function for all future operations.
func AddRawVisitHook(hookPoint boil.HookPoint, rawVisitHook RawVisitHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		rawVisitBeforeInsertHooks = append(rawVisitBeforeInsertHooks, rawVisitHook)
	case boil.BeforeUpdateHook:
		rawVisitBeforeUpdateHooks = append(rawVisitBeforeUpdateHooks, rawVisitHook)
	case boil.BeforeDeleteHook:
		rawVisitBeforeDeleteHooks = append(rawVisitBeforeDeleteHooks, rawVisitHook)
	case boil.BeforeUpsertHook:
		rawVisitBeforeUpsertHooks = append(rawVisitBeforeUpsertHooks, rawVisitHook)
	case boil.AfterInsertHook:
		rawVisitAfterInsertHooks = append(rawVisitAfterInsertHooks, rawVisitHook)
	case boil.AfterSelectHook:
		rawVisitAfterSelectHooks = append(rawVisitAfterSelectHooks, rawVisitHook)
	case boil.AfterUpdateHook:
		rawVisitAfterUpdateHooks = append(rawVisitAfterUpdateHooks, rawVisitHook)
	case boil.AfterDeleteHook:
		rawVisitAfterDeleteHooks = append(rawVisitAfterDeleteHooks, rawVisitHook)
	case boil.AfterUpsertHook:
		rawVisitAfterUpsertHooks = append(rawVisitAfterUpsertHooks, rawVisitHook)
	}
}

// One returns a single rawVisit record from the query.
func (q rawVisitQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RawVisit, error) {
	o := &RawVisit{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for raw_visits")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RawVisit records from the query.
func (q rawVisitQuery) All(ctx context.Context, exec boil.ContextExecutor) (RawVisitSlice, error) {
	var o []*RawVisit

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RawVisit slice")
	}

	if len(rawVisitAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RawVisit records in the query.
func (q rawVisitQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count raw_visits rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q rawVisitQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if raw_visits exists")
	}

	return count > 0, nil
}

// RawVisits retrieves all the records using an executor.
func RawVisits(mods ...qm.QueryMod) rawVisitQuery {
	mods = append(mods, qm.From("\"raw_visits\""))
	return rawVisitQuery{NewQuery(mods...)}
}

// FindRawVisit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRawVisit(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RawVisit, error) {
	rawVisitObj := &RawVisit{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"raw_visits\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, rawVisitObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from raw_visits")
	}

	if err = rawVisitObj.doAfterSelectHooks(ctx, exec); err != nil {
		return rawVisitObj, err
	}

	return rawVisitObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RawVisit) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no raw_visits provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rawVisitColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	rawVisitInsertCacheMut.RLock()
	cache, cached := rawVisitInsertCache[key]
	rawVisitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			rawVisitAllColumns,
			rawVisitColumnsWithDefault,
			rawVisitColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(rawVisitType, rawVisitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(rawVisitType, rawVisitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"raw_visits\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"raw_visits\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into raw_visits")
	}

	if !cached {
		rawVisitInsertCacheMut.Lock()
		rawVisitInsertCache[key] = cache
		rawVisitInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RawVisit.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RawVisit) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	rawVisitUpdateCacheMut.RLock()
	cache, cached := rawVisitUpdateCache[key]
	rawVisitUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			rawVisitAllColumns,
			rawVisitPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update raw_visits, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"raw_visits\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rawVisitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(rawVisitType, rawVisitMapping, append(wl, rawVisitPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update raw_visits row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for raw_visits")
	}

	if !cached {
		rawVisitUpdateCacheMut.Lock()
		rawVisitUpdateCache[key] = cache
		rawVisitUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q rawVisitQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for raw_visits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for raw_visits")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RawVisitSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rawVisitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"raw_visits\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rawVisitPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in rawVisit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all rawVisit")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RawVisit) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no raw_visits provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rawVisitColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	rawVisitUpsertCacheMut.RLock()
	cache, cached := rawVisitUpsertCache[key]
	rawVisitUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			rawVisitAllColumns,
			rawVisitColumnsWithDefault,
			rawVisitColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			rawVisitAllColumns,
			rawVisitPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert raw_visits, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rawVisitPrimaryKeyColumns))
			copy(conflict, rawVisitPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"raw_visits\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(rawVisitType, rawVisitMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(rawVisitType, rawVisitMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert raw_visits")
	}

	if !cached {
		rawVisitUpsertCacheMut.Lock()
		rawVisitUpsertCache[key] = cache
		rawVisitUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RawVisit record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RawVisit) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RawVisit provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rawVisitPrimaryKeyMapping)
	sql := "DELETE FROM \"raw_visits\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from raw_visits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for raw_visits")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q rawVisitQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no rawVisitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from raw_visits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for raw_visits")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RawVisitSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(rawVisitBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rawVisitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"raw_visits\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rawVisitPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rawVisit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for raw_visits")
	}

	if len(rawVisitAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RawVisit) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRawVisit(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RawVisitSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RawVisitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rawVisitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"raw_visits\".* FROM \"raw_visits\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rawVisitPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RawVisitSlice")
	}

	*o = slice

	return nil
}

// RawVisitExists checks if the RawVisit row exists.
func RawVisitExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"raw_visits\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if raw_visits exists")
	}

	return exists, nil
}
