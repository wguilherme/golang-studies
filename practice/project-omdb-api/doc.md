## Executando projeto com variável de ambiente do OMDB

```bash
OMDB_API_KEY=<key> go run .
```

### Buscando dados do OMDB

```bash
 curl -X GET \
  http://localhost:8080/?s=Blade
```

#### Exemplo de resposta:

```json
{"data":{"Search":[{"Title":"Blade Runner","Year":"1982","imdbID":"tt0083658","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BOWQ4YTBmNTQtMDYxMC00NGFjLTkwOGQtNzdhNmY1Nzc1MzUxXkEyXkFqcGc@._V1_SX300.jpg"},{"Title":"Blade Runner 2049","Year":"2017","imdbID":"tt1856101","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNzA1Njg4NzYxOV5BMl5BanBnXkFtZTgwODk5NjU3MzI@._V1_SX300.jpg"},{"Title":"Blade","Year":"1998","imdbID":"tt0120611","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNzAzMmY3OWMtNDgyMS00Y2U4LTlmM2UtY2YwMmM0MDI5ODJmXkEyXkFqcGc@._V1_SX300.jpg"},{"Title":"Blade II","Year":"2002","imdbID":"tt0187738","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMGE5ZmY2NzEtZTEyMi00MWIyLThmOWYtYzJkOTQ0Y2U3ZWU1XkEyXkFqcGc@._V1_SX300.jpg"},{"Title":"Blade: Trinity","Year":"2004","imdbID":"tt0359013","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMjE0Nzg2MzI3MF5BMl5BanBnXkFtZTYwMjExODQ3._V1_SX300.jpg"},{"Title":"Sling Blade","Year":"1996","imdbID":"tt0117666","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNDM5MjhhMDgtMmVmYy00ODNlLWI1OGMtMzAzNmY5OGRmNjk0XkEyXkFqcGc@._V1_SX300.jpg"},{"Title":"Dragon Blade","Year":"2015","imdbID":"tt3672840","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTk0MjgxOTQ5MF5BMl5BanBnXkFtZTgwODA3NTUwNjE@._V1_SX300.jpg"},{"Title":"Blade of the Immortal","Year":"2017","imdbID":"tt5084170","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BZjUzODFjYmYtYzA0Yi00ODk2LTg2OGUtYWExYjc2NzYyNTc1XkEyXkFqcGc@._V1_SX300.jpg"},{"Title":"Shinobi: Heart Under Blade","Year":"2005","imdbID":"tt0475723","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMjQ0ZDgzMDUtOTE5MC00MTUwLWFmM2EtOGRiNzQ3NWNmMWQzXkEyXkFqcGc@._V1_SX300.jpg"},{"Title":"Fate/stay night [Unlimited Blade Works]","Year":"2014–2015","imdbID":"tt3621796","Type":"series","Poster":"https://m.media-amazon.com/images/M/MV5BZTNmZWI3ZWYtNjkzYi00ZTFkLWI0MzctMjc2YzlmNGRkNzRiXkEyXkFqcGc@._V1_SX300.jpg"}],"totalResults":"349","Response":"True"}}
```