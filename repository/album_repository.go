package respository

import (
	"database/sql"
	"github.com/charan678/go-rest-example/domain"
	"github.com/charan678/go-rest-example/providers"
	"log"
)

type albumRespository struct {
	sqliteProvider providers.ISqliteProvider
}

func NewAlbumRepository(sqliteProvider providers.ISqliteProvider) *albumRespository {
	return &albumRespository{
		sqliteProvider: sqliteProvider,
	}
}

func (this *albumRespository) GetAlbums() ([]*domain.Album ,error) {
	db, err := this.sqliteProvider.Conn()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from albums;")
	if err != nil {
		return nil, err
	}
	return mapDaoToAlbums(rows), nil
}

func mapDaoToAlbums(rows *sql.Rows) []*domain.Album {
	albums := []*domain.Album{}
	for rows.Next() {
		var id string
		var price string
		var title string
		err := rows.Scan(&id, &title, &price)
		if err != nil {
			log.Fatal(err)
		}
		album := domain.NewAlbum()
		albums = append(albums, album)
	}
	return albums
}

