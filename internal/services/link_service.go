package services

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"gorm.io/gorm" // Nécessaire pour la gestion spécifique de gorm.ErrRecordNotFound

	"github.com/axellelanca/urlshortener/internal/models"
	"github.com/axellelanca/urlshortener/internal/repository" // Importe le package repository
)

// Définition du jeu de caractères pour la génération des codes courts.
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// LinkService est une structure qui fournit des méthodes pour la logique métier des liens.
// Elle détient linkRepo qui est une référence vers une interface LinkRepository.
type LinkService struct {
	linkRepo repository.LinkRepository // Interface pour accéder aux méthodes du repository
}

// NewLinkService crée et retourne une nouvelle instance de LinkService.
func NewLinkService(linkRepo repository.LinkRepository) *LinkService {
	return &LinkService{
		linkRepo: linkRepo,
	}
}

// GenerateShortCode génère un code court aléatoire d'une longueur spécifiée.
func (s *LinkService) GenerateShortCode(length int) (string, error) {
	// Génère un code court aléatoire sécurisé de la longueur spécifiée
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[n.Int64()]
	}
	return string(b), nil
}

// CreateLink crée un nouveau lien raccourci.
func (s *LinkService) CreateLink(longURL string) (*models.Link, error) {
	// TODO: Implémenter la logique de création de lien avec unicité du shortcode
	return nil, nil
}

// GetLinkByShortCode récupère un lien via son code court.
func (s *LinkService) GetLinkByShortCode(shortCode string) (*models.Link, error) {
	// TODO: Utiliser le repository pour récupérer le lien
	return nil, nil
}

// GetLinkStats récupère les statistiques pour un lien donné (nombre total de clics).
func (s *LinkService) GetLinkStats(shortCode string) (*models.Link, int, error) {
	// TODO: Récupérer le lien et compter les clics
	return nil, 0, nil
}

