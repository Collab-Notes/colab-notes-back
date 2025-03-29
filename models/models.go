package models

import (
	"time"

	"github.com/google/uuid"
)

// USUARIOS
type User struct {
	ID               uuid.UUID         `gorm:"type:uuid;primaryKey"`
	Name             string            `gorm:"type:varchar(100);not null"`
	Email            string            `gorm:"type:varchar(150);unique;not null"`
	CreatedAt        time.Time         `gorm:"autoCreateTime"`
	Vaults           []Vault           `gorm:"foreignKey:OwnerID"` // Relación: owns (tiene cajas)
	VaultPermissions []VaultPermission `gorm:"foreignKey:UserID"`  // Relación: puede acceder a cajas
	NotePermissions  []NotePermission  `gorm:"foreignKey:UserID"`  // Relación: puede editar notas
}

// CAJAS (VAULTS)
type Vault struct {
	ID          uint              `gorm:"primaryKey"`
	OwnerID     uuid.UUID         `gorm:"type:uuid;not null"`
	Name        string            `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	Owner       User              `gorm:"foreignKey:OwnerID"`
	Notes       []Note            `gorm:"foreignKey:VaultID"` // Relación: contiene notas
	Permissions []VaultPermission `gorm:"foreignKey:VaultID"` // Relación: permisos asignados
}

// NOTAS (NOTES)
type Note struct {
	ID              uint             `gorm:"primaryKey"`
	VaultID         uint             `gorm:"not null"`
	OwnerID         uuid.UUID        `gorm:"type:uuid;not null"`
	Title           string           `gorm:"type:varchar(255);not null"`
	Content         string           `gorm:"type:text;not null"`
	IsPublic        bool             `gorm:"default:false"`
	CreatedAt       time.Time        `gorm:"autoCreateTime"`
	UpdatedAt       time.Time        `gorm:"autoUpdateTime"`
	Vault           Vault            `gorm:"foreignKey:VaultID"`
	Owner           User             `gorm:"foreignKey:OwnerID"`
	NoteTags        []NoteTag        `gorm:"foreignKey:NoteID"` // Relación: etiquetado
	Attachments     []NoteAttachment `gorm:"foreignKey:NoteID"` // Relación: archivos adjuntos
	NotePermissions []NotePermission `gorm:"foreignKey:NoteID"` // Relación: permisos de nota
}

// PERMISOS DE CAJAS (VAULT_PERMISSIONS)
type VaultPermission struct {
	ID          uint      `gorm:"primaryKey"`
	VaultID     uint      `gorm:"not null"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	AccessLevel string    `gorm:"type:varchar(50);not null"`
	Vault       Vault     `gorm:"foreignKey:VaultID"`
	User        User      `gorm:"foreignKey:UserID"`
}

// PERMISOS DE NOTAS (NOTE_PERMISSIONS)
type NotePermission struct {
	ID          uint      `gorm:"primaryKey"`
	NoteID      uint      `gorm:"not null"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	AccessLevel string    `gorm:"type:varchar(50);not null"`
	Note        Note      `gorm:"foreignKey:NoteID"`
	User        User      `gorm:"foreignKey:UserID"`
}

// ETIQUETAS (TAGS)
type Tag struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"type:varchar(50);not null"`
	NoteTags []NoteTag `gorm:"foreignKey:TagID"` // Relación: notas asociadas
}

// TABLA INTERMEDIA DE ETIQUETAS PARA NOTAS (NOTE_TAGS)
// Clave compuesta formada por NoteID y TagID
type NoteTag struct {
	NoteID uint `gorm:"primaryKey;not null"`
	TagID  uint `gorm:"primaryKey;not null"`
	Note   Note `gorm:"foreignKey:NoteID"`
	Tag    Tag  `gorm:"foreignKey:TagID"`
}

// ARCHIVOS ADJUNTOS DE NOTAS (NOTE_ATTACHMENTS)
type NoteAttachment struct {
	ID         uint      `gorm:"primaryKey"`
	NoteID     uint      `gorm:"not null"`
	FileURL    string    `gorm:"type:text;not null"`
	LineNumber int       `gorm:"not null"`
	UploadedAt time.Time `gorm:"autoCreateTime"`
	Note       Note      `gorm:"foreignKey:NoteID"`
}
