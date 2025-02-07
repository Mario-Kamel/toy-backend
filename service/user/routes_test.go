package user

func TestUserServiceHandlers (t *testing.T) {
	userStore := &mockUserStore{}
}

type mockUserStore struct{}

func(m *mockUserStore) GetUserByEmail