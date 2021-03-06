package libsql

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"database/sql"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// SqlStmtMock implements sqlStmt
type SqlStmtMock struct {
	t minimock.Tester

	funcClose          func() (err error)
	inspectFuncClose   func()
	afterCloseCounter  uint64
	beforeCloseCounter uint64
	CloseMock          mSqlStmtMockClose

	funcExec          func(ctx context.Context, args ...interface{}) (r1 sql.Result, err error)
	inspectFuncExec   func(ctx context.Context, args ...interface{})
	afterExecCounter  uint64
	beforeExecCounter uint64
	ExecMock          mSqlStmtMockExec

	funcQuery          func(ctx context.Context, args ...interface{}) (s1 sqlRows, err error)
	inspectFuncQuery   func(ctx context.Context, args ...interface{})
	afterQueryCounter  uint64
	beforeQueryCounter uint64
	QueryMock          mSqlStmtMockQuery
}

// NewSqlStmtMock returns a mock for sqlStmt
func NewSqlStmtMock(t minimock.Tester) *SqlStmtMock {
	m := &SqlStmtMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mSqlStmtMockClose{mock: m}

	m.ExecMock = mSqlStmtMockExec{mock: m}
	m.ExecMock.callArgs = []*SqlStmtMockExecParams{}

	m.QueryMock = mSqlStmtMockQuery{mock: m}
	m.QueryMock.callArgs = []*SqlStmtMockQueryParams{}

	return m
}

type mSqlStmtMockClose struct {
	mock               *SqlStmtMock
	defaultExpectation *SqlStmtMockCloseExpectation
	expectations       []*SqlStmtMockCloseExpectation
}

// SqlStmtMockCloseExpectation specifies expectation struct of the sqlStmt.Close
type SqlStmtMockCloseExpectation struct {
	mock *SqlStmtMock

	results *SqlStmtMockCloseResults
	Counter uint64
}

// SqlStmtMockCloseResults contains results of the sqlStmt.Close
type SqlStmtMockCloseResults struct {
	err error
}

// Expect sets up expected params for sqlStmt.Close
func (mmClose *mSqlStmtMockClose) Expect() *mSqlStmtMockClose {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("SqlStmtMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &SqlStmtMockCloseExpectation{}
	}

	return mmClose
}

// Inspect accepts an inspector function that has same arguments as the sqlStmt.Close
func (mmClose *mSqlStmtMockClose) Inspect(f func()) *mSqlStmtMockClose {
	if mmClose.mock.inspectFuncClose != nil {
		mmClose.mock.t.Fatalf("Inspect function is already set for SqlStmtMock.Close")
	}

	mmClose.mock.inspectFuncClose = f

	return mmClose
}

// Return sets up results that will be returned by sqlStmt.Close
func (mmClose *mSqlStmtMockClose) Return(err error) *SqlStmtMock {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("SqlStmtMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &SqlStmtMockCloseExpectation{mock: mmClose.mock}
	}
	mmClose.defaultExpectation.results = &SqlStmtMockCloseResults{err}
	return mmClose.mock
}

//Set uses given function f to mock the sqlStmt.Close method
func (mmClose *mSqlStmtMockClose) Set(f func() (err error)) *SqlStmtMock {
	if mmClose.defaultExpectation != nil {
		mmClose.mock.t.Fatalf("Default expectation is already set for the sqlStmt.Close method")
	}

	if len(mmClose.expectations) > 0 {
		mmClose.mock.t.Fatalf("Some expectations are already set for the sqlStmt.Close method")
	}

	mmClose.mock.funcClose = f
	return mmClose.mock
}

// Close implements sqlStmt
func (mmClose *SqlStmtMock) Close() (err error) {
	mm_atomic.AddUint64(&mmClose.beforeCloseCounter, 1)
	defer mm_atomic.AddUint64(&mmClose.afterCloseCounter, 1)

	if mmClose.inspectFuncClose != nil {
		mmClose.inspectFuncClose()
	}

	if mmClose.CloseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClose.CloseMock.defaultExpectation.Counter, 1)

		mm_results := mmClose.CloseMock.defaultExpectation.results
		if mm_results == nil {
			mmClose.t.Fatal("No results are set for the SqlStmtMock.Close")
		}
		return (*mm_results).err
	}
	if mmClose.funcClose != nil {
		return mmClose.funcClose()
	}
	mmClose.t.Fatalf("Unexpected call to SqlStmtMock.Close.")
	return
}

// CloseAfterCounter returns a count of finished SqlStmtMock.Close invocations
func (mmClose *SqlStmtMock) CloseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.afterCloseCounter)
}

// CloseBeforeCounter returns a count of SqlStmtMock.Close invocations
func (mmClose *SqlStmtMock) CloseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.beforeCloseCounter)
}

// MinimockCloseDone returns true if the count of the Close invocations corresponds
// the number of defined expectations
func (m *SqlStmtMock) MinimockCloseDone() bool {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	return true
}

// MinimockCloseInspect logs each unmet expectation
func (m *SqlStmtMock) MinimockCloseInspect() {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to SqlStmtMock.Close")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to SqlStmtMock.Close")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to SqlStmtMock.Close")
	}
}

type mSqlStmtMockExec struct {
	mock               *SqlStmtMock
	defaultExpectation *SqlStmtMockExecExpectation
	expectations       []*SqlStmtMockExecExpectation

	callArgs []*SqlStmtMockExecParams
	mutex    sync.RWMutex
}

// SqlStmtMockExecExpectation specifies expectation struct of the sqlStmt.Exec
type SqlStmtMockExecExpectation struct {
	mock    *SqlStmtMock
	params  *SqlStmtMockExecParams
	results *SqlStmtMockExecResults
	Counter uint64
}

// SqlStmtMockExecParams contains parameters of the sqlStmt.Exec
type SqlStmtMockExecParams struct {
	ctx  context.Context
	args []interface{}
}

// SqlStmtMockExecResults contains results of the sqlStmt.Exec
type SqlStmtMockExecResults struct {
	r1  sql.Result
	err error
}

// Expect sets up expected params for sqlStmt.Exec
func (mmExec *mSqlStmtMockExec) Expect(ctx context.Context, args ...interface{}) *mSqlStmtMockExec {
	if mmExec.mock.funcExec != nil {
		mmExec.mock.t.Fatalf("SqlStmtMock.Exec mock is already set by Set")
	}

	if mmExec.defaultExpectation == nil {
		mmExec.defaultExpectation = &SqlStmtMockExecExpectation{}
	}

	mmExec.defaultExpectation.params = &SqlStmtMockExecParams{ctx, args}
	for _, e := range mmExec.expectations {
		if minimock.Equal(e.params, mmExec.defaultExpectation.params) {
			mmExec.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmExec.defaultExpectation.params)
		}
	}

	return mmExec
}

// Inspect accepts an inspector function that has same arguments as the sqlStmt.Exec
func (mmExec *mSqlStmtMockExec) Inspect(f func(ctx context.Context, args ...interface{})) *mSqlStmtMockExec {
	if mmExec.mock.inspectFuncExec != nil {
		mmExec.mock.t.Fatalf("Inspect function is already set for SqlStmtMock.Exec")
	}

	mmExec.mock.inspectFuncExec = f

	return mmExec
}

// Return sets up results that will be returned by sqlStmt.Exec
func (mmExec *mSqlStmtMockExec) Return(r1 sql.Result, err error) *SqlStmtMock {
	if mmExec.mock.funcExec != nil {
		mmExec.mock.t.Fatalf("SqlStmtMock.Exec mock is already set by Set")
	}

	if mmExec.defaultExpectation == nil {
		mmExec.defaultExpectation = &SqlStmtMockExecExpectation{mock: mmExec.mock}
	}
	mmExec.defaultExpectation.results = &SqlStmtMockExecResults{r1, err}
	return mmExec.mock
}

//Set uses given function f to mock the sqlStmt.Exec method
func (mmExec *mSqlStmtMockExec) Set(f func(ctx context.Context, args ...interface{}) (r1 sql.Result, err error)) *SqlStmtMock {
	if mmExec.defaultExpectation != nil {
		mmExec.mock.t.Fatalf("Default expectation is already set for the sqlStmt.Exec method")
	}

	if len(mmExec.expectations) > 0 {
		mmExec.mock.t.Fatalf("Some expectations are already set for the sqlStmt.Exec method")
	}

	mmExec.mock.funcExec = f
	return mmExec.mock
}

// When sets expectation for the sqlStmt.Exec which will trigger the result defined by the following
// Then helper
func (mmExec *mSqlStmtMockExec) When(ctx context.Context, args ...interface{}) *SqlStmtMockExecExpectation {
	if mmExec.mock.funcExec != nil {
		mmExec.mock.t.Fatalf("SqlStmtMock.Exec mock is already set by Set")
	}

	expectation := &SqlStmtMockExecExpectation{
		mock:   mmExec.mock,
		params: &SqlStmtMockExecParams{ctx, args},
	}
	mmExec.expectations = append(mmExec.expectations, expectation)
	return expectation
}

// Then sets up sqlStmt.Exec return parameters for the expectation previously defined by the When method
func (e *SqlStmtMockExecExpectation) Then(r1 sql.Result, err error) *SqlStmtMock {
	e.results = &SqlStmtMockExecResults{r1, err}
	return e.mock
}

// Exec implements sqlStmt
func (mmExec *SqlStmtMock) Exec(ctx context.Context, args ...interface{}) (r1 sql.Result, err error) {
	mm_atomic.AddUint64(&mmExec.beforeExecCounter, 1)
	defer mm_atomic.AddUint64(&mmExec.afterExecCounter, 1)

	if mmExec.inspectFuncExec != nil {
		mmExec.inspectFuncExec(ctx, args...)
	}

	mm_params := &SqlStmtMockExecParams{ctx, args}

	// Record call args
	mmExec.ExecMock.mutex.Lock()
	mmExec.ExecMock.callArgs = append(mmExec.ExecMock.callArgs, mm_params)
	mmExec.ExecMock.mutex.Unlock()

	for _, e := range mmExec.ExecMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.err
		}
	}

	if mmExec.ExecMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmExec.ExecMock.defaultExpectation.Counter, 1)
		mm_want := mmExec.ExecMock.defaultExpectation.params
		mm_got := SqlStmtMockExecParams{ctx, args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmExec.t.Errorf("SqlStmtMock.Exec got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmExec.ExecMock.defaultExpectation.results
		if mm_results == nil {
			mmExec.t.Fatal("No results are set for the SqlStmtMock.Exec")
		}
		return (*mm_results).r1, (*mm_results).err
	}
	if mmExec.funcExec != nil {
		return mmExec.funcExec(ctx, args...)
	}
	mmExec.t.Fatalf("Unexpected call to SqlStmtMock.Exec. %v %v", ctx, args)
	return
}

// ExecAfterCounter returns a count of finished SqlStmtMock.Exec invocations
func (mmExec *SqlStmtMock) ExecAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExec.afterExecCounter)
}

// ExecBeforeCounter returns a count of SqlStmtMock.Exec invocations
func (mmExec *SqlStmtMock) ExecBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExec.beforeExecCounter)
}

// Calls returns a list of arguments used in each call to SqlStmtMock.Exec.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmExec *mSqlStmtMockExec) Calls() []*SqlStmtMockExecParams {
	mmExec.mutex.RLock()

	argCopy := make([]*SqlStmtMockExecParams, len(mmExec.callArgs))
	copy(argCopy, mmExec.callArgs)

	mmExec.mutex.RUnlock()

	return argCopy
}

// MinimockExecDone returns true if the count of the Exec invocations corresponds
// the number of defined expectations
func (m *SqlStmtMock) MinimockExecDone() bool {
	for _, e := range m.ExecMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExec != nil && mm_atomic.LoadUint64(&m.afterExecCounter) < 1 {
		return false
	}
	return true
}

// MinimockExecInspect logs each unmet expectation
func (m *SqlStmtMock) MinimockExecInspect() {
	for _, e := range m.ExecMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SqlStmtMock.Exec with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecCounter) < 1 {
		if m.ExecMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SqlStmtMock.Exec")
		} else {
			m.t.Errorf("Expected call to SqlStmtMock.Exec with params: %#v", *m.ExecMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExec != nil && mm_atomic.LoadUint64(&m.afterExecCounter) < 1 {
		m.t.Error("Expected call to SqlStmtMock.Exec")
	}
}

type mSqlStmtMockQuery struct {
	mock               *SqlStmtMock
	defaultExpectation *SqlStmtMockQueryExpectation
	expectations       []*SqlStmtMockQueryExpectation

	callArgs []*SqlStmtMockQueryParams
	mutex    sync.RWMutex
}

// SqlStmtMockQueryExpectation specifies expectation struct of the sqlStmt.Query
type SqlStmtMockQueryExpectation struct {
	mock    *SqlStmtMock
	params  *SqlStmtMockQueryParams
	results *SqlStmtMockQueryResults
	Counter uint64
}

// SqlStmtMockQueryParams contains parameters of the sqlStmt.Query
type SqlStmtMockQueryParams struct {
	ctx  context.Context
	args []interface{}
}

// SqlStmtMockQueryResults contains results of the sqlStmt.Query
type SqlStmtMockQueryResults struct {
	s1  sqlRows
	err error
}

// Expect sets up expected params for sqlStmt.Query
func (mmQuery *mSqlStmtMockQuery) Expect(ctx context.Context, args ...interface{}) *mSqlStmtMockQuery {
	if mmQuery.mock.funcQuery != nil {
		mmQuery.mock.t.Fatalf("SqlStmtMock.Query mock is already set by Set")
	}

	if mmQuery.defaultExpectation == nil {
		mmQuery.defaultExpectation = &SqlStmtMockQueryExpectation{}
	}

	mmQuery.defaultExpectation.params = &SqlStmtMockQueryParams{ctx, args}
	for _, e := range mmQuery.expectations {
		if minimock.Equal(e.params, mmQuery.defaultExpectation.params) {
			mmQuery.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmQuery.defaultExpectation.params)
		}
	}

	return mmQuery
}

// Inspect accepts an inspector function that has same arguments as the sqlStmt.Query
func (mmQuery *mSqlStmtMockQuery) Inspect(f func(ctx context.Context, args ...interface{})) *mSqlStmtMockQuery {
	if mmQuery.mock.inspectFuncQuery != nil {
		mmQuery.mock.t.Fatalf("Inspect function is already set for SqlStmtMock.Query")
	}

	mmQuery.mock.inspectFuncQuery = f

	return mmQuery
}

// Return sets up results that will be returned by sqlStmt.Query
func (mmQuery *mSqlStmtMockQuery) Return(s1 sqlRows, err error) *SqlStmtMock {
	if mmQuery.mock.funcQuery != nil {
		mmQuery.mock.t.Fatalf("SqlStmtMock.Query mock is already set by Set")
	}

	if mmQuery.defaultExpectation == nil {
		mmQuery.defaultExpectation = &SqlStmtMockQueryExpectation{mock: mmQuery.mock}
	}
	mmQuery.defaultExpectation.results = &SqlStmtMockQueryResults{s1, err}
	return mmQuery.mock
}

//Set uses given function f to mock the sqlStmt.Query method
func (mmQuery *mSqlStmtMockQuery) Set(f func(ctx context.Context, args ...interface{}) (s1 sqlRows, err error)) *SqlStmtMock {
	if mmQuery.defaultExpectation != nil {
		mmQuery.mock.t.Fatalf("Default expectation is already set for the sqlStmt.Query method")
	}

	if len(mmQuery.expectations) > 0 {
		mmQuery.mock.t.Fatalf("Some expectations are already set for the sqlStmt.Query method")
	}

	mmQuery.mock.funcQuery = f
	return mmQuery.mock
}

// When sets expectation for the sqlStmt.Query which will trigger the result defined by the following
// Then helper
func (mmQuery *mSqlStmtMockQuery) When(ctx context.Context, args ...interface{}) *SqlStmtMockQueryExpectation {
	if mmQuery.mock.funcQuery != nil {
		mmQuery.mock.t.Fatalf("SqlStmtMock.Query mock is already set by Set")
	}

	expectation := &SqlStmtMockQueryExpectation{
		mock:   mmQuery.mock,
		params: &SqlStmtMockQueryParams{ctx, args},
	}
	mmQuery.expectations = append(mmQuery.expectations, expectation)
	return expectation
}

// Then sets up sqlStmt.Query return parameters for the expectation previously defined by the When method
func (e *SqlStmtMockQueryExpectation) Then(s1 sqlRows, err error) *SqlStmtMock {
	e.results = &SqlStmtMockQueryResults{s1, err}
	return e.mock
}

// Query implements sqlStmt
func (mmQuery *SqlStmtMock) Query(ctx context.Context, args ...interface{}) (s1 sqlRows, err error) {
	mm_atomic.AddUint64(&mmQuery.beforeQueryCounter, 1)
	defer mm_atomic.AddUint64(&mmQuery.afterQueryCounter, 1)

	if mmQuery.inspectFuncQuery != nil {
		mmQuery.inspectFuncQuery(ctx, args...)
	}

	mm_params := &SqlStmtMockQueryParams{ctx, args}

	// Record call args
	mmQuery.QueryMock.mutex.Lock()
	mmQuery.QueryMock.callArgs = append(mmQuery.QueryMock.callArgs, mm_params)
	mmQuery.QueryMock.mutex.Unlock()

	for _, e := range mmQuery.QueryMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmQuery.QueryMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmQuery.QueryMock.defaultExpectation.Counter, 1)
		mm_want := mmQuery.QueryMock.defaultExpectation.params
		mm_got := SqlStmtMockQueryParams{ctx, args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmQuery.t.Errorf("SqlStmtMock.Query got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmQuery.QueryMock.defaultExpectation.results
		if mm_results == nil {
			mmQuery.t.Fatal("No results are set for the SqlStmtMock.Query")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmQuery.funcQuery != nil {
		return mmQuery.funcQuery(ctx, args...)
	}
	mmQuery.t.Fatalf("Unexpected call to SqlStmtMock.Query. %v %v", ctx, args)
	return
}

// QueryAfterCounter returns a count of finished SqlStmtMock.Query invocations
func (mmQuery *SqlStmtMock) QueryAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmQuery.afterQueryCounter)
}

// QueryBeforeCounter returns a count of SqlStmtMock.Query invocations
func (mmQuery *SqlStmtMock) QueryBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmQuery.beforeQueryCounter)
}

// Calls returns a list of arguments used in each call to SqlStmtMock.Query.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmQuery *mSqlStmtMockQuery) Calls() []*SqlStmtMockQueryParams {
	mmQuery.mutex.RLock()

	argCopy := make([]*SqlStmtMockQueryParams, len(mmQuery.callArgs))
	copy(argCopy, mmQuery.callArgs)

	mmQuery.mutex.RUnlock()

	return argCopy
}

// MinimockQueryDone returns true if the count of the Query invocations corresponds
// the number of defined expectations
func (m *SqlStmtMock) MinimockQueryDone() bool {
	for _, e := range m.QueryMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.QueryMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterQueryCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcQuery != nil && mm_atomic.LoadUint64(&m.afterQueryCounter) < 1 {
		return false
	}
	return true
}

// MinimockQueryInspect logs each unmet expectation
func (m *SqlStmtMock) MinimockQueryInspect() {
	for _, e := range m.QueryMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SqlStmtMock.Query with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.QueryMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterQueryCounter) < 1 {
		if m.QueryMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SqlStmtMock.Query")
		} else {
			m.t.Errorf("Expected call to SqlStmtMock.Query with params: %#v", *m.QueryMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcQuery != nil && mm_atomic.LoadUint64(&m.afterQueryCounter) < 1 {
		m.t.Error("Expected call to SqlStmtMock.Query")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SqlStmtMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCloseInspect()

		m.MinimockExecInspect()

		m.MinimockQueryInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SqlStmtMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *SqlStmtMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseDone() &&
		m.MinimockExecDone() &&
		m.MinimockQueryDone()
}
