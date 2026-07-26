package proxy

import "fmt"

type UserRole string

const (
	RoleGuest    UserRole = "guest"
	RoleCustomer UserRole = "customer"
	RoleWaiter   UserRole = "waiter"
	RoleAdmin    UserRole = "admin"
)

type User struct {
	ID       int
	Name     string
	Role     UserRole
	IsActive bool
}

// защищенный сервис
type SecureService struct {
	users map[int]User
}

func NewSecureService() *SecureService {
	return &SecureService{
		users: map[int]User{
			1: {ID: 1, Name: "Алексей", Role: RoleWaiter, IsActive: true},
			2: {ID: 2, Name: "Мария", Role: RoleAdmin, IsActive: true},
			3: {ID: 3, Name: "Иван", Role: RoleCustomer, IsActive: true},
			4: {ID: 4, Name: "Гость", Role: RoleGuest, IsActive: true},
		},
	}
}

func (s *SecureService) GetUser(id int) User {
	return s.users[id]
}

// действие доступное только администратору
func (s *SecureService) AdminAction(userID int, action string) string {
	user := s.GetUser(userID)
	return fmt.Sprintf("Администратор %s выполнил действие: %s", user.Name, action)
}

// действие доступное официанту и администратору
func (s *SecureService) WaiterAction(userID int, action string) string {
	user := s.GetUser(userID)
	return fmt.Sprintf("Официант %s выполнил действие: %s", user.Name, action)
}

// действие доступное клиенту
func (s *SecureService) CustomerAction(userID int, action string) string {
	user := s.GetUser(userID)
	return fmt.Sprintf("Клиент %s выполнил действие: %s", user.Name, action)
}

// ПРОКСИ ДЛЯ АВТОРИЗАЦИИ

// прокси для проверки прав доступа
type AuthProxy struct {
	service *SecureService
}

func NewAuthProxy(service *SecureService) *AuthProxy {
	return &AuthProxy{service: service}
}

// проверяет доступ пользователя
func (p *AuthProxy) CheckAccess(userID int, requiredRole UserRole) (User, error) {
	user := p.service.GetUser(userID)

	if !user.IsActive {
		return user, fmt.Errorf("пользователь %s неактивен", user.Name)
	}

	if requiredRole == RoleAdmin && user.Role != RoleAdmin {
		return user, fmt.Errorf("пользователь %s не имеет прав администратора", user.Name)
	}

	if requiredRole == RoleWaiter && user.Role != RoleWaiter && user.Role != RoleAdmin {
		return user, fmt.Errorf("пользователь %s не имеет прав официанта", user.Name)
	}

	if requiredRole == RoleCustomer && user.Role != RoleCustomer && user.Role != RoleWaiter && user.Role != RoleAdmin {
		return user, fmt.Errorf("пользователь %s не является клиентом", user.Name)
	}

	fmt.Printf("Доступ разрешён для %s (роль: %s)\n", user.Name, user.Role)
	return user, nil
}

// действие администратора (с проверкой прав)
func (p *AuthProxy) AdminAction(userID int, action string) (string, error) {
	_, err := p.CheckAccess(userID, RoleAdmin)
	if err != nil {
		return "", err
	}
	return p.service.AdminAction(userID, action), nil
}

// действие официанта (с проверкой прав)
func (p *AuthProxy) WaiterAction(userID int, action string) (string, error) {
	_, err := p.CheckAccess(userID, RoleWaiter)
	if err != nil {
		return "", err
	}
	return p.service.WaiterAction(userID, action), nil
}

// действие клиента (с проверкой прав)
func (p *AuthProxy) CustomerAction(userID int, action string) (string, error) {
	_, err := p.CheckAccess(userID, RoleCustomer)
	if err != nil {
		return "", err
	}
	return p.service.CustomerAction(userID, action), nil
}
