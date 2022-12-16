package bd

import (
	"context"
	"time"

	"github.com/agustinLoag/restAPiTwitt/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitt")
	col := db.Collection("usuarios")

	filter := bson.M{"email": email}
	var resultado models.Usuario

	err := col.FindOne(ctx, filter).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
