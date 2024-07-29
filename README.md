 # CEO Track service

## How to run on production

```bash
cp .env.example .env
# Edit the .env file
sudo docker-compose up -d --build
echo "0 0 * * 1 sudo docker-compose up -d --build" | sudo tee -a /etc/crontab
```
Data persist in the `./ceo_track.db` file, that'll be created in the root of the project on the first launch.

## How to run on development

```bash
cp .env.example .env
# Edit the .env file
clear; go run ./cmd/app/
```

Goal: track when a CEO is changes in xlsx files

- Detect if it's the first launch using the database
- If it's the first launch
  - Download files
  - Read files
  - Insert data into the database
- Download files
- Read files
- Compare data with the database
- Get the changes
- Update the database
- Send notifications about the changes