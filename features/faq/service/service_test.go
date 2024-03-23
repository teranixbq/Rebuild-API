package service

import (
	"errors"
	"testing"
	"recything/features/faq/entity"
	"recything/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllFaqs(t *testing.T) {
	mockRepo := new(mocks.FaqRepositoryInterface)
	faqService := NewFaqService(mockRepo)

	// Mock data
	mockData := []entity.FaqCore{
		{Id: 1, Title: "Cara saya melakukan penukaran point", Description: "Berikut cara melakukan penukaran point"},
		{Id: 2, Title: "Cara saya bergabung ke dalam komunitas", Description: "Berikut cara bergabung ke dalam komunitas"},
		{Id: 3, Title: "Cara saya melaporkan sampah tidak pada tempatnya", Description: "Berikut cara melaporkan sampah tidak pada tempatnya"},
	}

	// Mock repository response
	mockRepo.On("GetFaqs").
		Return(mockData, nil)

	// Test case
	faqs, err := faqService.GetFaqs()

	assert.NoError(t, err)
	assert.Len(t, faqs, len(mockData))
	mockRepo.AssertExpectations(t)
}

func TestGetAllFaqsError(t *testing.T) {
    mockRepo := new(mocks.FaqRepositoryInterface)
    faqService := NewFaqService(mockRepo)

    // Mock repository response with an error
    mockRepo.On("GetFaqs").Return(nil, errors.New("some error"))

    // Test case
    faqs, err := faqService.GetFaqs()

    assert.Error(t, err)
    assert.Nil(t, faqs)
    mockRepo.AssertExpectations(t)
}

func TestGetFaqById(t *testing.T) {
	mockRepo := new(mocks.FaqRepositoryInterface)
	faqService := NewFaqService(mockRepo)

	// Mock data
	mockData := entity.FaqCore{
		Id: 1, Title: "Cara saya melakukan penukaran point", Description: "Berikut cara melakukan penukaran point",
	}

	mockRepo.On("GetFaqsById", mock.AnythingOfType("uint")).Return(mockData, nil)

	faq, err := faqService.GetFaqsById(1)

	assert.NoError(t, err)
	assert.Equal(t, mockData.Id, faq.Id)
	mockRepo.AssertExpectations(t)
}

func TestGetFaqByIdNonExistentId(t *testing.T) {
    mockRepo := new(mocks.FaqRepositoryInterface)
    faqService := NewFaqService(mockRepo)

    // Mock repository response with an error for a non-existent ID
    mockRepo.On("GetFaqsById", mock.AnythingOfType("uint")).Return(entity.FaqCore{}, errors.New("not found"))

    // Test case
    faq, err := faqService.GetFaqsById(999)

    assert.Error(t, err)
    assert.EqualError(t, err, "not found") // Check for the specific error message
    assert.Equal(t, entity.FaqCore{}, faq)   // Check for the specific zero-value of FaqCore
    mockRepo.AssertExpectations(t)
}
