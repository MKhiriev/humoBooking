package services

import (
	"go-booking-system/internal/database"
	"go-booking-system/internal/models"
	"go-booking-system/pkg"
	"time"
)

type Service struct {
	BookingService    BookingServiceInterface
	RoomService       RoomServiceInterface
	UserService       UserServiceInterface
	AuthService       AuthServiceInterface
	RoleService       RoleServiceInterface
	RouteService      RouteServiceInterface
	ScopeService      ScopeServiceInterface
	PermissionService PermissionServiceInterface
}

func NewService(db *database.Database) *Service {
	bookingService := NewBookingService(db.BookingRepository)
	roomService := NewRoomService(db.RoomRepository)

	roleService := NewRoleService(db.RoleRepository)
	routeService := NewRouteService(db.RouteRepository)
	scopeService := NewScopeService(db.ScopeRepository)
	permissionService := NewPermissionService(db.PermissionRepository)

	return &Service{
		BookingService:    bookingService,
		RoomService:       roomService,
		UserService:       NewUserService(db.UserRepository),
		AuthService:       NewAuthService(db.UserRepository, roleService, routeService, scopeService, permissionService, bookingService, roomService),
		RoleService:       roleService,
		RouteService:      routeService,
		ScopeService:      scopeService,
		PermissionService: permissionService,
	}
}

type AuthServiceInterface interface {
	Create(user models.User) (models.User, error)
	UpdatePassword(userId int, password string) (models.User, error)
	UpdateUsername(userId int, username string) (models.User, error)
	UpdateRole(userId int, roleId int) (models.User, error)
	CheckIfUserExistsAndPasswordIsCorrect(username string, password string) (models.User, error)
	CheckPermissions(destination string, recordType string, recordId string, subject string, roleString string) (bool, error)
	GeneratePasswordHash(password string) string
	GenerateTokens(user models.User, identity pkg.IPAddressIdentity) (accessToken pkg.JWTToken, refreshToken pkg.JWTToken)
	ValidateAccessToken(encodedToken string, ipAddress string) *JWTTokenValidator
	ValidateRefreshToken(encodedToken string, ipAddress string) *JWTTokenValidator
}

type BookingServiceInterface interface {
	CheckIfRoomAvailable(roomId int, dateTimeStart time.Time, dateTimeEnd time.Time) (bool, error)
	BookRoom(userId int, roomId int, dateTimeStart time.Time, dateTimeEnd time.Time, createdBy int) (models.Booking, error)
	GetAll() []models.Booking
	GetBookingById(bookingId int) (models.Booking, error)
	GetBookingsByRoomId(roomId int) ([]models.Booking, error)
	GetBookingsByRoomIdAndBookingTime(roomId int, dateTimeStart time.Time, dateTimeEnd time.Time) ([]models.Booking, error)
	Update(booking models.Booking) (models.Booking, error)
	Delete(bookingId int) (bool, error)
}

type RoomServiceInterface interface {
	Create(room models.Room) (models.Room, error)
	GetAll() []models.Room
	GetRoomById(roomId int) (models.Room, error)
	Update(room models.Room) (models.Room, error)
	Delete(roomId int) (bool, error)
}

type UserServiceInterface interface {
	GetAll() []models.User
	GetUserById(userId int) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(userId int) (bool, error)
}

type RoleServiceInterface interface {
	Create(role models.Role) (models.Role, error)
	GetAll() []models.Role
	GetRoleById(roleId int) (models.Role, error)
	Update(role models.Role) (models.Role, error)
	Delete(roleId int) (bool, error)
}

type RouteServiceInterface interface {
	Create(route models.Route) (models.Route, error)
	GetAll() []models.Route
	GetRouteById(routeId int) (models.Route, error)
	GetRouteByURL(url string) (models.Route, error)
	Update(route models.Route) (models.Route, error)
	Delete(routeId int) (bool, error)
}

type ScopeServiceInterface interface {
	Create(scope models.Scope) (models.Scope, error)
	GetAll() []models.Scope
	GetScopeById(scopeId int) (models.Scope, error)
	Update(scope models.Scope) (models.Scope, error)
	Delete(scopeId int) (bool, error)
}

type PermissionServiceInterface interface {
	Create(permission models.Permission) (models.Permission, error)
	GetAll() []models.Permission
	GetPermissionsByRoleId(roleId int) ([]models.Permission, error)
	GetPermissionsByRouteId(routeId int) ([]models.Permission, error)
	GetPermissionsByRoleIdAndRouteId(roleId int, routeId int) ([]models.Permission, error)
	Update(permission models.Permission) (models.Permission, error)
	Delete(roleId int, routeId int) (bool, error)
}
