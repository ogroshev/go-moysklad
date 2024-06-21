package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Application Серверное приложение.
//
// Ключевое слово: application
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-serwernye-prilozheniq
type Application struct {
	AccountID *uuid.UUID `json:"accountId,omitempty"` // ID учетной записи
	ID        *uuid.UUID `json:"id,omitempty"`        // ID Серверного приложения
	Name      *string    `json:"name,omitempty"`      // Наименование Серверного приложения
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные Серверного приложения
	AppUID    *uuid.UUID `json:"appUid,omitempty"`    // UID Серверного приложения
}

// GetAccountID возвращает ID учетной записи.
func (application Application) GetAccountID() uuid.UUID {
	return Deref(application.AccountID)
}

// GetID возвращает ID Серверного приложения.
func (application Application) GetID() uuid.UUID {
	return Deref(application.ID)
}

// GetName возвращает Наименование Серверного приложения.
func (application Application) GetName() string {
	return Deref(application.Name)
}

// GetMeta возвращает Метаданные Серверного приложения.
func (application Application) GetMeta() Meta {
	return Deref(application.Meta)
}

// GetAppUID возвращает UID Серверного приложения.
func (application Application) GetAppUID() uuid.UUID {
	return Deref(application.AppUID)
}

// String реализует интерфейс [fmt.Stringer].
func (application Application) String() string {
	return Stringify(application)
}

// MetaType возвращает тип сущности.
func (Application) MetaType() MetaType {
	return MetaTypeApplication
}

// ApplicationService сервис для работы с серверными приложениями.
type ApplicationService interface {
	// GetList выполняет запрос на получение списка сущностей установленных приложений.
	GetList(ctx context.Context, params ...*Params) (*List[Application], *resty.Response, error)

	// GetByID выполняет запрос на получение сущности установленного приложения.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Application, *resty.Response, error)
}

// NewApplicationService возвращает сервис для работы с серверными приложениями.
func NewApplicationService(client *Client) ApplicationService {
	e := NewEndpoint(client, "entity/application")
	return newMainService[Application, any, any, any](e)
}
