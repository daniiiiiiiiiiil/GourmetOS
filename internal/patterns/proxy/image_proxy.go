package proxy

import (
	"fmt"
	"sync"
	"time"
)

type Image struct {
	ID       int
	Name     string
	URL      string
	Width    int
	Height   int
	Format   string
	Data     []byte
	IsLoaded bool
}

// сервис для работы с изображениями
type RealImageService struct {
	images map[int]Image
}

func NewRealImageService() *RealImageService {
	return &RealImageService{
		images: map[int]Image{
			1: {ID: 1, Name: "маргарита.jpg", URL: "/images/pizza/margarita.jpg", Width: 800, Height: 600, Format: "jpeg", Data: []byte{}, IsLoaded: false},
			2: {ID: 2, Name: "карбонара.jpg", URL: "/images/pasta/carbonara.jpg", Width: 800, Height: 600, Format: "jpeg", Data: []byte{}, IsLoaded: false},
			3: {ID: 3, Name: "окономияки.jpg", URL: "/images/pizza/okonomiyaki.jpg", Width: 800, Height: 600, Format: "jpeg", Data: []byte{}, IsLoaded: false},
		},
	}
}

// загружает изображение (тяжелая операция)
func (s *RealImageService) LoadImage(id int) (*Image, error) {
	img, exists := s.images[id]
	if !exists {
		return nil, fmt.Errorf("изображение #%d не найдено", id)
	}

	if !img.IsLoaded {
		fmt.Printf("агрузка изображения '%s'...\n", img.Name)
		time.Sleep(500 * time.Millisecond)
		img.Data = []byte(fmt.Sprintf("image_data_%d", id))
		img.IsLoaded = true
		s.images[id] = img
		fmt.Printf("Изображение загружено (%d байт)\n", len(img.Data))
	}
	return &img, nil
}

func (s *RealImageService) GetImageInfo(id int) (string, error) {
	img, exists := s.images[id]
	if !exists {
		return "", fmt.Errorf("изображение #%d не найдено", id)
	}
	return fmt.Sprintf("%s (%dx%d, %s)", img.Name, img.Width, img.Height, img.Format), nil
}

// ПРОКСИ ДЛЯ ИЗОБРАЖЕНИЙ

type ImageProxy struct {
	id        int
	service   *RealImageService
	realImage *Image
	mu        sync.RWMutex
}

func NewImageProxy(id int, service *RealImageService) *ImageProxy {
	return &ImageProxy{
		id:      id,
		service: service,
	}
}

func (p *ImageProxy) GetInfo() (string, error) {
	return p.service.GetImageInfo(p.id)
}

func (p *ImageProxy) GetImage() (*Image, error) {
	p.mu.RLock()
	if p.realImage != nil {
		defer p.mu.RUnlock()
		fmt.Printf("Изображение #%d уже загружено\n", p.id)
		return p.realImage, nil
	}
	p.mu.RUnlock()

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.realImage != nil {
		return p.realImage, nil
	}

	img, err := p.service.LoadImage(p.id)
	if err != nil {
		return nil, err
	}
	p.realImage = img
	return p.realImage, nil
}

func (p *ImageProxy) GetName() string {
	img, exists := p.service.images[p.id]
	if !exists {
		return "неизвестно"
	}
	return img.Name
}

// проверяет, загружено ли изображение
func (p *ImageProxy) IsLoaded() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.realImage != nil && p.realImage.IsLoaded
}
