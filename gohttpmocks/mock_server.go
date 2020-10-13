package gohttpmocks

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"

	"github.com/ToteEmmanuel/go-httpclient/core"
)

var (
	MockedServer = mockServer{
		mocks:      make(map[string]*Mock),
		httpClient: &httpClientMock{},
	}
)

type mockServer struct {
	enabled     bool
	mocks       map[string]*Mock
	serverMutex sync.Mutex
	httpClient  core.HTTPClient
}

func (m *mockServer) Start() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()
	m.enabled = true
}

func (m *mockServer) Stop() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()
	m.enabled = false
}

func (m *mockServer) IsEnabled() bool {
	return m.enabled
}

func (m *mockServer) GetMockedClient() core.HTTPClient {
	return m.httpClient
}

func (m *mockServer) AddMock(mock Mock) {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()
	key := m.getMockKey(mock.Method, mock.URL, mock.RequestBody)
	m.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	return hex.EncodeToString(hasher.Sum(nil))
}
func (*mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return ""
}

func (m *mockServer) DeleteMocks() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()
	m.mocks = make(map[string]*Mock)
}
