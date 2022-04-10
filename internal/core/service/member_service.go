package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/core/port"
	"strings"
)

type memberService struct {
	memberRepository port.MemberRepository
}

func NewMemberService(memberRepository port.MemberRepository) MemberService {
	return &memberService{memberRepository: memberRepository}
}

func (m memberService) NewMember(requester NewMemberRequester) (*MemberResponse, error) {
	argon := argon2.DefaultConfig()

	if len(requester.Username) >= 5 && len(requester.Username) <= 24 {
		requester.Username = strings.ToLower(requester.Username)
	} else {
		return nil, errors.New("the username must be more than 6 letters and must not be more than 24")
	}

	if len(requester.Firstname) >= 2 && len(requester.Firstname) <= 64 {
		requester.Firstname = strings.ToLower(requester.Firstname)
	} else {
		return nil, errors.New("the firstname must be more than 2 letters and must not be more than 64")
	}

	if len(requester.Lastname) >= 2 && len(requester.Lastname) <= 64 {
		requester.Lastname = strings.ToLower(requester.Lastname)
	} else {
		return nil, errors.New("the lastname must be more than 2 letters and must not be more than 64")
	}

	if len(requester.Password) >= 6 {
		encoded, err := argon.HashEncoded([]byte(requester.Password))
		if err != nil {
			return nil, err
		}

		requester.Password = string(encoded)
	} else {
		return nil, errors.New("the password must be more than 6 letters")
	}

	member := port.Member{
		Username:  requester.Username,
		Firstname: requester.Firstname,
		Lastname:  requester.Lastname,
		Password:  requester.Password,
	}

	newMember, err := m.memberRepository.CreateMember(member)
	if err != nil {
		return nil, err
	}

	memberResponse := MemberResponse{
		ID:        newMember.ID,
		CreatedAt: newMember.CreatedAt,
		Username:  newMember.Username,
		Firstname: newMember.Firstname,
		Lastname:  newMember.Lastname,
	}

	return &memberResponse, nil
}

func (m memberService) ReadMembers() ([]*MemberResponse, error) {
	members, err := m.memberRepository.GetAllMember()
	if err != nil {
		return nil, err
	}

	var responses []*MemberResponse
	for _, member := range members {
		responses = append(responses, &MemberResponse{
			ID:        member.ID,
			CreatedAt: member.CreatedAt,
			Username:  member.Username,
			Firstname: member.Firstname,
			Lastname:  member.Lastname,
		})
	}

	return responses, nil
}

func (m memberService) ReadMemberById(uuid uuid.UUID) (*MemberResponse, error) {
	member, err := m.memberRepository.GetMemberById(uuid)
	if err != nil {
		return nil, err
	}

	response := MemberResponse{
		ID:        member.ID,
		CreatedAt: member.CreatedAt,
		Username:  member.Username,
		Firstname: member.Firstname,
		Lastname:  member.Lastname,
	}

	return &response, nil
}

func (m memberService) DeleteMemberById(uuid uuid.UUID) error {
	if err := m.memberRepository.DropMemberById(uuid); err != nil {
		return errors.New("the uuid could be not found")
	}

	return nil
}
