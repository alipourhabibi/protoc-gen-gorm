// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/feature_demo/demo_types.proto

package example

import context "context"
import errors "errors"
import time "time"

import auth1 "github.com/infobloxopen/atlas-app-toolkit/auth"
import gateway1 "github.com/infobloxopen/atlas-app-toolkit/gateway"
import go_uuid1 "github.com/satori/go.uuid"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import postgres1 "github.com/jinzhu/gorm/dialects/postgres"
import pq1 "github.com/lib/pq"
import ptypes1 "github.com/golang/protobuf/ptypes"
import types1 "github.com/infobloxopen/protoc-gen-gorm/types"

import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import user "github.com/infobloxopen/protoc-gen-gorm/example/user"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type TestTypesORM struct {
	ANestedObjectTypeWithIDId *uint32
	Array                     pq1.StringArray
	Array2                    pq1.StringArray
	BecomesInt                string
	CreatedAt                 time.Time
	JsonField                 *postgres1.Jsonb `gorm:"type:jsonb"`
	NullableUuid              *go_uuid1.UUID   `gorm:"type:uuid"`
	OptionalString            *string
	ThingsTypeWithIDId        *uint32
	TypeWithIdId              uint32
	Uuid                      go_uuid1.UUID `gorm:"type:uuid"`
}

// TableName overrides the default tablename generated by GORM
func (TestTypesORM) TableName() string {
	return "smorgasbord"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TestTypes) ToORM(ctx context.Context) (TestTypesORM, error) {
	to := TestTypesORM{}
	var err error
	if prehook, ok := interface{}(m).(TestTypesWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	// Repeated type []int32 is not an ORMable message type
	if m.OptionalString != nil {
		v := m.OptionalString.Value
		to.OptionalString = &v
	}
	to.BecomesInt = TestTypesStatus_name[int32(m.BecomesInt)]
	if m.Uuid != nil {
		to.Uuid, err = go_uuid1.FromString(m.Uuid.Value)
		if err != nil {
			return to, err
		}
	} else {
		to.Uuid = go_uuid1.Nil
	}
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
	}
	to.TypeWithIdId = m.TypeWithIdId
	if m.JsonField != nil {
		to.JsonField = &postgres1.Jsonb{[]byte(m.JsonField.Value)}
	}
	if m.NullableUuid != nil {
		tempUUID, uErr := go_uuid1.FromString(m.NullableUuid.Value)
		if uErr != nil {
			return to, uErr
		}
		to.NullableUuid = &tempUUID
	}
	if posthook, ok := interface{}(m).(TestTypesWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TestTypesORM) ToPB(ctx context.Context) (TestTypes, error) {
	to := TestTypes{}
	var err error
	if prehook, ok := interface{}(m).(TestTypesWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	// Repeated type []int32 is not an ORMable message type
	if m.OptionalString != nil {
		to.OptionalString = &google_protobuf1.StringValue{Value: *m.OptionalString}
	}
	to.BecomesInt = TestTypesStatus(TestTypesStatus_value[m.BecomesInt])
	to.Uuid = &types1.UUID{Value: m.Uuid.String()}
	if to.CreatedAt, err = ptypes1.TimestampProto(m.CreatedAt); err != nil {
		return to, err
	}
	to.TypeWithIdId = m.TypeWithIdId
	if m.JsonField != nil {
		to.JsonField = &types1.JSONValue{Value: string(m.JsonField.RawMessage)}
	}
	if m.NullableUuid != nil {
		to.NullableUuid = &types1.UUIDValue{Value: m.NullableUuid.String()}
	}
	if posthook, ok := interface{}(m).(TestTypesWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TestTypes the arg will be the target, the caller the one being converted from

// TestTypesBeforeToORM called before default ToORM code
type TestTypesWithBeforeToORM interface {
	BeforeToORM(context.Context, *TestTypesORM) error
}

// TestTypesAfterToORM called after default ToORM code
type TestTypesWithAfterToORM interface {
	AfterToORM(context.Context, *TestTypesORM) error
}

// TestTypesBeforeToPB called before default ToPB code
type TestTypesWithBeforeToPB interface {
	BeforeToPB(context.Context, *TestTypes) error
}

// TestTypesAfterToPB called after default ToPB code
type TestTypesWithAfterToPB interface {
	AfterToPB(context.Context, *TestTypes) error
}

type TypeWithIDORM struct {
	ANestedObject     *TestTypesORM `gorm:"foreignkey:ANestedObjectTypeWithIDId;association_foreignkey:Id"`
	Address           *types1.Inet  `gorm:"type:inet"`
	Id                uint32
	IntPointId        *uint32
	Ip                string          `gorm:"column:ip_addr"`
	MultiAccountTypes []*JoinTable    `gorm:"foreignkey:TypeWithIDID"`
	Point             *IntPointORM    `gorm:"foreignkey:IntPointId;association_foreignkey:Id"`
	SecretInt         int32           `gorm:"-"`
	Things            []*TestTypesORM `gorm:"foreignkey:ThingsTypeWithIDId;association_foreignkey:Id"`
	User              *user.UserORM   `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	UserId            *string
}

// TableName overrides the default tablename generated by GORM
func (TypeWithIDORM) TableName() string {
	return "type_with_ids"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TypeWithID) ToORM(ctx context.Context) (TypeWithIDORM, error) {
	to := TypeWithIDORM{}
	var err error
	if prehook, ok := interface{}(m).(TypeWithIDWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Ip = m.Ip
	for _, v := range m.Things {
		if v != nil {
			if tempThings, cErr := v.ToORM(ctx); cErr == nil {
				to.Things = append(to.Things, &tempThings)
			} else {
				return to, cErr
			}
		} else {
			to.Things = append(to.Things, nil)
		}
	}
	if m.ANestedObject != nil {
		tempANestedObject, err := m.ANestedObject.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.ANestedObject = &tempANestedObject
	}
	if m.Point != nil {
		tempPoint, err := m.Point.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.Point = &tempPoint
	}
	if m.User != nil {
		tempUser, err := m.User.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	if m.Address != nil {
		if to.Address, err = types1.ParseInet(m.Address.Value); err != nil {
			return to, err
		}
	}
	if posthook, ok := interface{}(m).(TypeWithIDWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TypeWithIDORM) ToPB(ctx context.Context) (TypeWithID, error) {
	to := TypeWithID{}
	var err error
	if prehook, ok := interface{}(m).(TypeWithIDWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Ip = m.Ip
	for _, v := range m.Things {
		if v != nil {
			if tempThings, cErr := v.ToPB(ctx); cErr == nil {
				to.Things = append(to.Things, &tempThings)
			} else {
				return to, cErr
			}
		} else {
			to.Things = append(to.Things, nil)
		}
	}
	if m.ANestedObject != nil {
		tempANestedObject, err := m.ANestedObject.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.ANestedObject = &tempANestedObject
	}
	if m.Point != nil {
		tempPoint, err := m.Point.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.Point = &tempPoint
	}
	if m.User != nil {
		tempUser, err := m.User.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	if m.Address != nil && m.Address.IPNet != nil {
		to.Address = &types1.InetValue{Value: m.Address.String()}
	}
	if posthook, ok := interface{}(m).(TypeWithIDWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TypeWithID the arg will be the target, the caller the one being converted from

// TypeWithIDBeforeToORM called before default ToORM code
type TypeWithIDWithBeforeToORM interface {
	BeforeToORM(context.Context, *TypeWithIDORM) error
}

// TypeWithIDAfterToORM called after default ToORM code
type TypeWithIDWithAfterToORM interface {
	AfterToORM(context.Context, *TypeWithIDORM) error
}

// TypeWithIDBeforeToPB called before default ToPB code
type TypeWithIDWithBeforeToPB interface {
	BeforeToPB(context.Context, *TypeWithID) error
}

// TypeWithIDAfterToPB called after default ToPB code
type TypeWithIDWithAfterToPB interface {
	AfterToPB(context.Context, *TypeWithID) error
}

type MultiaccountTypeWithIDORM struct {
	AccountID string
	Id        uint64
	SomeField string
}

// TableName overrides the default tablename generated by GORM
func (MultiaccountTypeWithIDORM) TableName() string {
	return "multiaccount_type_with_ids"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *MultiaccountTypeWithID) ToORM(ctx context.Context) (MultiaccountTypeWithIDORM, error) {
	to := MultiaccountTypeWithIDORM{}
	var err error
	if prehook, ok := interface{}(m).(MultiaccountTypeWithIDWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.SomeField = m.SomeField
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return to, err
	}
	to.AccountID = accountID
	if posthook, ok := interface{}(m).(MultiaccountTypeWithIDWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *MultiaccountTypeWithIDORM) ToPB(ctx context.Context) (MultiaccountTypeWithID, error) {
	to := MultiaccountTypeWithID{}
	var err error
	if prehook, ok := interface{}(m).(MultiaccountTypeWithIDWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.SomeField = m.SomeField
	if posthook, ok := interface{}(m).(MultiaccountTypeWithIDWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type MultiaccountTypeWithID the arg will be the target, the caller the one being converted from

// MultiaccountTypeWithIDBeforeToORM called before default ToORM code
type MultiaccountTypeWithIDWithBeforeToORM interface {
	BeforeToORM(context.Context, *MultiaccountTypeWithIDORM) error
}

// MultiaccountTypeWithIDAfterToORM called after default ToORM code
type MultiaccountTypeWithIDWithAfterToORM interface {
	AfterToORM(context.Context, *MultiaccountTypeWithIDORM) error
}

// MultiaccountTypeWithIDBeforeToPB called before default ToPB code
type MultiaccountTypeWithIDWithBeforeToPB interface {
	BeforeToPB(context.Context, *MultiaccountTypeWithID) error
}

// MultiaccountTypeWithIDAfterToPB called after default ToPB code
type MultiaccountTypeWithIDWithAfterToPB interface {
	AfterToPB(context.Context, *MultiaccountTypeWithID) error
}

type MultiaccountTypeWithoutIDORM struct {
	AccountID string
	SomeField string
}

// TableName overrides the default tablename generated by GORM
func (MultiaccountTypeWithoutIDORM) TableName() string {
	return "multiaccount_type_without_ids"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *MultiaccountTypeWithoutID) ToORM(ctx context.Context) (MultiaccountTypeWithoutIDORM, error) {
	to := MultiaccountTypeWithoutIDORM{}
	var err error
	if prehook, ok := interface{}(m).(MultiaccountTypeWithoutIDWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.SomeField = m.SomeField
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return to, err
	}
	to.AccountID = accountID
	if posthook, ok := interface{}(m).(MultiaccountTypeWithoutIDWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *MultiaccountTypeWithoutIDORM) ToPB(ctx context.Context) (MultiaccountTypeWithoutID, error) {
	to := MultiaccountTypeWithoutID{}
	var err error
	if prehook, ok := interface{}(m).(MultiaccountTypeWithoutIDWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.SomeField = m.SomeField
	if posthook, ok := interface{}(m).(MultiaccountTypeWithoutIDWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type MultiaccountTypeWithoutID the arg will be the target, the caller the one being converted from

// MultiaccountTypeWithoutIDBeforeToORM called before default ToORM code
type MultiaccountTypeWithoutIDWithBeforeToORM interface {
	BeforeToORM(context.Context, *MultiaccountTypeWithoutIDORM) error
}

// MultiaccountTypeWithoutIDAfterToORM called after default ToORM code
type MultiaccountTypeWithoutIDWithAfterToORM interface {
	AfterToORM(context.Context, *MultiaccountTypeWithoutIDORM) error
}

// MultiaccountTypeWithoutIDBeforeToPB called before default ToPB code
type MultiaccountTypeWithoutIDWithBeforeToPB interface {
	BeforeToPB(context.Context, *MultiaccountTypeWithoutID) error
}

// MultiaccountTypeWithoutIDAfterToPB called after default ToPB code
type MultiaccountTypeWithoutIDWithAfterToPB interface {
	AfterToPB(context.Context, *MultiaccountTypeWithoutID) error
}

type PrimaryUUIDTypeORM struct {
	Child *ExternalChildORM `gorm:"foreignkey:PrimaryUUIDTypeId;association_foreignkey:Id"`
	Id    *go_uuid1.UUID    `gorm:"type:uuid"`
}

// TableName overrides the default tablename generated by GORM
func (PrimaryUUIDTypeORM) TableName() string {
	return "primary_uuid_types"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *PrimaryUUIDType) ToORM(ctx context.Context) (PrimaryUUIDTypeORM, error) {
	to := PrimaryUUIDTypeORM{}
	var err error
	if prehook, ok := interface{}(m).(PrimaryUUIDTypeWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if m.Id != nil {
		tempUUID, uErr := go_uuid1.FromString(m.Id.Value)
		if uErr != nil {
			return to, uErr
		}
		to.Id = &tempUUID
	}
	if m.Child != nil {
		tempChild, err := m.Child.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.Child = &tempChild
	}
	if posthook, ok := interface{}(m).(PrimaryUUIDTypeWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *PrimaryUUIDTypeORM) ToPB(ctx context.Context) (PrimaryUUIDType, error) {
	to := PrimaryUUIDType{}
	var err error
	if prehook, ok := interface{}(m).(PrimaryUUIDTypeWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if m.Id != nil {
		to.Id = &types1.UUIDValue{Value: m.Id.String()}
	}
	if m.Child != nil {
		tempChild, err := m.Child.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.Child = &tempChild
	}
	if posthook, ok := interface{}(m).(PrimaryUUIDTypeWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type PrimaryUUIDType the arg will be the target, the caller the one being converted from

// PrimaryUUIDTypeBeforeToORM called before default ToORM code
type PrimaryUUIDTypeWithBeforeToORM interface {
	BeforeToORM(context.Context, *PrimaryUUIDTypeORM) error
}

// PrimaryUUIDTypeAfterToORM called after default ToORM code
type PrimaryUUIDTypeWithAfterToORM interface {
	AfterToORM(context.Context, *PrimaryUUIDTypeORM) error
}

// PrimaryUUIDTypeBeforeToPB called before default ToPB code
type PrimaryUUIDTypeWithBeforeToPB interface {
	BeforeToPB(context.Context, *PrimaryUUIDType) error
}

// PrimaryUUIDTypeAfterToPB called after default ToPB code
type PrimaryUUIDTypeWithAfterToPB interface {
	AfterToPB(context.Context, *PrimaryUUIDType) error
}

// DefaultCreateTestTypes executes a basic gorm create call
func DefaultCreateTestTypes(ctx context.Context, in *TestTypes, db *gorm1.DB) (*TestTypes, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateTestTypes")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultListTestTypes executes a gorm list call
func DefaultListTestTypes(ctx context.Context, db *gorm1.DB, req interface{}) ([]*TestTypes, error) {
	ormResponse := []TestTypesORM{}
	f, s, p, fs, err := getCollectionOperators(req)
	if err != nil {
		return nil, err
	}
	db, err = gorm2.ApplyCollectionOperators(db, &TestTypesORM{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	in := TestTypes{}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	db = db.Where(&ormParams)
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*TestTypes{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateTypeWithID executes a basic gorm create call
func DefaultCreateTypeWithID(ctx context.Context, in *TypeWithID, db *gorm1.DB) (*TypeWithID, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateTypeWithID")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadTypeWithID executes a basic gorm read call
func DefaultReadTypeWithID(ctx context.Context, in *TypeWithID, db *gorm1.DB) (*TypeWithID, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadTypeWithID")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := TypeWithIDORM{}
	if err = db.Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateTypeWithID executes a basic gorm update call
func DefaultUpdateTypeWithID(ctx context.Context, in *TypeWithID, db *gorm1.DB) (*TypeWithID, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateTypeWithID")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteTypeWithID(ctx context.Context, in *TypeWithID, db *gorm1.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteTypeWithID")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&TypeWithIDORM{}).Error
	return err
}

// DefaultStrictUpdateTypeWithID clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateTypeWithID(ctx context.Context, in *TypeWithID, db *gorm1.DB) (*TypeWithID, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTypeWithID")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	count := 1
	err = db.Model(&ormObj).Where("id=?", ormObj.Id).Count(&count).Error
	if err != nil {
		return nil, err
	}
	filterANestedObject := TestTypesORM{}
	if ormObj.Id == 0 {
		return nil, errors.New("Can't do overwriting update with no Id value for TypeWithIDORM")
	}
	filterANestedObject.ANestedObjectTypeWithIDId = new(uint32)
	*filterANestedObject.ANestedObjectTypeWithIDId = ormObj.Id
	if err = db.Where(filterANestedObject).Delete(TestTypesORM{}).Error; err != nil {
		return nil, err
	}
	filterThings := TestTypesORM{}
	if ormObj.Id == 0 {
		return nil, errors.New("Can't do overwriting update with no Id value for TypeWithIDORM")
	}
	filterThings.ThingsTypeWithIDId = new(uint32)
	*filterThings.ThingsTypeWithIDId = ormObj.Id
	if err = db.Where(filterThings).Delete(TestTypesORM{}).Error; err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		err = gateway1.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

// DefaultListTypeWithID executes a gorm list call
func DefaultListTypeWithID(ctx context.Context, db *gorm1.DB, req interface{}) ([]*TypeWithID, error) {
	ormResponse := []TypeWithIDORM{}
	f, s, p, fs, err := getCollectionOperators(req)
	if err != nil {
		return nil, err
	}
	db, err = gorm2.ApplyCollectionOperators(db, &TypeWithIDORM{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	in := TypeWithID{}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	db = db.Where(&ormParams)
	db = db.Order("id")
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*TypeWithID{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateMultiaccountTypeWithID executes a basic gorm create call
func DefaultCreateMultiaccountTypeWithID(ctx context.Context, in *MultiaccountTypeWithID, db *gorm1.DB) (*MultiaccountTypeWithID, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateMultiaccountTypeWithID")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadMultiaccountTypeWithID executes a basic gorm read call
func DefaultReadMultiaccountTypeWithID(ctx context.Context, in *MultiaccountTypeWithID, db *gorm1.DB) (*MultiaccountTypeWithID, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadMultiaccountTypeWithID")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := MultiaccountTypeWithIDORM{}
	if err = db.Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateMultiaccountTypeWithID executes a basic gorm update call
func DefaultUpdateMultiaccountTypeWithID(ctx context.Context, in *MultiaccountTypeWithID, db *gorm1.DB) (*MultiaccountTypeWithID, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateMultiaccountTypeWithID")
	}
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if exists, err := DefaultReadMultiaccountTypeWithID(ctx, &MultiaccountTypeWithID{Id: in.GetId()}, db); err != nil {
		return nil, err
	} else if exists == nil {
		return nil, errors.New("MultiaccountTypeWithID not found")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormObj.AccountID = accountID
	db = db.Where(&MultiaccountTypeWithIDORM{AccountID: accountID})
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteMultiaccountTypeWithID(ctx context.Context, in *MultiaccountTypeWithID, db *gorm1.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteMultiaccountTypeWithID")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&MultiaccountTypeWithIDORM{}).Error
	return err
}

// DefaultStrictUpdateMultiaccountTypeWithID clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateMultiaccountTypeWithID(ctx context.Context, in *MultiaccountTypeWithID, db *gorm1.DB) (*MultiaccountTypeWithID, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateMultiaccountTypeWithID")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	count := 1
	err = db.Model(&ormObj).Where("id=?", ormObj.Id).Count(&count).Error
	if err != nil {
		return nil, err
	}
	db = db.Where(&MultiaccountTypeWithIDORM{AccountID: ormObj.AccountID})
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		err = gateway1.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

// DefaultListMultiaccountTypeWithID executes a gorm list call
func DefaultListMultiaccountTypeWithID(ctx context.Context, db *gorm1.DB, req interface{}) ([]*MultiaccountTypeWithID, error) {
	ormResponse := []MultiaccountTypeWithIDORM{}
	f, s, p, fs, err := getCollectionOperators(req)
	if err != nil {
		return nil, err
	}
	db, err = gorm2.ApplyCollectionOperators(db, &MultiaccountTypeWithIDORM{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	in := MultiaccountTypeWithID{}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	db = db.Where(&ormParams)
	db = db.Order("id")
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*MultiaccountTypeWithID{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateMultiaccountTypeWithoutID executes a basic gorm create call
func DefaultCreateMultiaccountTypeWithoutID(ctx context.Context, in *MultiaccountTypeWithoutID, db *gorm1.DB) (*MultiaccountTypeWithoutID, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateMultiaccountTypeWithoutID")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultListMultiaccountTypeWithoutID executes a gorm list call
func DefaultListMultiaccountTypeWithoutID(ctx context.Context, db *gorm1.DB, req interface{}) ([]*MultiaccountTypeWithoutID, error) {
	ormResponse := []MultiaccountTypeWithoutIDORM{}
	f, s, p, fs, err := getCollectionOperators(req)
	if err != nil {
		return nil, err
	}
	db, err = gorm2.ApplyCollectionOperators(db, &MultiaccountTypeWithoutIDORM{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	in := MultiaccountTypeWithoutID{}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	db = db.Where(&ormParams)
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*MultiaccountTypeWithoutID{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreatePrimaryUUIDType executes a basic gorm create call
func DefaultCreatePrimaryUUIDType(ctx context.Context, in *PrimaryUUIDType, db *gorm1.DB) (*PrimaryUUIDType, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreatePrimaryUUIDType")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadPrimaryUUIDType executes a basic gorm read call
func DefaultReadPrimaryUUIDType(ctx context.Context, in *PrimaryUUIDType, db *gorm1.DB) (*PrimaryUUIDType, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadPrimaryUUIDType")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := PrimaryUUIDTypeORM{}
	if err = db.Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdatePrimaryUUIDType executes a basic gorm update call
func DefaultUpdatePrimaryUUIDType(ctx context.Context, in *PrimaryUUIDType, db *gorm1.DB) (*PrimaryUUIDType, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdatePrimaryUUIDType")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeletePrimaryUUIDType(ctx context.Context, in *PrimaryUUIDType, db *gorm1.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeletePrimaryUUIDType")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == nil || *ormObj.Id == go_uuid1.Nil {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&PrimaryUUIDTypeORM{}).Error
	return err
}

// DefaultStrictUpdatePrimaryUUIDType clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdatePrimaryUUIDType(ctx context.Context, in *PrimaryUUIDType, db *gorm1.DB) (*PrimaryUUIDType, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdatePrimaryUUIDType")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	count := 1
	err = db.Model(&ormObj).Where("id=?", ormObj.Id).Count(&count).Error
	if err != nil {
		return nil, err
	}
	filterChild := ExternalChildORM{}
	if ormObj.Id == nil || *ormObj.Id == go_uuid1.Nil {
		return nil, errors.New("Can't do overwriting update with no Id value for PrimaryUUIDTypeORM")
	}
	filterChild.PrimaryUUIDTypeId = *ormObj.Id
	if err = db.Where(filterChild).Delete(ExternalChildORM{}).Error; err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		err = gateway1.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

// DefaultListPrimaryUUIDType executes a gorm list call
func DefaultListPrimaryUUIDType(ctx context.Context, db *gorm1.DB, req interface{}) ([]*PrimaryUUIDType, error) {
	ormResponse := []PrimaryUUIDTypeORM{}
	f, s, p, fs, err := getCollectionOperators(req)
	if err != nil {
		return nil, err
	}
	db, err = gorm2.ApplyCollectionOperators(db, &PrimaryUUIDTypeORM{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	in := PrimaryUUIDType{}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	db = db.Where(&ormParams)
	db = db.Order("id")
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*PrimaryUUIDType{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
