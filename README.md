 # CEO Track service

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

## How to run

```bash
go run .\cmd\app\
```