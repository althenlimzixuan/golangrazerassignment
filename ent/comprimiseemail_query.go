// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hello/ent/comprimiseemail"
	"hello/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ComprimiseEmailQuery is the builder for querying ComprimiseEmail entities.
type ComprimiseEmailQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.ComprimiseEmail
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ComprimiseEmailQuery builder.
func (ceq *ComprimiseEmailQuery) Where(ps ...predicate.ComprimiseEmail) *ComprimiseEmailQuery {
	ceq.predicates = append(ceq.predicates, ps...)
	return ceq
}

// Limit adds a limit step to the query.
func (ceq *ComprimiseEmailQuery) Limit(limit int) *ComprimiseEmailQuery {
	ceq.limit = &limit
	return ceq
}

// Offset adds an offset step to the query.
func (ceq *ComprimiseEmailQuery) Offset(offset int) *ComprimiseEmailQuery {
	ceq.offset = &offset
	return ceq
}

// Order adds an order step to the query.
func (ceq *ComprimiseEmailQuery) Order(o ...OrderFunc) *ComprimiseEmailQuery {
	ceq.order = append(ceq.order, o...)
	return ceq
}

// First returns the first ComprimiseEmail entity from the query.
// Returns a *NotFoundError when no ComprimiseEmail was found.
func (ceq *ComprimiseEmailQuery) First(ctx context.Context) (*ComprimiseEmail, error) {
	nodes, err := ceq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{comprimiseemail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ceq *ComprimiseEmailQuery) FirstX(ctx context.Context) *ComprimiseEmail {
	node, err := ceq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ComprimiseEmail ID from the query.
// Returns a *NotFoundError when no ComprimiseEmail ID was found.
func (ceq *ComprimiseEmailQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ceq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{comprimiseemail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ceq *ComprimiseEmailQuery) FirstIDX(ctx context.Context) int {
	id, err := ceq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ComprimiseEmail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ComprimiseEmail entity is not found.
// Returns a *NotFoundError when no ComprimiseEmail entities are found.
func (ceq *ComprimiseEmailQuery) Only(ctx context.Context) (*ComprimiseEmail, error) {
	nodes, err := ceq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{comprimiseemail.Label}
	default:
		return nil, &NotSingularError{comprimiseemail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ceq *ComprimiseEmailQuery) OnlyX(ctx context.Context) *ComprimiseEmail {
	node, err := ceq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ComprimiseEmail ID in the query.
// Returns a *NotSingularError when exactly one ComprimiseEmail ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ceq *ComprimiseEmailQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ceq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{comprimiseemail.Label}
	default:
		err = &NotSingularError{comprimiseemail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ceq *ComprimiseEmailQuery) OnlyIDX(ctx context.Context) int {
	id, err := ceq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ComprimiseEmails.
func (ceq *ComprimiseEmailQuery) All(ctx context.Context) ([]*ComprimiseEmail, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ceq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ceq *ComprimiseEmailQuery) AllX(ctx context.Context) []*ComprimiseEmail {
	nodes, err := ceq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ComprimiseEmail IDs.
func (ceq *ComprimiseEmailQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ceq.Select(comprimiseemail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ceq *ComprimiseEmailQuery) IDsX(ctx context.Context) []int {
	ids, err := ceq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ceq *ComprimiseEmailQuery) Count(ctx context.Context) (int, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ceq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ceq *ComprimiseEmailQuery) CountX(ctx context.Context) int {
	count, err := ceq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ceq *ComprimiseEmailQuery) Exist(ctx context.Context) (bool, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ceq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ceq *ComprimiseEmailQuery) ExistX(ctx context.Context) bool {
	exist, err := ceq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ComprimiseEmailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ceq *ComprimiseEmailQuery) Clone() *ComprimiseEmailQuery {
	if ceq == nil {
		return nil
	}
	return &ComprimiseEmailQuery{
		config:     ceq.config,
		limit:      ceq.limit,
		offset:     ceq.offset,
		order:      append([]OrderFunc{}, ceq.order...),
		predicates: append([]predicate.ComprimiseEmail{}, ceq.predicates...),
		// clone intermediate query.
		sql:  ceq.sql.Clone(),
		path: ceq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Email string `json:"Email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ComprimiseEmail.Query().
//		GroupBy(comprimiseemail.FieldEmail).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ceq *ComprimiseEmailQuery) GroupBy(field string, fields ...string) *ComprimiseEmailGroupBy {
	group := &ComprimiseEmailGroupBy{config: ceq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ceq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Email string `json:"Email,omitempty"`
//	}
//
//	client.ComprimiseEmail.Query().
//		Select(comprimiseemail.FieldEmail).
//		Scan(ctx, &v)
//
func (ceq *ComprimiseEmailQuery) Select(field string, fields ...string) *ComprimiseEmailSelect {
	ceq.fields = append([]string{field}, fields...)
	return &ComprimiseEmailSelect{ComprimiseEmailQuery: ceq}
}

func (ceq *ComprimiseEmailQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ceq.fields {
		if !comprimiseemail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ceq.path != nil {
		prev, err := ceq.path(ctx)
		if err != nil {
			return err
		}
		ceq.sql = prev
	}
	return nil
}

func (ceq *ComprimiseEmailQuery) sqlAll(ctx context.Context) ([]*ComprimiseEmail, error) {
	var (
		nodes = []*ComprimiseEmail{}
		_spec = ceq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ComprimiseEmail{config: ceq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ceq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ceq *ComprimiseEmailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ceq.querySpec()
	return sqlgraph.CountNodes(ctx, ceq.driver, _spec)
}

func (ceq *ComprimiseEmailQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ceq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ceq *ComprimiseEmailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comprimiseemail.Table,
			Columns: comprimiseemail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comprimiseemail.FieldID,
			},
		},
		From:   ceq.sql,
		Unique: true,
	}
	if fields := ceq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comprimiseemail.FieldID)
		for i := range fields {
			if fields[i] != comprimiseemail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ceq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ceq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ceq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ceq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, comprimiseemail.ValidColumn)
			}
		}
	}
	return _spec
}

func (ceq *ComprimiseEmailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ceq.driver.Dialect())
	t1 := builder.Table(comprimiseemail.Table)
	selector := builder.Select(t1.Columns(comprimiseemail.Columns...)...).From(t1)
	if ceq.sql != nil {
		selector = ceq.sql
		selector.Select(selector.Columns(comprimiseemail.Columns...)...)
	}
	for _, p := range ceq.predicates {
		p(selector)
	}
	for _, p := range ceq.order {
		p(selector, comprimiseemail.ValidColumn)
	}
	if offset := ceq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ceq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ComprimiseEmailGroupBy is the group-by builder for ComprimiseEmail entities.
type ComprimiseEmailGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cegb *ComprimiseEmailGroupBy) Aggregate(fns ...AggregateFunc) *ComprimiseEmailGroupBy {
	cegb.fns = append(cegb.fns, fns...)
	return cegb
}

// Scan applies the group-by query and scans the result into the given value.
func (cegb *ComprimiseEmailGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cegb.path(ctx)
	if err != nil {
		return err
	}
	cegb.sql = query
	return cegb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cegb *ComprimiseEmailGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cegb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cegb *ComprimiseEmailGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: ComprimiseEmailGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cegb *ComprimiseEmailGroupBy) StringsX(ctx context.Context) []string {
	v, err := cegb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cegb *ComprimiseEmailGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cegb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{comprimiseemail.Label}
	default:
		err = fmt.Errorf("ent: ComprimiseEmailGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cegb *ComprimiseEmailGroupBy) StringX(ctx context.Context) string {
	v, err := cegb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cegb *ComprimiseEmailGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: ComprimiseEmailGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cegb *ComprimiseEmailGroupBy) IntsX(ctx context.Context) []int {
	v, err := cegb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cegb *ComprimiseEmailGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cegb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{comprimiseemail.Label}
	default:
		err = fmt.Errorf("ent: ComprimiseEmailGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cegb *ComprimiseEmailGroupBy) IntX(ctx context.Context) int {
	v, err := cegb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cegb *ComprimiseEmailGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: ComprimiseEmailGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cegb *ComprimiseEmailGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cegb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cegb *ComprimiseEmailGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cegb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{comprimiseemail.Label}
	default:
		err = fmt.Errorf("ent: ComprimiseEmailGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cegb *ComprimiseEmailGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cegb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cegb *ComprimiseEmailGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: ComprimiseEmailGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cegb *ComprimiseEmailGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cegb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cegb *ComprimiseEmailGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cegb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{comprimiseemail.Label}
	default:
		err = fmt.Errorf("ent: ComprimiseEmailGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cegb *ComprimiseEmailGroupBy) BoolX(ctx context.Context) bool {
	v, err := cegb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cegb *ComprimiseEmailGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cegb.fields {
		if !comprimiseemail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cegb *ComprimiseEmailGroupBy) sqlQuery() *sql.Selector {
	selector := cegb.sql
	columns := make([]string, 0, len(cegb.fields)+len(cegb.fns))
	columns = append(columns, cegb.fields...)
	for _, fn := range cegb.fns {
		columns = append(columns, fn(selector, comprimiseemail.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(cegb.fields...)
}

// ComprimiseEmailSelect is the builder for selecting fields of ComprimiseEmail entities.
type ComprimiseEmailSelect struct {
	*ComprimiseEmailQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ces *ComprimiseEmailSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ces.prepareQuery(ctx); err != nil {
		return err
	}
	ces.sql = ces.ComprimiseEmailQuery.sqlQuery(ctx)
	return ces.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ces *ComprimiseEmailSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ces.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ces *ComprimiseEmailSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: ComprimiseEmailSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ces *ComprimiseEmailSelect) StringsX(ctx context.Context) []string {
	v, err := ces.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ces *ComprimiseEmailSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ces.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{comprimiseemail.Label}
	default:
		err = fmt.Errorf("ent: ComprimiseEmailSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ces *ComprimiseEmailSelect) StringX(ctx context.Context) string {
	v, err := ces.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ces *ComprimiseEmailSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: ComprimiseEmailSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ces *ComprimiseEmailSelect) IntsX(ctx context.Context) []int {
	v, err := ces.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ces *ComprimiseEmailSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ces.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{comprimiseemail.Label}
	default:
		err = fmt.Errorf("ent: ComprimiseEmailSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ces *ComprimiseEmailSelect) IntX(ctx context.Context) int {
	v, err := ces.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ces *ComprimiseEmailSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: ComprimiseEmailSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ces *ComprimiseEmailSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ces.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ces *ComprimiseEmailSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ces.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{comprimiseemail.Label}
	default:
		err = fmt.Errorf("ent: ComprimiseEmailSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ces *ComprimiseEmailSelect) Float64X(ctx context.Context) float64 {
	v, err := ces.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ces *ComprimiseEmailSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: ComprimiseEmailSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ces *ComprimiseEmailSelect) BoolsX(ctx context.Context) []bool {
	v, err := ces.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ces *ComprimiseEmailSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ces.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{comprimiseemail.Label}
	default:
		err = fmt.Errorf("ent: ComprimiseEmailSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ces *ComprimiseEmailSelect) BoolX(ctx context.Context) bool {
	v, err := ces.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ces *ComprimiseEmailSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ces.sqlQuery().Query()
	if err := ces.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ces *ComprimiseEmailSelect) sqlQuery() sql.Querier {
	selector := ces.sql
	selector.Select(selector.Columns(ces.fields...)...)
	return selector
}
