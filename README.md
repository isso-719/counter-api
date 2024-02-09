# counter-api

Access counter API for isso blog.

This code explained in the blog post [here](https://isso.cc/blog/clean_arch_counter_api).

## Getting Started

1. Clone the repository
```bash
git clone git@github.com:isso-719/counter-api.git
```

2. Install the dependencies
```bash
cd counter-api
go mod download
```

3. Auth on Google Cloud
```bash
gcloud auth application-default login
```

4. Create .env file (REGION is optional, but if you deploy with deploy script, it's required) 
```bash
cp .env.example .env

vi .env
(edit the .env file with your values)
```

5. Run the application
```bash
go run cmd/main.go
```

## Deploy

1. Set Google Cloud project
```bash
gcloud config set project <project-id>
```

2. Deploy
```bash
make deploy
```
