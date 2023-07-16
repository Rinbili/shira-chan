// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"shira-chan-dev/ent/order"
	"shira-chan-dev/ent/predicate"
	"shira-chan-dev/ent/user"
)

// OrderWhereInput represents a where input for filtering Order queries.
type OrderWhereInput struct {
	Predicates []predicate.Order  `json:"-"`
	Not        *OrderWhereInput   `json:"not,omitempty"`
	Or         []*OrderWhereInput `json:"or,omitempty"`
	And        []*OrderWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "title" field predicates.
	Title             *string  `json:"title,omitempty"`
	TitleNEQ          *string  `json:"titleNEQ,omitempty"`
	TitleIn           []string `json:"titleIn,omitempty"`
	TitleNotIn        []string `json:"titleNotIn,omitempty"`
	TitleGT           *string  `json:"titleGT,omitempty"`
	TitleGTE          *string  `json:"titleGTE,omitempty"`
	TitleLT           *string  `json:"titleLT,omitempty"`
	TitleLTE          *string  `json:"titleLTE,omitempty"`
	TitleContains     *string  `json:"titleContains,omitempty"`
	TitleHasPrefix    *string  `json:"titleHasPrefix,omitempty"`
	TitleHasSuffix    *string  `json:"titleHasSuffix,omitempty"`
	TitleEqualFold    *string  `json:"titleEqualFold,omitempty"`
	TitleContainsFold *string  `json:"titleContainsFold,omitempty"`

	// "content" field predicates.
	Content             *string  `json:"content,omitempty"`
	ContentNEQ          *string  `json:"contentNEQ,omitempty"`
	ContentIn           []string `json:"contentIn,omitempty"`
	ContentNotIn        []string `json:"contentNotIn,omitempty"`
	ContentGT           *string  `json:"contentGT,omitempty"`
	ContentGTE          *string  `json:"contentGTE,omitempty"`
	ContentLT           *string  `json:"contentLT,omitempty"`
	ContentLTE          *string  `json:"contentLTE,omitempty"`
	ContentContains     *string  `json:"contentContains,omitempty"`
	ContentHasPrefix    *string  `json:"contentHasPrefix,omitempty"`
	ContentHasSuffix    *string  `json:"contentHasSuffix,omitempty"`
	ContentEqualFold    *string  `json:"contentEqualFold,omitempty"`
	ContentContainsFold *string  `json:"contentContainsFold,omitempty"`

	// "contact" field predicates.
	Contact             *string  `json:"contact,omitempty"`
	ContactNEQ          *string  `json:"contactNEQ,omitempty"`
	ContactIn           []string `json:"contactIn,omitempty"`
	ContactNotIn        []string `json:"contactNotIn,omitempty"`
	ContactGT           *string  `json:"contactGT,omitempty"`
	ContactGTE          *string  `json:"contactGTE,omitempty"`
	ContactLT           *string  `json:"contactLT,omitempty"`
	ContactLTE          *string  `json:"contactLTE,omitempty"`
	ContactContains     *string  `json:"contactContains,omitempty"`
	ContactHasPrefix    *string  `json:"contactHasPrefix,omitempty"`
	ContactHasSuffix    *string  `json:"contactHasSuffix,omitempty"`
	ContactEqualFold    *string  `json:"contactEqualFold,omitempty"`
	ContactContainsFold *string  `json:"contactContainsFold,omitempty"`

	// "type" field predicates.
	Type             *string  `json:"type,omitempty"`
	TypeNEQ          *string  `json:"typeNEQ,omitempty"`
	TypeIn           []string `json:"typeIn,omitempty"`
	TypeNotIn        []string `json:"typeNotIn,omitempty"`
	TypeGT           *string  `json:"typeGT,omitempty"`
	TypeGTE          *string  `json:"typeGTE,omitempty"`
	TypeLT           *string  `json:"typeLT,omitempty"`
	TypeLTE          *string  `json:"typeLTE,omitempty"`
	TypeContains     *string  `json:"typeContains,omitempty"`
	TypeHasPrefix    *string  `json:"typeHasPrefix,omitempty"`
	TypeHasSuffix    *string  `json:"typeHasSuffix,omitempty"`
	TypeEqualFold    *string  `json:"typeEqualFold,omitempty"`
	TypeContainsFold *string  `json:"typeContainsFold,omitempty"`

	// "is_closed" field predicates.
	IsClosed    *bool `json:"isClosed,omitempty"`
	IsClosedNEQ *bool `json:"isClosedNEQ,omitempty"`

	// "is_finished" field predicates.
	IsFinished    *bool `json:"isFinished,omitempty"`
	IsFinishedNEQ *bool `json:"isFinishedNEQ,omitempty"`

	// "evaluation" field predicates.
	Evaluation       *float64  `json:"evaluation,omitempty"`
	EvaluationNEQ    *float64  `json:"evaluationNEQ,omitempty"`
	EvaluationIn     []float64 `json:"evaluationIn,omitempty"`
	EvaluationNotIn  []float64 `json:"evaluationNotIn,omitempty"`
	EvaluationGT     *float64  `json:"evaluationGT,omitempty"`
	EvaluationGTE    *float64  `json:"evaluationGTE,omitempty"`
	EvaluationLT     *float64  `json:"evaluationLT,omitempty"`
	EvaluationLTE    *float64  `json:"evaluationLTE,omitempty"`
	EvaluationIsNil  bool      `json:"evaluationIsNil,omitempty"`
	EvaluationNotNil bool      `json:"evaluationNotNil,omitempty"`

	// "hope_at" field predicates.
	HopeAt      *int64  `json:"hopeAt,omitempty"`
	HopeAtNEQ   *int64  `json:"hopeAtNEQ,omitempty"`
	HopeAtIn    []int64 `json:"hopeAtIn,omitempty"`
	HopeAtNotIn []int64 `json:"hopeAtNotIn,omitempty"`
	HopeAtGT    *int64  `json:"hopeAtGT,omitempty"`
	HopeAtGTE   *int64  `json:"hopeAtGTE,omitempty"`
	HopeAtLT    *int64  `json:"hopeAtLT,omitempty"`
	HopeAtLTE   *int64  `json:"hopeAtLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *int64  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *int64  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []int64 `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []int64 `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *int64  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *int64  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *int64  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *int64  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *int64  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *int64  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []int64 `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []int64 `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *int64  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *int64  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *int64  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *int64  `json:"updatedAtLTE,omitempty"`

	// "requester" edge predicates.
	HasRequester     *bool             `json:"hasRequester,omitempty"`
	HasRequesterWith []*UserWhereInput `json:"hasRequesterWith,omitempty"`

	// "receiver" edge predicates.
	HasReceiver     *bool             `json:"hasReceiver,omitempty"`
	HasReceiverWith []*UserWhereInput `json:"hasReceiverWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *OrderWhereInput) AddPredicates(predicates ...predicate.Order) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the OrderWhereInput filter on the OrderQuery builder.
func (i *OrderWhereInput) Filter(q *OrderQuery) (*OrderQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyOrderWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyOrderWhereInput is returned in case the OrderWhereInput is empty.
var ErrEmptyOrderWhereInput = errors.New("ent: empty predicate OrderWhereInput")

// P returns a predicate for filtering orders.
// An error is returned if the input is empty or invalid.
func (i *OrderWhereInput) P() (predicate.Order, error) {
	var predicates []predicate.Order
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, order.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Order, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, order.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Order, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, order.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, order.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, order.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, order.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, order.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, order.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, order.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, order.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, order.IDLTE(*i.IDLTE))
	}
	if i.Title != nil {
		predicates = append(predicates, order.TitleEQ(*i.Title))
	}
	if i.TitleNEQ != nil {
		predicates = append(predicates, order.TitleNEQ(*i.TitleNEQ))
	}
	if len(i.TitleIn) > 0 {
		predicates = append(predicates, order.TitleIn(i.TitleIn...))
	}
	if len(i.TitleNotIn) > 0 {
		predicates = append(predicates, order.TitleNotIn(i.TitleNotIn...))
	}
	if i.TitleGT != nil {
		predicates = append(predicates, order.TitleGT(*i.TitleGT))
	}
	if i.TitleGTE != nil {
		predicates = append(predicates, order.TitleGTE(*i.TitleGTE))
	}
	if i.TitleLT != nil {
		predicates = append(predicates, order.TitleLT(*i.TitleLT))
	}
	if i.TitleLTE != nil {
		predicates = append(predicates, order.TitleLTE(*i.TitleLTE))
	}
	if i.TitleContains != nil {
		predicates = append(predicates, order.TitleContains(*i.TitleContains))
	}
	if i.TitleHasPrefix != nil {
		predicates = append(predicates, order.TitleHasPrefix(*i.TitleHasPrefix))
	}
	if i.TitleHasSuffix != nil {
		predicates = append(predicates, order.TitleHasSuffix(*i.TitleHasSuffix))
	}
	if i.TitleEqualFold != nil {
		predicates = append(predicates, order.TitleEqualFold(*i.TitleEqualFold))
	}
	if i.TitleContainsFold != nil {
		predicates = append(predicates, order.TitleContainsFold(*i.TitleContainsFold))
	}
	if i.Content != nil {
		predicates = append(predicates, order.ContentEQ(*i.Content))
	}
	if i.ContentNEQ != nil {
		predicates = append(predicates, order.ContentNEQ(*i.ContentNEQ))
	}
	if len(i.ContentIn) > 0 {
		predicates = append(predicates, order.ContentIn(i.ContentIn...))
	}
	if len(i.ContentNotIn) > 0 {
		predicates = append(predicates, order.ContentNotIn(i.ContentNotIn...))
	}
	if i.ContentGT != nil {
		predicates = append(predicates, order.ContentGT(*i.ContentGT))
	}
	if i.ContentGTE != nil {
		predicates = append(predicates, order.ContentGTE(*i.ContentGTE))
	}
	if i.ContentLT != nil {
		predicates = append(predicates, order.ContentLT(*i.ContentLT))
	}
	if i.ContentLTE != nil {
		predicates = append(predicates, order.ContentLTE(*i.ContentLTE))
	}
	if i.ContentContains != nil {
		predicates = append(predicates, order.ContentContains(*i.ContentContains))
	}
	if i.ContentHasPrefix != nil {
		predicates = append(predicates, order.ContentHasPrefix(*i.ContentHasPrefix))
	}
	if i.ContentHasSuffix != nil {
		predicates = append(predicates, order.ContentHasSuffix(*i.ContentHasSuffix))
	}
	if i.ContentEqualFold != nil {
		predicates = append(predicates, order.ContentEqualFold(*i.ContentEqualFold))
	}
	if i.ContentContainsFold != nil {
		predicates = append(predicates, order.ContentContainsFold(*i.ContentContainsFold))
	}
	if i.Contact != nil {
		predicates = append(predicates, order.ContactEQ(*i.Contact))
	}
	if i.ContactNEQ != nil {
		predicates = append(predicates, order.ContactNEQ(*i.ContactNEQ))
	}
	if len(i.ContactIn) > 0 {
		predicates = append(predicates, order.ContactIn(i.ContactIn...))
	}
	if len(i.ContactNotIn) > 0 {
		predicates = append(predicates, order.ContactNotIn(i.ContactNotIn...))
	}
	if i.ContactGT != nil {
		predicates = append(predicates, order.ContactGT(*i.ContactGT))
	}
	if i.ContactGTE != nil {
		predicates = append(predicates, order.ContactGTE(*i.ContactGTE))
	}
	if i.ContactLT != nil {
		predicates = append(predicates, order.ContactLT(*i.ContactLT))
	}
	if i.ContactLTE != nil {
		predicates = append(predicates, order.ContactLTE(*i.ContactLTE))
	}
	if i.ContactContains != nil {
		predicates = append(predicates, order.ContactContains(*i.ContactContains))
	}
	if i.ContactHasPrefix != nil {
		predicates = append(predicates, order.ContactHasPrefix(*i.ContactHasPrefix))
	}
	if i.ContactHasSuffix != nil {
		predicates = append(predicates, order.ContactHasSuffix(*i.ContactHasSuffix))
	}
	if i.ContactEqualFold != nil {
		predicates = append(predicates, order.ContactEqualFold(*i.ContactEqualFold))
	}
	if i.ContactContainsFold != nil {
		predicates = append(predicates, order.ContactContainsFold(*i.ContactContainsFold))
	}
	if i.Type != nil {
		predicates = append(predicates, order.TypeEQ(*i.Type))
	}
	if i.TypeNEQ != nil {
		predicates = append(predicates, order.TypeNEQ(*i.TypeNEQ))
	}
	if len(i.TypeIn) > 0 {
		predicates = append(predicates, order.TypeIn(i.TypeIn...))
	}
	if len(i.TypeNotIn) > 0 {
		predicates = append(predicates, order.TypeNotIn(i.TypeNotIn...))
	}
	if i.TypeGT != nil {
		predicates = append(predicates, order.TypeGT(*i.TypeGT))
	}
	if i.TypeGTE != nil {
		predicates = append(predicates, order.TypeGTE(*i.TypeGTE))
	}
	if i.TypeLT != nil {
		predicates = append(predicates, order.TypeLT(*i.TypeLT))
	}
	if i.TypeLTE != nil {
		predicates = append(predicates, order.TypeLTE(*i.TypeLTE))
	}
	if i.TypeContains != nil {
		predicates = append(predicates, order.TypeContains(*i.TypeContains))
	}
	if i.TypeHasPrefix != nil {
		predicates = append(predicates, order.TypeHasPrefix(*i.TypeHasPrefix))
	}
	if i.TypeHasSuffix != nil {
		predicates = append(predicates, order.TypeHasSuffix(*i.TypeHasSuffix))
	}
	if i.TypeEqualFold != nil {
		predicates = append(predicates, order.TypeEqualFold(*i.TypeEqualFold))
	}
	if i.TypeContainsFold != nil {
		predicates = append(predicates, order.TypeContainsFold(*i.TypeContainsFold))
	}
	if i.IsClosed != nil {
		predicates = append(predicates, order.IsClosedEQ(*i.IsClosed))
	}
	if i.IsClosedNEQ != nil {
		predicates = append(predicates, order.IsClosedNEQ(*i.IsClosedNEQ))
	}
	if i.IsFinished != nil {
		predicates = append(predicates, order.IsFinishedEQ(*i.IsFinished))
	}
	if i.IsFinishedNEQ != nil {
		predicates = append(predicates, order.IsFinishedNEQ(*i.IsFinishedNEQ))
	}
	if i.Evaluation != nil {
		predicates = append(predicates, order.EvaluationEQ(*i.Evaluation))
	}
	if i.EvaluationNEQ != nil {
		predicates = append(predicates, order.EvaluationNEQ(*i.EvaluationNEQ))
	}
	if len(i.EvaluationIn) > 0 {
		predicates = append(predicates, order.EvaluationIn(i.EvaluationIn...))
	}
	if len(i.EvaluationNotIn) > 0 {
		predicates = append(predicates, order.EvaluationNotIn(i.EvaluationNotIn...))
	}
	if i.EvaluationGT != nil {
		predicates = append(predicates, order.EvaluationGT(*i.EvaluationGT))
	}
	if i.EvaluationGTE != nil {
		predicates = append(predicates, order.EvaluationGTE(*i.EvaluationGTE))
	}
	if i.EvaluationLT != nil {
		predicates = append(predicates, order.EvaluationLT(*i.EvaluationLT))
	}
	if i.EvaluationLTE != nil {
		predicates = append(predicates, order.EvaluationLTE(*i.EvaluationLTE))
	}
	if i.EvaluationIsNil {
		predicates = append(predicates, order.EvaluationIsNil())
	}
	if i.EvaluationNotNil {
		predicates = append(predicates, order.EvaluationNotNil())
	}
	if i.HopeAt != nil {
		predicates = append(predicates, order.HopeAtEQ(*i.HopeAt))
	}
	if i.HopeAtNEQ != nil {
		predicates = append(predicates, order.HopeAtNEQ(*i.HopeAtNEQ))
	}
	if len(i.HopeAtIn) > 0 {
		predicates = append(predicates, order.HopeAtIn(i.HopeAtIn...))
	}
	if len(i.HopeAtNotIn) > 0 {
		predicates = append(predicates, order.HopeAtNotIn(i.HopeAtNotIn...))
	}
	if i.HopeAtGT != nil {
		predicates = append(predicates, order.HopeAtGT(*i.HopeAtGT))
	}
	if i.HopeAtGTE != nil {
		predicates = append(predicates, order.HopeAtGTE(*i.HopeAtGTE))
	}
	if i.HopeAtLT != nil {
		predicates = append(predicates, order.HopeAtLT(*i.HopeAtLT))
	}
	if i.HopeAtLTE != nil {
		predicates = append(predicates, order.HopeAtLTE(*i.HopeAtLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, order.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, order.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, order.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, order.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, order.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, order.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, order.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, order.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, order.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, order.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, order.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, order.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, order.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, order.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, order.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, order.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	if i.HasRequester != nil {
		p := order.HasRequester()
		if !*i.HasRequester {
			p = order.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasRequesterWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasRequesterWith))
		for _, w := range i.HasRequesterWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasRequesterWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, order.HasRequesterWith(with...))
	}
	if i.HasReceiver != nil {
		p := order.HasReceiver()
		if !*i.HasReceiver {
			p = order.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasReceiverWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasReceiverWith))
		for _, w := range i.HasReceiverWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasReceiverWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, order.HasReceiverWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyOrderWhereInput
	case 1:
		return predicates[0], nil
	default:
		return order.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "uname" field predicates.
	Uname             *string  `json:"uname,omitempty"`
	UnameNEQ          *string  `json:"unameNEQ,omitempty"`
	UnameIn           []string `json:"unameIn,omitempty"`
	UnameNotIn        []string `json:"unameNotIn,omitempty"`
	UnameGT           *string  `json:"unameGT,omitempty"`
	UnameGTE          *string  `json:"unameGTE,omitempty"`
	UnameLT           *string  `json:"unameLT,omitempty"`
	UnameLTE          *string  `json:"unameLTE,omitempty"`
	UnameContains     *string  `json:"unameContains,omitempty"`
	UnameHasPrefix    *string  `json:"unameHasPrefix,omitempty"`
	UnameHasSuffix    *string  `json:"unameHasSuffix,omitempty"`
	UnameEqualFold    *string  `json:"unameEqualFold,omitempty"`
	UnameContainsFold *string  `json:"unameContainsFold,omitempty"`

	// "passwd" field predicates.
	Passwd             *string  `json:"passwd,omitempty"`
	PasswdNEQ          *string  `json:"passwdNEQ,omitempty"`
	PasswdIn           []string `json:"passwdIn,omitempty"`
	PasswdNotIn        []string `json:"passwdNotIn,omitempty"`
	PasswdGT           *string  `json:"passwdGT,omitempty"`
	PasswdGTE          *string  `json:"passwdGTE,omitempty"`
	PasswdLT           *string  `json:"passwdLT,omitempty"`
	PasswdLTE          *string  `json:"passwdLTE,omitempty"`
	PasswdContains     *string  `json:"passwdContains,omitempty"`
	PasswdHasPrefix    *string  `json:"passwdHasPrefix,omitempty"`
	PasswdHasSuffix    *string  `json:"passwdHasSuffix,omitempty"`
	PasswdEqualFold    *string  `json:"passwdEqualFold,omitempty"`
	PasswdContainsFold *string  `json:"passwdContainsFold,omitempty"`

	// "phone" field predicates.
	Phone             *string  `json:"phone,omitempty"`
	PhoneNEQ          *string  `json:"phoneNEQ,omitempty"`
	PhoneIn           []string `json:"phoneIn,omitempty"`
	PhoneNotIn        []string `json:"phoneNotIn,omitempty"`
	PhoneGT           *string  `json:"phoneGT,omitempty"`
	PhoneGTE          *string  `json:"phoneGTE,omitempty"`
	PhoneLT           *string  `json:"phoneLT,omitempty"`
	PhoneLTE          *string  `json:"phoneLTE,omitempty"`
	PhoneContains     *string  `json:"phoneContains,omitempty"`
	PhoneHasPrefix    *string  `json:"phoneHasPrefix,omitempty"`
	PhoneHasSuffix    *string  `json:"phoneHasSuffix,omitempty"`
	PhoneEqualFold    *string  `json:"phoneEqualFold,omitempty"`
	PhoneContainsFold *string  `json:"phoneContainsFold,omitempty"`

	// "is_admin" field predicates.
	IsAdmin    *bool `json:"isAdmin,omitempty"`
	IsAdminNEQ *bool `json:"isAdminNEQ,omitempty"`

	// "is_secretary" field predicates.
	IsSecretary    *bool `json:"isSecretary,omitempty"`
	IsSecretaryNEQ *bool `json:"isSecretaryNEQ,omitempty"`

	// "is_active" field predicates.
	IsActive    *bool `json:"isActive,omitempty"`
	IsActiveNEQ *bool `json:"isActiveNEQ,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *int64  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *int64  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []int64 `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []int64 `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *int64  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *int64  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *int64  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *int64  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *int64  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *int64  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []int64 `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []int64 `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *int64  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *int64  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *int64  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *int64  `json:"updatedAtLTE,omitempty"`

	// "requested" edge predicates.
	HasRequested     *bool              `json:"hasRequested,omitempty"`
	HasRequestedWith []*OrderWhereInput `json:"hasRequestedWith,omitempty"`

	// "received" edge predicates.
	HasReceived     *bool              `json:"hasReceived,omitempty"`
	HasReceivedWith []*OrderWhereInput `json:"hasReceivedWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("ent: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Uname != nil {
		predicates = append(predicates, user.UnameEQ(*i.Uname))
	}
	if i.UnameNEQ != nil {
		predicates = append(predicates, user.UnameNEQ(*i.UnameNEQ))
	}
	if len(i.UnameIn) > 0 {
		predicates = append(predicates, user.UnameIn(i.UnameIn...))
	}
	if len(i.UnameNotIn) > 0 {
		predicates = append(predicates, user.UnameNotIn(i.UnameNotIn...))
	}
	if i.UnameGT != nil {
		predicates = append(predicates, user.UnameGT(*i.UnameGT))
	}
	if i.UnameGTE != nil {
		predicates = append(predicates, user.UnameGTE(*i.UnameGTE))
	}
	if i.UnameLT != nil {
		predicates = append(predicates, user.UnameLT(*i.UnameLT))
	}
	if i.UnameLTE != nil {
		predicates = append(predicates, user.UnameLTE(*i.UnameLTE))
	}
	if i.UnameContains != nil {
		predicates = append(predicates, user.UnameContains(*i.UnameContains))
	}
	if i.UnameHasPrefix != nil {
		predicates = append(predicates, user.UnameHasPrefix(*i.UnameHasPrefix))
	}
	if i.UnameHasSuffix != nil {
		predicates = append(predicates, user.UnameHasSuffix(*i.UnameHasSuffix))
	}
	if i.UnameEqualFold != nil {
		predicates = append(predicates, user.UnameEqualFold(*i.UnameEqualFold))
	}
	if i.UnameContainsFold != nil {
		predicates = append(predicates, user.UnameContainsFold(*i.UnameContainsFold))
	}
	if i.Passwd != nil {
		predicates = append(predicates, user.PasswdEQ(*i.Passwd))
	}
	if i.PasswdNEQ != nil {
		predicates = append(predicates, user.PasswdNEQ(*i.PasswdNEQ))
	}
	if len(i.PasswdIn) > 0 {
		predicates = append(predicates, user.PasswdIn(i.PasswdIn...))
	}
	if len(i.PasswdNotIn) > 0 {
		predicates = append(predicates, user.PasswdNotIn(i.PasswdNotIn...))
	}
	if i.PasswdGT != nil {
		predicates = append(predicates, user.PasswdGT(*i.PasswdGT))
	}
	if i.PasswdGTE != nil {
		predicates = append(predicates, user.PasswdGTE(*i.PasswdGTE))
	}
	if i.PasswdLT != nil {
		predicates = append(predicates, user.PasswdLT(*i.PasswdLT))
	}
	if i.PasswdLTE != nil {
		predicates = append(predicates, user.PasswdLTE(*i.PasswdLTE))
	}
	if i.PasswdContains != nil {
		predicates = append(predicates, user.PasswdContains(*i.PasswdContains))
	}
	if i.PasswdHasPrefix != nil {
		predicates = append(predicates, user.PasswdHasPrefix(*i.PasswdHasPrefix))
	}
	if i.PasswdHasSuffix != nil {
		predicates = append(predicates, user.PasswdHasSuffix(*i.PasswdHasSuffix))
	}
	if i.PasswdEqualFold != nil {
		predicates = append(predicates, user.PasswdEqualFold(*i.PasswdEqualFold))
	}
	if i.PasswdContainsFold != nil {
		predicates = append(predicates, user.PasswdContainsFold(*i.PasswdContainsFold))
	}
	if i.Phone != nil {
		predicates = append(predicates, user.PhoneEQ(*i.Phone))
	}
	if i.PhoneNEQ != nil {
		predicates = append(predicates, user.PhoneNEQ(*i.PhoneNEQ))
	}
	if len(i.PhoneIn) > 0 {
		predicates = append(predicates, user.PhoneIn(i.PhoneIn...))
	}
	if len(i.PhoneNotIn) > 0 {
		predicates = append(predicates, user.PhoneNotIn(i.PhoneNotIn...))
	}
	if i.PhoneGT != nil {
		predicates = append(predicates, user.PhoneGT(*i.PhoneGT))
	}
	if i.PhoneGTE != nil {
		predicates = append(predicates, user.PhoneGTE(*i.PhoneGTE))
	}
	if i.PhoneLT != nil {
		predicates = append(predicates, user.PhoneLT(*i.PhoneLT))
	}
	if i.PhoneLTE != nil {
		predicates = append(predicates, user.PhoneLTE(*i.PhoneLTE))
	}
	if i.PhoneContains != nil {
		predicates = append(predicates, user.PhoneContains(*i.PhoneContains))
	}
	if i.PhoneHasPrefix != nil {
		predicates = append(predicates, user.PhoneHasPrefix(*i.PhoneHasPrefix))
	}
	if i.PhoneHasSuffix != nil {
		predicates = append(predicates, user.PhoneHasSuffix(*i.PhoneHasSuffix))
	}
	if i.PhoneEqualFold != nil {
		predicates = append(predicates, user.PhoneEqualFold(*i.PhoneEqualFold))
	}
	if i.PhoneContainsFold != nil {
		predicates = append(predicates, user.PhoneContainsFold(*i.PhoneContainsFold))
	}
	if i.IsAdmin != nil {
		predicates = append(predicates, user.IsAdminEQ(*i.IsAdmin))
	}
	if i.IsAdminNEQ != nil {
		predicates = append(predicates, user.IsAdminNEQ(*i.IsAdminNEQ))
	}
	if i.IsSecretary != nil {
		predicates = append(predicates, user.IsSecretaryEQ(*i.IsSecretary))
	}
	if i.IsSecretaryNEQ != nil {
		predicates = append(predicates, user.IsSecretaryNEQ(*i.IsSecretaryNEQ))
	}
	if i.IsActive != nil {
		predicates = append(predicates, user.IsActiveEQ(*i.IsActive))
	}
	if i.IsActiveNEQ != nil {
		predicates = append(predicates, user.IsActiveNEQ(*i.IsActiveNEQ))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, user.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, user.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, user.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, user.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, user.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, user.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, user.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, user.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, user.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, user.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, user.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, user.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, user.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, user.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, user.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, user.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	if i.HasRequested != nil {
		p := user.HasRequested()
		if !*i.HasRequested {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasRequestedWith) > 0 {
		with := make([]predicate.Order, 0, len(i.HasRequestedWith))
		for _, w := range i.HasRequestedWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasRequestedWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasRequestedWith(with...))
	}
	if i.HasReceived != nil {
		p := user.HasReceived()
		if !*i.HasReceived {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasReceivedWith) > 0 {
		with := make([]predicate.Order, 0, len(i.HasReceivedWith))
		for _, w := range i.HasReceivedWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasReceivedWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasReceivedWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}
