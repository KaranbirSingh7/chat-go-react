# Chat App with Go and React

## TODO

- [ ] Ability to add dynamic ports config
- [ ] Add proxy to frontend


## Running

### Docker

```bash
docker-compose up
```

Frontend(React) will be accessible at: http://localhost:80/

Backend(Go) will be accessible at: http://localhost:8000/

## Developing

### Backend

```bash
cd backend
go mod download
go run main.go
```

### Frontend

```bash
cd frontend
npm install
npm start
```