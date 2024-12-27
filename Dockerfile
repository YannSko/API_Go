# image de base Golang
FROM golang:1.23-alpine

# directory de work
WORKDIR /app

# Copier les fichiers Go dans le répertoire de travail du conteneur
COPY . .

# Instal les dépendances
RUN go mod tidy

# Exposer le port 8080 pour l'API Go
EXPOSE 8080

# Démarrer l'application Go
CMD ["go", "run", "main.go"]
