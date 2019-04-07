package storage

import (
	"backend/models"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	//
	// Import support for PostgreSQL
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/zerolog/log"
)

//
// DB is the local handle to gorm
var _db *gorm.DB

//
// NewGORM create storages for GORM
func NewGORM() (Organisation, Party, Payment) {
	log.Info().Msg("Establishing connection to database")
	var err error
	var done bool
	var count int
	for !done {
		count++
		_db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", os.Getenv("DB_HOST"), 5432, os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")))
		if err != nil {
			if _db != nil {
				_db.Close()
			}
		} else {
			err = _db.DB().Ping()
		}
		if err != nil {
			if count > 5 {
				log.Fatal().Msgf("To many connection attempts. %v", err.Error())
			}
			log.Warn().Msgf("Waiting for database: %v", err.Error())
			time.Sleep(5 * time.Second)
		} else {
			done = true
		}
	}
	log.Info().Msg("Database connection established")
	log.Info().Msg("GORM setting up logging")
	_db.SetLogger(GORMZerolog{})
	_db.LogMode(true)
	log.Info().Msg("Database create schema")
	if !_db.HasTable(&models.OrganisationDB{}) {
		if _db.CreateTable(&models.OrganisationDB{}).Error != nil {
			log.Fatal().Msgf("Unable to create organisation table. %v", _db.Error)
		}
	}
	if !_db.HasTable(&models.PartyDB{}) {
		if _db.CreateTable(&models.PartyDB{}).Error != nil {
			log.Fatal().Msgf("Unable to create party table. %v", _db.Error)
		}
	}
	if !_db.HasTable(&models.PaymentDB{}) {
		if _db.CreateTable(&models.PaymentDB{}).Error != nil {
			log.Fatal().Msgf("Unable to create payment table. %v", _db.Error)
		}
	}
	if !_db.HasTable(&models.PaymentDB{}) {
		if _db.CreateTable(&models.PaymentDB{}).Error != nil {
			log.Fatal().Msgf("Unable to create payment table. %v", _db.Error)
		}
	}
	if !_db.HasTable(&models.ChargeDB{}) {
		if _db.CreateTable(&models.ChargeDB{}).Error != nil {
			log.Fatal().Msgf("Unable to create charge table. %v", _db.Error)
		}
	}
	//
	//
	organisation := &OrganisationGORM{db: _db}
	party := &PartyGORM{db: _db}
	payment := &PaymentGORM{db: _db, party: party}
	return organisation, party, payment
}

//
// Close the global handle to gorm
func Close() {
	_db.Close()
}
