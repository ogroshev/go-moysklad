package moysklad

import "encoding/json"

// Operation общие поля операций, к которым привязан платёж.
type Operation struct {
	Group     *Group         `json:"group,omitempty"`     // Отдел сотрудника
	Meta      *Meta          `json:"meta,omitempty"`      // Метаданные операции
	Name      *string        `json:"name,omitempty"`      // Наименование операции
	LinkedSum *float64       `json:"linkedSum,omitempty"` // Сумма, оплаченную по данному документу
	AccountID *string        `json:"accountId,omitempty"` // ID учётной записи
	ID        *string        `json:"id,omitempty"`        // ID операции
	raw       []byte         // сырые данные для последующей конвертации в нужный тип
	Payments  Slice[Payment] `json:"payments,omitempty"` // Список ссылок на связанные платежи
}

func newOperation(metaOwner MetaOwner) *Operation {
	meta := metaOwner.GetMeta()
	return &Operation{Meta: &meta}
}

// GetGroup возвращает Отдел сотрудника.
func (operation Operation) GetGroup() Group {
	return Deref(operation.Group)
}

// GetMeta возвращает Метаданные операции.
func (operation Operation) GetMeta() Meta {
	return Deref(operation.Meta)
}

// GetName возвращает наименование операции.
func (operation Operation) GetName() string {
	return Deref(operation.Name)
}

// GetPayments возвращает список ссылок на связанные платежи.
func (operation Operation) GetPayments() Slice[Payment] {
	return operation.Payments
}

// GetLinkedSum возвращает Сумму, оплаченную по данному документу.
func (operation Operation) GetLinkedSum() float64 {
	return Deref(operation.LinkedSum)
}

// GetAccountID возвращает ID учётной записи.
func (operation Operation) GetAccountID() string {
	return Deref(operation.AccountID)
}

// GetID возвращает ID операции.
func (operation Operation) GetID() string {
	return Deref(operation.ID)
}

// SetMeta устанавливает Метаданные операции.
func (operation *Operation) SetMeta(meta *Meta) *Operation {
	operation.Meta = meta
	return operation
}

// SetLinkedSum устанавливает Сумму, оплаченную по данному документу.
func (operation *Operation) SetLinkedSum(linkedSum float64) *Operation {
	operation.LinkedSum = &linkedSum
	return operation
}

// OperationConverter описывает метод, возвращающий [Operation].
type OperationConverter interface {
	AsOperation() *Operation
}

// OperationInConverter описывает метод, возвращающий [Operation] для объектов [CashIn] и [PaymentIn].
//
// Метод должны реализовывать:
//   - CustomerOrder (Заказ покупателя)
//   - PurchaseReturn (Возврат поставщику)
//   - Demand (Отгрузка)
//   - InvoiceOut (Счет покупателю)
//   - CommissionReportIn (Полученный отчет комиссионера)
//   - RetailShift (Смена)
type OperationInConverter interface {
	AsOperationIn() *Operation
}

// OperationOutConverter описывает метод, возвращающий [Operation] для объектов [CashOut] и [PaymentOut].
//
// Метод должны реализовывать:
//   - SalesReturn (Возврат покупателя)
//   - Supply (Приемка)
//   - InvoiceIn (Счет поставщика)
//   - PurchaseOrder (Заказ поставщику)
//   - CommissionReportOut (Выданный отчет комиссионера)
type OperationOutConverter interface {
	AsOperationOut() *Operation
}

// String реализует интерфейс [fmt.Stringer].
func (operation Operation) String() string {
	return Stringify(operation.Meta)
}

// MetaType возвращает код сущности.
func (operation Operation) MetaType() MetaType {
	return operation.Meta.GetType()
}

// Raw реализует интерфейс [RawMetaTyper].
func (operation Operation) Raw() []byte {
	return operation.raw
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
func (operation *Operation) UnmarshalJSON(data []byte) error {
	type alias Operation
	var t alias

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.raw = data

	*operation = Operation(t)
	return nil
}

// IsCustomerOrder возвращает true, если объект имеет код сущности [MetaTypeCustomerOrder].
func (operation Operation) IsCustomerOrder() bool {
	return operation.Meta.GetType() == MetaTypeCustomerOrder
}

// IsPurchaseReturn возвращает true, если объект имеет код сущности [MetaTypePurchaseReturn].
func (operation Operation) IsPurchaseReturn() bool {
	return operation.Meta.GetType() == MetaTypePurchaseReturn
}

// IsDemand возвращает true, если объект имеет код сущности [MetaTypeDemand].
func (operation Operation) IsDemand() bool {
	return operation.Meta.GetType() == MetaTypeDemand
}

// IsInvoiceOut возвращает true, если объект имеет код сущности [MetaTypeInvoiceOut].
func (operation Operation) IsInvoiceOut() bool {
	return operation.Meta.GetType() == MetaTypeInvoiceOut
}

// IsRetailShift возвращает true, если объект имеет код сущности [MetaTypeRetailShift].
func (operation Operation) IsRetailShift() bool {
	return operation.Meta.GetType() == MetaTypeRetailShift
}

// IsCommissionReportIn возвращает true, если объект имеет код сущности [MetaTypeCommissionReportIn].
func (operation Operation) IsCommissionReportIn() bool {
	return operation.Meta.GetType() == MetaTypeCommissionReportIn
}

// IsSalesReturn возвращает true, если объект имеет код сущности [MetaTypeSalesReturn].
func (operation Operation) IsSalesReturn() bool {
	return operation.Meta.GetType() == MetaTypeSalesReturn
}

// IsSupply возвращает true, если объект имеет код сущности [MetaTypeSupply].
func (operation Operation) IsSupply() bool {
	return operation.Meta.GetType() == MetaTypeSupply
}

// IsInvoiceIn возвращает true, если объект имеет код сущности [MetaTypeInvoiceIn].
func (operation Operation) IsInvoiceIn() bool {
	return operation.Meta.GetType() == MetaTypeInvoiceIn
}

// IsPurchaseOrder возвращает true, если объект имеет код сущности [MetaTypePurchaseOrder].
func (operation Operation) IsPurchaseOrder() bool {
	return operation.Meta.GetType() == MetaTypePurchaseOrder
}

// IsCommissionReportOut возвращает true, если объект имеет код сущности [MetaTypeCommissionReportOut].
func (operation Operation) IsCommissionReportOut() bool {
	return operation.Meta.GetType() == MetaTypeCommissionReportOut
}

// AsCustomerOrder пытается привести объект к типу [CustomerOrder].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CustomerOrder] или nil в случае неудачи.
func (operation Operation) AsCustomerOrder() *CustomerOrder {
	return UnmarshalAsType[CustomerOrder](operation)
}

// AsPurchaseReturn пытается привести объект к типу [PurchaseReturn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PurchaseReturn] или nil в случае неудачи.
func (operation Operation) AsPurchaseReturn() *PurchaseReturn {
	return UnmarshalAsType[PurchaseReturn](operation)
}

// AsDemand пытается привести объект к типу [Demand].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Demand] или nil в случае неудачи.
func (operation Operation) AsDemand() *Demand {
	return UnmarshalAsType[Demand](operation)
}

// AsInvoiceOut пытается привести объект к типу [InvoiceOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [InvoiceOut] или nil в случае неудачи.
func (operation Operation) AsInvoiceOut() *InvoiceOut {
	return UnmarshalAsType[InvoiceOut](operation)
}

// AsRetailShift пытается привести объект к типу [RetailShift].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailShift] или nil в случае неудачи.
func (operation Operation) AsRetailShift() *RetailShift {
	return UnmarshalAsType[RetailShift](operation)
}

// AsCommissionReportIn пытается привести объект к типу [CommissionReportIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CommissionReportIn] или nil в случае неудачи.
func (operation Operation) AsCommissionReportIn() *CommissionReportIn {
	return UnmarshalAsType[CommissionReportIn](operation)
}

// AsSalesReturn пытается привести объект к типу [SalesReturn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [SalesReturn] или nil в случае неудачи.
func (operation Operation) AsSalesReturn() *SalesReturn {
	return UnmarshalAsType[SalesReturn](operation)
}

// AsSupply пытается привести объект к типу [Supply].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Supply] или nil в случае неудачи.
func (operation Operation) AsSupply() *Supply {
	return UnmarshalAsType[Supply](operation)
}

// AsInvoiceIn пытается привести объект к типу [InvoiceIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [InvoiceIn] или nil в случае неудачи.
func (operation Operation) AsInvoiceIn() *InvoiceIn {
	return UnmarshalAsType[InvoiceIn](operation)
}

// AsPurchaseOrder пытается привести объект к типу [PurchaseOrder].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PurchaseOrder] или nil в случае неудачи.
func (operation Operation) AsPurchaseOrder() *PurchaseOrder {
	return UnmarshalAsType[PurchaseOrder](operation)
}

// AsCommissionReportOut пытается привести объект к типу [CommissionReportOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CommissionReportOut] или nil в случае неудачи.
func (operation Operation) AsCommissionReportOut() *CommissionReportOut {
	return UnmarshalAsType[CommissionReportOut](operation)
}

// Operations список операций.
type Operations Slice[Operation]

// Push Привязка платежей к документам.
//
// Принимает множество объектов, реализующих интерфейс [OperationConverter].
func (operations *Operations) Push(elements ...*Operation) *Operations {
	for _, operation := range elements {
		*operations = append(*operations, operation)
	}
	return operations
}

// FilterCustomerOrder фильтрует список по типу [CustomerOrder].
func (operations Operations) FilterCustomerOrder() Slice[CustomerOrder] {
	return filterType[CustomerOrder](operations)
}

// FilterPurchaseReturn фильтрует список по типу [PurchaseReturn].
func (operations Operations) FilterPurchaseReturn() Slice[PurchaseReturn] {
	return filterType[PurchaseReturn](operations)
}

// FilterDemand фильтрует список по типу [Demand].
func (operations Operations) FilterDemand() Slice[Demand] {
	return filterType[Demand](operations)
}

// FilterInvoiceOut фильтрует список по типу [InvoiceOut].
func (operations Operations) FilterInvoiceOut() Slice[InvoiceOut] {
	return filterType[InvoiceOut](operations)
}

// FilterCommissionReportIn фильтрует список по типу [CommissionReportIn].
func (operations Operations) FilterCommissionReportIn() Slice[CommissionReportIn] {
	return filterType[CommissionReportIn](operations)
}

// FilterSalesReturn фильтрует список по типу [SalesReturn].
func (operations Operations) FilterSalesReturn() Slice[SalesReturn] {
	return filterType[SalesReturn](operations)
}

// FilterSupply фильтрует список по типу [Supply].
func (operations Operations) FilterSupply() Slice[Supply] {
	return filterType[Supply](operations)
}

// FilterInvoiceIn фильтрует список по типу [InvoiceIn].
func (operations Operations) FilterInvoiceIn() Slice[InvoiceIn] {
	return filterType[InvoiceIn](operations)
}

// FilterPurchaseOrder фильтрует список по типу [PurchaseOrder].
func (operations Operations) FilterPurchaseOrder() Slice[PurchaseOrder] {
	return filterType[PurchaseOrder](operations)
}

// FilterCommissionReportOut фильтрует список по типу [CommissionReportOut].
func (operations Operations) FilterCommissionReportOut() Slice[CommissionReportOut] {
	return filterType[CommissionReportOut](operations)
}

// FilterRetailShift фильтрует список по типу [RetailShift].
func (operations Operations) FilterRetailShift() Slice[RetailShift] {
	return filterType[RetailShift](operations)
}
