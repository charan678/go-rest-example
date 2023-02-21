package dao

type AlbumDAO struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}


func NewAlbumDao(id, title, artist string, price float64) *AlbumDAO {
	return &AlbumDAO{ID: id, Title: title, Artist: artist, Price: price}
}
